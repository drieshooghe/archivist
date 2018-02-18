package main

import(
        "strings"
        "strconv"
)

func getMarkdown(path, timestamp string, di []DirInfo) string{
        returnString := "# " + strings.ToUpper(path) + " # \n"
        returnString = returnString + "created: *" + timestamp  + "* \n\n"

        for _, l := range di {
                line := strings.Repeat("     ", l.level) + " * "
                line = line + " **" + l.name
                if l.isDir {
                        line = line + "/"
                }
                line = line + "** "
                line = line + "\t*" + strconv.FormatInt(l.size, 10) + "kB*"
		returnString = returnString + line + "\n"
	}

	return returnString
}
