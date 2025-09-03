package handlers

import (
	"log"
	"net/http"
	"os"

	"goserver/internal/services"

	"github.com/gin-gonic/gin"
)

type FileHandler struct{}

func NewFileHandler() *FileHandler {
	return &FileHandler{}
}

// GET /api/v1/files/structure
func (h *FileHandler) GetStructure(c *gin.Context) {
	structure, err := services.GetFileStructure()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, structure)
}

// GET /api/v1/files
func (h *FileHandler) GetYearMakes(c *gin.Context) {
	yearMakes, err := services.GetYearMakes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, yearMakes)
}

// GET /api/v1/files/:yearMake
func (h *FileHandler) GetModels(c *gin.Context) {
	yearMake := c.Param("yearMake")
	models, err := services.GetModels(yearMake)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, models)
}

// GET /api/v1/files/:yearMake/:model
func (h *FileHandler) GetFiles(c *gin.Context) {
	yearMake := c.Param("yearMake")
	model := c.Param("model")
	files, err := services.GetFiles(yearMake, model)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, files)
}

func (h *FileHandler) DownloadFile(c *gin.Context) {
	yearMake := c.Param("yearMake")
	model := c.Param("model")
	fileName := c.Param("fileName")
	parentDir := c.Param("parentDir") // Will be empty string if not in URL path

	var filePath string
	if parentDir != "" {
		// Route: /:yearMake/:model/download/:parentDir/:fileName
		filePath = services.GetFilePath(yearMake, model, fileName, parentDir)
	} else {
		// Route: /:yearMake/:model/download/:fileName
		filePath = services.GetFilePath(yearMake, model, fileName, "")
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.File(filePath)
}

// GET /api/v1/files/:yearMake/:model/download-directory/:dirName
func (h *FileHandler) DownloadDirectory(c *gin.Context) {
	yearMake := c.Param("yearMake")
	model := c.Param("model")
	dirName := c.Param("dirName")

	dirPath := services.GetDirPath(yearMake, model, dirName)
	log.Printf("Downloading directory: %s", dirPath)
	zipPath, err := services.CreateZipFromDirectory(dirPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer os.Remove(zipPath) // Clean up temp file

	c.Header("Content-Disposition", "attachment; filename="+dirName+".zip")
	c.File(zipPath)
}

// GET /api/v1/files/:yearMake/:model/download-all
func (h *FileHandler) DownloadAll(c *gin.Context) {
	yearMake := c.Param("yearMake")
	model := c.Param("model")

	dirPath := services.GetDirPath(yearMake, model, "")
	zipPath, err := services.CreateZipFromDirectory(dirPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer os.Remove(zipPath) // Clean up temp file

	zipName := yearMake + "_" + model + "_all.zip"
	c.Header("Content-Disposition", "attachment; filename="+zipName)
	c.File(zipPath)
}

// POST /api/v1/files/:yearMake/:model/download-selected
func (h *FileHandler) DownloadSelected(c *gin.Context) {
	yearMake := c.Param("yearMake")
	model := c.Param("model")

	var req struct {
		FileNames []string `json:"fileNames"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	zipPath, err := services.CreateZipFromFiles(yearMake, model, req.FileNames)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer os.Remove(zipPath) // Clean up temp file

	zipName := yearMake + "_" + model + "_selected.zip"
	c.Header("Content-Disposition", "attachment; filename="+zipName)
	c.File(zipPath)
}
