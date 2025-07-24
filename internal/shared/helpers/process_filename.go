package helpers

import "path/filepath"

func ProcessFilename(filename string) string {
	ext := filepath.Ext(filename)               // ".txt"
	base := filepath.Base(filename)             // "file.name.txt"
	nameWithoutExt := base[:len(base)-len(ext)] // "file.name"

	return nameWithoutExt + "_ascii" + ext
}
