package utility

import "fmt"

func FormatSize(size int64) string {
	sizes := []string{"B", "KB", "MB", "GB", "TB"}
	i := 0
	floatSize := float64(size)

	for floatSize >= 1024 && i < len(sizes)-1 {
		floatSize /= 1024
		i++
	}

	return fmt.Sprintf("%.2f %s", floatSize, sizes[i])
}
