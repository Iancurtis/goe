package geo_path

import (
	"path/filepath"
	"strings"
)

// FileBaseName 指定文件名去除后缀名
func FileBaseName(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}
