package encoding

import (
	"testing"
)

func TestAsciiToString(t *testing.T) {
	expected := "Hello, World!"
	ascii := []int{72, 101, 108, 108, 111, 44, 32, 87, 111, 114, 108, 100, 33}
	result := AsciiToString(ascii)
	if expected != result {
		t.Errorf("AsciiToString(%v) = %v, want %v", ascii, result, expected)
	}
}

func TestDecodeBase64(t *testing.T) {
	expected := []int{72, 101, 108, 108, 111, 44, 32, 87, 111, 114, 108, 100, 33}
	str := "SGVsbG8sIFdvcmxkIQ=="
	actual := DecodeBase64(str)
	if len(expected) != len(actual) {
		t.Errorf("DecodeBase64(%v) = %v, want %v", str, actual, expected)
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Errorf("DecodeBase64(%v) = %v, want %v", str, actual, expected)
		}
	}

	str = `ijj"kkkk`
	actual = DecodeBase64(str)
	if actual != nil {
		t.Errorf("DecodeBase64(%v) = %v, want %v", str, actual, nil)
	}
}

func TestDecodeGob(t *testing.T) {
	expected := "Hello, World!"
	array := []int{16, 12, 0, 13, 72, 101, 108, 108, 111, 44, 32, 87, 111, 114, 108, 100, 33}
	actual := DecodeGob(array)
	if expected != actual {
		t.Errorf("DecodeGob(%v) = %v, want %v", array, actual, expected)
	}

	array = []int{-1}
	actual = DecodeGob(array)
	if actual != "" {
		t.Errorf("DecodeGob(%v) = %v, want %v", array, actual, "")
	}
}

func TestDecodeHex(t *testing.T) {
	expected := []int{72, 101, 108, 108, 111, 44, 32, 87, 111, 114, 108, 100, 33}
	str := "48656c6c6f2c20576f726c6421"
	actual := DecodeHex(str)
	if len(expected) != len(actual) {
		t.Errorf("DecodeHex(%v) = %v, want %v", str, actual, expected)
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Errorf("DecodeHex(%v) = %v, want %v", str, actual, expected)
		}
	}

	str = "abfK0"
	actual = DecodeHex(str)
	if actual != nil {
		t.Errorf("DecodeHex(%v) = %v, want %v", str, actual, nil)
	}
}

func TestEncodeBase64(t *testing.T) {
	expected := "SGVsbG8sIFdvcmxkIQ=="
	array := []int{72, 101, 108, 108, 111, 44, 32, 87, 111, 114, 108, 100, 33}
	actual := EncodeBase64(array)
	if expected != actual {
		t.Errorf("EncodeBase64(%v) = %v, want %v", array, actual, expected)
	}
}

func TestEncodeGob(t *testing.T) {
	expected := []int{16, 12, 0, 13, 72, 101, 108, 108, 111, 44, 32, 87, 111, 114, 108, 100, 33}
	str := "Hello, World!"
	actual := EncodeGob(str)
	if len(expected) != len(actual) {
		t.Errorf("EncodeGob(%v) = %v, want %v", str, actual, expected)
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Errorf("EncodeGob(%v) = %v, want %v", str, actual, expected)
		}
	}
}

func TestEncodeHex(t *testing.T) {
	expected := "48656c6c6f2c20576f726c6421"
	array := []int{72, 101, 108, 108, 111, 44, 32, 87, 111, 114, 108, 100, 33}
	actual := EncodeHex(array)
	if expected != actual {
		t.Errorf("EncodeHex(%v) = %v, want %v", array, actual, expected)
	}
}

func TestStringToAscii(t *testing.T) {
	expected := []int{72, 101, 108, 108, 111, 44, 32, 87, 111, 114, 108, 100, 33}
	str := "Hello, World!"
	result := StringToAscii(str)
	if len(expected) != len(result) {
		t.Errorf("StringToAscii(%v) = %v, want %v", str, result, expected)
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != result[i] {
			t.Errorf("StringToAscii(%v) = %v, want %v", str, result, expected)
		}
	}
}
