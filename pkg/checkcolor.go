package pkg

import "log"

// Check the color availble or not
func CheckColor(s []string) bool{
	for _, r := range s {
		if GetColor(r) == "" {
			log.Fatal("Error: Color selected not available!")
		}
	}
	return true
}
