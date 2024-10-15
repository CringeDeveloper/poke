package main

type Paths struct {
	Next     *string
	Previous *string
}

func newPaths() *Paths {
	base := "https://pokeapi.co/api/v2/location-area/"
	return &Paths{Next: &base, Previous: nil}
}
