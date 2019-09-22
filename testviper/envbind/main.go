/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func main() {

	os.Setenv("KIM_ID", "18")

	viper.SetEnvPrefix("kim") // will be uppercased automatically
	viper.BindEnv("id")


	id := viper.Get("id")

	fmt.Println(id)

	id2 := os.Getenv("KIM_ID")

	fmt.Println(id2)

}
