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

func initializeTree() *Tree {
	return &Tree{}
}

func (t *Tree) AddEntry(entry TreeEntry) {
	t.Entries = append(t.Entries, entry)
}

func (t *Tree) GetEntry(name string) *TreeEntry {
	for _, entry := range t.Entries {
		if entry.Name == name {
			return &entry
		}
	}
	return nil
}
