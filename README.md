## Control Music.app from the terminal. 

### Features:
- Browse your existing music library from the terminal
- Supports album, playlist, and song playback
- Full playback controls (play, pause, skip)
- Quick actions with CLI args (try `m help`)

### Benefits:
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

# Optional: Remove the repo
cd .. && rm -rf m

# Run the program
m
```

### Notes:
- This is a side project. Please only use it if you are comfortable in the terminal.
- A temporary playlist is created in Apple Music when playing an album. This is due to limitations inherent to AppleScript. It is safe to delete.

