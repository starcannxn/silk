package main

import (
	"fmt"
	"io"
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

	// Save artwork if available, otherwise use placeholder
	artworkPath := filepath.Join(outputDir, "artwork.jpg")
	if track.Artwork != "" {
		if err := saveArtwork(track.Artwork, outputDir); err != nil {
			fmt.Printf("Warning: failed to save artwork: %v\n", err)
		}
	} else {
		// Copy placeholder when no artwork
		if err := copyPlaceholder(artworkPath); err != nil {
			fmt.Printf("Warning: failed to copy placeholder: %v\n", err)
		}
	}

	return nil
}

// saveArtwork saves the artwork image to output/artwork.jpg
func saveArtwork(artworkData string, outputDir string) error {
	if artworkData == "" {
		return nil // No artwork to save
	}

	// artworkData is raw binary data
	artworkPath := filepath.Join(outputDir, "artwork.jpg")
	if err := os.WriteFile(artworkPath, []byte(artworkData), 0644); err != nil {
		return fmt.Errorf("failed to write artwork: %v", err)
	}

	return nil
}

// copyPlaceholder copies placeholder.jpg to the output folder
func copyPlaceholder(destPath string) error {
	// Open source file
	srcFile, err := os.Open("placeholder.jpg")
	if err != nil {
		return fmt.Errorf("placeholder.jpg not found: %v", err)
	}
	defer srcFile.Close()

	// Create destination file
	destFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	// Copy contents
	_, err = io.Copy(destFile, srcFile)
	return err
}
