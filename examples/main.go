package main

import (
	"image"
	"image/png"
	"log"
	"math"
	"os"

	"github.com/danibezoff/perspective-transform/perspective"
)

const srcImgFile = "gopher.png"

func main() {
	srcImg := loadPng(srcImgFile)
	dstImg := image.NewRGBA(srcImg.Bounds())

	srcPoints := [8]float64{
		0, 0, 100, 0, 100, 100, 0, 100,
	}
	dstPoints := [8]float64{
		0, 0, 20, 0, 200, 200, 0, 20,
	}

	p := perspective.New(srcPoints, dstPoints)

	for x := dstImg.Bounds().Min.X; x < dstImg.Bounds().Max.X; x++ {
		for y := dstImg.Bounds().Min.Y; y < dstImg.Bounds().Max.Y; y++ {
			srcX, srcY := p.TransformInv(float64(x), float64(y))
			c := srcImg.At(int(math.Round(srcX)), int(math.Round(srcY)))
			dstImg.Set(x, y, c)
		}
	}

	savePng("new-gopher.png", dstImg)
}

func loadPng(filename string) image.Image {
	reader, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	img, err := png.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	return img
}

func savePng(filename string, img image.Image) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	png.Encode(file, img)
}
