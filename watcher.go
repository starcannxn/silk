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
	fmt.Println("Started watching for track changes...")

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
		// No track playing or error - that's okay, just skip
		return
	}

	// Check if track has changed
	if w.hasTrackChanged(track) {
		fmt.Printf("Track changed: %s - %s\n", track.Artist, track.Title)

		// Save to files
		if err := SaveTrackToFile(track); err != nil {
			fmt.Printf("Error saving track: %v\n", err)
		} else {
			fmt.Printf("Updated files: %s - %s\n", track.Artist, track.Title)
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
