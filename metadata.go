package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"syscall"
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
	// PowerShell command to get current media session info
	psCommand := `Add-Type -AssemblyName System.Runtime.WindowsRuntime; $asTaskGeneric = ([System.WindowsRuntimeSystemExtensions].GetMethods() | Where-Object { $_.Name -eq 'AsTask' -and $_.GetParameters().Count -eq 1 -and $_.GetParameters()[0].ParameterType.Name -eq 'IAsyncOperation' + [char]0x0060 + '1' })[0]; Function Await($WinRtTask, $ResultType) { $asTask = $asTaskGeneric.MakeGenericMethod($ResultType); $netTask = $asTask.Invoke($null, @($WinRtTask)); $netTask.Wait(-1) | Out-Null; $netTask.Result }; [Windows.Media.Control.GlobalSystemMediaTransportControlsSessionManager,Windows.Media.Control,ContentType=WindowsRuntime] | Out-Null; $sessionManager = Await ([Windows.Media.Control.GlobalSystemMediaTransportControlsSessionManager]::RequestAsync()) ([Windows.Media.Control.GlobalSystemMediaTransportControlsSessionManager]); $session = $sessionManager.GetCurrentSession(); if ($session) { $mediaProperties = Await ($session.TryGetMediaPropertiesAsync()) ([Windows.Media.Control.GlobalSystemMediaTransportControlsSessionMediaProperties]); $playbackInfo = $session.GetPlaybackInfo(); $output = @{ Title = $mediaProperties.Title; Artist = $mediaProperties.Artist; Album = $mediaProperties.AlbumTitle; IsPlaying = ($playbackInfo.PlaybackStatus -eq 4) }; [Console]::OutputEncoding = [System.Text.Encoding]::UTF8; $output | ConvertTo-Json -Compress }`

	// Execute PowerShell command with UTF-8 output encoding
	cmd := exec.Command("powershell", "-NoProfile", "-NonInteractive", "-OutputFormat", "Text", "-Command", psCommand)
	cmd.Env = append(os.Environ(), "POWERSHELL_OUTPUT_ENCODING=utf-8")

	// Hide the PowerShell window (important for tray versions)
	if runtime.GOOS == "windows" {
		cmd.SysProcAttr = &syscall.SysProcAttr{
			HideWindow:    true,
			CreationFlags: 0x08000000, // CREATE_NO_WINDOW
		}
	}

	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to get media info: %v", err)
	}

	// Parse JSON output
	var result struct {
		Title     string `json:"Title"`
		Artist    string `json:"Artist"`
		Album     string `json:"Album"`
		IsPlaying bool   `json:"IsPlaying"`
	}

	if err := json.Unmarshal(output, &result); err != nil {
		return nil, fmt.Errorf("failed to parse media info: %v", err)
	}

	track := &Track{
		Title:     result.Title,
		Artist:    result.Artist,
		Album:     result.Album,
		Artwork:   "",
		IsPlaying: result.IsPlaying,
	}

	// Try to fetch artwork from Last.fm (don't fail if it doesn't work)
	if result.Artist != "" && result.Title != "" {
		artworkData, err := FetchArtwork(result.Artist, result.Title)
		if err == nil && len(artworkData) > 0 {
			track.Artwork = string(artworkData)
		}
	}

	return track, nil
}

// getCurrentTrackLinux reads metadata from MPRIS (Linux)
func getCurrentTrackLinux() (*Track, error) {
	// Placeholder - we'll implement this later
	return nil, fmt.Errorf("Linux support coming soon")
}
