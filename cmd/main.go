package main

import (
	"fmt"
	"os"

	"github.com/galkowskit/generate-cover-image/pkg/logo"
)

func main() {
	title := os.Args[1]
	subtitle := os.Args[2]

	err := logo.DrawLogo(title, subtitle)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}
