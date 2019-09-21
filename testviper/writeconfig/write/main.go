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
	err := viper.WriteConfig()

	if err != nil {
		panic(err)
	}

	err = viper.WriteConfigAs(".config/config.yaml")
	if err != nil {
		panic(err)
	}

	// viper.SafeWriteConfigAs("/path/to/my/.config") // will error since it has already been written

	err = viper.WriteConfigAs("config/jsonconfig.json")
	if err != nil {
		panic(err)
	}
}
