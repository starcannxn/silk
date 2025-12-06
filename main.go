package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// API endpoint to get current track
	http.HandleFunc("/api/current", handleCurrentTrack)

	// Serve static files from the web directory
	http.Handle("/", http.FileServer(http.Dir("./web")))

	port := "8080"
	fmt.Printf("Server starting on http://localhost:%s\n", port)

	// Start the web server
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
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
		// Log error but don't fail the request
		fmt.Printf("Warning: failed to save track to file: %v\n", err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(track)
}
