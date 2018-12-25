package handler

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

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
