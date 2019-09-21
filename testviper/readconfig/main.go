/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigName("testconfig") // name of config file (without extension)
	viper.AddConfigPath("/etc/appname/")   // path to look for the config file in
	viper.AddConfigPath("config")   // path to look for the config file in
	viper.AddConfigPath("$HOME/.appname")  // call multiple times to add many search paths
	viper.AddConfigPath(".")               // optionally look for config in the working directory

	if err := viper.ReadInConfig(); err != nil { // Find and read the config file

		// Handle errors reading the config file
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
		} else {
			// Config file was found but another error was produced
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	}

	test := viper.Get("Test")

	fmt.Println(test)
}
