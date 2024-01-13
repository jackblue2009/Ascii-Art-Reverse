package pkg

import (
	"errors"
	"strings"
)

/*
Check the Validate of the color given & return the needed name
*/

func ColorName(color string) (string, error) {
	if len(color) < 8 { // if the length is less than 8, return error
		return "", errors.New("error")
	}
	if color[:8] != "--color=" { // if the string first 8 chars are not "--color=" then return error
		return "", errors.New("error")
	}
	col := strings.ToLower(color[8:])
	hex := ""
	colname := ""
	if strings.Contains(col, "hsl") { // if the color is hsl
		if strings.Count(col, ", ") != 2 {
			return "", errors.New("error")
		}
		h, s, l := StringsToFloat64(col)
		r, b, g := HSLToRGB(h, s, l)
		hex = RGBToHex(int(r), int(b), int(g))
		colname = HexToColor(hex)
		return colname, nil
	} else if strings.Contains(col, "rgb") { // if the color is rgb
		if strings.Count(col, ", ") != 2 {
			return "", errors.New("error")
		}
		r, b, g := StringsToInt(col)
		hex = RGBToHex(int(r), int(b), int(g))
		colname = HexToColor(hex)
		return colname, nil
	} else if strings.Contains(col, "#") { // if the color is hex
		colname = HexToColor(strings.ToUpper(col))
		return colname, nil
	}
	return col, nil
}
