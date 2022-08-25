package helper

import (
	"github.com/go-demo-1/define"
	"path"
	"path/filepath"
	"runtime"
)

// GetCurrentAbPathByCaller 获取当前绝对路径
func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}


// GetUploadsPath 默认的uploads路径
func GetUploadsPath() (uploads string) {
	dir := getCurrentAbPathByCaller()
	uploads = filepath.Join(dir, define.DefaultUploadPath)
	return
}
