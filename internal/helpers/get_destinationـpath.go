package helpers

import (
	"path"
)

func GetDestinationPath(base, username, dir string) string {
	return path.Join(base, username, dir)
}
