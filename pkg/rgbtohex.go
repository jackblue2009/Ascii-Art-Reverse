package pkg

import (
	"strings"
	"fmt"
	"log"
)
// Convert RGB to HEXS
func RGBToHex(r, g, b int) string {
	if r > 255 || g > 255 || b > 255 {
		log.Fatal("Error: invalid number for RGB values")
	}
	return strings.ToUpper(fmt.Sprintf("#%02x%02x%02x", r, g, b))
}
