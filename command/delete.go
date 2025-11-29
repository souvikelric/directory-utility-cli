package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/souvikelric/dirclean/models"
)

func DeleteFiles(path string) error {
	// Implementation for deleting files goes here
	err := os.Remove(path)
	if err != nil {
		fmt.Println("Error deleting file:", err)
	}
	return nil
}

func ConfirmAndDeleteFiles(dir string, files []models.FileInfo, confirmCheck bool) {
	if confirmCheck {
		var response string
		fmt.Print("Are you sure you want to delete the listed files? (y/n): ")
		fmt.Scanln(&response)
		if response != "y" && response != "Y" {
			fmt.Println("Deletion cancelled.")
			return
		}
	}
	for _, file := range files {
		err := os.Remove(filepath.Join(dir,file.Name))
		if err != nil {
			fmt.Println("Error deleting file:", err)
		} else {
			if confirmCheck {
				fmt.Println("Deleted:", file.Name)
			}
		}
	}
}

