package scan

import (
	"fmt"
	"os"
	"path/filepath"
)

// GetFilesInDirectory scans the given directory and returns a slice of FileInfo
// containing information about each file and subdirectory, including sizes for directories.
func GetDirSize(path string) int64 {
	var totalSize int64 = 0

	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			totalSize += info.Size()
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}

	return totalSize
}