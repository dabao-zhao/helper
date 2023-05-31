package xfile

import (
	"fmt"
	"os"
)

func IsDir(path string) bool {
	var (
		fileInfo os.FileInfo
		err      error
	)
	fileInfo, err = os.Stat(path)

	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

func IsFile(path string) bool {
	var (
		fileInfo os.FileInfo
		err      error
	)
	fileInfo, err = os.Stat(path)

	if err != nil {
		return false
	}
	return !fileInfo.IsDir()
}

func FileExists(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}

func CreateIfNotExist(file string) (*os.File, error) {
	if _, err := os.Stat(file); !os.IsNotExist(err) {
		return nil, fmt.Errorf("%s already exist", file)
	}

	return os.Create(file)
}

func SameFile(path1, path2 string) (bool, error) {
	stat1, err := os.Stat(path1)
	if err != nil {
		return false, err
	}

	stat2, err := os.Stat(path2)
	if err != nil {
		return false, err
	}

	return os.SameFile(stat1, stat2), nil
}
