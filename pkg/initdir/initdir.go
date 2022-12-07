package initdir

import (
	"fmt"
	"os"
)

func initializeDir() {
	// Create the .sgit/objects directory

	if err := os.MkdirAll(".sgit/objects", 0755); err != nil {
		fmt.Println("Error: ", err)
	}

	// Create the sub directories
	if err := os.Mkdir(".sgit/objects/info", 0755); err != nil {
		fmt.Println("Error: ", err)
	}

	if err := os.Mkdir(".sgit/objects/pack", 0755); err != nil {
		fmt.Println("Error: ", err)
	}
}

func removeDir() {
	// Remove the .sgit directory
	if err := os.RemoveAll(".sgit"); err != nil {
		fmt.Println("Error: ", err)
	}
}
