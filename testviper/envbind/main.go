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

	viper.SetEnvPrefix("kim") // will be uppercased automatically
	viper.BindEnv("id")

	os.Setenv("KIM_ID", "18")

	id := viper.Get("id")

	fmt.Println(id)

}
