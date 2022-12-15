package object

import (
	"testing"

	"github.com/Okabe-Junya/sgit/pkg/dir"
)

func TestMain(m *testing.M) {
	// Setup
	m.Run()
	// Teardown
	dir.RemoveDir("testdata")
}

func TestStoreBlob(t *testing.T) {
	rootdir := "testdata"
	content := "hello world"
	StoreBlob(content, rootdir)

}
