## ENCODING LIBRARY FOR ECLA

# Candidate functions :

|   Func Name   |               Prototype               |                Description                 | Comments |
|:-------------:|:-------------------------------------:|:------------------------------------------:|:--------:|
| AsciiToString | AsciiToString(intArr []int) string {} | Converts array of ASCII values to a string |   N/A    |
| DecodeBase64  |   DecodeBase64(str string) []int {}   |  Decodes a base64 string to ASCII values   |   N/A    |
|   DecodeGob   |    DecodeGob(intArr []int) string     |    Decodes a gob int array to a string     |   N/A    |
|   DecodeHex   |    DecodeHex(str string) []int {}     |    Decodes a hex string to ASCII values    |   N/A    |
| EncodeBase64  |   EncodeBase64(intArr []int) string   |  Encodes ASCII values to a base64 string   |   N/A    |
|   EncodeGob   |      EncodeGob(str string) []int      |     Encodes string to a gob int array      |   N/A    |
|   EncodeHex   |    EncodeHex(intArr []int) string     |    Encodes ASCII values to a hex string    |   N/A    |
| StringToAscii |  StringToAscii(str string) []int {}   |      Converts string to ASCII values       |   N/A    |