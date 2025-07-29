package utils

import (
	"os"
)

/**
* Error logging method.
 */

// if this is true, a log file (m.txt) will be created in the event of a (handled) error
// disabled for distribution
const ENABLE_LOGS = false

func Log(msg string) {
	if !ENABLE_LOGS {
		return
	}
	f, _ := os.OpenFile("m.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	divider := "\n(MESSAGE) "
	_, _ = f.WriteString(divider + msg + "\n")
}
