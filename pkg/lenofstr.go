package pkg

//check the lenght of printed ascii art for the text
func LenOfStr(s string, lines []string) int {
	var str string
	for _, r := range s {
		str += GetAscii(r, 1, lines)
	}
	return len(str)
}