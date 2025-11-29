package utility

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/souvikelric/dirclean/models"
	scan "github.com/souvikelric/dirclean/scanner"
)

func GetAllFilesInDir(dirPath string, sortBy string) []models.FileInfo {
	downloads_files, err := os.ReadDir(dirPath)
	files := []models.FileInfo{}

	var size int64 = 0

	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	for _, file := range downloads_files {
		fileInfo, err := file.Info()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		
		if file.IsDir(){
			size = scan.GetDirSize(filepath.Join(dirPath,file.Name()))
		} else {
			size = fileInfo.Size()
		}

		files = append(files, models.FileInfo{
			Name:          file.Name(),
			Size:          size,
			IsDir:         file.IsDir(),
			FormattedSize: FormatSize(size),
			LastModified: fileInfo.ModTime(),
		})
	}
	SortFilesByField(files,sortBy)

	return files
}

func PrintFilesInfo(files []models.FileInfo) {
	for _, file := range files{
		c:=color.New(color.Italic)
		icon := c.Sprint("[file]")
		if file.IsDir {
			icon := c.Sprint("[dir]")
			if file.Name[0] == '.' {
				color.HiYellow("%s %s - %s (%s)\n", icon, file.Name, file.FormattedSize,file.LastModified.Format("2006-01-02 15:04:05"))
			} else {
				color.Magenta("%s %s - %s (%s)\n", icon, file.Name, file.FormattedSize,file.LastModified.Format("2006-01-02 15:04:05"))
			}
		} else{
			color.Cyan("%s %s - %s (%s)\n", icon, file.Name, file.FormattedSize,file.LastModified.Format("2006-01-02 15:04:05"))
		}	
	}
}

