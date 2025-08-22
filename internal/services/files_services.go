package services

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
)

const BASE_PATH = "DiscoveryDrawings" // Adjust this path as needed

// Add a File struct for better type safety
type FileInfo struct {
	Name            string `json:"name"`
	Type            string `json:"type"`
	Size            int64  `json:"size"`
	FileCount       int64  `json:"fileCount"`
	Modified        string `json:"modified"`
	Extension       string `json:"extension"`
	IsDownloadable  bool   `json:"isDownloadable"`
	ParentDirectory string `json:"parentDirectory"`
	FullPath        string `json:"fullPath"`
}

type YearMakeStructure struct {
	YearMake   string   `json:"yearMake"`
	Models     []string `json:"models"`
	ModelCount int      `json:"modelCount"`
	Error      string   `json:"error,omitempty"`
}

type DirectoryStructure struct {
	TotalYearMakes int                 `json:"totalYearMakes"`
	TotalModels    int                 `json:"totalModels"`
	Structure      []YearMakeStructure `json:"structure"`
}

// Add a helper function to get the absolute base path
func getAbsBasePath() (string, error) {
	return filepath.Abs(BASE_PATH)
}

// GetFileStructure returns the directory structure
func GetFileStructure() (*DirectoryStructure, error) {
	// Read the base directory
	entries, err := os.ReadDir(BASE_PATH)
	if err != nil {
		return nil, fmt.Errorf("failed to read base directory: %v", err)
	}

	// Filter for directories only and get names
	var yearMakes []string
	for _, entry := range entries {
		if entry.IsDir() {
			yearMakes = append(yearMakes, entry.Name())
		}
	}

	// Sort the yearMakes
	sort.Strings(yearMakes)

	// Build complete structure with models for each yearMake
	var structure []YearMakeStructure
	totalModels := 0

	for _, yearMake := range yearMakes {
		yearMakeStruct := processYearMake(yearMake)
		structure = append(structure, yearMakeStruct)
		totalModels += yearMakeStruct.ModelCount
	}

	return &DirectoryStructure{
		TotalYearMakes: len(structure),
		TotalModels:    totalModels,
		Structure:      structure,
	}, nil
}

func processYearMake(yearMake string) YearMakeStructure {
	yearMakePath := filepath.Join(BASE_PATH, yearMake)

	modelEntries, err := os.ReadDir(yearMakePath)
	if err != nil {
		fmt.Printf("Error reading models for %s: %v\n", yearMake, err)
		return YearMakeStructure{
			YearMake:   yearMake,
			Models:     []string{},
			ModelCount: 0,
			Error:      "Failed to read models",
		}
	}

	var models []string
	for _, entry := range modelEntries {
		if entry.IsDir() {
			models = append(models, entry.Name())
		}
	}

	sort.Strings(models)

	return YearMakeStructure{
		YearMake:   yearMake,
		Models:     models,
		ModelCount: len(models),
	}
}

// GetYearMakes returns available year/make combinations
func GetYearMakes() ([]string, error) {
	entries, err := os.ReadDir(BASE_PATH)
	if err != nil {
		return nil, err
	}

	var yearMakes []string
	for _, entry := range entries {
		if entry.IsDir() {
			yearMakes = append(yearMakes, entry.Name())
		}
	}

	sort.Strings(yearMakes)
	return yearMakes, nil
}

// GetModels returns available models for a year/make
func GetModels(yearMake string) ([]string, error) {
	yearMakePath := filepath.Join(BASE_PATH, yearMake)

	// Validate path safety
	absBasePath, err := getAbsBasePath()
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute base path: %v", err)
	}

	absYearMakePath, err := filepath.Abs(yearMakePath)
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path: %v", err)
	}

	if !strings.HasPrefix(absYearMakePath, absBasePath) {
		return nil, fmt.Errorf("invalid path")
	}

	entries, err := os.ReadDir(yearMakePath)
	if err != nil {
		return nil, err
	}

	var models []string
	for _, entry := range entries {
		if entry.IsDir() {
			models = append(models, entry.Name())
		}
	}

	sort.Strings(models)
	return models, nil
}

// calculateDirectorySize uses goroutines to calculate directory size in parallel
func calculateDirectorySize(dirPath string) (int64, int64) {
	var totalSize int64
	var fileCount int64
	var wg sync.WaitGroup

	// Get entries in the directory
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return 0, 0
	}

	// Process each top-level entry in a separate goroutine
	for _, entry := range entries {
		wg.Add(1)
		go func(e os.DirEntry) {
			defer wg.Done()

			entryPath := filepath.Join(dirPath, e.Name())

			if e.IsDir() {
				// For directories, walk them
				err := filepath.Walk(entryPath, func(path string, info os.FileInfo, err error) error {
					if err != nil {
						return nil // Skip files we can't access
					}
					if !info.IsDir() {
						atomic.AddInt64(&totalSize, info.Size())
						atomic.AddInt64(&fileCount, 1)
					}
					return nil
				})

				if err != nil {
					// Just continue if there's an error
				}
			} else {
				// For files, get size directly
				info, err := e.Info()
				if err != nil {
					return
				}
				atomic.AddInt64(&totalSize, info.Size())
				atomic.AddInt64(&fileCount, 1)
			}
		}(entry)
	}

	wg.Wait()
	return totalSize, fileCount
}

// GetFiles returns files AND directories for a specific year/make/model
func GetFiles(yearMake, model string) ([]FileInfo, error) {
	basePath := filepath.Join(BASE_PATH, yearMake, model)

	// Validate path safety
	_, err := getAbsBasePath()
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute base path: %v", err)
	}

	absModelPath, err := filepath.Abs(basePath)
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path: %v", err)
	}

	absBaseDir, err := filepath.Abs(BASE_PATH)
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute base directory: %v", err)
	}

	if !strings.HasPrefix(absModelPath, absBaseDir) {
		return nil, fmt.Errorf("invalid path")
	}

	var files []FileInfo

	err = filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip the root directory itself
		if path == basePath {
			return nil
		}

		relPath, _ := filepath.Rel(basePath, path)

		// Get the parent directory relative to the model directory
		parentDir := filepath.Dir(relPath)
		if parentDir == "." {
			parentDir = "" // Root of model directory
		}

		if info.IsDir() {
			dirSize, fileCount := calculateDirectorySize(path)
			// Handle directories
			files = append(files, FileInfo{
				Name:            info.Name(),
				Type:            "directory",
				Size:            dirSize,
				FileCount:       fileCount,
				Modified:        info.ModTime().Format("2006-01-02T15:04:05Z07:00"),
				Extension:       "", // Directories don't have extensions
				IsDownloadable:  dirSize > 0,
				ParentDirectory: parentDir,
				FullPath:        relPath,
			})
		} else {
			// Handle files
			ext := strings.ToLower(filepath.Ext(info.Name()))

			files = append(files, FileInfo{
				Name:            info.Name(),
				Type:            "file",
				Size:            info.Size(),
				FileCount:       1, // Each file is counted as one
				Modified:        info.ModTime().Format("2006-01-02T15:04:05Z07:00"),
				Extension:       ext,
				IsDownloadable:  true,
				ParentDirectory: parentDir,
				FullPath:        relPath,
			})
		}
		return nil
	})

	// Sort results: directories first, then files, both alphabetically
	sort.Slice(files, func(i, j int) bool {
		if files[i].Type != files[j].Type {
			return files[i].Type == "directory" // directories come first
		}
		return files[i].Name < files[j].Name
	})

	return files, err
}

func GetFilePath(yearMake, model, fileName, parentDir string) string {
	if parentDir != "" {
		return filepath.Join(BASE_PATH, yearMake, model, parentDir, fileName)
	}
	return filepath.Join(BASE_PATH, yearMake, model, fileName)
}

func GetDirPath(yearMake, model, dir string) string {
	return filepath.Join(BASE_PATH, yearMake, model, dir)
}

// CreateZipFromDirectory creates a zip file from a directory
func CreateZipFromDirectory(dirPath string) (string, error) {
	zipPath := dirPath + ".zip"
	zipFile, err := os.Create(zipPath)
	if err != nil {
		return "", err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	err = filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		relPath, _ := filepath.Rel(dirPath, path)
		zipEntry, err := zipWriter.Create(relPath)
		if err != nil {
			return err
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(zipEntry, file)
		return err
	})

	return zipPath, err
}

// CreateZipFromFiles creates a zip file from selected files
func CreateZipFromFiles(yearMake, model string, fileNames []string) (string, error) {
	basePath := filepath.Join(BASE_PATH, yearMake, model)
	zipPath := fmt.Sprintf("/tmp/%s_%s_selected.zip", yearMake, model)

	zipFile, err := os.Create(zipPath)
	if err != nil {
		return "", err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for _, fileName := range fileNames {
		parts := strings.Split(fileName, "__")
		fileBase := parts[1]
		dir := parts[0]
		filePath := filepath.Join(basePath, dir, fileBase)
		f, err := os.Open(filePath)
		if err != nil {
			continue // Skip missing files
		}
		defer f.Close()

		zipEntry, err := zipWriter.Create(fileName)
		if err != nil {
			continue
		}

		_, err = io.Copy(zipEntry, f)
		if err != nil {
			continue
		}
	}

	return zipPath, nil
}
