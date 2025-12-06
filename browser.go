package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"time"
)

// OpenBrowser opens the default browser to the specified URL
func OpenBrowser(url string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	case "darwin": // macOS
		cmd = exec.Command("open", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	default:
		return fmt.Errorf("unsupported platform")
	}

	return cmd.Start()
}

// OpenBrowserDelayed opens the browser after a short delay (to let server start)
func OpenBrowserDelayed(url string, delay time.Duration) {
	time.Sleep(delay)
	if err := OpenBrowser(url); err != nil {
		fmt.Printf("Note: Could not auto-open browser: %v\n", err)
		fmt.Printf("Please manually open: %s\n", url)
	}
}
