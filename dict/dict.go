package dict

import (
	"unsafe"
)

//Dictht struct
type Dictht struct {
	Table    **Entry
	Size     uint32
	SizeMark uint32
	Used     uint32
}

//Entry struct
type Entry struct {
	Key  unsafe.Pointer
	Next *Entry
}

//Dict struct real
type Dict struct {
	Ht        [2]Dictht
	RehashIdx int
}
