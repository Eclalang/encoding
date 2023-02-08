package encodinglibtest

import (
	"github.com/Eclalang/encoding"
	"testing"
)

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
