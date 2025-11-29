package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"slices"

	cmd "github.com/souvikelric/dirclean/command"
	"github.com/souvikelric/dirclean/models"
	"github.com/souvikelric/dirclean/utility"
)

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
	var dir string
	flag.StringVar(&dir,"dir","","(not used) specify directory path")
	var command string
	flag.StringVar(&command,"cmd","list","command to execute: move/delete")

	flag.Parse()

	var downloads_path string

	if dir == ""{
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

		downloads_path = filepath.Join(home_dir,"Downloads")
	} else {
		downloads_path = dir
	}

	all_files := utility.GetAllFilesInDir(downloads_path,sort_by)
	slices.Reverse(all_files)

	if only_dirs {
		filtered_dirs := []models.FileInfo{}
		for _, file := range all_files {
			if file.IsDir {
				filtered_dirs = append(filtered_dirs, file)
			}
		}
		all_files = filtered_dirs
	} else {

		if only_files {
			filtered_files := []models.FileInfo{}
			for _, file := range all_files {
				if !file.IsDir {
					filtered_files = append(filtered_files, file)
				}
			}
			all_files = filtered_files
		}

		if ext != "all" {
			ext = "." + ext
			filtered_files := []models.FileInfo{}
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
	utility.PrintFilesInfo(all_files)
	fmt.Println()

	if command == "delete" || command == "del" {
		cmd.ConfirmAndDeleteFiles(dir, all_files)
	}

}