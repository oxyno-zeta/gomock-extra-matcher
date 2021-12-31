package extra

import (
	"fmt"
	"regexp"

	"github.com/golang/mock/gomock"
)

type stringRegexpMatcher struct {
	reg *regexp.Regexp
}

func (s *stringRegexpMatcher) String() string {
	return fmt.Sprintf("input matching regexp %s", s.reg.String())
}

func (s *stringRegexpMatcher) Matches(x interface{}) bool {
	// Try to cast input as string
	st, ok := x.(string)
	// Check if it is a string
	if !ok {
		return false
	}

	return s.reg.MatchString(st)
}

// Will return a new string regexp matcher.
func StringRegexpMatcher(regexSt string) gomock.Matcher {
	return &stringRegexpMatcher{
		reg: regexp.MustCompile(regexSt),
	}
}
