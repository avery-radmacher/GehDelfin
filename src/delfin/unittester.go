package main

import (
	"dcipher"
	"fmt"
)

// entry point for unit tester
func main() {
	fmt.Println("Unit tester")
	c := dcipher.NewCipher([16]byte{}) //{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255})
	c.InternalInfo()
	for i := 0; i < 128; i++ {
		fmt.Printf("%d", c.Tick())
	}
	fmt.Println()
	c.InternalInfo()
}
