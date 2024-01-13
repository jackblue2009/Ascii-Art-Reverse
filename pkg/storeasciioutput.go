package pkg

/*
PrintAsciiArtWord function takes two slices of string as parameters, iterates through s and uses lines
in GetAscii() to properly print ascii art characters
*/

func StoreAsciiOutput(str, lines []string) string {
	var result string
	for _, v := range str { //MOVE THROUGH ARRAY ELEMENT
		if v == "" { // skip empty strings with new line
			result += "\n"
		} else {
			for i := 1; i < 9; i++ { //Print line by line
				for _, r := range v { // print line of character
					result += (GetAscii(r, i, lines))
				}
				result += "\n"
			}
		}
	}
	return result
}
