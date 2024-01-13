package main

import (
	"ascii-art-reverse/pkg"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	// check the terminal size
	terminalWidth := pkg.GetTerminalWidth()
	var reversefilelist []string
	align, outputfile, reverse, color, strToBeColored, str, banner, err := pkg.ArgsList(os.Args[1:])
	if err != nil { // check if the args match the format or no
		if err == errors.New("Error") {
			pkg.PrintUsageMsg()
		} else {
			pkg.PrintUsageMsgReverse()
		}
		return
	}
	if reverse != "" {
		if align != "" || color != nil || outputfile != "" || str != "" {
			pkg.PrintUsageMsgReverse()
			pkg.PrintError("Error: you can't use reverse flag with other flags!")
			return
		}
		var reversefile string
		reversefile, banner = pkg.ReverseFile(reverse)
		reversefilelist = strings.Split(reversefile, "\n")
	}
	// check the output file givin
	if outputfile != "" {
		if align != "" || color != nil || reverse != "" {
			pkg.PrintError("Error: can't use output flag with other flags!")
			return
		}
		if len(outputfile) > 4 {
			if !strings.HasSuffix(outputfile, ".txt") {
				pkg.PrintError("Error: invalid output file format")
				return
			}
		} else {
			pkg.PrintError("Error: invalid output file format")
			return
		}
	}
	// check the align givin
	if align == "" {
		align = "left"
	}
	if align != "left" && align != "right" && align != "center" && align != "justify" {
		pkg.PrintError("Error: invalid align")
		return
	}
	// check if no color givin
	if color == nil {
		color = []string{"default"}
	}
	// Checking the givin color
	pkg.CheckColor(color)
	// Checking the givin text is printable or not
	pkg.CheckString(str)
	strToPrint := strings.Split(str, "\\n")
	// if no identifing for the text to be colored that mean all the text to be colored
	if strToBeColored == nil {
		strToBeColored = strToPrint
		primeColor := color[0]
		for i := 0; i < len(strToBeColored)-1; i++ {
			color = append(color, primeColor)
		}
	}
	// Store string value from specified txt file in lines variable
	lines := strings.Split(pkg.GetBannerFile("banners/"+banner+".txt"), "\n")
	// Printing ascii art word
	var space []int
	// process for the text here before printing it
	for _, s := range strToPrint {
		// check the lenghtof each santance
		sraw := pkg.LenOfStr(s, lines)
		if sraw > terminalWidth {
			pkg.PrintUsageMsg()
			pkg.PrintError("The input text can't be aligned because it doesn't fit the terminal size. Please try again.")
			return
		}
		// check the align
		switch align {
		case "right":
			space = append(space, (terminalWidth - sraw))
		case "center":
			space = append(space, (terminalWidth-sraw)/2)
		case "justify":
			count := strings.Count(s, " ")
			if count == 0 {
				space = append(space, 0)
			} else {
				space = append(space, (terminalWidth-sraw)/count)
			}
		default:
			space = append(space, 0)
		}
	}
	var result string
	if outputfile != "" {
		result = pkg.StoreAsciiOutput(strToPrint, lines)
	} else if reverse != "" {
		result = pkg.ReverseAsciiArt(reversefilelist, lines)
	} else {
		result = pkg.StoreAsciiArt(strToPrint, strToBeColored, color, lines, space, terminalWidth, align)
	}
	if outputfile == "" { // if no output file
		fmt.Println(result)
	} else { // if there is output file
		err := os.WriteFile(outputfile, []byte(result), 0666)
		if err != nil {
			fmt.Println(err)
		}
	}
}
