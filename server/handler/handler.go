package handler

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

// Upload handles file upload request
func Upload(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]

	uuid := c.PostForm("uuid")

	for _, file := range files {

		ext := filepath.Ext(file.Filename)

		err := c.SaveUploadedFile(file, "images/"+uuid+ext)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

	}
	c.JSON(http.StatusOK, gin.H{"message": "success!!"})
}

// Delete handles file delete request.
func Delete(c *gin.Context) {
	uuid := c.Param("uuid")
	targets, err := filepath.Glob(fmt.Sprintf("./images/%v.*", uuid))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	for _, target := range targets {
		err := os.Remove(target)
		if err != nil {
			fmt.Println(err)
		}
	}
	c.Status(http.StatusNoContent)
}

// File has file's info.
type File struct {
	Path string `json:"path"`
	Size int64  `json:"size"`
}

func dirwalk(dir string) (files []File, err error) {
	files = make([]File, 0)

	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if strings.Contains(info.Name(), "DS_Store") {
			return nil
		}
		path = strings.Replace(path, "images/", "http://localhost:8888/files/", 1)
		size := info.Size()
		f := File{Path: path, Size: size}
		files = append(files, f)
		return nil
	})
	if err != nil {
		return
	}

	return
}

// List handles listing files request.
func List(c *gin.Context) {
	files, err := dirwalk("./images")
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, files)
}
