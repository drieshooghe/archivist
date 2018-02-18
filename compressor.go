package archivist

import (
    "github.com/jservice-rvbd/archiver"
    "os"
    "fmt"
)

var validFormats = []string{"zip", "tar.gz"}

type Compressor struct {}

func (c Compressor) isValidFormat(formatPtr string) bool {
        return contains(validFormats, formatPtr)
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
                        err := archiver.Zip.Make(path + ".zip", []string{path})
                        if err != nil {
		                fmt.Println(err)
		                os.Exit(1)
	                }
                case "tar.gz":
                        err := archiver.TarGz.Make(path + ".tar.gz", []string{path})
                        if err != nil {
		                fmt.Println(err)
		                os.Exit(1)
	                }

        }
        
        return nil
}
