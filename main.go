package main

import (
	"filters/picture"
	"filters/picture/filter"
	"fmt"
	"log"

	"github.com/disintegration/gift"
)

func main() {

	var generators = []filter.Filter{

		{Name: "sepia", Params: []any{90.0}},
		{Name: "convolution", Params: []any{[]float64{1.0, 3.0, 5.0, 3.4, 5.5, 5.5}, true, true, false, 2.0}},
		{Name: "sepia", Params: []any{50.0}},
	}

	var filters []gift.Filter
	for _, filter := range generators {
		f, err := filter.Generate()
		if err != nil {
			log.Fatalf("error generating filter %s: %s", filter.Name, err.Error())

		}
		filters = append(filters, f)

	}
	fmt.Println(len(filters))

	car := picture.NewPicture("testdata/car.png", filters[0])
	fish := picture.NewPicture("testdata/fish.png", filters[1])
	sunflower := picture.NewPicture("testdata/sunflower.png", filters[2])

	pics := []*picture.Picture{car, fish, sunflower}

	paths := []string{"testdata/results/teste1.png", "testdata/results/teste2.png", "testdata/results/teste3.png"}

	for idx, pic := range pics {
		pic.Apply()
		pic.Save(paths[idx])
	}

}
