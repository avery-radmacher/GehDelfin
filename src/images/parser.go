package images

import (
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"os"
)

// Test runs a test of the image code so far
func Test() {
	fmt.Println("images test")
	reader, err := os.Open(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	bounds := m.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := m.At(x, y).RGBA()
			m.(*image.RGBA).Set(x, y, color.RGBA{byte(r >> 15 << 7), byte(g >> 15 << 7), byte(b >> 15 << 7), byte(a >> 15 << 7)})
			color := r>>13&4 | g>>14&2 | b>>15
			if color == 4 {
				fmt.Printf("%04d.%04d: \n", x, y)
			}
		}
	}
}
