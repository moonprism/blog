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
