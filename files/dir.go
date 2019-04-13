package files

import (
	"io/ioutil"
	"log"
)

func GetAllFileNamesUnder(path string) ([]string, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	var fileNames []string
	for _, v := range files {
		if !v.IsDir() {
			fileNames = append(fileNames, v.Name())
		} else {
			log.Println("[GetAllFileNames() :warning: path contains directory]")
		}
	}
	return fileNames, nil
}
