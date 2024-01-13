package pkg

import (
	"strings"
	"strconv"
)

// Conver the String from the argement to use it for RGB
func StringsToInt(str string) (a, b, c int) {
	numList := strings.Split(str[4:len(str)-1], ", ")
	var list [3]int
	for i, v := range numList {
		num, _ := strconv.Atoi(v)
		list[i] = num
	}
	return list[0], list[1], list[2]
}
