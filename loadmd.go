package main

import (
	"bytes"
	"embed"
	"fmt"
	"log"
	"os"

	"charm.land/glamour/v2"
	"github.com/charmbracelet/colorprofile"
)

//go:embed recipes/example.md
var f embed.FS

func main() {
	f, err := os.Open("recipes/example.md")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	var buf bytes.Buffer
	if _, readErr := buf.ReadFrom(f); readErr != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", readErr)
		os.Exit(1)
	}

	w := colorprofile.NewWriter(os.Stdout, os.Environ())

	fmt.Fprintf(&buf, "\n\nBy the way, this was rendered as _%s._\n", w.Profile)

	r, err := glamour.NewTermRenderer(glamour.WithEnvironmentConfig())
	if err != nil {
		log.Fatal(err)
	}

	md, err := r.RenderBytes(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "%s\n", md)
}
