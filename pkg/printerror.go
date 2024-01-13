package pkg

import (
	"fmt"
	"log"
)

//
// This function is to be called when error is found or when command line arguments number
// is not the correct number of arg required by the program
//

func PrintError(s string) {
	log.Fatal(s)
}

func PrintUsageMsg() {
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
	fmt.Println("EX: go run . --reverse=<fileName>")
}

func PrintUsageMsgReverse() {
	fmt.Println("Usage: go run . [OPTION]")
	fmt.Println("EX: go run . --reverse=<fileName>")
}
