package repl

import "os"

func commandExit(cfg *replConfig) error {
	os.Exit(0)
	return nil
}
