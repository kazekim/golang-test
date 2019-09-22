/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package main

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	yaml "gopkg.in/yaml.v2"
)

func main() {

	v := viper.New()
	v.SetConfigName("testconfig")
	v.AddConfigPath(".")

	if err := v.ReadInConfig(); err != nil { // Find and read the config file
		panic(err)
	}

	s := yamlStringSettings(v)
	fmt.Println(s)
}

func yamlStringSettings(v *viper.Viper) string {
	c := v.AllSettings()
	bs, err := yaml.Marshal(c)
	if err != nil {
		log.Fatalf("unable to marshal config to YAML: %v", err)
	}
	return string(bs)
}
