package pkg

import (
	"strings"
)

/*
PrintAsciiArtWord function takes two slices of string as parameters, iterates through s and uses lines
in GetAscii() to properly print ascii art characters
*/

func StoreAsciiArt(str, strcol, color, lines []string, space []int, terminalWidth int, align string) string {
	var result string
	colstring := ColorStringMap(strcol, color)
	// declair the first and last index of the slice with -1 to unable to use it
	first := -1
	last := -1
	// color to be default if no color usable
	colorname := "default"
	c := 0
	count := 0
	vlen := 0
	for n, v := range str { //MOVE THROUGH ARRAY ELEMENT
		if align == "justify" {
			count = strings.Count(v, " ")
			vlen = LenOfStr(v, lines)
		}
		if v == "" { // skip empty strings with new line
			result += "\n"
		} else {
			for i := 1; i < 9; i++ { //Print line by line
				for k, r := range v { // print line of character
					if k >= last {
						first = -1
						last = -1
						colorname = "default"
					}
					if first == -1 {
						for s, col := range colstring { // loop to indentify the index start & end & color
							if len(v[k:]) >= len(s) {
								if v[k:k+len(s)] == s && s != "" {
									first = k
									last = k + len(s)
									colorname = col
								}
							}
						}
					}
					//check the align type if not justify the spaces will be printed pefor the ascii art
					if align != "justify" && k == 0 {
						result += strings.Repeat(" ", space[n])
					}
					// if align is justify the spaces will be printed with " "
					if align == "justify" && r == ' ' {
						c++
						if count == c {
							lastspace := terminalWidth - ((c - 1) * space[n]) - vlen
							result += strings.Repeat(" ", lastspace)
						} else {
							result += strings.Repeat(" ", space[n])
						}
					}
					if k >= first && k < last { // print the character with the color
						result += (GetColor(colorname) + GetAscii(r, i, lines) + GetColor("default"))
					} else { // print the character without the color
						result += (GetAscii(r, i, lines))
					}
				}
				// return the value to normal state
				c = 0
				first = -1
				last = -1
				colorname = "default"
				result += "\n"
			}
		}
	}
	return result
}
