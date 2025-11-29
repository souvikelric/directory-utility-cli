package cmd

import (
	"os"
	"path/filepath"
)

func TidyFiles(dirPath string) error {
	// Implementation for tidying files goes here
	files,err := os.ReadDir(dirPath)

	if err != nil {
		return err
	}

	for _,file := range files {
		// get all image files and move them to Images folder
		// get all document files and move them to Documents folder
		// get all video files and move them to Videos folder
		// get all audio files and move them to Audio folder
		// get all other files and move them to Others folder

		if !file.IsDir() {
			// create folders if not exist
			// move files accordingly
			os.MkdirAll(dirPath+"/Images", os.ModePerm)
			os.MkdirAll(dirPath+"/Documents", os.ModePerm)
			os.MkdirAll(dirPath+"/Videos", os.ModePerm)
			os.MkdirAll(dirPath+"/Audio", os.ModePerm)
			os.MkdirAll(dirPath+"/Others", os.ModePerm)

			ext := filepath.Ext(file.Name())
			switch ext {
			case ".jpg", ".jpeg", ".png", ".gif", ".bmp":
				os.Rename(filepath.Join(dirPath, file.Name()), filepath.Join(dirPath, "Images", file.Name()))
			case ".pdf", ".doc", ".docx", ".txt", ".xls", ".xlsx", ".ppt", ".pptx":
				os.Rename(filepath.Join(dirPath, file.Name()), filepath.Join(dirPath, "Documents", file.Name()))
			case ".mp4", ".mkv", ".avi", ".mov":
				os.Rename(filepath.Join(dirPath, file.Name()), filepath.Join(dirPath, "Videos", file.Name()))
			case ".mp3", ".wav", ".aac":
				os.Rename(filepath.Join(dirPath, file.Name()), filepath.Join(dirPath, "Audio", file.Name()))
			default:
				os.Rename(filepath.Join(dirPath, file.Name()), filepath.Join(dirPath, "Others", file.Name()))
			}
		}
	}
	return nil
}