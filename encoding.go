package encoding

import (
	"encoding/base64"
	_ "encoding/hex"
	"encoding/json"
	"errors"
)

// Encoding is an interface that defines the methods for encoding and decoding data.
type Encoding interface {
	Encode(data []byte) string
	Decode(encoded string) ([]byte, error)
}

// Base64Encoding is a struct that implements the Encoding interface using the base64 algorithm.
type Base64Encoding struct{}

func (b Base64Encoding) Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func (b Base64Encoding) Decode(encoded string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, errors.New("failed to decode data")
	}
	return data, nil
}

// HexEncoding is a struct that implements the Encoding interface using the hexadecimal encoding algorithm.
type HexEncoding struct{}

func (h HexEncoding) Encode(data interface{}) []byte {
	return []byte(data.(string))
}
func (h HexEncoding) Decode(encoded []byte) (interface{}, error) {
	return string(encoded), nil
}

// JSONEncoding is a struct that implements the Encoding interface using the JSON encoding algorithm.
type JSONEncoding struct{}

func (j JSONEncoding) Encode(data interface{}) string {
	json, _ := json.Marshal(data)
	return string(json)
}

func (j JSONEncoding) Decode(encoded string) ([]byte, error) {
	var data []byte
	err := json.Unmarshal([]byte(encoded), &data)
	if err != nil {
		return nil, errors.New("failed to decode data")
	}
	return data, nil
}

// XMLEncoding is a struct that implements the Encoding interface using the XML encoding algorithm.
type XMLEncoding struct{}

func (x XMLEncoding) Encode(data []byte) string {
	return string(data) // Assumes data is already in XML format
}

func (x XMLEncoding) Decode(encoded string) ([]byte, error) {
	return []byte(encoded), nil // Assumes encoded string is already in XML format
}

// // AsciiToString converts an array of ASCII values to a string
// func AsciiToString(intArr []int) string {
// 	var buf bytes.Buffer
// 	for _, v := range intArr {
// 		buf.WriteByte(byte(v))
// 	}
// 	return buf.String()
// }

// // StringToAscii converts a string to an array of ASCII values
// func StringToAscii(str string) []int {
// 	intArr := make([]int, len(str))
// 	for i, v := range str {
// 		intArr[i] = int(v)
// 	}
// 	return intArr
// }
