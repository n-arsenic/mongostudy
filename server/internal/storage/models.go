package storage

import "time"

type Sample struct {
	Name     string
	Category struct {
		RootCategory int
		SubCategory  []int
	}
	AnythingList   []string
	IntVolume      int
	BigFloatVolume float64
	Time           time.Time
}

type Category struct {
	Id   int
	Name string
}
