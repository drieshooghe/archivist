package archivist

import (
        "github.com/jservice-rvbd/archiver"
        "os"
        "fmt"
)

var ValidFormats = []string{"zip", "tar.gz"}

type Compressor struct {}

func (c Compressor) IsValidFormat(formatPtr string) bool {
        return contains(ValidFormats, formatPtr)
}

func contains(arr []string, str string) bool {
        for _, a := range arr {
                if a == str {
                        return true
                }
        }
        return false
}

func (c Compressor) compress(format, path string) error {

        switch format {

        case "zip":
                err := archiver.Zip.Make(path + "." + format, []string{path})
                if err != nil {
                        fmt.Println(err)
                        os.Exit(1)
                }
        case "tar.gz":
                err := archiver.TarGz.Make(path + "." + format, []string{path})
                if err != nil {
                        fmt.Println(err)
                        os.Exit(1)
                }
        }
        return nil
}
