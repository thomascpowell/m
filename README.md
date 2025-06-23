## Control Music.app from the terminal. 

### Features:
- Browse your existing Apple Music library
- Supports album, playlist, and song playback
- Quick actions with CLI args (try `m help`)

### Benefits:
- Faster (keyboard) navigation
- Supports small/narrow window sizes
- Far cooler (if you like the terminal)

### Technical Details:
- Communicates with Music.app via AppleScript
- Core functionality written in pure Go
- Caches library data locally for fast startup

### Installation:
- Have `git` and `go` installed
- Make sure `$HOME/.local/bin` is in your `$PATH`
```zsh
# Clone and enter the repo
git clone https://github.com/thomascpowell/m.git && cd m

# Build and install
go build -o m && install m "$HOME/.local/bin/m"

# Optional: Remove the repo
cd .. && rm -rf m

# Launch the TUI
m
```

### Notes:
- This is a side project. Please only use it if you are comfortable in the terminal.
- A temporary playlist is created in Apple Music when playing an album.
