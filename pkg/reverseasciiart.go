package pkg

import (
	"fmt"
	"os"
)

func ReverseAsciiArt(rlist, lines []string) string {
	var list []int
	k := 0
	count := 0
	for _, c := range rlist { // count valid line
		if len(c) > 1 {
			count++
		}
	}
	if count%8 != 0 { // /8 each line == 8 lines
		fmt.Println("Error: the file have less/more lines than usual!")
		os.Exit(1)
	}
	for k < len(rlist)-1 {
		// only need to check first 6 lines in each line
		if len(rlist[k]) == 0 { //if the line empty add "\n"
			list = append(list, -1)
			k++
		} else {
			reverselist := rlist[k : k+8]
			for i := 0; i < len(reverselist[1]); i++ {
				for j := 0; j < 95; j++ {
					for index := 0; index <= 7; index++ {
						// current line for the ascii to compare it with the same line in the reverse file
						current := lines[j*9+index+1]
						// length of the line to compare with
						endpoint := i + len(current)
						if len(current) <= len(reverselist[index][i:]) {
							if current != reverselist[index][i:endpoint] { // if not equal break the loop and check with another char
								break
							}
							if index == 7 { // if they reach to the line no. 5 and they equal add the number of the char
								if current == reverselist[index][i:i+len(current)] {
									list = append(list, j)
									i += len(current)
									j = -1
								}
							}
						} else { // this to check last char
							if current != reverselist[index][i:] {
								break
							}
							if index == 5 { // if they reach to the line no. 5 and they equal add the number of the char
								if current == reverselist[index][i:i+len(current)] {
									list = append(list, j)
									i += len(current)
									j = 0
								}
							}
						}
					}
				}
			}
			// this to go to second line
			k += 8
			if len(rlist)-k > 8 { // if there more than 8 lines available this mean there should '\n' between text
				list = append(list, -1)
			}
		}
	}
	result := ""
	for _, v := range list {
		if v == -1 { // add '\n'
			result += "\\n"
		} else { // add the char
			result += string(rune(v + 32))
		}
	}
	return result
}
