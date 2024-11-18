package tools

import (
	"os"
	"path/filepath"
)

func ListStaticFiles(resourceDir, subDir string, modules []string, diskPath bool) []string {
	var res []string
	for _, module := range modules {
		dirName := filepath.Join(resourceDir, "static", module, subDir)
		fileInfos, _ := os.ReadDir(dirName)
		for _, fi := range fileInfos {
			if !fi.IsDir() {
				fPath := filepath.Join("static", module, subDir, fi.Name())
				if diskPath {
					fPath = filepath.Join(resourceDir, fPath)
				}
				res = append(res, fPath)
			}
		}
	}
	return res
}
