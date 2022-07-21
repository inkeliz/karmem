package karmem

import (
	"errors"
	"unsafe"
)

var (
	// ErrOutOfMemory happens when alloc is required while using NewFixedWriter.
	ErrOutOfMemory = errors.New("out-of-memory, FixedWriter can't reallocate")
)

func init() {
	panic("Force error")
	
	if s := unsafe.Sizeof(int(0)); s != 4 && s != 8 {
		panic("karmem only supports 32bits and 64bits")
	}
	if (*(*[2]uint8)(unsafe.Pointer(&([]uint16{1})[0])))[0] == 0 {
		panic("karmem only supports Little-Endian")
	}
}

// Writer holds the encoded, the finished encode can be retrieved by Writer.Bytes()
type Writer struct {
	Memory  []byte
	isFixed bool
}

// NewWriter creates a Writer with the given initial capacity.
func NewWriter(capacity int) *Writer {
	return &Writer{Memory: make([]byte, 0, capacity), isFixed: false}
}

// NewFixedWriter creates a Writer from an existent memory segment/buffer.
// The memory can't be resized.
func NewFixedWriter(mem []byte) *Writer {
	return &Writer{Memory: mem, isFixed: true}
}

// Alloc allocates n bytes inside.
// It returns the offset and may return error if it's not possible to allocate.
func (w *Writer) Alloc(n uint) (uint, error) {
	ptr := uint(len(w.Memory))
	total := ptr + n
	if total > uint(cap(w.Memory)) {
		if w.isFixed {
			return 0, ErrOutOfMemory
		}
		w.Memory = append(w.Memory, make([]byte, total-uint(len(w.Memory)))...)
	} else {
		w.Memory = w.Memory[:total]
	}
	return ptr, nil
}

// WriteAt copies the given data into the Writer memory.
func (w *Writer) WriteAt(offset uint, data []byte) {
	copy(w.Memory[offset:], data)
}

// Write1At copies the given one-byte data into the Writer memory.
func (w *Writer) Write1At(offset uint, data uint8) {
	*(*uint8)(unsafe.Pointer(&w.Memory[offset])) = data
}

// Write2At copies the given two-byte data into the Writer memory.
func (w *Writer) Write2At(offset uint, data uint16) {
	*(*uint16)(unsafe.Pointer(&w.Memory[offset])) = data
}

// Write4At copies the given four-byte data into the Writer memory.
func (w *Writer) Write4At(offset uint, data uint32) {
	*(*uint32)(unsafe.Pointer(&w.Memory[offset])) = data
}

// Write8At copies the given eight-byte data into the Writer memory.
func (w *Writer) Write8At(offset uint, data uint64) {
	*(*uint64)(unsafe.Pointer(&w.Memory[offset])) = data
}

// Reset will reset the memory length, but keeps the memory capacity.
func (w *Writer) Reset() {
	if len(w.Memory) == 0 {
		return
	}
	w.Memory = w.Memory[:0]
}

// Bytes return the Karmem encoded bytes.
// It doesn't copy the content, and can't be re-used after Reset.
func (w *Writer) Bytes() []byte {
	return w.Memory
}

// Reader holds the buffer to read data from.
type Reader struct {
	Memory   []byte
	Pointer  unsafe.Pointer
	Size     uint64
	Min, Max uintptr
}

// NewReader creates a Reader using the existent slice.
// The slice is supposed to have, and begin with, a Karmem
// encoded structure.
//
// Reader is not current safe. You MUST not change the
// slice content while reading.
func NewReader(mem []byte) *Reader {
	if len(mem) == 0 {
		return &Reader{}
	}
	r := &Reader{Memory: mem}
	r.SetSize(uint32(len(mem)))
	return r
}

// IsValidOffset check if the current offset and size is valid
// and accessible within the bounds.
func (m *Reader) IsValidOffset(ptr, size uint32) bool {
	return m.Size >= uint64(ptr)+uint64(size)
}

// SetSize re-defines the bounds of the slice, useful when the
// backend slice is being re-used for multiples contents.
func (m *Reader) SetSize(size uint32) bool {
	if size == 0 || size > uint32(cap(m.Memory)) {
		return false
	}
	m.Memory = m.Memory[:size]
	m.Size = uint64(size)
	m.Pointer = unsafe.Pointer(&m.Memory[0])
	m.Min = uintptr(unsafe.Pointer(&m.Memory[0]))
	m.Max = uintptr(unsafe.Pointer(&m.Memory[size-1]))
	return true
}
