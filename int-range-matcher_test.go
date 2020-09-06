package extra

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestIntRangeMatcher(t *testing.T) {
	type args struct {
		lowerBound int
		upperBound int
	}
	tests := []struct {
		name string
		args args
		want gomock.Matcher
	}{
		{
			name: "constructor",
			args: args{lowerBound: 5, upperBound: 15},
			want: &intRangeMatcher{
				lowerBound: 5,
				upperBound: 15,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntRangeMatcher(tt.args.lowerBound, tt.args.upperBound); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntRangeMatcher() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intRange_String(t *testing.T) {
	type fields struct {
		lowerBound int
		upperBound int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "display",
			fields: fields{lowerBound: 5, upperBound: 15},
			want:   "it upper than 5 and lower than 15",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &intRangeMatcher{
				lowerBound: tt.fields.lowerBound,
				upperBound: tt.fields.upperBound,
			}
			if got := i.String(); got != tt.want {
				t.Errorf("intRange.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intRange_Matches(t *testing.T) {
	type fields struct {
		lowerBound int
		upperBound int
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
			name: "not an int",
			args: args{x: "string"},
			want: false,
		},
		{
			name: "not an int 2",
			args: args{x: true},
			want: false,
		},
		{
			name: "should match 0 case",
			args: args{x: 0},
			want: true,
		},
		{
			name:   "should match in range",
			args:   args{x: 5},
			fields: fields{lowerBound: 1, upperBound: 15},
			want:   true,
		},
		{
			name:   "shouldn't match not in range",
			args:   args{x: 25},
			fields: fields{lowerBound: 1, upperBound: 15},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &intRangeMatcher{
				lowerBound: tt.fields.lowerBound,
				upperBound: tt.fields.upperBound,
			}
			if got := i.Matches(tt.args.x); got != tt.want {
				t.Errorf("intRange.Matches() = %v, want %v", got, tt.want)
			}
		})
	}
}
