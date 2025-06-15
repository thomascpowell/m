## A CLI application for controlling Music.app on macOS. 

### Features:
- Clean TUI for browsing music library
- Playback controls (play, pause, next)
- Soon: CLI arguments for quick actions

### Benefits (compared to native app):
- Faster, keyboard based navigation
- Supports all window sizes
- Far cooler

### Technical Details:
- Written in Go with an MVC architecture
- Uses AppleScript to communicate with music.app
- Caches library data locally for fast startup

### Installation:
- Have `git` and `go` installed
- Make sure `$HOME/.local/bin` is in your `$PATH`

### Usage:
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
- Due to AppleScript limitations, this project creates a temporary playlist in Apple Music.

