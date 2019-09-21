/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {

	viper.SetDefault("ContentDir", "content")
	viper.SetDefault("LayoutDir", "layouts")
	viper.SetDefault("Taxonomies", map[string]string{"tag": "tags", "category": "categories"})

	ld := viper.Get("LayoutDir")

	fmt.Println(ld)


	viper.Set("Name", "Kim")

	viper.RegisterAlias("Name", "NickName")

	name := viper.Get("Name")

	fmt.Println(name)

	nickName := viper.Get("NickName")

	fmt.Println(nickName)
}