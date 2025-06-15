## A CLI application for controlling Music.app on macOS. 

### Features:
- Clean TUI for browsing music library
- Playback controls (play, pause, next)
- Soon: CLI arguments for quick actions

Benefits (compared to native app):
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
1. Clone and enter the repo: 
    ```zsh
    git clone https://github.com/thomascpowell/m.git && cd m
    ```
2. Build and install: 
    ```zsh
    go build -o m && install m "$HOME/.local/bin/m"
    ```
3. Run it: 
    ```zsh
    m
    ```


