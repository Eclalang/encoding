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

var (
	ErrInvalidBase64String = errors.New("invalid base64 string")
	ErrInvalidHexString    = errors.New("invalid hex string")
)

// EncodeBase64 encodes data using the base64 algorithm.
func EncodeBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// DecodeBase64 decodes data using the base64 algorithm.
func DecodeBase64(s string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, ErrInvalidBase64String
	}
	return data, nil
}

// EncodeHex encodes data using the hexadecimal encoding algorithm.
func EncodeHex(data []byte) string {
	return hex.EncodeToString(data)
}

// DecodeHex decodes data using the hexadecimal encoding algorithm.
func DecodeHex(s string) ([]byte, error) {
	data, err := hex.DecodeString(s)
	if err != nil {
		return nil, ErrInvalidHexString
	}
	return data, nil
}

// EncodeJSON encodes data using the JSON encoding algorithm.
func EncodeJSON(v interface{}) ([]byte, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// DecodeJSON decodes data using the JSON encoding algorithm.
func DecodeJSON(data []byte, v interface{}) error {
	err := json.Unmarshal(data, v)
	if err != nil {
		return err
	}
	return nil
}

// EncodeGob encodes data using the Gob encoding algorithm.
func EncodeGob(v interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := gob.NewEncoder(buf).Encode(v)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// DecodeGob decodes data using the Gob encoding algorithm.
func DecodeGob(data []byte, v interface{}) error {
	err := gob.NewDecoder(bytes.NewReader(data)).Decode(v)
	if err != nil {
		return err
	}
	return nil
}
