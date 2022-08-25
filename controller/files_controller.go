package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-demo-1/define"
	"github.com/go-demo-1/helper"
	"github.com/google/uuid"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func FilesController(c *gin.Context) {
	file, err := c.FormFile("raw")
	if err != nil {
		log.Fatal(err)
	}

	uploadAbsDir := helper.GetUploadsPath()
	filename := uuid.New().String()
	err = os.MkdirAll(uploadAbsDir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	newFileName := filename+filepath.Ext(file.Filename)
	fileErr := c.SaveUploadedFile(file, filepath.Join(uploadAbsDir, newFileName))
	if fileErr != nil {
		log.Fatal(fileErr)
	}
	c.JSON(http.StatusOK, gin.H{"url": "/" + define.DefaultUploadPath +  "/" + newFileName})
}
