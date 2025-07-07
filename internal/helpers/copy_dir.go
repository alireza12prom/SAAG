package helpers

import (
	"os"
	"path/filepath"
)

func CopyDir(src string, dst string) error {
	return filepath.WalkDir(src, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		rel, _ := filepath.Rel(src, path)
		target := filepath.Join(dst, rel)

		if d.IsDir() {
			return MakeDir(target)
		}

		if !IsRegular(path) {
			return nil
		}

		return CopyFile(path, target)
	})
}
