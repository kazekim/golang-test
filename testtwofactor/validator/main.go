/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package main

import (
	"bufio"
	"fmt"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"os"
)

func main() {

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "kazekim",
		AccountName: "tester@kazekim.com",
		Secret: []byte("YESTHISSHOULDBEVSLIDJL"),
		Algorithm: otp.AlgorithmSHA512,
		Digits: 8,
	})
	if err != nil {
		panic(err)
	}

	passCode := promptForPasscode()
	valid := totp.Validate(passCode, key.Secret())
	if valid {
		println("Valid passcode!")
		os.Exit(0)
	} else {
		println("Invalid passocde!")
		os.Exit(1)
	}
}

func promptForPasscode() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Passcode: ")
	text, _ := reader.ReadString('\n')
	return text
}