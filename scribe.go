package archivist

import(
        "time"
        "os"
        "fmt"
        "path/filepath"
)

type Describer struct {
        title, timestamp string
        totalsize int64
        content []DirInfo
        path string
        pathInfo os.FileInfo
}

func (d *Describer) setTitle(gTitle string){
        d.title = gTitle
}

func (d Describer) getTitle() string{
        return d.title
}

func (d *Describer) setTotalSize(gPath string){
        size, err := DirSize(gPath)
        if err != nil {
                fmt.Println(err)
		os.Exit(1)
        }
        d.totalsize = size
}

func (d Describer) getTotalSize() int64{
        return d.totalsize
}

func (d *Describer) setPath(gPath string){
        gPathInfo, err := os.Stat(gPath)
        if err != nil {
                fmt.Println("Whoops, it looks like the specified file or directory doesn't exist")
                fmt.Println(err)
                os.Exit(1)
        }
        d.pathInfo = gPathInfo
        d.path = gPath
}

func (d Describer) getPath() string{
        return d.path
}

func (d *Describer) setTimeStamp(){
        d.timestamp = time.Now().Local().Format(time.Stamp)
}

func (d Describer) getTimeStamp() string{
        return d.timestamp
}

func (d Describer) getPathInfo() os.FileInfo{
        return d.pathInfo
}

func (d *Describer) setContent(){

        switch mode := d.pathInfo.Mode(); {

        case mode.IsDir():
                listOfPaths := []DirInfo{}
                list(d.path, 0, &listOfPaths)
                d.content = listOfPaths

        case mode.IsRegular():
                listOfPaths := []DirInfo{}
                d.content = append(listOfPaths, DirInfo{
                        level: 0,
                        name: d.pathInfo.Name(),
                        size: d.pathInfo.Size(),
                        isDir: false,
                })
        }

}

func list(path string, level int, listOfPaths *[]DirInfo) {
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

                        item := DirInfo{
                                level: level,
                                name: fi.Name(),
                                size: fi.Size(),
                                isDir: true,
                        }

                        *listOfPaths = append(*listOfPaths, item)

                        level++
                        list(path + "/" +  fi.Name(), level, listOfPaths)
                        level--

                } else {

                        item := DirInfo{
                                level: level,
                                name: fi.Name(),
                                size: fi.Size(),
                                isDir: false,
                        }

                        *listOfPaths = append(*listOfPaths, item)

                }
        }
}

type DirInfo struct{
        level int
        name string
        size int64
        isDir bool
}

func DirSize(path string) (int64, error) {
    var size int64
    err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
        if !info.IsDir() {
            size += info.Size()
        }
        return err
    })
    return size, err
}
