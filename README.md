# Silk - Now Playing for OBS

A lightweight app that captures currently playing music metadata and saves it for use in OBS or other streaming software.

## Features

- ✅ **Auto-updates** - Monitors track changes every 3 seconds
- ✅ **Album artwork** - Fetches cover art from Last.fm
- ✅ **Clean web UI** - View what's currently playing in your browser
- ✅ **OBS-ready files** - Outputs to `nowplaying.txt` and `artwork.jpg`
- ✅ **Multiple players** - Supports Spotify, Foobar2000, YouTube, and VLC
- ✅ **Multiple modes** - Console, system tray, or background operation
- ⌚ **Linux Support** - Coming soon

## Download & Install

Three versions available:

| Version | Description | Use Case |
|---------|-------------|----------|
| **silk.exe** | Console + Browser | Debugging, seeing what's happening |
| **silk-tray.exe** | System Tray + Browser | Clean UI, auto-opens browser |
| **silk-tray-only.exe** | System Tray only | Background operation, no browser |

**Requirements:**
- The `web/` folder must be in the same directory as the executable
- Your `icon.ico` should be in the same directory (for tray versions)

## Quick Start

1. **Download** your preferred version
2. **Run** the executable
3. Files are saved to the `output/` folder

**For tray versions:**
- Right-click the tray icon for options
- "Open in Browser" to view the web UI at `http://localhost:5555`
- "Open Output Folder" to see generated files
- "Quit" to stop the app

## Using in OBS

Add a **Text (GDI+)** source:
- Check "Read from file"
- Browse to: `output/nowplaying.txt`

Add an **Image** source:
- Browse to: `output/artwork.jpg`

Both will update automatically as songs change!

## Supported Players

| Player | Setup Required |
|--------|----------------|
| Spotify | ✅ Works out of the box |
| Foobar2000 | ✅ Works out of the box |
| YouTube (Browser) | ✅ Works out of the box |
| VLC | ⚠️ Requires plugin ([see VLC_SETUP.md](VLC_SETUP.md)) |

## Configuration

**Port:**  
Default port is `5555`. Change it in the source code (`main.go`, `main_tray.go`, `main_tray_only.go`) and rebuild.

**Last.fm API Key:**  
A Last.fm API key is included for convenience. The app works out of the box!  
If you want to use your own key (optional), edit `artwork.go` and replace the `LASTFM_API_KEY` constant.

## Files & Folders

```
silk/
├── silk.exe               # Console version
├── silk-tray.exe          # Tray + browser version
├── silk-tray-only.exe     # Tray only version
├── icon.ico               # Tray icon (required for tray versions)
├── web/                   # Web UI files (required)
│   ├── index.html
│   └── placeholder.svg
├── output/                # Generated files (auto-created)
│   ├── nowplaying.txt     # "Artist - Title"
│   └── artwork.jpg        # Album cover
├── VLC_SETUP.md           # VLC setup instructions
└── README.md              # This file
```

## Building from Source

Requirements: Go 1.21+

**Build all versions:**
```bash
.\build.bat
```

Or build individually:
```bash
# Console version
go build -o silk.exe

# Tray + browser version
go build -tags tray -ldflags -H=windowsgui -o silk-tray.exe

# Tray only version
go build -tags trayonly -ldflags -H=windowsgui -o silk-tray-only.exe
```

**To add a custom icon (Windows):**
```bash
go install github.com/akavel/rsrc@latest
rsrc -ico icon.ico -o rsrc.syso
# Then build normally
```

## Troubleshooting

**No track detected:**
- Make sure music is actually playing
- Check if your player is in the supported list
- For VLC, see [VLC_SETUP.md](VLC_SETUP.md)

**No artwork:**
- Artwork is fetched from Last.fm
- May not be available for all tracks
- A placeholder will be shown if artwork isn't found

**Connection issues:**
- Check the connection indicator (top-right of web UI)
- Make sure no other app is using port 5555
- For tray versions, right-click icon and select "Open in Browser"

**Tray icon not showing:**
- Make sure `icon.ico` is in the same folder as the executable
- Check Windows system tray settings (hidden icons)

## License

MIT License - feel free to use and modify!

## Credits

Built with Go and powered by:
- Last.fm API for artwork
- Windows Media Control API for metadata
- [systray](https://github.com/getlantern/systray) for system tray functionality
<!-- - MPRIS for Linux support -->