/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package security

import (
	"github.com/spf13/cast"
	"strconv"
	"strings"
)

func DoEncode(raw interface{}) string {
	var val string
	switch raw := raw.(type) {
	case []string:
		val = strings.Join(raw, ",")
	case []int:
		var ss []string
		for _, r := range raw {
			ss = append(ss, strconv.Itoa(r))
		}
		val = strings.Join(ss, ",")
	default:
		val = cast.ToString(raw)
	}

	// Add Encode Algorithm here

	return val
}

func DoDecode(raw []byte) (interface{}, error) {
	s := string(raw)

	// Add Decode Algorithm here

	if strings.Contains(s, ",") {
		return strings.Split(s, ","), nil
	}

	return s, nil
}