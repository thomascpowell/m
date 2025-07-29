package scripts

/**
* Shared AppleScript wrappers.
 */

import (
	"bytes"
	"os/exec"
	"strings"
)

func Run(command string) (string, error) {
	cmd := exec.Command("osascript", "-e", command)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return strings.TrimRight(out.String(), "\r\n"), nil
}
