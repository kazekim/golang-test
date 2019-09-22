/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package main

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
)

func main() {


	viper.SetConfigType("yaml") // or viper.SetConfigType("YAML")

	// any approach to require this configuration into your program.
	var yamlExample = []byte(`
Hacker: true
name: Kim
hobbies:
- BoardGame
- Travel
- go
clothing:
  jacket: shirt
  trousers: jeans
eyes : brown
beard: false
`)

	err := viper.ReadConfig(bytes.NewBuffer(yamlExample))
	if err != nil {
		panic(err)
	}

	name := viper.Get("name") // this would be "Kim"

	fmt.Println(name)
}
