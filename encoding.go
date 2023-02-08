package encoding

// AsciiToString converts an array of ASCII values to a string
func AsciiToString(intArr []int) string {
	str := ""
	for i := 0; i < len(intArr); i++ {
		if intArr[i] > 127 {
			intArr[i] = 127
		}
		str += string(intArr[i])
	}
	return str
}

// StringToAscii converts a string to an array of ASCII values
func StringToAscii(str string) []int {
	intArr := make([]int, len(str))
	for i := 0; i < len(str); i++ {
		intArr[i] = int(str[i])
	}
	return intArr
}
