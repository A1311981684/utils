package files

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//Only care about files, skip every directory if exists
func GetAllFileNamesUnder(path string) ([]string, error) {
	//Read all content under the path
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	var fileNames []string
	//Classify type of each content
	for _, v := range files {
		if !v.IsDir() {
			//If it is a file type, add to the result
			fileNames = append(fileNames, v.Name())
		} else {
			log.Println("[GetAllFileNames() :warning: path contains directory]")
		}
	}
	return fileNames, nil
}

func GetCurrentProjectName() (string, error) {
	//First get the full path of current working directory
	dirPath, err := filepath.Abs(os.Args[0])
	if err != nil {
		return "", err
	}
	// "/home/app/app.exe", "/app/app.exe", "D:\\project\\app\\app.exe
	splitStr := strings.Split(dirPath, string(filepath.Separator))
	if len(splitStr) < 1 {
		return "", errors.New("ca not find a correct project name under this path: " + dirPath)
	}
	projectName := splitStr[len(splitStr)-2]
	log.Println("Project name is:", projectName)
	return projectName, nil
}