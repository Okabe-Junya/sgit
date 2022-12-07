package initdir

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Call the function we are testing
	initializeDir()

	// Run the tests
	m.Run()

	// Remove the .sgit directory
	removeDir()
}

func TestInitializeDir(t *testing.T) {
	// Check if the .sgit/objects directory exists
	_, err := os.Stat(".sgit/objects")
	if err != nil {
		t.Error("Error: ", err)
	}

	// Check if the .sgit/objects/info directory exists
	_, err = os.Stat(".sgit/objects/info")
	if err != nil {
		t.Error("Error: ", err)
	}

	// Check if the .sgit/objects/pack directory exists
	_, err = os.Stat(".sgit/objects/pack")
	if err != nil {
		t.Error("Error: ", err)
	}
}
