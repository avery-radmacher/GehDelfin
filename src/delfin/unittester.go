package main

import (
	"dcipher"
	"fmt"
)

// entry point for unit tester
func main() {
	fmt.Println("Unit tester")
	c := dcipher.NewCipher([16]byte{}) //{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255})
	//fmt.Println(c)
	m := make(map[byte]uint16)
	for i := 0; i < 256*1024; i++ {
		val := c.GetByte()
		//fmt.Printf("%2X ", val)
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
