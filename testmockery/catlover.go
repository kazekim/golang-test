/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package testmockery

import "github.com/kazekim/golang-test/testmockery/pet"

type CatLover struct {
	Cat pet.Cat
}

func (u *CatLover) FeedCat() error {
	return u.Cat.Eat("fish", 1)
}
