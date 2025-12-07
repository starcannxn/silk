# Silk - Now Playing for OBS

A lightweight app that captures currently playing music metadata and saves it for use in OBS or other streaming software.

## Features

- ✅ **Auto-updates** - Monitors track changes every 3 seconds
- ✅ **Album artwork** - Fetches cover art from Last.fm
- ✅ **Clean web UI** - View what's currently playing in your browser
- ✅ **OBS-ready files** - Outputs to `nowplaying.txt` and `artwork.jpg`
- ✅ **Multiple players** - Supports Spotify, Foobar2000, YouTube, and VLC
<!-- - ✅ **Cross-platform** - Windows (Linux support coming soon) -->

## Quick Start

1. **Download** `silk.exe`
2. **Run** the executable
3. Your browser will open automatically to `http://localhost:5555`
4. Files are saved to the `output/` folder

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

**Last.fm API Key:**  
The app uses Last.fm to fetch album artwork. If you want to use your own API key, edit `artwork.go` and replace the `LASTFM_API_KEY` constant.

## Files & Folders

```
silk/
├── silk.exe           # Main executable
├── web/               # Web UI files
│   ├── index.html
│   └── placeholder.svg
├── output/            # Generated files (auto-created)
│   ├── nowplaying.txt # "Artist - Title"
│   └── artwork.jpg    # Album cover
├── VLC_SETUP.md       # VLC setup instructions
└── README.md          # This file
```

## Building from Source

Requirements: Go 1.21+

```bash
go build -o silk.exe
```

To add a custom icon (Windows):
```bash
go install github.com/akavel/rsrc@latest
rsrc -ico icon.ico -o rsrc.syso
go build -o silk.exe
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

## License

MIT License - feel free to use and modify!

## Credits

Built with Go and powered by:
- Last.fm API for artwork
- Windows Media Control API for metadata
<!-- - MPRIS for Linux support -->