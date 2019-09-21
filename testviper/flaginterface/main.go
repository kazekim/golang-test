/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type myFlag struct {}
func (f myFlag) HasChanged() bool { return false }
func (f myFlag) Name() string { return "my-flag-name" }
func (f myFlag) ValueString() string { return "my-flag-value" }
func (f myFlag) ValueType() string { return "string" }

type myFlag2 struct {}
func (f myFlag2) HasChanged() bool { return false }
func (f myFlag2) Name() string { return "my-flag-name2" }
func (f myFlag2) ValueString() string { return "my-flag-value2" }
func (f myFlag2) ValueType() string { return "string" }

type myFlagSet struct {
	flag myFlag
	flag2 myFlag2
}

func (f myFlagSet) VisitAll(fn func(value viper.FlagValue)) {
	fn(f.flag)
	fn(f.flag2)
}

func main() {

	//Bind 3rd Party Flag
	viper.BindFlagValue("my-flag-name", myFlag{})

	i := viper.GetString("my-flag-name") // retrieve value from viper

	fmt.Println(i)


	//FlagValueSet
	fSet := myFlagSet{
		flag: myFlag{},
		flag2: myFlag2{},
	}
	viper.BindFlagValues(fSet)

	j := viper.Get("my-flag-name2")

	fmt.Println(j)
}
