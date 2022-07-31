package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/adrg/xdg"
)

// wanted reads the passed list of paths or exits with a usage message.
func wanted() []string {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Println("usage: nnn-select [paths...]")
		os.Exit(1)
	}

	var s []string
	for _, v := range flag.Args() {
		s = append(s, v)
	}
	return s
}

// existing selections from the nnn .selection file.
func existing(path string) []string {
	b, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return []string{}
	} else if err != nil {
		fmt.Fprintf(os.Stderr, "failed reading current selections: %v", err)
		os.Exit(1)
	}
	if len(b) == 0 {
		return []string{}
	}
	return strings.Split(string(b), "\x00")
}

// selection takes a list of wants and exists paths and returns a null character
// separated byte slice. If a path was found in one list but not the other it's
// included, but if it's found in both lists it's not included.
func selection(wants, exists []string) []byte {
	set := make(map[string]struct{})
	for _, e := range exists {
		set[e] = struct{}{}
	}
	for _, w := range wants {
		if _, ok := set[w]; ok {
			delete(set, w)
			continue
		}
		set[w] = struct{}{}
	}

	selections := make([]string, 0, len(set))
	for k := range set {
		selections = append(selections, k)
	}
	return []byte(strings.Join(selections, "\x00"))
}

// save the selections to path.
func save(path string, selections []byte) {
	if err := os.WriteFile(path, selections, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "failed writing selections: %v", err)
	}
}

func main() {
	wants := wanted()
	path := filepath.Join(xdg.ConfigHome, "nnn", ".selection")
	exists := existing(path)
	save(path, selection(wants, exists))
}
