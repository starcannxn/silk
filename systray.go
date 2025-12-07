package main

import (
	"fmt"
	"os"

	"github.com/getlantern/systray"
)

// StartSystemTray initializes and runs the system tray
func StartSystemTray(openBrowser bool, port string) {
	systray.Run(func() {
		onReady(openBrowser, port)
	}, onExit)
}

func onReady(openBrowser bool, port string) {
	// Load and set icon from file
	iconData, err := os.ReadFile("icon.ico")
	if err == nil {
		systray.SetIcon(iconData)
	}

	systray.SetTitle("Silk")
	systray.SetTooltip("Silk - Now Playing")

	// Menu items
	mOpen := systray.AddMenuItem("Open in Browser", "Open web interface")
	mOutput := systray.AddMenuItem("Open Output Folder", "Open folder with files")
	systray.AddSeparator()
	mStatus := systray.AddMenuItem("Running on localhost:"+port, "Server status")
	mStatus.Disable()
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Quit Silk")

	// Auto-open browser if requested
	if openBrowser {
		go func() {
			OpenBrowserDelayed(fmt.Sprintf("http://localhost:%s", port), 500)
		}()
	}

	// Handle menu clicks
	go func() {
		for {
			select {
			case <-mOpen.ClickedCh:
				OpenBrowser(fmt.Sprintf("http://localhost:%s", port))
			case <-mOutput.ClickedCh:
				openOutputFolder()
			case <-mQuit.ClickedCh:
				systray.Quit()
				os.Exit(0)
			}
		}
	}()
}

func onExit() {
	// Cleanup
}

func openOutputFolder() {
	// Open the output folder in file explorer
	OpenBrowser("output")
}
