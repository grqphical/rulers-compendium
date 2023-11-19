package test_util

import (
	"os"
	"path"
	"runtime"
)

// Workaround for directory issue where go uses the cwd of the package instead of the workspace root
func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}
