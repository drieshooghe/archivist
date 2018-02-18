package main

import (
        "fmt"
        "flag"
        "strings"
        "os"
        "io"
)

func main() {

        // Parse command line arguments 
        formatPtr := flag.String("format", "zip", "Archiving file format { " + strings.Join(validFormats," | ") + " }")
                flag.Parse()

        c := Compressor{}
        if ! c.isValidFormat(*formatPtr) {
                fmt.Println(*formatPtr + " is not allowed as an archiving format")
                os.Exit(1)
        }

        paths := flag.Args()

        // Make description for each specified path
        for _, path := range paths {
                d := Describer{}
                d.setTitle(path)
                d.setTimeStamp()
                d.setPath(path)
                d.setContent()

                md := getMarkdown(d.path, d.timestamp, d.content)

                f, err := os.Create(path + ".md")
                if err != nil {
                        fmt.Println(err)
			os.Exit(1)
                }
                defer f.Close()

                _, err = io.Copy(f, strings.NewReader(md))
                if err != nil {
                        fmt.Println(err)
			os.Exit(1)
                }
        }

        // Make archive for each path
        for _, path := range paths {
                err := c.compress(*formatPtr, path)
                if err != nil {
                        fmt.Println(err)
                        os.Exit(1)
                }
        }
}

