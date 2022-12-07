package dir

import (
	"fmt"
	"os"
)

func createDir(dirname string) {
	if _, err := os.Stat(dirname); os.IsNotExist(err) {
		if err := os.MkdirAll(dirname, 0755); err != nil {
			fmt.Println("Error: ", err)
		}
	}
}
