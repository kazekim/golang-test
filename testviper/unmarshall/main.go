/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Database Database
}

type Database struct {
	Username string
	Pass string `mapstructure:"Password"`
}

func main() {

	var c Config

	v := viper.New()
	v.SetConfigName("testconfig")
	v.AddConfigPath(".")

	if err := v.ReadInConfig(); err != nil { // Find and read the config file
		panic(err)
	}

	err := v.Unmarshal(&c)
	if err != nil {
		panic(err)
	}

	fmt.Println(c)
}
