package images

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
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
			m.(*image.RGBA).Set(x, y, color.RGBA{byte(r >> 15 * 255), byte(g >> 15 * 255), byte(b >> 15 * 255), byte(a)})
		}
	}

	outputFile, err := os.Create(os.Args[3])
	if err != nil {
		panic(err.Error())
	}
	encoder := png.Encoder{CompressionLevel: png.BestCompression}
	err = encoder.Encode(outputFile, m)
	if err != nil {
		panic(err.Error())
	}
}
