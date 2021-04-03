package utils

import (
	"fmt"
	"strconv"
)

const RESET = "\033[0m"

func ClearLine() {
	fmt.Print("\033[K")
}

func ClearAfterCursor() {
	fmt.Print("\033[J")
}

func MoveCursorTo(line, column int) {
	fmt.Printf("\033[%d;%df", line, column)
}

func ClearScreen() {
	MoveCursorTo(1, 1)
	ClearAfterCursor()
}

func AsHexRGB(hex string) string {
	r, _ := strconv.ParseInt(hex[0:2], 16, 64)
	g, _ := strconv.ParseInt(hex[2:4], 16, 64)
	b, _ := strconv.ParseInt(hex[4:6], 16, 64)
	return AsRGB(r, g, b)
}

func AsRGB(r, g, b int64) string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}
