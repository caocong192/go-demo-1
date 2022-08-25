package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-demo-1/define"
	"github.com/go-demo-1/helper"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func TextsController(c *gin.Context) {
	var json struct {
		Raw string `json:"raw"`
	}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		uploadAbsDir := helper.GetUploadsPath()
		filename := uuid.New().String()
		err = os.MkdirAll(uploadAbsDir, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}

		newFileName := filename+".txt"
		err = ioutil.WriteFile(filepath.Join(uploadAbsDir, newFileName), []byte(json.Raw), 0644)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{"url": "/" + define.DefaultUploadPath +  "/" +newFileName})
	}
}
