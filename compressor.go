package main

var validFormats = []string{"zip"}

type Compressor struct {
        format string
}

func (c Compressor) getFormat() string{
        return c.format
}

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
