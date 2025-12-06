package main

import (
	"fmt"
	"runtime"
)

// Track represents the currently playing track
type Track struct {
	Title     string `json:"title"`
	Artist    string `json:"artist"`
	Album     string `json:"album"`
	Artwork   string `json:"artwork"`
	IsPlaying bool   `json:"isPlaying"`
}

// GetCurrentTrack returns the currently playing track info
func GetCurrentTrack() (*Track, error) {
	// We'll implement platform-specific code here
	if runtime.GOOS == "windows" {
		return getCurrentTrackWindows()
	} else if runtime.GOOS == "linux" {
		return getCurrentTrackLinux()
	}

	return nil, fmt.Errorf("unsupported platform: %s", runtime.GOOS)
}

// getCurrentTrackWindows reads metadata from Windows Media Control API
func getCurrentTrackWindows() (*Track, error) {
	// Placeholder for now - we'll implement this next
	return &Track{
		Title:     "Test Song",
		Artist:    "Test Artist",
		Album:     "Test Album",
		Artwork:   "",
		IsPlaying: true,
	}, nil
}

// getCurrentTrackLinux reads metadata from MPRIS (Linux)
func getCurrentTrackLinux() (*Track, error) {
	// Placeholder - we'll implement this later
	return nil, fmt.Errorf("Linux support coming soon")
}
