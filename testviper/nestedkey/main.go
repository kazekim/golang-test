/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigName("testconfig")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil { // Find and read the config file
		panic(err)
	}
	username := viper.GetString("Database.Username")
	fmt.Println(username)


}
