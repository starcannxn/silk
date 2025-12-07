//go:build tray
// +build tray

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	port := "5555"

	// Start the watcher in a goroutine
	watcher := NewWatcher(3 * time.Second)
	go watcher.Start()

	// API endpoint to get current track
	http.HandleFunc("/api/current", handleCurrentTrack)

	// Serve output folder
	http.Handle("/output/", http.StripPrefix("/output/", http.FileServer(http.Dir("./output"))))

	// Serve static files from the web directory
	http.Handle("/", http.FileServer(http.Dir("./web")))

	// Start web server in background
	go func() {
		if err := http.ListenAndServe(":"+port, nil); err != nil {
			log.Fatal(err)
		}
	}()

	// Start system tray (with browser auto-open)
	StartSystemTray(true, port)
}

// handleCurrentTrack returns the currently playing track as JSON
func handleCurrentTrack(w http.ResponseWriter, r *http.Request) {
	track, err := GetCurrentTrack()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Save to file automatically
	if err := SaveTrackToFile(track); err != nil {
		fmt.Printf("Warning: failed to save track to file: %v\n", err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(track)
}
