package database

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

var filesMap = map[string]string{}

type File struct {
	name string
	path string
}
type FileMethods interface {
	CreatFile(name string) string
	ReadFile(name string)
	DeleteFile(name string) bool
	WriteToFile(name string, data interface{})
}

func (f *File) CreateFile(name string) string {

	var path string = "file couldn't be created"

	fileName := name + ".json"
	filePath := filepath.Join("files", fileName)
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			return "file cannot be created" + err.Error()
		}

		defer file.Close()
		path = filePath
		f.name = fileName
		f.path = path
	} else {
		path = f.path
	}
	filesMap[fileName] = path
	return path
}
func (f *File) ReadFile(name string) (map[string]interface{}, error) {
	var data map[string]interface{}
	fileName := "C:\\Users\\KCSP2494\\GolandProjects\\awesomeProject\\files\\fileBase.json"
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
func (f *File) DeleteFile(name string) bool {
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

func (f *File) WriteToFile(name string, data map[string]interface{}) error {

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	path := path.Join("files", name+".json")
	err = ioutil.WriteFile(path, jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}

func deleteFromFile(name string) bool {
	return false
}
