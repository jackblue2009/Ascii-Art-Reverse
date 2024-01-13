package pkg

/*
Identify the lines after for each character given in the input
*/
func GetAscii(ch rune, i int, StList []string) string {
	n := int(ch) - 32
	return StList[n*9+i]
}
