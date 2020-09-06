package extra

import (
	"reflect"
	"regexp"
	"testing"
)

func Test_stringRegexp_String(t *testing.T) {
	type fields struct {
		reg *regexp.Regexp
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "should display regex",
			fields: fields{
				reg: regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`),
			},
			want: `input matching regexp ^[a-z]+\[[0-9]+\]$`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &stringRegexpMatcher{
				reg: tt.fields.reg,
			}
			if got := s.String(); got != tt.want {
				t.Errorf("stringRegexp.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringRegexp_Matches(t *testing.T) {
	starStrFunc := func(s string) *string { return &s }
	type fields struct {
		reg *regexp.Regexp
	}
	type args struct {
		x interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "not a string",
			args: args{
				x: 1,
			},
			want: false,
		},
		{
			name: "not a string 2",
			args: args{
				x: starStrFunc("fake"),
			},
			want: false,
		},
		{
			name: "not matching regexp",
			fields: fields{
				reg: regexp.MustCompile("^a$"),
			},
			args: args{
				x: "0",
			},
		},
		{
			name: "matching regexp",
			fields: fields{
				reg: regexp.MustCompile("^a$"),
			},
			args: args{
				x: "a",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &stringRegexpMatcher{
				reg: tt.fields.reg,
			}
			if got := s.Matches(tt.args.x); got != tt.want {
				t.Errorf("stringRegexp.Matches() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringRegexpMatcher(t *testing.T) {
	type args struct {
		regexSt string
	}
	tests := []struct {
		name string
		args args
		want *stringRegexpMatcher
	}{
		{
			name: "constructor",
			args: args{
				regexSt: "^a$",
			},
			want: &stringRegexpMatcher{
				reg: regexp.MustCompile("^a$"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringRegexpMatcher(tt.args.regexSt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringRegexpMatcher() = %v, want %v", got, tt.want)
			}
		})
	}
}
