/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"image/png"
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
	// Convert TOTP key into a PNG
	var buf bytes.Buffer
	img, err := key.Image(200, 200)
	if err != nil {
		panic(err)
	}
	err = png.Encode(&buf, img)
	if err != nil {
		panic(err)
	}

	encoded := base64.StdEncoding.EncodeToString(buf.Bytes())

	fmt.Println(encoded)

	passcode := "880786"
	valid := totp.Validate(passcode, key.Secret())
	if valid {
		println("Valid passcode!")
		os.Exit(0)
	} else {
		println("Invalid passocde!")
		os.Exit(1)
	}
}
