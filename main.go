package main

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main(){
	s, err := GetCurrentPath()
	if err != nil {
		panic(err)
	}
	log.Println(s)
}

func GetCurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}

	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return "", errors.New(`error: Can't find "/" or "\".`)
	}
	return string(path[i-1:i+1]), nil
}