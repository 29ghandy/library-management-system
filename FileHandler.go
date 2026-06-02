package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type File struct {
	name string
	path string
}
type FileMethods interface {
	creatFile(name string) string
	readFile(name string)
	deleteFile(name string) bool
}

func (f *File) createFile(name string) string {

	var path string = "file couldn't be created"
	fileName := name + ".json"
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		file, err := os.Create(fileName)
		if err != nil {
			return "file cannot be created" + err.Error()
		}

		defer file.Close()
		path = file.Name()
		f.name = fileName
		f.path = path
	} else {
		path = f.path
	}

	return path
}
func (f *File) readFile(name string) (interface{}, error) {
	var data interface{}
	fileName := name + ".json"
	dataFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(dataFile, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (f *File) deleteFile(name string) bool {
	fileName := name + ".json"

	err := os.Remove(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return true
		}
		return false
	}

	return true

}
func getAllFiles() []File {
	var files []File
	var path string = "."
	filesInfo, err := ioutil.ReadDir(path)
	if err != nil {
		return files
	}
	for _, f := range filesInfo {
		if f.IsDir() && filepath.Ext(f.Name()) == ".json" {
			files = append(files, File{f.Name(), path + "/" + f.Name()})
		}
	}
	return files
}
