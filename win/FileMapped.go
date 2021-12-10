package win

import (
	"strings"
	"unsafe"

	"github.com/netfiredotnet/windigo/internal/util"
	"github.com/netfiredotnet/windigo/win/co"
)

// High-level abstraction to HFILEMAP, providing several operations.
//
// Created with FileMappedOpen().
type FileMapped struct {
	objFile  *File
	hMap     HFILEMAP
	pMem     HFILEMAPVIEW
	sz       int
	readOnly bool
}

// Opens a memory-mapped file, returning a new high-level FileMapped object.
//
// ⚠️ You must defer FileMapped.Close().
func FileMappedOpen(
	filePath string, desiredAccess co.FILE_OPEN) (*FileMapped, error) {

	objFile, err := FileOpen(filePath, desiredAccess)
	if err != nil {
		return nil, err
	}

	me := &FileMapped{
		objFile:  objFile,
		hMap:     HFILEMAP(0),
		pMem:     HFILEMAPVIEW(0),
		sz:       0,
		readOnly: desiredAccess == co.FILE_OPEN_READ_EXISTING,
	}

	if err := me.mapInMemory(); err != nil {
		me.Close()
		return nil, err
	}
	return me, nil
}

// Unmaps and releases the file resource.
func (me *FileMapped) Close() {
	if me.pMem != 0 {
		me.pMem.UnmapViewOfFile()
		me.pMem = 0
	}
	if me.hMap != 0 {
		me.hMap.CloseHandle()
		me.hMap = 0
	}
	me.objFile.Close()
	me.sz = 0
}

// Returns a slice to the memory-mapped bytes. The FileMapped object must remain
// open while the slice is being used.
//
// If you need to close the file right away, use CopyToBuffer() instead.
func (me *FileMapped) HotSlice() []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(me.pMem)), me.sz)
}

// Returns a new []byte with a copy of all data in the file.
func (me *FileMapped) ReadAll() []byte {
	return me.ReadChunk(0, me.sz)
}

// Returns a new []byte with a copy of data, start with offset, and with the
// undefined given undefined length.
func (me *FileMapped) ReadChunk(offset, length int) []byte {
	hotSlice := me.HotSlice()
	buf := make([]byte, length)
	copy(buf, hotSlice[offset:offset+length])
	return buf
}

// Parses the file content as text and returns the lines.
func (me *FileMapped) ReadLines() []string {
	allText := string(me.HotSlice())
	lines := strings.Split(allText, "\n")

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if line[len(line)-1] == '\r' {
			lines[i] = line[:len(line)-1] // trim trailing \r
		}
	}
	return lines
}

// Truncates or expands the file, according to the new size. Zero will empty the
// file.
//
// Internally, the file is unmapped, then remapped back into memory.
func (me *FileMapped) Resize(numBytes int) error {
	me.pMem.UnmapViewOfFile()
	me.hMap.CloseHandle()
	if err := me.objFile.Resize(numBytes); err != nil {
		return err
	}
	return me.mapInMemory()
}

// Retrieves the file size. This value is cached.
func (me *FileMapped) Size() int {
	return me.sz
}

func (me *FileMapped) mapInMemory() error {
	// Mapping into memory.
	pageFlags := util.Iif(me.readOnly,
		co.PAGE_READONLY, co.PAGE_READWRITE).(co.PAGE)

	var err error
	me.hMap, err = me.objFile.Hfile().
		CreateFileMapping(nil, pageFlags, co.SEC_NONE, 0, nil)
	if err != nil {
		return err
	}

	// Get pointer to data block.
	mapFlags := util.Iif(me.readOnly,
		co.FILE_MAP_READ, co.FILE_MAP_WRITE).(co.FILE_MAP)

	if me.pMem, err = me.hMap.MapViewOfFile(mapFlags, 0, 0); err != nil {
		return err
	}

	// Cache file size.
	me.sz = me.objFile.Size()

	return nil // file mapped successfully
}
