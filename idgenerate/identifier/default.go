/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package identifier

import (
	b64 "encoding/base64"
	"encoding/hex"
)

var DEFAULTKEY = DEFAULTPREFIX+"-mYcgJKQhRpfFdYof"
var DEFAULTPREFIX = "xent"

type Identifier struct {
	value string
	prefix *string
	secretKey string
}

func NewIdentifier() (*Identifier, error) {

	v, err := generateRandomString(16)
	if err != nil {
		return nil, err
	}

	return &Identifier{
		value: v,
		secretKey: DEFAULTKEY,
	}, nil
}

func (v *Identifier) SetPrefix(prefix string) {
	v.prefix = &prefix
}

func (v *Identifier) ID() string {
	return v.value
}

func (v *Identifier) GenerateNewKey() string {
	key := b64.StdEncoding.EncodeToString([]byte(v.value))
	return key
}

func (v *Identifier) GenerateNewSecret() (string, error) {

	s, err := cipherEncrypt([]byte(v.value), v.secretKey)
	if err != nil {
		return "", err
	}

	secret := hex.EncodeToString(s)

	return string(secret), nil
}

func IdentifierFromKey(key string) (*Identifier, error) {

	v, err := b64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, err
	}

	return &Identifier{
		value: string(v),
	}, nil
}

func IdentifierFromSecret(secret string) (*Identifier, error) {

	return IdentifierFromSecretAndKey(secret, DEFAULTKEY)
}

func IdentifierFromSecretAndKey(secret string, key string) (*Identifier, error) {

	s, err := hex.DecodeString(secret)
	if err != nil {
		return nil, err
	}

	v, err := cipherDecrypt(s, key)

	if err != nil {
		return nil, err
	}

	return &Identifier{
		value: string(v),
	}, nil
}

