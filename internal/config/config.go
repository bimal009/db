package config

const (
	PUT    byte = 1
	DELETE byte = 2
)

type Record struct {
	Type  byte
	Key   []byte
	Value []byte
}


type ColumnsSizes struct{
	
}