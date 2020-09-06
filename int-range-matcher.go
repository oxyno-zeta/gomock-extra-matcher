package extra

import (
	"fmt"

	"github.com/golang/mock/gomock"
)

type intRangeMatcher struct {
	lowerBound, upperBound int
}

func (i *intRangeMatcher) String() string {
	return fmt.Sprintf("it upper than %d and lower than %d", i.lowerBound, i.upperBound)
}

func (i *intRangeMatcher) Matches(x interface{}) bool {
	// Try to cast input as int
	inp, ok := x.(int)
	// Check
	if !ok {
		return false
	}

	return i.lowerBound <= inp && inp <= i.upperBound
}

func IntRangeMatcher(lowerBound, upperBound int) gomock.Matcher {
	return &intRangeMatcher{lowerBound: lowerBound, upperBound: upperBound}
}
