package commands

import (
	"os"
)

func commandExit(paths *Paths, arg string) error {
	os.Exit(0)
	return nil
}
