package commands

import (
	"os"
)

func commandExit(paths *Paths) error {
	os.Exit(0)
	return nil
}
