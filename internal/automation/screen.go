package automation

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/go-vgo/robotgo"
)

// CaptureScreenshot captures the full screen and saves it to a temporary location
// Returns the path to the saved screenshot file
func CaptureScreenshot() (string, error) {
	// Capture the full screen
	img := robotgo.CaptureImg()
	if img == nil {
		return "", fmt.Errorf("failed to capture screen: image is nil")
	}

	// Generate unique filename with timestamp
	timestamp := time.Now().Format("20060102_150405")
	filename := fmt.Sprintf("screenshot_%s.png", timestamp)

	// Get temporary directory
	tempDir := os.TempDir()
	filePath := filepath.Join(tempDir, filename)

	// Save the screenshot
	err := robotgo.Save(img, filePath)
	if err != nil {
		return "", fmt.Errorf("failed to save screenshot to %s: %w", filePath, err)
	}

	return filePath, nil
}

// CaptureScreenshotToPath captures the full screen and saves it to the specified path
// Returns the path to the saved screenshot file
func CaptureScreenshotToPath(outputPath string) error {
	// Validate output path
	if outputPath == "" {
		return fmt.Errorf("output path cannot be empty")
	}

	// Ensure directory exists
	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dir, err)
	}

	// Capture the full screen
	img := robotgo.CaptureImg()
	if img == nil {
		return fmt.Errorf("failed to capture screen: image is nil")
	}

	// Save the screenshot
	err := robotgo.Save(img, outputPath)
	if err != nil {
		return fmt.Errorf("failed to save screenshot to %s: %w", outputPath, err)
	}

	return nil
}

// GetScreenSize returns the screen dimensions
func GetScreenSize() (width, height int) {
	return robotgo.GetScreenSize()
}
