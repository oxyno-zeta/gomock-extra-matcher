package extra

import (
	"reflect"
	"testing"

	"go.uber.org/mock/gomock"
)

func Test_structMatcher_Field(t *testing.T) {
	type fields struct {
		fields []*sStorage
	}
	type args struct {
		fName   string
		matcher interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []*sStorage
	}{
		{
			name: "empty field name",
		},
		{
			name: "empty matcher",
			args: args{fName: "fake"},
			want: []*sStorage{{
				fName: "fake",
			}},
		},
		{
			name: "add already existing field",
			fields: fields{
				fields: []*sStorage{{
					fName: "fake",
					match: "value",
				}},
			},
			args: args{fName: "fake", matcher: gomock.Eq("value1")},
			want: []*sStorage{
				{fName: "fake", match: "value"},
				{fName: "fake", match: gomock.Eq("value1")},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &structMatcher{
				fields: tt.fields.fields,
			}
			f.Field(tt.args.fName, tt.args.matcher)
			if !reflect.DeepEqual(f.fields, tt.want) {
				t.Errorf("structMatcher.Field() = %v, want %v", f.fields, tt.want)
			}
		})
	}
}

func Test_structMatcher_String(t *testing.T) {
	type fields struct {
		fields []*sStorage
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "empty",
			want: "",
		},
		{
			name: "gomock matcher",
			fields: fields{
				fields: []*sStorage{{
					fName: "fake",
					match: gomock.Nil(),
				}},
			},
			want: "field fake must match is nil",
		},
		{
			name: "match value",
			fields: fields{
				fields: []*sStorage{{
					fName: "fake",
					match: "value",
				}},
			},
			want: "field fake must be equal to value",
		},
		{
			name: "gomock matcher and match value",
			fields: fields{
				fields: []*sStorage{{
					fName: "fake",
					match: gomock.Nil(),
				}, {
					fName: "fake2",
					match: "value",
				}},
			},
			want: "field fake must match is nil, field fake2 must be equal to value",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &structMatcher{
				fields: tt.fields.fields,
			}
			if got := f.String(); got != tt.want {
				t.Errorf("structMatcher.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_structMatcher_Matches(t *testing.T) {
	starStrFunc := func(s string) *string { return &s }
	type innerFakeStruct struct {
		Bo2  bool
		Stp2 *string
		St2  string
	}
	type fakeStruct struct {
		Bo     bool
		St     string
		Stp    *string
		I      int
		Mm     map[string]string
		InnerP *innerFakeStruct
		Inner  innerFakeStruct
	}

	type fields struct {
		fields []*sStorage
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
			name: "nil input",
			args: args{x: nil},
			want: false,
		},
		{
			name: "bool input",
			args: args{x: false},
			want: false,
		},
		{
			name: "string input",
			args: args{x: "string"},
			want: false,
		},
		{
			name: "string pointer input",
			args: args{x: starStrFunc("string")},
			want: false,
		},
		{
			name: "int input",
			args: args{x: 8},
			want: false,
		},
		{
			name: "map input",
			args: args{x: map[string]string{}},
			want: false,
		},
		{
			name: "func input",
			args: args{x: func() {}},
			want: false,
		},
		{
			name: "shouldn't match first level struct",
			fields: fields{
				fields: []*sStorage{
					{fName: "Stp", match: gomock.Any()},
					{fName: "Stp", match: starStrFunc("fake")},
					{fName: "St", match: "fake"},
				},
			},
			args: args{
				x: fakeStruct{},
			},
			want: false,
		},
		{
			name: "shouldn't match first level struct 2",
			fields: fields{
				fields: []*sStorage{
					{fName: "Stp", match: gomock.Any()},
					{fName: "Stp", match: starStrFunc("fake")},
					{fName: "St", match: "fake"},
				},
			},
			args: args{
				x: fakeStruct{Stp: starStrFunc("fake2")},
			},
			want: false,
		},
		{
			name: "should match first level struct",
			fields: fields{
				fields: []*sStorage{
					{fName: "Stp", match: gomock.Any()},
					{fName: "Stp", match: gomock.Eq(starStrFunc("fake"))},
					{fName: "St", match: "fake"},
				},
			},
			args: args{
				x: fakeStruct{
					Stp: starStrFunc("fake"),
					St:  "fake",
				},
			},
			want: true,
		},
		{
			name: "shouldn't match first level struct because field not found",
			fields: fields{
				fields: []*sStorage{
					{fName: "Stfake", match: "fake"},
				},
			},
			args: args{
				x: fakeStruct{
					Stp: starStrFunc("fake"),
					St:  "fake",
				},
			},
			want: false,
		},
		{
			name: "should match first level struct (pointer)",
			fields: fields{
				fields: []*sStorage{
					{fName: "Stp", match: gomock.Any()},
					{fName: "Stp", match: gomock.Eq(starStrFunc("fake"))},
					{fName: "St", match: "fake"},
				},
			},
			args: args{
				x: &fakeStruct{
					Stp: starStrFunc("fake"),
					St:  "fake",
				},
			},
			want: true,
		},
		{
			name: "should match first level struct (map)",
			fields: fields{
				fields: []*sStorage{
					{fName: "Stp", match: gomock.Any()},
					{fName: "Stp", match: gomock.Eq(starStrFunc("fake"))},
					{fName: "St", match: "fake"},
					{fName: "Mm", match: gomock.Eq(map[string]string{"fake1": "fake1"})},
				},
			},
			args: args{
				x: &fakeStruct{
					Stp: starStrFunc("fake"),
					St:  "fake",
					Mm:  map[string]string{"fake1": "fake1"},
				},
			},
			want: true,
		},
		{
			name: "should match first and second level struct",
			fields: fields{
				fields: []*sStorage{
					{fName: "Stp", match: gomock.Any()},
					{fName: "Stp", match: gomock.Eq(starStrFunc("fake"))},
					{fName: "St", match: "fake"},
					{
						fName: "Inner",
						match: StructMatcher().Field("St2", "fake2").Field("Stp2", gomock.Eq(starStrFunc("fake2"))),
					},
					{
						fName: "InnerP",
						match: StructMatcher().Field("St2", "fake2").Field("Stp2", gomock.Eq(starStrFunc("fake2"))),
					},
				},
			},
			args: args{
				x: &fakeStruct{
					Stp: starStrFunc("fake"),
					St:  "fake",
					Inner: innerFakeStruct{
						Stp2: starStrFunc("fake2"),
						St2:  "fake2",
					},
					InnerP: &innerFakeStruct{
						Stp2: starStrFunc("fake2"),
						St2:  "fake2",
					},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &structMatcher{
				fields: tt.fields.fields,
			}
			if got := f.Matches(tt.args.x); got != tt.want {
				t.Errorf("structMatcher.Matches() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStructMatcher(t *testing.T) {
	tests := []struct {
		name string
		want StMatcher
	}{
		{name: "constructor", want: &structMatcher{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StructMatcher(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StructMatcher() = %v, want %v", got, tt.want)
			}
		})
	}
}
