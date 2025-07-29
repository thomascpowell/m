package utils

import (
	"os"
)

/**
* Contains utility functions.
 */

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func GetHomePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return home
}

func GetGlobalCachePath() string {
	return GetHomePath() + "/.m.gob"
}
