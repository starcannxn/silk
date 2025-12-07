package main

import (
	"fmt"
	"time"
)

// Watcher continuously monitors for track changes
type Watcher struct {
	lastTrack *Track
	interval  time.Duration
	stopChan  chan bool
}

// NewWatcher creates a new track watcher
func NewWatcher(interval time.Duration) *Watcher {
	return &Watcher{
		interval: interval,
		stopChan: make(chan bool),
	}
}

// Start begins watching for track changes
func (w *Watcher) Start() {
	fmt.Println()

	ticker := time.NewTicker(w.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			w.checkAndUpdate()
		case <-w.stopChan:
			fmt.Println("Stopped watching")
			return
		}
	}
}

// Stop stops the watcher
func (w *Watcher) Stop() {
	w.stopChan <- true
}

// checkAndUpdate checks for track changes and updates files if needed
func (w *Watcher) checkAndUpdate() {
	track, err := GetCurrentTrack()
	if err != nil {
		// No track playing or error - update files to show disconnected state
		disconnectedTrack := &Track{
			Title:     "No track playing",
			Artist:    "",
			Album:     "",
			Artwork:   "", // Empty means placeholder will be used
			IsPlaying: false,
		}

		// Only update if state changed (avoid constant file writes)
		if w.lastTrack != nil && w.lastTrack.Title != "No track playing" {
			SaveTrackToFile(disconnectedTrack)
			w.lastTrack = disconnectedTrack
		} else if w.lastTrack == nil {
			// First run with no track
			SaveTrackToFile(disconnectedTrack)
			w.lastTrack = disconnectedTrack
		}
		return
	}

	// Check if track has changed
	if w.hasTrackChanged(track) {
		status := "▶"
		if !track.IsPlaying {
			status = "⏸"
		}
		fmt.Printf("%s %s - %s\n", status, track.Artist, track.Title)

		// Save to files
		if err := SaveTrackToFile(track); err != nil {
			fmt.Printf("   ✗ Error saving: %v\n", err)
		}

		w.lastTrack = track
	}
}

// hasTrackChanged checks if the current track is different from the last one
func (w *Watcher) hasTrackChanged(current *Track) bool {
	if w.lastTrack == nil {
		return true
	}

	// Compare title and artist (the key identifiers)
	return current.Title != w.lastTrack.Title ||
		current.Artist != w.lastTrack.Artist
}
