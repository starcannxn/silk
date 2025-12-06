package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// SaveTrackToFile writes the current track info to output/nowplaying.txt
func SaveTrackToFile(track *Track) error {
	// Create output directory if it doesn't exist
	outputDir := "output"
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %v", err)
	}

	// Format: "Title - Artist"
	content := fmt.Sprintf("%s - %s", track.Title, track.Artist)

	// Write to file
	filePath := filepath.Join(outputDir, "nowplaying.txt")
	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}

	return nil
}
