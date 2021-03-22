package utils

import (
	"bytes"
	"io/ioutil"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func GetExchangeCode(code string) string {
	if !('0' <= code[0] && code[0] <= '9') {
		return code
	}
	if code[0] == '6' {
		return "sh" + code
	}
	return "sz" + code
}

func UTF8ToGBK(src []byte) ([]byte, error) {
	dst, err := ioutil.ReadAll(
		transform.NewReader(bytes.NewReader(src), simplifiedchinese.GBK.NewDecoder()),
	)
	return dst, err
}
