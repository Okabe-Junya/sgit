package object

import "testing"

func TestMain(m *testing.M) {
	// Setup
	m.Run()
	// Teardown
	// TODO: Remove test directory
}

func TestStoreBlob(t *testing.T) {
	rootdir := "testdata"
	content := "hello world"
	StoreBlob(content, rootdir)

}
