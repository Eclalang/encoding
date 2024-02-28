package encoding

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"encoding/hex"
)

// AsciiToString converts an array of ASCII values to a string
func AsciiToString(intArr []int) string {
	var buf bytes.Buffer
	for _, v := range intArr {
		buf.WriteByte(byte(v))
	}
	return buf.String()
}

// DecodeBase64 decodes a base64 string to an array of ASCII values
func DecodeBase64(str string) []int {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return nil
	}
	intArr := make([]int, len(data))
	for i, v := range data {
		intArr[i] = int(v)
	}
	return intArr
}

// DecodeGob decodes a gob int array to a string
func DecodeGob(intArr []int) string {
	var buf bytes.Buffer
	for _, v := range intArr {
		buf.WriteByte(byte(v))
	}
	dec := gob.NewDecoder(&buf)
	var str string
	err := dec.Decode(&str)
	if err != nil {
		return ""
	}
	return str
}

// DecodeHex decodes a hex string to an array of ASCII values
func DecodeHex(str string) []int {
	data, err := hex.DecodeString(str)
	if err != nil {
		return nil
	}
	intArr := make([]int, len(data))
	for i, v := range data {
		intArr[i] = int(v)
	}
	return intArr
}

// EncodeBase64 encodes an array of ASCII values to a base64 string
func EncodeBase64(intArr []int) string {
	data := make([]byte, len(intArr))
	for i, v := range intArr {
		data[i] = byte(v)
	}
	return base64.StdEncoding.EncodeToString(data)
}

// EncodeGob encodes a string to a gob int array
func EncodeGob(str string) []int {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	_ = enc.Encode(str)
	tempo := buf.Bytes()
	intArr := make([]int, len(tempo))
	for i, v := range tempo {
		intArr[i] = int(v)
	}
	return intArr
}

// EncodeHex encodes an array of ASCII values to a hex string
func EncodeHex(intArr []int) string {
	data := make([]byte, len(intArr))
	for i, v := range intArr {
		data[i] = byte(v)
	}
	return hex.EncodeToString(data)
}

// StringToAscii converts a string to an array of ASCII values
func StringToAscii(str string) []int {
	intArr := make([]int, len(str))
	for i, v := range str {
		intArr[i] = int(v)
	}
	return intArr
}
