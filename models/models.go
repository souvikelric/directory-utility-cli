package models

import "time"

type FileInfo struct {
	Name string
	Size int64
	IsDir bool
	LastModified time.Time
	FormattedSize string
}