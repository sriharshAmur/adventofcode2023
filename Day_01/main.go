package main

import (
	"fmt"

	"github.com/sriharshAmur/adventofcode2023/internal"
)

func main() {
	test := false
	test = true
	lines, err := internal.ReadFileContents(test)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	A(lines)
}
