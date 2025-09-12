package dbpath

import (
	"os"
	"path/filepath"
)

func GetPath(name string) string {
	if len(os.Args) > 1 {
		return filepath.Join("../", name)
	} else {
		return filepath.Join(".", name)
	}
}
