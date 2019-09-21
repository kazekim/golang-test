/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigFile("config.yaml")


	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	writeConfig1()
	writeConfig2()
}

func writeConfig1() {

	viper.Set("Name", "Jirawat Kim")
	err := viper.WriteConfig()
	if err != nil {
		panic(err)
	}
}

func writeConfig2() {
	viper.Set("WebSite", "jirawat.kim")
	err := viper.WriteConfig()
	if err != nil {
		panic(err)
	}
}