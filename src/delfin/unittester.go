package main

import (
	"dcipher"
	"fmt"
)

// entry point for unit tester
func main() {
	fmt.Println("Unit tester")
	c := dcipher.NewCipher([16]byte{255})
	c.InternalInfo()
	for i := 0; i < 8; i++ {
		fmt.Println(c.GetByte())
	}
}
