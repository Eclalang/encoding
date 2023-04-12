## ENCODING LIBRARY FOR ECLA

# Candidate functions :

|   Func Name   |                   Prototype                   |                      Description                       | Comments |
|:-------------:|:---------------------------------------------:|:------------------------------------------------------:|:--------:|
| AsciiToString |     AsciiToString(intArr []int) string {}     |       Converts array of ASCII values to a string       |   N/A    |
| DecodeBase64  |   DecodeBase64(str string) ([]byte, error)    |        Decodes data using the base64 algorithm.        |   N/A    |
|   DecodeGob   |  DecodeGob(data []byte, v interface{}) error  |     Decodes data using the Gob encoding algorithm.     |   N/A    |
|   DecodeHex   |     DecodeHex(str string) ([]byte, error)     | Decodes data using the hexadecimal encoding algorithm. |   N/A    |
|  DecodeJSON   | DecodeJSON(data []byte, v interface{}) error  |    Decodes data using the JSON encoding algorithm.     |   N/A    |
| EncodeBase64  |       EncodeBase64(data []byte) string        |        Encodes data using the base64 algorithm.        |   N/A    |
|   EncodeGob   |   EncodeGob(v interface{}) ([]byte, error)    |     Encodes data using the Gob encoding algorithm.     |   N/A    |
|   EncodeHex   |         EncodeHex(data []byte) string         | Encodes data using the hexadecimal encoding algorithm. |   N/A    |
|  EncodeJSON   |   EncodeJSON(v interface{}) ([]byte, error)   |    Encodes data using the JSON encoding algorithm.     |   N/A    |
| StringToAscii |      StringToAscii(str string) []int {}       |            Converts string to ASCII values             |   N/A    |