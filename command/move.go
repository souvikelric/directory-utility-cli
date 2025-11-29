package cmd

import (
	"fmt"

	"github.com/souvikelric/dirclean/models"
)

func MoveFiles(downloads_path string, all_files []models.FileInfo, dest string) error {
	err := CopyFiles(downloads_path, all_files, dest)
			if err != nil {
				fmt.Println("Error moving files:", err)
				return err
			}
	ConfirmAndDeleteFiles(downloads_path,all_files,false)
	return nil
}