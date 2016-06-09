package tool

import (
	"strconv"
)

/*
*将字符串变为json字符串
 */
func StringtoJson(str string) string {
	rs := []rune(str)
	json := ""
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			json += string(r)

		} else {
			json += "\\u" + strconv.FormatInt(int64(rint), 16)
		}
	}
	return json
}
