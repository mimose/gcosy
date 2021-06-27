package lib

import (
	"bufio"
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
	if !strings.HasPrefix(fileName, "\\") {
		fileName = "\\" + fileName
	}
	fileObj, err := os.OpenFile(dirPath+fileName, fileFlag, 0644)
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
