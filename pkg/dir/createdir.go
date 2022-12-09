package dir

import (
	"fmt"
	"os"
)

func CreateDir(dirname string) {
	if _, err := os.Stat(dirname); os.IsNotExist(err) {
		if err := os.MkdirAll(dirname, 0755); err != nil {
			fmt.Println("Error: ", err)
		}
	}
}

func CreateFile(filename string) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		if _, err := os.Create(filename); err != nil {
			fmt.Println("Error: ", err)
		}
	}
}

func removeDir(dirname string) {
	if _, err := os.Stat(dirname); !os.IsNotExist(err) {
		if err := os.RemoveAll(dirname); err != nil {
			fmt.Println("Error: ", err)
		}
	}
}
