package reverse_string

func ReverseString(input string) (output string) {
	rr := []rune(input)
	for i := len(rr) - 1; i >= 0; i-- {
		output += string(rr[i])
	}
	return output
}
