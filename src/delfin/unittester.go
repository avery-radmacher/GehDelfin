package main

import (
	dcipher "cipher"
	"fmt"
	"header"
	"images"
	"os"
)

// entry point for unit tester
func main() {
	fmt.Println("Unit tester")
	if len(os.Args) == 1 {
		fmt.Println("no test specified")
		return
	}
	switch os.Args[1] {
	case "cipher":
		password := "delfin"
		if len(os.Args) > 2 {
			password = os.Args[2]
		}
		c := dcipher.NewCipherv2(password)
		m := make(map[byte]uint16)
		b := make([]byte, 256*1024)
		c.Crypt(b)
		for _, val := range b {
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
	case "image":
		images.Test()
	case "header":
		header.Test()
	}
}
