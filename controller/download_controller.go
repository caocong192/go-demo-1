package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-demo-1/helper"
	"net/http"
	"path/filepath"
)


func DownloadController(c *gin.Context) {
	if path := c.Param("path"); path != "" {
		target := filepath.Join(helper.GetUploadsPath(), path)
		c.Header("Content-Description", "File Transfer")
		c.Header("Content-Transfer-Encoding", "binary")
		c.Header("Content-Disposition", "attachment; filename="+path)
		c.Header("Content-Type", "application/octet-stream")
		c.File(target)
	} else {
		c.Status(http.StatusNotFound)
	}
}
