package main

import (
	"filters/picture"

	"github.com/disintegration/gift"
)

func main() {
	filters := make([]gift.Filter, 2)
	filters[0] = gift.Rotate180()
	filters[1] = gift.Grayscale()

	car := picture.NewPicture("testdata/car.png", filters...)
	fish := picture.NewPicture("testdata/fish.png", gift.Sepia(85))
	sunflower := picture.NewPicture("testdata/sunflower.png", gift.GaussianBlur(1.5))

	pics := []picture.Picture{car, fish, sunflower}

	paths := []string{"testdata/results/teste1.png", "testdata/results/teste2.png", "testdata/results/teste3.png"}

	for idx, pic := range pics {
		pic.Apply()
		pic.Save(paths[idx])
	}

}
