## ENCODING LIBRARY FOR ECLA

# Candidate functions :

| Name | Prototype | Description |
| --- | --- | --- |
| `AsciiToString` | `func AsciiToString(intArr []int) string` | Converts an array of ASCII values to a string. |
| `StringToAscii` | `func StringToAscii(str string) []int` | Converts a string to an array of ASCII values. |
| `EncodeBase64` | `func EncodeBase64(data []byte) string` | Encodes data using the base64 algorithm. |
| `DecodeBase64` | `func DecodeBase64(s string) ([]byte, error)` | Decodes data using the base64 algorithm. |
| `EncodeHex` | `func EncodeHex(data []byte) string` | Encodes data using the hexadecimal encoding algorithm. |
| `DecodeHex` | `func DecodeHex(s string) ([]byte, error)` | Decodes data using the hexadecimal encoding algorithm. |
| `EncodeJSON` | `func EncodeJSON(v interface{}) ([]byte, error)` | Encodes data using the JSON encoding algorithm. |
| `DecodeJSON` | `func DecodeJSON(data []byte, v interface{}) error` | Decodes data using the JSON encoding algorithm. |
| `EncodeGob` | `func EncodeGob(v interface{}) ([]byte, error)` | Encodes data using the Gob encoding algorithm. |
| `DecodeGob` | `func DecodeGob(data []byte, v interface{}) error` | Decodes data using the Gob encoding algorithm. |

