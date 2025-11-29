package utility

import (
	"fmt"
	"sort"

	"github.com/souvikelric/dirclean/models"
)

// sort by provided FileInfo field
func SortFilesByField(files []models.FileInfo, field string) {

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