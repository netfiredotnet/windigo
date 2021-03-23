package win

import (
	"fmt"
	"syscall"
	"unsafe"
)

type _StrT struct{}

// Wide char UTF-16 string conversion functions.
var Str _StrT

// Formats bytes into KB, MB, GB or TB.
func (_StrT) FmtBytes(numBytes uint64) string {
	switch {
	case numBytes < 1024:
		return fmt.Sprintf("%d bytes", numBytes)
	case numBytes < 1024*1024:
		return fmt.Sprintf("%.2f KB", float64(numBytes)/1024)
	case numBytes < 1024*1024*1024:
		return fmt.Sprintf("%.2f MB", float64(numBytes)/1024/1024)
	case numBytes < 1024*1024*1024*1024:
		return fmt.Sprintf("%.2f GB", float64(numBytes)/1024/1024/1024)
	default:
		return fmt.Sprintf("%.2f TB", float64(numBytes)/1024/1024/1024/1024)
	}
}

// Converts a null-terminated *uint16 to string.
//
// Copied from syscall_windows.go, utf16PtrToString() private function.
func (_StrT) FromUint16Ptr(p *uint16) string {
	if p == nil {
		return ""
	}

	// Find null terminator.
	pRun := unsafe.Pointer(p)
	sLen := 0
	for *(*uint16)(pRun) != 0 {
		pRun = unsafe.Pointer(uintptr(pRun) + unsafe.Sizeof(*p)) // pRun++
		sLen++
	}

	slice := Str.ptrToSlice(p, sLen) // create slice without terminating null
	return Str.FromUint16Slice(slice)
}

// Converts a multi null-terminated *uint16 to []string.
//
// Source must have 2 terminating nulls.
func (_StrT) FromUint16PtrMulti(p *uint16) []string {
	values := make([]string, 0)
	if p == nil {
		return values
	}

	pRun := unsafe.Pointer(p)
	sLen := 0
	for {
		if *(*uint16)(pRun) == 0 { // terminating null found
			if sLen == 0 {
				break // two terminating nulls
			}

			slice := Str.ptrToSlice(p, sLen) // create slice without terminating null
			values = append(values, Str.FromUint16Slice(slice))

			pRun = unsafe.Pointer(uintptr(pRun) + unsafe.Sizeof(*p)) // pRun++
			p = (*uint16)(pRun)
			sLen = 0

		} else {
			pRun = unsafe.Pointer(uintptr(pRun) + unsafe.Sizeof(*p)) // pRun++
			sLen++
		}
	}

	return values
}

// Converts a null-terminated []uint16 to string.
//
// Simple wrapper to syscall.UTF16ToString().
func (_StrT) FromUint16Slice(s []uint16) string {
	return syscall.UTF16ToString(s)
}

// Converts string to null-terminated *uint16.
//
// Ideal to pass strings to syscalls. We won't return an uintptr right away
// because it has no pointer semantics, it's just a number, so pointed memory
// can be garbage-collected.
//
// https://stackoverflow.com/a/51188315
//
// Wrapper to syscall.UTF16PtrFromString(). Panics on error.
func (_StrT) ToUint16Ptr(s string) *uint16 {
	pstr, lerr := syscall.UTF16PtrFromString(s)
	if lerr != nil {
		panic(fmt.Sprintf("StrToPtr failed \"%s\": %s",
			s, lerr))
	}
	return pstr
}

// Converts string to null-terminated *uint16, or nil if source is empty.
//
// Wrapper to syscall.UTF16PtrFromString(). Panics on error.
func (_StrT) ToUint16PtrBlankIsNil(s string) *uint16 {
	if s != "" {
		return Str.ToUint16Ptr(s)
	}
	return nil
}

// Converts []string to multi null-terminated *uint16.
//
// Memory block will have 2 terminating nulls.
func (_StrT) ToUint16PtrMulti(ss []string) *uint16 {
	slice := Str.ToUint16SliceMulti(ss)
	return &slice[0]
}

// Converts string to null-terminated []uint16.
//
// Wrapper to syscall.UTF16FromString(). Panics on error.
func (_StrT) ToUint16Slice(s string) []uint16 {
	sli, lerr := syscall.UTF16FromString(s)
	if lerr != nil {
		panic(fmt.Sprintf("StrToSlice failed \"%s\": %s",
			s, lerr))
	}
	return sli
}

// Converts []string to multi null-terminated []uint16.
//
// Returned slice will have 2 terminating nulls.
func (_StrT) ToUint16SliceMulti(ss []string) []uint16 {
	totalLen := 0
	for _, s := range ss {
		totalLen += len(s) + 1 // also count terminating null
	}

	buf := make([]uint16, 0, totalLen)

	for _, s := range ss {
		buf = append(buf, Str.ToUint16Slice(s)...)
	}
	buf = append(buf, 0) // 2nd terminating null

	return buf
}

// Converts *uint16 to []uint16, with the given length.
func (_StrT) ptrToSlice(ptr *uint16, length int) []uint16 {
	// https://stackoverflow.com/a/43592538
	// https://golang.org/pkg/internal/unsafeheader/#Slice
	var sliceMem = struct { // slice memory layout
		addr unsafe.Pointer
		len  int
		cap  int
	}{unsafe.Pointer(ptr), length, length}

	return *(*[]uint16)(unsafe.Pointer(&sliceMem)) // convert to slice itself
}
