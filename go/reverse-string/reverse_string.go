package reverse

func Reverse(input string) (reversed string) {
	for _, char := range input {
		reversed = string(char) + reversed
	}
	return
}
