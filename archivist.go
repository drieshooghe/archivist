package archivist

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type archivist struct{}

func (a archivist) describe(path string) error {
	d := describer{}
	d.setTitle(path)
	d.setTimeStamp()
	d.setPath(path)
	d.setContent()
	d.setTotalSize(path)

	md := getMarkdown(d.path, d.timestamp, d.totalsize, d.content)

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

func (a archivist) make(format, path string) error {
	c := compressor{}
	err := c.compress(format, path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return nil
}
