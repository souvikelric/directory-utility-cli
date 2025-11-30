package cmd

import (
	"os"
	"path/filepath"
)

// function to check if provided folder name exists
// if it doesnt it creates the folder
func checkPath(folderName string){
	_,err := os.Stat(folderName)
	if os.IsNotExist(err){
		os.MkdirAll(folderName,os.ModePerm)
	}
}

func TidyFiles(dirPath string) error {
	files,err := os.ReadDir(dirPath)

	if err != nil {
		return err
	}

	for _,file := range files {

		if !file.IsDir() {
			// create folders if not exist
			// move files accordingly

			ext := filepath.Ext(file.Name())
			switch ext {
			case ".jpg", ".jpeg", ".png", ".gif", ".bmp":
				checkPath(dirPath+"/Images")
				os.Rename(filepath.Join(dirPath, file.Name()), filepath.Join(dirPath, "Images", file.Name()))
			case ".pdf", ".doc", ".docx", ".txt", ".xls", ".xlsx", ".ppt", ".pptx":
				checkPath(dirPath+"/Documents")
				os.Rename(filepath.Join(dirPath, file.Name()), filepath.Join(dirPath, "Documents", file.Name()))
			case ".mp4", ".mkv", ".avi", ".mov":
				checkPath(dirPath+"/Videos")
				os.Rename(filepath.Join(dirPath, file.Name()), filepath.Join(dirPath, "Videos", file.Name()))
			case ".mp3", ".wav", ".aac":
				checkPath(dirPath+"/Audio")
				os.Rename(filepath.Join(dirPath, file.Name()), filepath.Join(dirPath, "Audio", file.Name()))
			default:
				checkPath(dirPath+"/Others")
				os.Rename(filepath.Join(dirPath, file.Name()), filepath.Join(dirPath, "Others", file.Name()))
			}
		}
	}
	return nil
}