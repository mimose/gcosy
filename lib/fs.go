package lib

import (
	"bufio"
	"errors"
	"io/ioutil"
	"os"
	"strings"
	"syscall"
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

func Write(dirPath, fileName string, buf []byte, append bool) (string, error) {
	var fileFlag int
	if append {
		fileFlag = os.O_RDWR | os.O_CREATE | os.O_APPEND
	} else {
		fileFlag = os.O_RDWR | os.O_CREATE | os.O_TRUNC
	}
	if !DirExists(dirPath) {
		oldMask := syscall.Umask(0)
		defer syscall.Umask(oldMask)
		if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
			return "", err
		}
	}
	fileFullPath := CompletePath(dirPath, fileName)
	fileObj, err := os.OpenFile(fileFullPath, fileFlag, os.ModePerm)
	if err == nil {
		defer fileObj.Close()
		writeObj := bufio.NewWriterSize(fileObj, 4096)

		_, err = writeObj.Write(buf)
		if err == nil {
			err = writeObj.Flush()
		}
		return fileFullPath, err
	}
	return fileFullPath, err
}

func Read(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ioutil.ReadAll(file)
}

func ReadDirAllFiles(dirPath string) ([][]byte, error) {
	if !DirExists(dirPath) {
		return nil, errors.New("dir is not exists")
	}
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}
	//var filesContents []string
	var fileBytes [][]byte
	for _, file := range files {
		byte, err := Read(CompletePath(dirPath, file.Name()))
		if err != nil {
			return nil, err
		}
		fileBytes = append(fileBytes, byte)
	}
	return fileBytes, nil
}

func DeleteFile(filePath string) error {
	return os.Remove(filePath)
}

func CompletePath(path ...string) string {
	var fullPath = path[0]
	for i := range path {
		if i > 0 {
			thisPath := path[i]
			if !strings.HasPrefix(thisPath, "/") {
				thisPath = "/" + thisPath
			}
			fullPath = strings.Join([]string{fullPath, thisPath}, "")
		}
	}
	return fullPath
}
