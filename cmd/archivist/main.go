package main

import (
        "fmt"
        "os"
        "flag"
        "strings"
        "github.com/drieshooghe/archivist"
)

func main() {

        // Parse command line arguments 
        formatPtr := flag.String("format", "zip", "Archiving file format { " + strings.Join(archivist.ValidFormats," | ") + " }")
        flag.Parse()

        c := archivist.Compressor{}
        if ! c.IsValidFormat(*formatPtr) {
                fmt.Println(*formatPtr + " is not allowed as an archiving format")
                os.Exit(1)
        }

        paths := flag.Args()

        a := archivist.Archivist{}

        // Make description file
        for _, path := range paths {
               a.Describe(path)
        }
        // Make archives
        for _, path := range paths {
                a.Make(*formatPtr, path)
        }

}

