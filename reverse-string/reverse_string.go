package reverse

func Reverse(input string) string {
	in := []rune(input)
	runeLen := len(in)
	out := make([]rune, runeLen)
	for i := 0; i < runeLen; i++ {

		out[i] = in[runeLen-1-i]
	}
	return string(out)

}
