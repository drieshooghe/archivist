package main

import (
        "fmt"
        "os"
        "github.com/drieshooghe/archivist"
)

func main() {

        // Parse command line arguments 
        formatPtr := flag.String("format", "zip", "Archiving file format { " + strings.Join(validFormats," | ") + " }")
                flag.Parse()

        c := archivist.Compressor{}
        if ! c.isValidFormat(*formatPtr) {
                fmt.Println(*formatPtr + " is not allowed as an archiving format")
                os.Exit(1)
        }

        paths := flag.Args()

        // Make description file
        for _, path := range paths {
                _, err := archivist.Describe(path)
                if err != nil {
                        fmt.Println(err)
                        os.Exit(1)
                }
        
        // Make archives
        for _, path := range paths {
                _, err := archivist.Make(*formatPtr, path)
                if err != nil {
                        fmt.Println(err)
                        os.Exit(1)
                }
        }

}

