package extra

import (
	"fmt"
	"reflect"

	"go.uber.org/mock/gomock"
)

// StMatcher interface improved to add functions.
type StMatcher interface {
	gomock.Matcher
	// Add a matcher for a specific field
	Field(fName string, match interface{}) StMatcher
}

type sStorage struct {
	match interface{}
	fName string
}

type structMatcher struct {
	fields []*sStorage
}

func (f *structMatcher) String() string {
	str := ""

	// Loop over all fields
	for in, v := range f.fields {
		// Add separator for display
		if in > 0 {
			str += ", "
		}

		str += fmt.Sprintf("field %s", v.fName)

		// Try to cast to a matcher interface
		m, ok := v.match.(gomock.Matcher)
		// Check if cast is ok
		if ok {
			str += fmt.Sprintf(" must match %s", m.String())
		} else {
			str += fmt.Sprintf(" must be equal to %v", v.match)
		}
	}

	return str
}

func (f *structMatcher) Field(fName string, match interface{}) StMatcher {
	// Check if field name exists
	if fName == "" {
		return f
	}
	// Field name exists => add data
	f.fields = append(f.fields, &sStorage{fName: fName, match: match})
	// Return
	return f
}

func (f *structMatcher) Matches(x interface{}) bool {
	// Check if x is nil
	if x == nil {
		return false
	}

	// Value of interface input
	rval := reflect.ValueOf(x)
	rkind := rval.Kind()
	// Check if reflect value is supported or not
	if rkind != reflect.Struct && rkind != reflect.Ptr {
		return false
	}

	// Default case
	res := len(f.fields) != 0

	// Create reflect indirect
	indirect := reflect.Indirect(rval)
	// Loop over all fields
	for _, v := range f.fields {
		// Try to get field value
		fval := indirect.FieldByName(v.fName)
		// Check if field doesn't exist
		if !fval.IsValid() {
			// In this case returning false is ok
			// This case appears when the field isn't found in structure
			return false
		}

		// Get data from field
		data := fval.Interface()

		// Try to cast a gomock matcher
		m, ok := v.match.(gomock.Matcher)
		// Check if cast is ok
		if ok {
			// Run matcher
			res = res && m.Matches(data)
		} else {
			res = res && v.match == data
		}

		// Check result
		if !res {
			// If result isn't true at this step, stop now
			return false
		}
	}

	// Default case
	return res
}

// StructMatcher will return a new struct matcher.
func StructMatcher() StMatcher { return &structMatcher{} }
