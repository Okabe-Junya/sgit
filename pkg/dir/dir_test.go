package dir

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Call the function we are testing
	InitializeDir()

	// Run the tests
	m.Run()

	// Remove the .sgit directory
	RemoveAllDir()
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

func TestCreateDir(t *testing.T) {
	// Create a new directory
	createDir(".sgit/test")

	// Check if the .sgit/test directory exists
	_, err := os.Stat(".sgit/test")
	if err != nil {
		t.Error("Error: ", err)
	}
	removeDir(".sgit/test")
}
