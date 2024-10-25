package commands

type Paths struct {
	Next     *string
	Previous *string
}

func NewPaths() *Paths {
	return &Paths{Next: nil, Previous: nil}
}
