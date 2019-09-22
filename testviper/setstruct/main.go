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
	Password string
}

func main() {

	config := Config{
		Database: Database{
			Username: "admin",
			Password: "P@ssW0rd",
		},
	}

	viper.Set("Config", config)

	c := viper.Get("Config")
	fmt.Println(c)
}
