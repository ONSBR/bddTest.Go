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
	FileHandler struct{}

	FileContent struct {
		Filename string
		Path     string
		Content  string
	}

	FileHandlerInterface interface {
		FindFiles(string) ([]string, error)
		ReadFile(string) (string, error)
		ReadFiles(string) (map[string]FileContent, error)
		WriteFile(string, string) error
		DoesFileExists(string) bool
		BackupFile(string) error
	}
)

func (this *FileHandler) FindFiles(pattern string) (files []string, err error) {
	path := pattern
	return filepath.Glob(path)
}

func (this *FileHandler) ReadFile(filename string) (content string, err error) {
	bytes, errR := ioutil.ReadFile(filename)
	if errR != nil {
		err = errR
		return
	}
	content = string(bytes)
	return
}

func (this *FileHandler) ReadFiles(pattern string) (files map[string]FileContent, err error) {
	filenames, errFind := this.FindFiles(pattern)
	if errFind != nil {
		err = errFind
		return
	}
	files = map[string]FileContent{}
	for _, filename := range filenames {
		content, errRead := this.ReadFile(filename)
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

func (this *FileHandler) WriteFile(filename string, content string) error {
	data := []byte(content)
	err := ioutil.WriteFile(filename, data, 0644)
	return err
}

func (this *FileHandler) DoesFileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func (this *FileHandler) BackupFile(filename string) error {
	_, err := os.Stat(filename + ".bkp")
	if err == nil {
		_ = os.Remove(filename + ".bkp")
	}
	return os.Rename(filename, filename+".bkp")
}
