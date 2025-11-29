package cmd

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/souvikelric/dirclean/models"
)

func copyDirectory(srcDir, destDir string) error {
	if srcDir == "" || destDir == "" || strings.Contains(srcDir,".") {
		return nil
	}
	entries, err := os.ReadDir(srcDir)
	if err != nil {
		return err
	}

	err = os.MkdirAll(destDir, os.ModePerm)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(srcDir, entry.Name())
		destPath := filepath.Join(destDir, entry.Name())

		if entry.IsDir() {
			err = copyDirectory(srcPath, destPath)
			if err != nil {
				return err
			}
		} else {
			srcFile, err := os.Open(srcPath)
			if err != nil {
				return err
			}
			defer srcFile.Close()

			destFile, err := os.Create(destPath)
			if err != nil {
				return err
			}
			defer destFile.Close()

			_, err = io.Copy(destFile, srcFile)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func CopyFiles(dir string, srcFiles []models.FileInfo, destPath string) error {

	if destPath == "" {
		return errors.New("destination path cannot be empty")
	}
	for _, file := range srcFiles {
		srcFilePath := filepath.Join(dir, file.Name)
		destFilePath := filepath.Join(destPath, file.Name)

		_,err := os.Stat(destFilePath)
		if !os.IsExist(err) {
			//create destination directory if not exists
			err := os.MkdirAll(destPath, os.ModePerm)
			if err != nil {
				return err
			}	
		}

		//copy directories and all files and subdirectories inside it
		if file.IsDir {
			err := copyDirectory(srcFilePath, destFilePath)
			if err != nil {
				return err
			}
			continue
		}

		srcFile, err := os.Open(srcFilePath)
		if err != nil {
			return err
		}
		defer srcFile.Close()

		destFile, err := os.Create(destFilePath)
		if err != nil {
			return err
		}
		defer destFile.Close()

		_, err = io.Copy(destFile, srcFile)
		if err != nil {
			return err
		}
	}
	return nil
}