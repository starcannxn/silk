# VLC Setup Instructions

VLC (non-Microsoft Store version) requires a small plugin to work with Silk and Windows media controls.

## Quick Setup:

1. Download the VLC SMTC plugin from: https://github.com/spmn/vlc-win10smtc/releases
2. Extract the ZIP file
3. Copy the `.dll` file to your VLC plugins folder:
   - Usually: `C:\Program Files\VideoLAN\VLC\plugins`
   - Or: `C:\Program Files (x86)\VideoLAN\VLC\plugins`
4. Open VLC and go to: **Tools → Preferences**
5. Click **"Show settings: All"** at the bottom left
6. Navigate to: **Interface → Control interfaces**
7. Check the box for **"Win10 SMTC..."**
8. Click **Save** and restart VLC

That's it! VLC will now work with:
- Silk (showing current track)
- Windows media keyboard controls (play/pause buttons)
- Windows taskbar media controls

## Notes:
- **Microsoft Store VLC**: Works out of the box, no setup needed
- **Regular VLC**: Needs the plugin above (one-time setup)
- **Other players**: Spotify, Foobar2000, and YouTube work with no setup