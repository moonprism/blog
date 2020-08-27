package utils

import (
	"io"
	"mime/multipart"
	"os"
)

func FileCopy(file multipart.File, path string) error {
	// Destination
	dst, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	_, err = io.Copy(dst, file)
	return err
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
