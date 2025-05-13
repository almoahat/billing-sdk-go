package sdk

import (
	"io"
	"log"
)

// ConfigureLogger sets log flags and output destination.
// Use nil for output to leave it unchanged.
func ConfigureLogger(out io.Writer, disableTimestamps bool) {
	if disableTimestamps {
		log.SetFlags(0)
	} else {
		log.SetFlags(log.LstdFlags)
	}

	if out != nil {
		log.SetOutput(out)
	}
}
