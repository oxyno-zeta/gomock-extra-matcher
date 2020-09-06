package extra

import (
	"fmt"
	"reflect"

	"github.com/golang/mock/gomock"
)

// MMatcher interface improved to add functions.
type MMatcher interface {
	gomock.Matcher
	// Add a matcher for a specific key
	Key(key interface{}, match interface{}) MMatcher
}

type mStorage struct {
	key   interface{}
	match interface{}
}

type mapMatcher struct {
	keys []*mStorage
}

func (m *mapMatcher) String() string {
	str := ""

	// Loop over all keys
	for in, v := range m.keys {
		// Add separator for display
		if in > 0 {
			str += ", "
		}

		// Check if key is a gomock matcher
		m, ok := v.key.(gomock.Matcher)
		if ok {
			str += fmt.Sprintf("key %s", m.String())
		} else {
			str += fmt.Sprintf("key %v", v.key)
		}

		// Try to cast to a matcher interface
		m, ok = v.match.(gomock.Matcher)
		// Check if cast is ok
		if ok {
			str += fmt.Sprintf(" must match %s", m.String())
		} else {
			str += fmt.Sprintf(" must be equal to %v", v.match)
		}
	}

	return str
}

func (m *mapMatcher) Matches(x interface{}) bool {
	// Check if x is nil
	if x == nil {
		return false
	}

	// Value of interface input
	rval := reflect.ValueOf(x)
	rkind := rval.Kind()
	// Check if reflect value is supported or not
	if rkind != reflect.Map {
		return false
	}

	// Default case
	res := len(m.keys) != 0

	// Create reflect indirect
	// indirect := reflect.Indirect(rval)
	// Loop over all matcher keys
	for _, kk := range m.keys {
		// Store if matcher key can be found
		matchKeyFound := false
		// Loop over map keys
		for _, kVal := range rval.MapKeys() {
			// Get key data
			keyD := kVal.Interface()
			// Get reflect value from key
			rv := rval.MapIndex(kVal)
			// Get data from key
			val := rv.Interface()
			// Check if matcher key is matching current key
			if !isMatchingData(kk.key, keyD) {
				// Skip this key
				continue
			}

			// Set match key as found
			matchKeyFound = true

			// Check map key value is matching
			res = res && isMatchingData(kk.match, val)

			// Break the loop at this step
			// No need to continue to check map values
			break
		}

		// Check if match key was found
		if !matchKeyFound {
			// Match key wasn't found, it is an error
			return false
		}

		// Check result
		if !res {
			// If result isn't true at this step, stop now
			return false
		}
	}

	return res
}

func isMatchingData(matchKey, keyData interface{}) bool {
	// Check if given key in matcher is a gomock matcher
	mk, ok := matchKey.(gomock.Matcher)
	if ok {
		return mk.Matches(keyData)
	}

	return matchKey == keyData
}

func (m *mapMatcher) Key(key interface{}, match interface{}) MMatcher {
	// Check if key exists
	if key == nil {
		return m
	}
	// Key name exists => add data
	m.keys = append(m.keys, &mStorage{key: key, match: match})
	// Return
	return m
}

// MapMatcher will return a new map matcher.
func MapMatcher() MMatcher { return &mapMatcher{} }
