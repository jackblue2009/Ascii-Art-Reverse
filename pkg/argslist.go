package pkg

import (
	"errors"
	"strings"
)

/*
Segiment the given string into a parts (colors, text to be colored, text and possible errors)
*/

func ArgsList(s []string) (align, outputfile, reverse string, color, strcol []string, str, banner string, err error) {
	// check if the last argument is banner name so we take the first
	outputfile = ""
	align = ""
	reverse = ""
	banner = "standard"
	if len(s) == 0 {
		return "", "", "", nil, nil, "", "", errors.New("Error")
	}
	if s[len(s)-1] == "standard" || s[len(s)-1] == "shadow" || s[len(s)-1] == "thinkertoy" || s[len(s)-1] == "3d" {
		banner = s[len(s)-1]
		s = s[:len(s)-1]
	}
	for i, r := range s {
		if strings.HasPrefix(r, "--align") { // any index contain `--align` flag
			if len(r) > 8 {
				align = strings.ToLower(r[8:])
			} else {
				return "", "", "", nil, nil, "", "", errors.New("Error")
			}
		} else if strings.HasPrefix(r, "--output") { // any index contain `--output` flag
			if len(r) > 9 {
				outputfile = r[9:]
			} else {
				return "", "", "", nil, nil, "", "", errors.New("Error")
			}
		} else if strings.HasPrefix(r, "--color") { // any index contain `--color` flag will be taking
			colorname, err := ColorName(r)
			if err != nil {
				return "", "", "", nil, nil, "", "", errors.New("Error")
			}
			color = append(color, colorname)
			if i < len(s)-1 && strings.Contains(s[i+1], "--color") {
				strcol = append(strcol, "")
			}
		} else if strings.HasPrefix(r, "--reverse") {
			if len(r) > 10 {
				reverse = r[10:]
			} else {
				return "", "", "", nil, nil, "", "", errors.New("ErrorReverse")
			}
		} else if strings.HasPrefix(r, "--") {
			return "", "", "", nil, nil, "", "", errors.New("Error")
		} else if i < len(s)-1 { // any index before last index and not contain `--color` flag will be taking as text to be colored
			if len(strcol) >= len(color) {
				return "", "", "", nil, nil, "", "", errors.New("Error")
			} else {
				strcol = append(strcol, r)
			}
		} else if i == len(s)-1 {
			str = r
		}
	}
	if str == "" && reverse == "" {
		return "", "", "", nil, nil, "", "", errors.New("Error")
	}
	return align, outputfile, reverse, color, strcol, str, banner, nil
}
