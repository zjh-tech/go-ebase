package util

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func GetAllFile(fileList *[]string, pathname string) error {
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		return err
	}

	for _, fi := range rd {
		if fi.IsDir() {
			GetAllFile(fileList, pathname+fi.Name()+"/")
		} else {
			*fileList = append(*fileList, pathname+fi.Name())
		}
	}
	return err
}

func GetExeDirectory() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}
	return strings.Replace(dir, "\\", "/", -1), nil
}
