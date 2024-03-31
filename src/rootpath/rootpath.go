package rootpath

import (
	"os"
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)

	Root = filepath.Join(filepath.Dir(b), "../..", "src")
)

func Init() {
	args := os.Args[1:]
	if len(args) == 0 {
		Root = filepath.Join(filepath.Dir(Root), "/app")
	}
}
