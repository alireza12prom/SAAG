package helpers

import "os"

func RemoveDir(dir string) error {
	if !CheckDirExists(dir) {
		return nil
	}

	err := os.RemoveAll(dir)
	if err != nil {
		return err
	}

	return nil
}
