package main

import (
	"color-preview/utils"
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

func main() {
	color := "FFFFFF"
	if len(os.Args) < 2 {
		// TODO: random
	} else {
		color = strings.TrimPrefix(os.Args[1], "#")
	}

	columns, _, err := term.GetSize(0)
	if err != nil {
		columns = 120
	}

	utils.ClearScreen()
	line := strings.Repeat("█", columns)
	for i := 0; i < 10; i++ {
		fmt.Printf("%s%s\n", utils.AsHexRGB(color), line)
	}
	fmt.Printf("%s#%s\n", utils.RESET, color)
}
