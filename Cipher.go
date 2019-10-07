package main

import (
	"fmt"
	"os"
)

// Test Run unit test on cipher algorithm
func Test() {
	fmt.Println("Cipher.go")
	if len(os.Args) != 2 {
		fmt.Println("Must supply hexadecimal password")
		return
	}

	for index, char := range os.Args[1] {
		fmt.Printf("%2d: %c: %2d\n", index, char, hexVal(char))
	}
}

func hexVal(c rune) int {
	switch c {
	case '1':
		return 1
	case '2':
		return 2
	case '3':
		return 3
	case '4':
		return 4
	case '5':
		return 5
	case '6':
		return 6
	case '7':
		return 7
	case '8':
		return 8
	case '9':
		return 9
	case 'A', 'a':
		return 10
	case 'B', 'b':
		return 11
	case 'C', 'c':
		return 12
	case 'D', 'd':
		return 13
	case 'E', 'e':
		return 14
	case 'F', 'f':
		return 15
	default:
		return 0
	}
}
