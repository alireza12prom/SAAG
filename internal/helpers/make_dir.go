package helpers

import "os"

func MakeDir(dir string) error {
	if CheckDirExists(dir) {
		return nil
	}

	err := os.MkdirAll(dir, 0o755)
	if err != nil {
		return err
	}

	return nil
}
