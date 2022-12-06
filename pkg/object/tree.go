package object

type Tree struct {
	Entries []TreeEntry
}

type TreeEntry struct {
	Mode string            // XXXXXX: should be a number
	Type EnumTreeEntryType // blob, tree, commit
	Hash string            // SHA1
	Name string            // filename
}

type EnumTreeEntryType int

const (
	blob EnumTreeEntryType = iota
	tree
	commit
)

func (t EnumTreeEntryType) String() string {
	switch t {
	case blob:
		return "blob"
	case tree:
		return "tree"
	case commit:
		return "commit"
	default:
		return ""
	}
}
