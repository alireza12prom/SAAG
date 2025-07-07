package helpers

import (
	"fmt"
	"io"
	"os"
)

func CopyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println("Error opening source file: %w", err)
		return nil
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		fmt.Println("Error creating destination file: %w", err)
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		fmt.Println("Error copying data: %w", err)
		return err
	}

	return nil
}
