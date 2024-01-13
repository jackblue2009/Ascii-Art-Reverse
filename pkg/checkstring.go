package pkg

import "log"

// check the input string
func CheckString(s string) bool {
	for _, r := range s {
		if r < 32 || r > 126 {
			log.Fatal("Error: invalid CHARACTER: " + string(r))
		}
	}
	return true
}
