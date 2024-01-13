package pkg

import (
	"fmt"
	"os"
	"regexp"
)

// get the name of file and identify the banner
func ReverseFile(reverse string) (reversefile, banner string) {
	file, err := os.ReadFile(reverse)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	reversefile = string(file)
	// regexp to check the input
	var shadowRegx = regexp.MustCompile(`^[_|\s]+$`)
	var thinkertoyRegx = regexp.MustCompile(`^['-/O\\_o|\s]+$`)
	// check which one match with the input
	if shadowRegx.Match(file) {
		banner = "shadow"
	} else if thinkertoyRegx.Match(file) {
		banner = "thinkertoy"
	} else {
		banner = "standard"
	}
	return reversefile, banner
}
