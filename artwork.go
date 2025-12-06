package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

// Last.fm API response
type lastFmResponse struct {
	Track struct {
		Album struct {
			Image []struct {
				Text string `json:"#text"`
				Size string `json:"size"`
			} `json:"image"`
		} `json:"album"`
	} `json:"track"`
}

const LASTFM_API_KEY = "05ccb98f8e36d2b971de7f028ca7dbd7" // Replace this with your actual API key

// FetchArtwork fetches album artwork from Last.fm API
func FetchArtwork(artist, track string) ([]byte, error) {
	if artist == "" || track == "" {
		return nil, fmt.Errorf("artist and track required")
	}

	// Build Last.fm API URL
	baseURL := "http://ws.audioscrobbler.com/2.0/"
	params := url.Values{}
	params.Add("method", "track.getInfo")
	params.Add("api_key", LASTFM_API_KEY)
	params.Add("artist", artist)
	params.Add("track", track)
	params.Add("format", "json")

	apiURL := baseURL + "?" + params.Encode()

	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch track info: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Last.fm returned status %d", resp.StatusCode)
	}

	var result lastFmResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to parse Last.fm response: %v", err)
	}

	// Find the largest available image
	var imageURL string
	for _, img := range result.Track.Album.Image {
		if img.Size == "extralarge" || img.Size == "large" {
			imageURL = img.Text
			break
		}
	}

	// If no large image, try medium or small
	if imageURL == "" {
		for _, img := range result.Track.Album.Image {
			if img.Text != "" {
				imageURL = img.Text
				break
			}
		}
	}

	if imageURL == "" {
		return nil, fmt.Errorf("no artwork found")
	}

	// Download the image
	imgResp, err := client.Get(imageURL)
	if err != nil {
		return nil, fmt.Errorf("failed to download artwork: %v", err)
	}
	defer imgResp.Body.Close()

	if imgResp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to download artwork: status %d", imgResp.StatusCode)
	}

	imgData, err := io.ReadAll(imgResp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read artwork: %v", err)
	}

	return imgData, nil
}
