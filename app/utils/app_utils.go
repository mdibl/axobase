package utils

import (
	"io/ioutil"
	"os"

	"github.com/Jeffail/gabs"
)

func ReadMarkdownPath() (string, error) {
	var path = ""
	dir, _ := os.Getwd()
	content, err := ioutil.ReadFile(dir + "/conf/paths.json")
	if err != nil {
		return path, err
	}

	jsonParsed, err := gabs.ParseJSON(content)
	if err != nil {
		return path, err
	}

	return jsonParsed.Path("markdown").Data().(string), nil
}
