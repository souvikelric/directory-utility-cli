package command

import (
	"fmt"
	"os"
)

func DeleteFiles(path string) error {
	// Implementation for deleting files goes here
	err := os.Remove(path)
	if err != nil {
		fmt.Println("Error deleting file:", err)
	}
	return nil
}