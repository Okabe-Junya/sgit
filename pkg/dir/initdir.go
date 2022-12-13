package dir

func InitializeDir() {
	// Create the .sgit/objects directory
	CreateDir(".sgit/objects")

	// Create the sub directories
	CreateDir(".sgit/objects/info")
	CreateDir(".sgit/objects/pack")
}
