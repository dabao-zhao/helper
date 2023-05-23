package xfile

import "os"

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
