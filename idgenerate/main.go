/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package main

import (
	"fmt"
	"github.com/kazekim/golang-test/idgenerate/identifier"
)

func main() {

	v, err := identifier.NewIdentifier()
	if err != nil {
		panic(err)
	}

	v.SetPrefix("xent")

	key := v.GenerateNewKey()
	secret, err := v.GenerateNewSecret()
	if err != nil {
		panic(err)
	}

	fmt.Println(v.ID())
	fmt.Println(key)
	fmt.Println(secret)

	v, err = identifier.IdentifierFromKey(key)
	if err != nil {
		panic(err)
	}

	fmt.Println(v.ID())

	v, err = identifier.IdentifierFromSecret(secret)
	if err != nil {
		panic(err)
	}
}
