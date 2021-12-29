package extra

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_orMatcher_String(t *testing.T) {
	type fields struct {
		matchers []gomock.Matcher
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "empty",
			fields: fields{},
			want:   "the \"or\" matcher will return false because list is empty",
		},
		{
			name: "1 element", // Even if it weird
			fields: fields{
				matchers: []gomock.Matcher{
					gomock.Any(),
				},
			},
			want: "(is anything)",
		},
		{
			name: "multiple elements",
			fields: fields{
				matchers: []gomock.Matcher{
					gomock.Any(),
					gomock.Eq(true),
				},
			},
			want: "(is anything) or (is equal to true (bool))",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			om := &orMatcher{
				matchers: tt.fields.matchers,
			}
			if got := om.String(); got != tt.want {
				t.Errorf("orMatcher.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_orMatcher_Matches(t *testing.T) {
	type fields struct {
		matchers []gomock.Matcher
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
			name: "empty",
			fields: fields{
				matchers: []gomock.Matcher{},
			},
			args: args{x: false},
			want: false,
		},
		{
			name: "1 element (ok)", // Even if it is weird
			fields: fields{
				matchers: []gomock.Matcher{
					gomock.Any(),
				},
			},
			args: args{x: false},
			want: true,
		},
		{
			name: "1 element (ko)", // Even if it is weird
			fields: fields{
				matchers: []gomock.Matcher{
					gomock.Eq(true),
				},
			},
			args: args{x: false},
			want: false,
		},
		{
			name: "multiple elements (all ok)",
			fields: fields{
				matchers: []gomock.Matcher{
					gomock.Any(),
					gomock.Any(),
				},
			},
			args: args{x: false},
			want: true,
		},
		{
			name: "multiple elements (first ok, second ko)",
			fields: fields{
				matchers: []gomock.Matcher{
					gomock.Any(),
					gomock.Eq(true),
				},
			},
			args: args{x: false},
			want: true,
		},
		{
			name: "multiple elements (first ko, second ok)",
			fields: fields{
				matchers: []gomock.Matcher{
					gomock.Eq(true),
					gomock.Any(),
				},
			},
			args: args{x: false},
			want: true,
		},
		{
			name: "multiple elements (first ko, second ko)",
			fields: fields{
				matchers: []gomock.Matcher{
					gomock.Eq(true),
					gomock.Eq(true),
				},
			},
			args: args{x: false},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			om := &orMatcher{
				matchers: tt.fields.matchers,
			}
			if got := om.Matches(tt.args.x); got != tt.want {
				t.Errorf("orMatcher.Matches() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrMatcher(t *testing.T) {
	type args struct {
		matchers []gomock.Matcher
	}
	tests := []struct {
		name string
		args args
		want gomock.Matcher
	}{
		{
			name: "init",
			args: args{matchers: []gomock.Matcher{}},
			want: &orMatcher{matchers: []gomock.Matcher{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OrMatcher(tt.args.matchers...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrMatcher() = %v, want %v", got, tt.want)
			}
		})
	}
}
