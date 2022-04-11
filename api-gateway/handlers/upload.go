package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
)

func SaveAndResizeImage(c *gin.Context) {
	srcFilePath, err := SaveImage(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusCreated, srcFilePath)
	go ResizeImage(srcFilePath)
}

func SaveImage(c *gin.Context) (string, error) {
	target := c.Request.FormValue("target")
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		return "", err
	}
	filename := header.Filename
	ext := filepath.Ext(filename)
	targetFile := fmt.Sprintf("%s%s%s", "../realty", target, ext)
	_, err = os.Stat(filepath.Dir(targetFile))
	if os.IsNotExist(err) {
		err = os.MkdirAll(filepath.Dir(targetFile), 0644)
		if err != nil {
			return "", err
		}
	}

	out, err := os.Create(targetFile)
	if err != nil {
		return "", err
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		return "", err
	}
	return targetFile, nil
}

func ResizeImage(srcFilePath string) error {
	srcImage, err := imaging.Open(srcFilePath, imaging.AutoOrientation(true))
	if err != nil {
		return err
	}

	extension := filepath.Ext(srcFilePath)
	filepath128 := getDestFilePath(srcFilePath, "s", extension)
	dstImage128 := imaging.Resize(srcImage, 128, 128, imaging.CatmullRom)
	err = imaging.Save(dstImage128, filepath128)
	if err != nil {
		return err
	}

	filepath400 := getDestFilePath(srcFilePath, "m", extension)
	dstImage400 := imaging.Resize(srcImage, 400, 300, imaging.Linear)
	err = imaging.Save(dstImage400, filepath400)
	if err != nil {
		return err
	}

	filepath1200 := getDestFilePath(srcFilePath, "x", extension)
	dstImage1200 := imaging.Resize(srcImage, 1200, 800, imaging.Linear)
	err = imaging.Save(dstImage1200, filepath1200)
	if err != nil {
		return err
	}

	return nil
}

func getDestFilePath(srcFilePath string, size string, extension string) string {
	return fmt.Sprintf("%s_%s%s", fileNameWithoutExtension(srcFilePath), size, extension)
}

func fileNameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}
