package object

import "testing"

func TestStoreBlob(t *testing.T) {
	content := "hello world"
	storeBlob(content)
}