package extra

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_mapMatcher_Matches(t *testing.T) {
	starStrFunc := func(s string) *string { return &s }
	type fakeStruct struct {
		Bo2 bool
	}
	type fields struct {
		keys []*mStorage
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
			name: "struct pointer input",
			args: args{x: &fakeStruct{}},
			want: false,
		},
		{
			name: "struct input",
			args: args{x: fakeStruct{}},
			want: false,
		},
		{
			name: "func input",
			args: args{x: func() {}},
			want: false,
		},
		{
			name: "key not found on map[*string][string]",
			fields: fields{
				keys: []*mStorage{
					{key: "fake", match: "data"},
				},
			},
			args: args{
				x: map[*string]string{
					starStrFunc("fake"): "data",
				},
			},
			want: false,
		},
		{
			name: "key not found on map[string][string]",
			fields: fields{
				keys: []*mStorage{
					{key: "fake2", match: "data"},
				},
			},
			args: args{
				x: map[string]string{
					"fake": "data",
				},
			},
			want: false,
		},
		{
			name: "shouldn't match with no keys in matcher",
			args: args{
				x: map[*string]string{
					starStrFunc("fake"): "data",
				},
			},
			want: false,
		},
		{
			name: "shouldn't match map[*string][string] with value",
			fields: fields{
				keys: []*mStorage{
					{key: gomock.Eq(starStrFunc("fake")), match: "data-fake"},
				},
			},
			args: args{
				x: map[*string]string{
					starStrFunc("fake"): "data",
				},
			},
			want: false,
		},
		{
			name: "shouldn't match map[string][string] with value",
			fields: fields{
				keys: []*mStorage{
					{key: "fake", match: "data-fake"},
				},
			},
			args: args{
				x: map[string]string{
					"fake": "data",
				},
			},
			want: false,
		},
		{
			name: "should match map[*string][string] with value",
			fields: fields{
				keys: []*mStorage{
					{key: gomock.Eq(starStrFunc("fake")), match: "data"},
				},
			},
			args: args{
				x: map[*string]string{
					starStrFunc("fake"): "data",
				},
			},
			want: true,
		},
		{
			name: "should match map[*string][string] with value and a key ingored",
			fields: fields{
				keys: []*mStorage{
					{key: gomock.Eq(starStrFunc("fake")), match: "data"},
				},
			},
			args: args{
				x: map[*string]string{
					starStrFunc("fake"):  "data",
					starStrFunc("fake2"): "data2",
				},
			},
			want: true,
		},
		{
			name: "should match map[*string][*string] with a gomock matcher",
			fields: fields{
				keys: []*mStorage{
					{key: gomock.Eq(starStrFunc("fake")), match: gomock.Eq(starStrFunc("data"))},
				},
			},
			args: args{
				x: map[*string]*string{
					starStrFunc("fake"): starStrFunc("data"),
				},
			},
			want: true,
		},
		{
			name: "should match map[string][*struct] with a gomock matcher",
			fields: fields{
				keys: []*mStorage{
					{key: "fake", match: gomock.Eq(&fakeStruct{
						Bo2: true,
					})},
				},
			},
			args: args{
				x: map[string]*fakeStruct{
					"fake": {Bo2: true},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mapMatcher{
				keys: tt.fields.keys,
			}
			if got := m.Matches(tt.args.x); got != tt.want {
				t.Errorf("mapMatcher.Matches() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mapMatcher_String(t *testing.T) {
	type fields struct {
		keys []*mStorage
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
			name: "with gomock matcher as key and value",
			fields: fields{
				keys: []*mStorage{
					{key: gomock.Any(), match: gomock.Nil()},
				},
			},
			want: "key is anything must match is nil",
		},
		{
			name: "with gomock matcher as key and value as value",
			fields: fields{
				keys: []*mStorage{
					{key: gomock.Any(), match: "data"},
				},
			},
			want: "key is anything must be equal to data",
		},
		{
			name: "with value as key and value as value",
			fields: fields{
				keys: []*mStorage{
					{key: "fake", match: "data"},
				},
			},
			want: "key fake must be equal to data",
		},
		{
			name: "with value as key and value as value",
			fields: fields{
				keys: []*mStorage{
					{key: "fake", match: "data"},
				},
			},
			want: "key fake must be equal to data",
		},
		{
			name: "with value as key and gomock matcher as value",
			fields: fields{
				keys: []*mStorage{
					{key: "fake", match: gomock.Any()},
				},
			},
			want: "key fake must match is anything",
		},
		{
			name: "with multiple keys",
			fields: fields{
				keys: []*mStorage{
					{key: "fake", match: gomock.Any()},
					{key: "fake2", match: gomock.Any()},
				},
			},
			want: "key fake must match is anything, key fake2 must match is anything",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mapMatcher{
				keys: tt.fields.keys,
			}
			if got := m.String(); got != tt.want {
				t.Errorf("mapMatcher.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mapMatcher_Key(t *testing.T) {
	type fields struct {
		keys []*mStorage
	}
	type args struct {
		key   interface{}
		match interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []*mStorage
	}{
		{
			name: "should ingore when key is nil",
			want: nil,
		},
		{
			name: "empty matcher",
			args: args{key: "fake"},
			want: []*mStorage{
				{key: "fake"},
			},
		},
		{
			name: "add already existing key",
			fields: fields{
				keys: []*mStorage{{
					key:   "fake",
					match: "value",
				}},
			},
			args: args{key: "fake", match: gomock.Eq("value1")},
			want: []*mStorage{
				{key: "fake", match: "value"},
				{key: "fake", match: gomock.Eq("value1")},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mapMatcher{
				keys: tt.fields.keys,
			}
			m.Key(tt.args.key, tt.args.match)
			if !reflect.DeepEqual(m.keys, tt.want) {
				t.Errorf("mapMatcher.Key() = %v, want %v", m.keys, tt.want)
			}
		})
	}
}
