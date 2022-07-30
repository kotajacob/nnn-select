package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/adrg/xdg"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: nnn-select [path]")
		os.Exit(1)
	}
	selection := os.Args[1]

	// Get old selection.
	sPath := filepath.Join(xdg.ConfigHome, "nnn", ".selection")
	current, err := os.ReadFile(sPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed reading selections: %v", err)
		os.Exit(1)
	}

	// Check if already selected.
	old := strings.Split(string(current), "\x00")
	for _, v := range old {
		if v == selection {
			os.Exit(0)
		}
	}

	// Append new selection to file.
	if len(current) > 0 {
		current = append(current, []byte("\x00")...)
	}
	current = append(current, []byte(selection)...)
	if err := os.WriteFile(sPath, current, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "failed writing selections: %v", err)
	}
}
