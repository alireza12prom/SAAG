package helpers

import "os"

func CheckDirExists(dir string) bool {
	_, err := os.Stat(dir)
	return !os.IsNotExist(err)
}
