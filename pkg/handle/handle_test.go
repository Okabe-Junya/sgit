package handle

import (
	"os"
	"testing"
)

func init() {
	// Set the command line arguments
	os.Args = []string{"sgit", "add", "file1", "file2"}
}

func TestParseArgs(t *testing.T) {
	// Parse the command line arguments
	parseArgs()
}
