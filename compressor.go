package archivist

import (
	"fmt"
	"os"

	archivingtool "github.com/mholt/archiver"
)

var validFormats = []string{"zip", "tar.gz"}

type compressor struct{}

func (c compressor) isValidFormat(formatPtr string) bool {
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

func (c compressor) compress(format, path string) error {

	switch format {

	case "zip":
		err := archivingtool.Zip.Make(path+"."+format, []string{path})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "tar.gz":
		err := archivingtool.TarGz.Make(path+"."+format, []string{path})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	return nil
}
