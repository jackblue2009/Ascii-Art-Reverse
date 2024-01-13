package pkg

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

/* 
This function executes the "tput cols" command in the command line,
then converts the output to a string and returns it as an integer.
 */
func GetTerminalWidth() int {
	cmd := exec.Command("tput", "cols")
	cmd.Stdin = os.Stdin

	out, err := cmd.Output() 
	if err != nil {
		log.Fatal("Unable to get Cli width. Please try again.", err)
	}

	outputStr := string(out)
	parts := strings.TrimSpace(outputStr)

	width, _ := strconv.Atoi(parts)

	return width
}
