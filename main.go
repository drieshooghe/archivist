package main

import (
        "fmt"
)

func main() {
        
        d := Describer{}
        d.setTitle("testdirectory")
        d.setTimeStamp()
        d.setPath("./testdir")
        d.setContent()
        fmt.Println(d.content)

        //packer := Compressor{format:"zip"}
        //fmt.Println(packer.isValidFormat("zip"))
}

