package dir

import (
	"fmt"
	"os"
)

func RemoveDir(dirname string) {
	if _, err := os.Stat(dirname); !os.IsNotExist(err) {
		if err := os.RemoveAll(dirname); err != nil {
			fmt.Println("Error: ", err)
		}
	}
}

func removeFile(filename string) {
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		if err := os.Remove(filename); err != nil {
			fmt.Println("Error: ", err)
		}
	}
}

func RemoveAllDir() {
	// Remove the .sgit directory
	if err := os.RemoveAll(".sgit"); err != nil {
		fmt.Println("Error: ", err)
	}
}
