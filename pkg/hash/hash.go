package hash

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func hashFile(file string) (dirname string, filename string) {
	// Open the file
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// Create a new hash interface to write to
	h := sha256.New()

	// Copy the file in the hash interface and check for any error
	if _, err := io.Copy(h, f); err != nil {
		fmt.Println("Error: ", err)
	}

	// Get the 40 bytes hash
	hashInBytes := h.Sum(nil)[:40]

	dirname = fmt.Sprintf("%x", hashInBytes[:2])
	filename = fmt.Sprintf("%x", hashInBytes[2:])
	return
}
