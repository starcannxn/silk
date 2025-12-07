package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	// Start the watcher in a goroutine (runs in background)
	watcher := NewWatcher(3 * time.Second) // Check every 3 seconds
	go watcher.Start()

	// API endpoint to get current track
	http.HandleFunc("/api/current", handleCurrentTrack)

	// Serve output folder (for artwork and text files)
	http.Handle("/output/", http.StripPrefix("/output/", http.FileServer(http.Dir("./output"))))

	// Serve static files from the web directory
	http.Handle("/", http.FileServer(http.Dir("./web")))

	port := "5555"
	url := fmt.Sprintf("http://localhost:%s", port)

	// Nice startup message
	fmt.Println("╔════════════════════════════════════════╗")
	fmt.Println("║          Silk - Now Playing            ║")
	fmt.Println("╚════════════════════════════════════════╝")
	fmt.Println()
	fmt.Printf("→ Server running at %s\n", url)
	fmt.Println("→ Auto-updating every 3 seconds")
	fmt.Println("→ Files saved to './output' folder")
	fmt.Println()
	fmt.Println("Press Ctrl+C to stop")
	fmt.Println("────────────────────────────────────────")

	// Auto-open browser after a short delay
	go OpenBrowserDelayed(url, 500*time.Millisecond)

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
