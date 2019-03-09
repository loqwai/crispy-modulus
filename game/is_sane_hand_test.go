package game_test

import (
	"fmt"

	"github.com/onsi/gomega/types"
)

func BeASaneHand() types.GomegaMatcher {
	return &_SaneHandMatcher{}
}

type _SaneHandMatcher struct{}

func (m *_SaneHandMatcher) Match(actual interface{}) (success bool, err error) {
	cards, ok := actual.([]int)
	if !ok {
		return false, fmt.Errorf("BeSaneHand expects an []int")
	}

	for i := 0; i < len(cards); i++ {
		for j := 0; j < len(cards); j++ {
			if i == j {
				continue
			}
			if cards[i] == cards[j] {
				return false, nil
			}
		}
	}
	return true, nil
}

func (m *_SaneHandMatcher) FailureMessage(actual interface{}) string {
	return fmt.Sprintf("Expected %v to only have one of each type of card", actual)
}

func (m *_SaneHandMatcher) NegatedFailureMessage(actual interface{}) string {
	return fmt.Sprintf("Expected %v not to only have one of each type of card", actual)
}
