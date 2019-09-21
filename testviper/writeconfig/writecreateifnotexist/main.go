/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package main

import (
	"github.com/spf13/viper"
)

func main() {

	viper.Set("Name", "Jirawat Kim")

	writeConfig()
}


func writeConfig() {

	viper.SetConfigFile("config.yaml")
	err := viper.WriteConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			err := viper.SafeWriteConfig()
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}

}