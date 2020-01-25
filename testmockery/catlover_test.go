/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package testmockery

import (
	"errors"
	"github.com/kazekim/golang-test/testmockery/mocks"
	"testing"
)

func TestFeedCat(t *testing.T) {
	c := &mocks.Cat{}

	cl := &CatLover{Cat:c}

	// Expect Do to be called once with 1 and "abc" as parameters, and return nil from the mocked call.
	c.On("Eat",  "fish", 1).Return(errors.New("cat is still hungry")).Once()

	err := cl.FeedCat()
	if err != nil {
		t.Error(err.Error())
	}
	c.AssertExpectations(t)
}