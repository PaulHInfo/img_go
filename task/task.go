package task

import (
	"os"
	"path/filepath"
	"strings"
)

type Tasker interface {
	Process()
}

type sirCtx struct {
	SrcDir string
	DstDir string
	files  []string
}

func BuildFileList(SrcDir string) []string {
	files := []string{}
	filepath.Walk(SrcDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() || !strings.HasSuffix(path, "jpeg") {
			return nil
		}
		files = append(files, path)
		return nil
	})
	return files
}
