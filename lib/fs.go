package lib

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
)

// FileExists returns true if the given file exists
func FileExists(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}

// DirExists returns true if the given file exists and is a directory
func DirExists(dir string) bool {
	stat, err := os.Stat(dir)
	if err != nil {
		return false
	}
	return stat.IsDir()
}

func Write(dirPath, fileName string, buf []byte, append bool) error {
	var fileFlag int
	if append {
		fileFlag = os.O_RDWR | os.O_CREATE | os.O_APPEND
	} else {
		fileFlag = os.O_RDWR | os.O_CREATE | os.O_TRUNC
	}
	if !DirExists(dirPath) {
		if err := os.MkdirAll(dirPath, 0644); err != nil {
			return err
		}
	}
	fileObj, err := os.OpenFile(CompletePath(dirPath, fileName), fileFlag, 0644)
	if err == nil {
		defer fileObj.Close()
		writeObj := bufio.NewWriterSize(fileObj, 4096)

		_, err = writeObj.Write(buf)
		if err == nil {
			err = writeObj.Flush()
		}
		return err
	}
	return err
}

func Read(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(file)
}

func CompletePath(path ...string) string {
	var fullPath string = path[0]
	for i := range path {
		if i > 0 {
			thisPath := path[i]
			if !strings.HasPrefix(thisPath, "\\") {
				thisPath = "\\" + thisPath
			}
			fullPath = strings.Join([]string{fullPath, thisPath}, "")
		}
	}
	return fullPath
}
