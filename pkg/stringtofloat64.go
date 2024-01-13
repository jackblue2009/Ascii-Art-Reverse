package pkg

import (
	"strings"
	"strconv"
)

// Conver the String from the argement to use it for HSL
func StringsToFloat64(str string) (a, b, c float64) {
	numList := strings.Split(str[4:len(str)-1], ", ")
	var list [3]float64
	for i, v := range numList {
		s := strings.Trim(v, "%")
		num, _ := strconv.Atoi(s)
		if i != 0 {
			var numf float64 = float64(num) / 100.0
			list[i] = numf
		} else {
			list[i] = float64(num)
		}
	}
	return list[0], list[1], list[2]
}
