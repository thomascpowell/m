## A CLI application for controlling Music.app on macOS. 

### Features:
- TUI for browsing music library
- Playback controls (play, pause, next)
- Soon: CLI arguments for quick actions

### Benefits (vs native app):
- Faster (keyboard based) navigation
- Supports small/narrow window sizes

### Technical Details:
- Communicates with music.app via AppleScript
- Caches library data locally for fast startup

### Installation:
- Have `git` and `go` installed
- Make sure `$HOME/.local/bin` is in your `$PATH`
```zsh
# Clone and enter the repo
git clone https://github.com/thomascpowell/m.git && cd m

# Build and install
go build -o m && install m "$HOME/.local/bin/m"

# Run the program
m

# Optional: Remove the repo
cd .. && rm -rf m
```

### Notes:
- This is a side project. Please only use it if you are comfortable in the terminal.
- A temporary playlist is created in Apple Music when playing an album. This is due to limitations inherent to AppleScript. It is safe to delete.

