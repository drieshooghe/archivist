package archivist

import (
        "fmt"
        "strings"
        "os"
        "io"
)

type Archivist struct {}

func (a Archivist) Describe(path string) error{
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

                return nil
}

func (a Archivist) Make(format, path string) error{
                c := Compressor{}
                err := c.compress(format, path)
                if err != nil {
                        fmt.Println(err)
                        os.Exit(1)
                }

                return nil
}
