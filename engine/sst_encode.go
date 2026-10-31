package engine

type SSTRecord struct {
	Deleted bool
	Key     []byte
	Value   []byte
}

type SSTIndexRecord struct {
	Key    string
	Offset uint64
}

type SSTFooter struct {
	IndexOffset   uint64
	IndexSize     uint64
	EntryCount    uint64
	IndexChecksum uint32
}
