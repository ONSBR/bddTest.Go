package util

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const sep = string(os.PathSeparator)

var log = GetLogger("util.fileHandler")

type (
	//FileHandler structure representation for file handling operations
	FileHandler struct{}

	//FileContent represents a file
	FileContent struct {
		Filename string
		Path     string
		Content  string
	}

	iFileHandler interface {
		FindFiles(string) ([]string, error)
		ReadFile(string) (string, error)
		ReadFiles(string) (map[string]FileContent, error)
		WriteFile(string, string) error
		DoesFileExists(string) bool
		BackupFile(string) error
		RemoveFile(string) error
	}
)

//FindFiles list all files based on a folder pattern
func (fileHandler *FileHandler) FindFiles(pattern string) (files []string, err error) {
	path := pattern
	return filepath.Glob(path)
}

//ReadFile read a single file and puts content on content ref parameter
func (fileHandler *FileHandler) ReadFile(filename string) (content string, err error) {
	bytes, errR := ioutil.ReadFile(filename)
	if errR != nil {
		err = errR
		return
	}
	content = string(bytes)
	return
}

//ReadFiles reads all files based on a folder pattern and return a list of file contents
func (fileHandler *FileHandler) ReadFiles(pattern string) (files map[string]FileContent, err error) {
	filenames, errFind := fileHandler.FindFiles(pattern)
	if errFind != nil {
		err = errFind
		return
	}
	files = map[string]FileContent{}
	for _, filename := range filenames {
		content, errRead := fileHandler.ReadFile(filename)
		if errRead != nil {
			err = errRead
			return
		}
		idx := strings.LastIndex(filename, sep)

		path := filename[:idx]
		name := filename[idx+1:]
		files[filename] = FileContent{Filename: name, Path: path, Content: content}
	}
	return
}

//WriteFile writes content to a single file
func (fileHandler *FileHandler) WriteFile(filename string, content string) error {
	data := []byte(content)
	_, err := os.Stat(filename)
	if err == nil {
		_ = os.Remove(filename)
	}
	err = ioutil.WriteFile(filename, data, 0644)
	return err
}

//DoesFileExists helper function to check if a file exists
func (fileHandler *FileHandler) DoesFileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

//BackupFile helper function to save a copy of a single file with .bkp extension
func (fileHandler *FileHandler) BackupFile(filename string) error {
	_, err := os.Stat(filename + ".bkp")
	if err == nil {
		_ = os.Remove(filename + ".bkp")
	}
	return os.Rename(filename, filename+".bkp")
}

//RemoveFile helper function exclude a single file
func (fileHandler *FileHandler) RemoveFile(filename string) error {
	_, err := os.Stat(filename)
	if err == nil {
		err = os.Remove(filename)
	}
	return err
}

//NewFileHandler basic constructor
func NewFileHandler() *FileHandler {
	return &FileHandler{}
}
