package encoding

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"encoding/hex"
	"encoding/json"
	"errors"
)

// AsciiToString converts an array of ASCII values to a string
func AsciiToString(intArr []int) string {
	var buf bytes.Buffer
	for _, v := range intArr {
		buf.WriteByte(byte(v))
	}
	return buf.String()
}

// StringToAscii converts a string to an array of ASCII values
func StringToAscii(str string) []int {
	intArr := make([]int, len(str))
	for i, v := range str {
		intArr[i] = int(v)
	}
	return intArr
}

