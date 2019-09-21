/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil { // Find and read the config file
		panic(err)
	}

	name := viper.Get("Name")

	fmt.Println(name)

}
