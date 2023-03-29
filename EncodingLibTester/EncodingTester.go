package main

import (
	"fmt"

	encoding "github.com/Eclalang/encoding"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func main() {
	// Test Base64Encoding
	base64Encoding := encoding.Base64Encoding{}
	data := []byte("Hello, world!")
	encoded := base64Encoding.Encode(data)
	fmt.Printf("Base64 encoded string: %s\n", encoded)
	decoded, err := base64Encoding.Decode(encoded)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Base64 decoded string: %s\n", string(decoded))

	// Test HexEncoding
	hexEncoding := encoding.HexEncoding{}
	data_string := "Hello, world!"
	encoded_string := hexEncoding.Encode(data_string)
	fmt.Printf("Hex encoded string: %s\n", encoded)
	decoded_string, err := hexEncoding.Decode(encoded_string)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hex decoded string: %s\n", decoded_string)

	// Test JSONEncoding
	jsonEncoding := encoding.JSONEncoding{}
	book1 := Book{"The Hobbit", "J.R.R. Tolkien"}
	book2 := Book{"The Lord of the Rings", "J.R.R. Tolkien"}
	books := []Book{book1, book2}
	encoded = jsonEncoding.Encode(books)
	fmt.Printf("JSON encoded string: %s\n", encoded)
	decoded, err = jsonEncoding.Decode(encoded)
	if err != nil {
		panic(err)
	}
	fmt.Printf("JSON decoded string: %s\n", string(decoded))

	// Test XMLEncoding
	xmlEncoding := encoding.XMLEncoding{}
	data = []byte("<hello>world</hello>")
	encoded = xmlEncoding.Encode(data)
	fmt.Printf("XML encoded string: %s\n", encoded)
	decoded, err = xmlEncoding.Decode(encoded)
	if err != nil {
		panic(err)
	}
	fmt.Printf("XML decoded string: %s\n", string(decoded))
}

// func TestAsciiToString(t *testing.T) {
// 	expected := "Hello, World!"
// 	ascii := []int{72, 101, 108, 108, 111, 44, 32, 87, 111, 114, 108, 100, 33}
// 	result := encoding.AsciiToString(ascii)
// 	if expected != result {
// 		t.Errorf("AsciiToString(%v) = %v, want %v", ascii, result, expected)
// 	}
// }

// func TestStringToAscii(t *testing.T) {
// 	expected := []int{72, 101, 108, 108, 111, 44, 32, 87, 111, 114, 108, 100, 33}
// 	str := "Hello, World!"
// 	result := encoding.StringToAscii(str)
// 	if len(expected) != len(result) {
// 		t.Errorf("StringToAscii(%v) = %v, want %v", str, result, expected)
// 	}
// 	for i := 0; i < len(expected); i++ {
// 		if expected[i] != result[i] {
// 			t.Errorf("StringToAscii(%v) = %v, want %v", str, result, expected)
// 		}
// 	}
// }
