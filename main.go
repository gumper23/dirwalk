package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Printf("Number of arguments: [%d]\n", len(os.Args[1:]))
	if len(os.Args[1:]) == 0 {
		fmt.Fprintf(os.Stderr, "Usage: %s <dir1> <dir2> ... <dirn>\n", os.Args[0])
		os.Exit(1)
	}

	for _, dir := range os.Args[1:] {
		err := DirWalk(dir, 0)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s", err.Error())
			os.Exit(1)
		}
	}
}

func DirWalk(dir string, indent int) (err error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	// Basis case
	if len(entries) == 0 {
		return nil
	}

	for _, entry := range entries {
		if entry.IsDir() {
			fmt.Printf("%s[%s]\n", strings.Repeat("  ", indent), entry.Name())
		} else {
			fmt.Printf("%s%s\n", strings.Repeat("  ", indent), entry.Name())
		}
		if entry.IsDir() {
			err = DirWalk(filepath.Join(dir, entry.Name()), indent+1)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
