package tools

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetFileList(path string) []string {
	files := make([]string, 0)
	PthSep := string(os.PathSeparator)
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			files = append(files, path+PthSep)
		} else {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
	return files
}
