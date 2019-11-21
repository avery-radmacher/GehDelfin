package images

import (
	dcipher "cipher"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
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
	n := image.NewNRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := m.At(x, y).RGBA()
			rs, gs, bs := r>>8&3, g>>8&3, b>>8&3
			if (x+y)%16 == 0 {
				fmt.Printf("%04d.%04d: %02b%02b%02b\n", y, x, rs, gs, bs)
			}
			n.Set(x, y, color.RGBA{byte(r >> 15 * 255), byte(g >> 15 * 255), byte(b >> 15 * 255), 255})
		}
	}

	outputFile, err := os.Create(os.Args[3])
	if err != nil {
		log.Fatal(err)
	}
	encoder := png.Encoder{CompressionLevel: png.BestCompression}
	err = encoder.Encode(outputFile, n)
	if err != nil {
		log.Fatal(err)
	}
}

func loadImage(reader io.Reader, cipher dcipher.Cipher) (img image.Image, data []byte, err error) {
	return
}

func writeImage(img image.Image, data []byte, cipher dcipher.Cipher, writer io.Writer) (err error) {
	return
}
