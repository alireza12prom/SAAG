package helpers

import "os"

func IsRegular(src string) bool {
	info, err := os.Stat(src)
	if err != nil {
		return false
	}
	return info.Mode().IsRegular()
}
