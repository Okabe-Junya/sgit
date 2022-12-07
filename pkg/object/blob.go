package object

import (
	"crypto/sha1"
	"strconv"
)

func storeBlob(content string) {
	header := "blob " + strconv.Itoa(len(content)) + "\x00"

	store := header + content
	sha1_store := sha1.Sum([]byte(store))
	// TODO: zlibで圧縮して格納する
}
