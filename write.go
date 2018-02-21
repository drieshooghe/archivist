package archivist

import(
        "strings"
        "strconv"
)

func getMarkdown(path, timestamp string, totalSize int64, di []DirInfo) string{
        returnString := "# " + strings.ToUpper(path) + " # \n"
        returnString = returnString + "created: *" + timestamp  + "* \n"
        returnString = returnString + "uncompressed size: *" + strconv.FormatFloat(float64(totalSize)/float64(1048576), 'f', 2, 64) + "MB* \n\n"

        for _, l := range di {
                line := strings.Repeat("     ", l.level) + " * "
                line = line + " **" + l.name
                if l.isDir {
                        line = line + "/"
                }
                line = line + "** "
                line = line + "\t*" + strconv.FormatFloat(float64(l.size)/float64(10485760), 'f', 2, 64,) + "MB*"
		returnString = returnString + line + "\n"
	}

	return returnString
}
