package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func ReadMarkdownFileBody(filePath string) (string, error) {
	var err error
	var fileBody string

	// Open the file for buffer based read.
	fileBuf, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Defer file handle closing.
	defer func() {
		if err = fileBuf.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// Create a file scanner for reading the lines of the file.
	scanner := bufio.NewScanner(fileBuf)

	// Read the file line by line, ignore the front-matter.
	ignoreLines := false
	for scanner.Scan() {
		line := scanner.Text()
		if ignoreLines && strings.Contains(line, "---") {
			ignoreLines = false
		} else if strings.Contains(line, "---") {
			ignoreLines = true
			continue
		}

		if !ignoreLines {
			fileBody += fmt.Sprintf("%s\n", line)
		}
	}

	return fileBody, nil
}

func InsertImgURL(fileBody, url string) string {
	re := regexp.MustCompile(`\{{2}url\}{2}`)
	return re.ReplaceAllString(fileBody, url)
}
