package main

import (
	"bytes"
	"encoding/gob"
	"reflect"
	"testing"

	encoding "github.com/Eclalang/encoding"
)

func TestEncodeDecodeString(t *testing.T) {
	input := "hello world"
	encoded := encoding.EncodeBase64([]byte(input))
	decoded, err := encoding.DecodeBase64(encoded)
	if err != nil {
		t.Errorf("DecodeString error: %v", err)
	}
	if string(decoded) != input {
		t.Errorf("DecodeString result = %s; want %s", decoded, input)
	}
}

func TestEncodeDecodeHex(t *testing.T) {
	input := "hello world"
	encoded := encoding.EncodeHex([]byte(input))
	decoded, err := encoding.DecodeHex(encoded)
	if err != nil {
		t.Errorf("DecodeHex error: %v", err)
	}
	if string(decoded) != input {
		t.Errorf("DecodeHex result = %s; want %s", decoded, input)
	}
}

func TestEncodeDecodeJSON(t *testing.T) {
	input := map[string]string{
		"hello": "world",
		"foo":   "bar",
	}
	encoded, err := encoding.EncodeJSON(input)
	if err != nil {
		t.Errorf("EncodeJSON error: %v", err)
	}
	var decoded map[string]string
	err = encoding.DecodeJSON(encoded, &decoded)
	if err != nil {
		t.Errorf("DecodeJSON error: %v", err)
	}
	if !bytes.Equal(encoded, []byte(`{"foo":"bar","hello":"world"}`)) {
		t.Errorf("EncodeJSON result = %s; want %s", encoded, []byte(`{"foo":"bar","hello":"world"}`))
	}
	if !reflect.DeepEqual(decoded, input) {
		t.Errorf("DecodeJSON result = %v; want %v", decoded, input)
	}
}

func TestEncodeDecodeGob(t *testing.T) {
	input := map[string]string{
		"hello": "world",
		"foo":   "bar",
	}
	encoded, err := encoding.EncodeGob(input)
	if err != nil {
		t.Errorf("EncodeGob error: %v", err)
	}
	var decoded map[string]string
	err = encoding.DecodeGob(encoded, &decoded)
	if err != nil {
		t.Errorf("DecodeGob error: %v", err)
	}
	buf := new(bytes.Buffer)
	err = gob.NewEncoder(buf).Encode(input)
	if err != nil {
		t.Errorf("gob encode error: %v", err)
	}
	if !bytes.Equal(encoded, buf.Bytes()) {
		t.Errorf("EncodeGob result = \n %v; want \n %v", encoded, buf.Bytes())
	}
	if !reflect.DeepEqual(decoded, input) {
		t.Errorf("DecodeGob result = %v; want %v", decoded, input)
	}
}

func TestAsciiToString(t *testing.T) {
	expected := "Hello, World!"
	ascii := []int{72, 101, 108, 108, 111, 44, 32, 87, 111, 114, 108, 100, 33}
	result := encoding.AsciiToString(ascii)
	if expected != result {
		t.Errorf("AsciiToString(%v) = %v, want %v", ascii, result, expected)
	}
}

func TestStringToAscii(t *testing.T) {
	expected := []int{72, 101, 108, 108, 111, 44, 32, 87, 111, 114, 108, 100, 33}
	str := "Hello, World!"
	result := encoding.StringToAscii(str)
	if len(expected) != len(result) {
		t.Errorf("StringToAscii(%v) = %v, want %v", str, result, expected)
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != result[i] {
			t.Errorf("StringToAscii(%v) = %v, want %v", str, result, expected)
		}
	}
}
