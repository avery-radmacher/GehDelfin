package main

import (
	dcipher "cipher"
	"fmt"
	"os"
)

// entry point for unit tester
func main() {
	fmt.Println("Unit tester")
	password := "delfin"
	if len(os.Args) > 1 {
		password = os.Args[1]
	}
	c := dcipher.NewCipherv2(password)
	m := make(map[byte]uint16)
	for i := 0; i < 256*1024; i++ {
		val := c.GetByte()
		m[val]++
	}
	fmt.Println()
	for i := byte(0); i < 16; i++ {
		for j := byte(0); j < 16; j++ {
			fmt.Printf("%2X: %4d | ", i*16+j, m[byte(i*16+j)])
		}
		fmt.Println()
	}
	fmt.Println(c)
}
