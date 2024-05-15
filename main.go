package main

import (
	"filters/picture"

	"github.com/disintegration/gift"
)

func main() {
	filters := make([]gift.Filter, 2)
	filters[0] = gift.Invert()
	filters[1] = gift.Grayscale()

	leao := picture.NewPicture("testdata/car.png", filters...)
	dinheiro := picture.NewPicture("testdata/fish.png", filters...)
	dinossauro := picture.NewPicture("testdata/sunflower.png", filters...)

	pics := []picture.Picture{leao, dinheiro, dinossauro}

	paths := []string{"testdata/teste1.png", "testdata/teste2.png", "testdata/teste3.png"}

	for idx, pic := range pics {
		pic.Apply()
		pic.Save(paths[idx])
	}

}
