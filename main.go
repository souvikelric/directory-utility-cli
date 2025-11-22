package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"time"

	"github.com/fatih/color"
)

type FileInfo struct {
	Name string
	Size int64
	IsDir bool
	LastModified time.Time
	FormattedSize string
}

func formatSize(size int64) string {
	sizes := []string{"B", "KB", "MB", "GB", "TB"}
	i := 0
	floatSize := float64(size)

	for floatSize >= 1024 && i < len(sizes)-1 {
		floatSize /= 1024
		i++
	}

	return fmt.Sprintf("%.2f %s", floatSize, sizes[i])
}

//get size of directory by summing sizes of all files within it
func getDirSize(path string) int64 {
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

// sort by provided FileInfo Size field
func sortFilesByField(files []FileInfo, field string) {

	switch field { 
	case "size":
		sort.Slice(files,func(i,j int) bool {
		return files[i].Size < files[j].Size
		})
	case "name":
		sort.Slice(files,func(i,j int) bool {
		return files[i].Name < files[j].Name
		})
	case "date":
		// sort by file modified time
		sort.Slice(files,func(i,j int) bool {
		return files[i].LastModified.Before(files[j].LastModified)
		})
	default:
		fmt.Println("Invalid sort field. Supported fields are: size, name, date")
	}
	
}

func getAllFilesInDir(dirPath string, sortBy string) []FileInfo {
	downloads_files, err := os.ReadDir(dirPath)
	files := []FileInfo{}

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
			size = getDirSize(filepath.Join(dirPath,file.Name()))
		} else {
			size = fileInfo.Size()
		}

		files = append(files, FileInfo{
			Name:          file.Name(),
			Size:          size,
			IsDir:         file.IsDir(),
			FormattedSize: formatSize(size),
			LastModified: fileInfo.ModTime(),
		})
	}
	sortFilesByField(files,sortBy)

	return files
}

func printFilesInfo(files []FileInfo) {
	for _, file := range files{
		c:=color.New(color.Italic)
		icon := c.Sprint("[file]")
		if file.IsDir {
			icon := c.Sprint("[dir]")
			color.Magenta("%s %s - %s (%s)\n", icon, file.Name, file.FormattedSize,file.LastModified.Format("2006-01-02 15:04:05"))
		} else{
			color.Cyan("%s %s - %s (%s)\n", icon, file.Name, file.FormattedSize,file.LastModified.Format("2006-01-02 15:04:05"))
		}	
	}
}

func main(){

	//defining flag variables to be accepted from command line
	var top int64
	flag.Int64Var(&top,"top",0,"lists top n values")
	var ext string
	flag.StringVar(&ext,"ext","all","filter by file extension")
	var only_dirs bool
	flag.BoolVar(&only_dirs,"ld",false,"list only directories")
	var only_files bool
	flag.BoolVar(&only_files,"lf",false,"list only files")
	var sort_by string
	flag.StringVar(&sort_by,"sort","size","sort by field: size/date/name")

	flag.Parse()

	home_dir , err := os.UserHomeDir()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	//fmt.Println("User Home Directory:", home_dir)
	user_dir_files, err := os.ReadDir(home_dir)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	//fmt.Println("Files and Directories in User Home Directory:")
	found_downloads := false
	for _, file := range user_dir_files {
		if file.IsDir() && file.Name() == "Downloads" {
			found_downloads = true
		}
	}
	if !found_downloads {
		fmt.Println("Downloads directory not found in user home directory.")
		return
	}

	downloads_path := filepath.Join(home_dir,"Downloads")
	all_files := getAllFilesInDir(downloads_path,sort_by)
	slices.Reverse(all_files)

	if only_dirs {
		filtered_dirs := []FileInfo{}
		for _, file := range all_files {
			if file.IsDir {
				filtered_dirs = append(filtered_dirs, file)
			}
		}
		all_files = filtered_dirs
	} else {

		if only_files {
			filtered_files := []FileInfo{}
			for _, file := range all_files {
				if !file.IsDir {
					filtered_files = append(filtered_files, file)
				}
			}
			all_files = filtered_files
		}

		if ext != "all" {
			ext = "." + ext
			filtered_files := []FileInfo{}
			for _, file := range all_files {
				if !file.IsDir && filepath.Ext(file.Name) == ext{
					filtered_files = append(filtered_files, file)
				}
			}
			all_files = filtered_files
		}
	}

	if top > 0 && int64(len(all_files)) > top {
		all_files = all_files[:top]
	}

	fmt.Println()
	printFilesInfo(all_files)
	fmt.Println()

}