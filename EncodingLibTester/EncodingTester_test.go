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
	encoded := encoding.EncodeToString([]byte(input))
	decoded, err := encoding.DecodeString(encoded)
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
