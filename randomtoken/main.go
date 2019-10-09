/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package main

import (
	"fmt"
	"github.com/kazekim/golang-test/randomtoken/gentoken"
)

func main() {
	token, err := gentoken.GenerateRandomStringPrefix("xent", 34)
	if err != nil {
		panic(err)
	}

	fmt.Println(token)

}
