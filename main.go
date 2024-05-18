package main

import (
	"filters/picture"
	"filters/picture/filter"
	"log"

	"sync"

	"github.com/disintegration/gift"
)

func main() {
	var wg sync.WaitGroup
	var generators = []filter.Filter{

		{Name: "sepia", Params: []any{90.0}},
		{Name: "hue", Params: []any{500.6}},
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

	car := picture.NewPicture("testdata/car.png", filters[0])
	fish := picture.NewPicture("testdata/fish.png", filters[1])
	sunflower := picture.NewPicture("testdata/sunflower.png", filters[2])

	pics := []*picture.Picture{car, fish, sunflower}

	paths := []string{"testdata/results/teste1.png", "testdata/results/teste2.png", "testdata/results/teste3.png"}
	wg.Add(len(pics))
	for idx, pic := range pics {
		if pic == nil {
			log.Fatal("nil Picture object!")
			continue
		}
		go func() {
			pic.Apply()
			pic.Save(paths[idx])
			defer wg.Done()
		}()
	}
	wg.Wait()
}
