package extra

import (
	"fmt"

	"go.uber.org/mock/gomock"
)

type orMatcher struct {
	matchers []gomock.Matcher
}

func (om *orMatcher) String() string {
	if len(om.matchers) == 0 {
		return "the \"or\" matcher will return false because list is empty"
	}

	// Initialize string
	str := ""

	// Loop over matchers
	for i, m := range om.matchers {
		// Ignore the first item
		if i > 0 {
			str += " or "
		}

		// Concat matcher string
		str += fmt.Sprintf("(%s)", m.String())
	}

	return str
}

func (om *orMatcher) Matches(x interface{}) bool {
	// Check empty case
	if len(om.matchers) == 0 {
		return false
	}

	// Loop over matchers
	for _, m := range om.matchers {
		// Check if matcher is ok
		if m.Matches(x) {
			// Matches so ... End
			return true
		}
	}

	// No match until now
	// Or is false
	return false
}

// OrMatcher will return a new Or matcher.
func OrMatcher(matchers ...gomock.Matcher) gomock.Matcher {
	return &orMatcher{matchers: matchers}
}
