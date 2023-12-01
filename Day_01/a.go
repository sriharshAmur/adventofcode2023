package main

import (
	"fmt"
	"strconv"
)

func A(lines []string) {
	sum := 0
	for _, line := range lines {
		first, last := "", ""
		for _, letter := range line {
			if letter < 65 || letter > 122 {
				if first == "" {
					first = string(letter)
				} else {
					last = string(letter)
				}
			}
		}
		if last == "" {
			last = first
		}
		digit, err := strconv.Atoi(first + "" + last)

		if err != nil {
			fmt.Printf("Faild to get number for %v. The error is %s", first+""+last, err)
			return
		}

		sum += digit

	}
	fmt.Println(sum)
}
