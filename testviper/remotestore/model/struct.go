/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package model

import (
	"github.com/kazekim/golang-test/testviper/remotestore/security"
	"io"
	"io/ioutil"
)

type Decode struct{
	Access Access
	Redis Redis
	Deeply Deeply
	Providers []string
	Lucky Lucky
}

type Access struct {
	Token string
}

type Redis struct {
	Addr string
	Password string
}

type Deeply struct {
	Nested Nested
}

type Nested struct {
	Config NestedConfig
}

type NestedConfig struct {
	Wow string
}

type Lucky struct {
	Numbers []int
}


func (d Decode) Decode(r io.Reader) (interface{}, error) {
	raw, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return security.DoDecode(raw)
}