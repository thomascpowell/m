package scripts

import (
	"fmt"
)

/**
* Wrapper for use in CLI commands.
*/

// function standardizes CLI mode output
// progressive (tense): string for ongoing action, e.g. pausing
// success: success message
func RunAsCli(progressive string, success string, fn func() error) error {
	fmt.Printf("%s... ", progressive)
	err := fn()
	if err != nil {
		fmt.Printf("✗")
	} else {
		fmt.Printf("%s", success)
	}
	fmt.Println()
	return err
}
