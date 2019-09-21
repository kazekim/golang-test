/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package main

import (
	"github.com/spf13/viper"
)

func main() {

	viper.Set("Name", "Jirawat Kim")

	viper.SetConfigFile("config/config.yaml")

	err := viper.SafeWriteConfig()
	if err != nil {
		panic(err)
	}

	err = viper.SafeWriteConfigAs("jsonconfig.json")
	if err != nil {
		panic(err)
	}

}
