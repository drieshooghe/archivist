package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/drieshooghe/archivist"
)

func main() {

	// Parse command line arguments
	formatPtr := flag.String("format", "zip", "Archiving file format { "+strings.Join(archivist.ValidFormats, " | ")+" }")
	flag.Parse()

	c := archivist.Compressor{}
	if !c.IsValidFormat(*formatPtr) {
		fmt.Println(*formatPtr + " is not allowed as an archiving format")
		os.Exit(1)
	}

	paths := flag.Args()

	a := archivist.Archivist{}
	format := *formatPtr

	// Make description file
	for _, path := range paths {
		err := a.Describe(path)
		if err != nil {
			fmt.Println(format + " is not allowed as an archiving format")
			os.Exit(1)
		}
	}

	// Make archives
	for _, path := range paths {
		err := a.Make(format, path)
		if err != nil {
			fmt.Println(format + " is not allowed as an archiving format")
			os.Exit(1)
		}

	}

}
