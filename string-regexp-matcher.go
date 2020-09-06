package extra

import (
	"fmt"
	"regexp"

	"github.com/golang/mock/gomock"
)

type stringRegexp struct {
	reg *regexp.Regexp
}

func (s *stringRegexp) String() string {
	return fmt.Sprintf("input matching regexp %s", s.reg.String())
}

func (s *stringRegexp) Matches(x interface{}) bool {
	// Try to cast input as string
	st, ok := x.(string)
	// Check if it is a string
	if !ok {
		return false
	}

	return s.reg.Match([]byte(st))
}

// Will return a new string regexp matcher
func StringRegexpMatcher(regexSt string) gomock.Matcher {
	return &stringRegexp{
		reg: regexp.MustCompile(regexSt),
	}
}
