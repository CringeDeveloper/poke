package commands

type Paths struct {
	Next     *string
	Previous *string
}

func NewPaths() *Paths {
	base := "https://pokeapi.co/api/v2/location-area/"
	return &Paths{Next: &base, Previous: nil}
}
