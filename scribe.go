package archivist

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type describer struct {
	title, timestamp string
	totalsize        int64
	content          []dirInfo
	path             string
	pathInfo         os.FileInfo
}

func (d *describer) setTitle(gTitle string) {
	d.title = gTitle
}

func (d describer) getTitle() string {
	return d.title
}

func (d *describer) setTotalSize(gPath string) {
	size, err := dirSize(gPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	d.totalsize = size
}

func (d describer) getTotalSize() int64 {
	return d.totalsize
}

func (d *describer) setPath(gPath string) {
	gPathInfo, err := os.Stat(gPath)
	if err != nil {
		fmt.Println("Whoops, it looks like the specified file or directory doesn't exist")
		fmt.Println(err)
		os.Exit(1)
	}
	d.pathInfo = gPathInfo
	d.path = gPath
}

func (d describer) getPath() string {
	return d.path
}

func (d *describer) setTimeStamp() {
	d.timestamp = time.Now().Local().Format(time.Stamp)
}

func (d describer) getTimeStamp() string {
	return d.timestamp
}

func (d describer) getPathInfo() os.FileInfo {
	return d.pathInfo
}

func (d *describer) setContent() {

	switch mode := d.pathInfo.Mode(); {

	case mode.IsDir():
		listOfPaths := []dirInfo{}
		list(d.path, 0, &listOfPaths)
		d.content = listOfPaths

	case mode.IsRegular():
		listOfPaths := []dirInfo{}
		d.content = append(listOfPaths, dirInfo{
			level: 0,
			name:  d.pathInfo.Name(),
			size:  d.pathInfo.Size(),
			isDir: false,
		})
	}

}

func list(path string, level int, listOfPaths *[]dirInfo) {
	dir, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer dir.Close()
	fi, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, fi := range fi {
		if fi.IsDir() {

			item := dirInfo{
				level: level,
				name:  fi.Name(),
				size:  fi.Size(),
				isDir: true,
			}

			*listOfPaths = append(*listOfPaths, item)

			level++
			list(path+"/"+fi.Name(), level, listOfPaths)
			level--

		} else {

			item := dirInfo{
				level: level,
				name:  fi.Name(),
				size:  fi.Size(),
				isDir: false,
			}

			*listOfPaths = append(*listOfPaths, item)

		}
	}
}

type dirInfo struct {
	level int
	name  string
	size  int64
	isDir bool
}

func dirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}
