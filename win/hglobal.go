package win

import (
	"syscall"
	"unsafe"

	"github.com/netfiredotnet/windigo/internal/proc"
	"github.com/netfiredotnet/windigo/win/co"
	"github.com/netfiredotnet/windigo/win/errco"
)

// A handle to a global memory block.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hglobal
type HGLOBAL HANDLE

// With co.GMEM_FIXED, the handle itself is the pointer to the memory block.
// With co.GMEM_MOVEABLE, you must call HGLOBAL.GlobalLock() to retrieve the
// pointer.
//
// ⚠️ You must defer HGLOBAL.GlobalFree().
//
// Example:
//
//  hMem := win.GlobalAlloc(co.GMEM_FIXED|co.GMEM_ZEROINIT, 50)
//  defer hMem.GlobalFree()
//
//  sliceMem := unsafe.Slice((*byte)(unsafe.Pointer(hMem)), 50)
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-globalalloc
func GlobalAlloc(uFlags co.GMEM, dwBytes uint64) HGLOBAL {
	ret, _, err := syscall.Syscall(proc.GlobalAlloc.Addr(), 2,
		uintptr(uFlags), uintptr(dwBytes), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HGLOBAL(ret)
}

// Allocs a null-terminated *uint16.
//
// With co.GMEM_FIXED, the handle itself is the pointer to the memory block.
// With co.GMEM_MOVEABLE, you must call HGLOBAL.GlobalLock() to retrieve the
// pointer.
//
// ⚠️ You must defer HGLOBAL.GlobalFree().
//
// Example:
//
//  hMem := win.GlobalAllocStr(co.GMEM_FIXED, "my text")
//  defer hMem.GlobalFree()
//
//  pChars := (*uint16)(unsafe.Pointer(hMem))
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-globalalloc
func GlobalAllocStr(uFlags co.GMEM, s string) HGLOBAL {
	sliceStr := Str.ToNativeSlice(s)
	numBytesSliceStr := len(sliceStr) * int(unsafe.Sizeof(sliceStr[0]))

	hMem := GlobalAlloc(uFlags, uint64(numBytesSliceStr))
	sliceMem := unsafe.Slice((*uint16)(unsafe.Pointer(hMem)), len(sliceStr))
	copy(sliceMem, sliceStr)
	return hMem
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-globalflags
func (hGlobal HGLOBAL) GlobalFlags() co.GMEM {
	ret, _, err := syscall.Syscall(proc.GlobalFlags.Addr(), 1,
		uintptr(hGlobal), 0, 0)
	if ret == _GMEM_INVALID_HANDLE {
		panic(errco.ERROR(err))
	}
	return co.GMEM(ret)
}

// This method is safe to be called if hGlobal is zero.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-globalfree
func (hGlobal HGLOBAL) GlobalFree() {
	ret, _, err := syscall.Syscall(proc.GlobalFree.Addr(), 1,
		uintptr(hGlobal), 0, 0)
	if ret != 0 {
		panic(errco.ERROR(err))
	}
}

// If you called GlobalAlloc() with co.GMEM_FIXED, you don't need to call this
// function. The handle itself is the pointer to the memory block.
//
// ⚠️ You must defer HGLOBAL.GlobalUnlock(). After that, the slice must not be
// used.
//
// Example:
//
//  hMem := win.GlobalAlloc(co.GMEM_FIXED, 50)
//  defer hMem.GlobalFree()
//
//  lockedPtr := hMem.GlobalLock()
//  defer hMem.GlobalUnlock()
//
//  sliceMem := unsafe.Slice((*byte)(lockedPtr), 50)
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-globallock
func (hGlobal HGLOBAL) GlobalLock() unsafe.Pointer {
	ret, _, err := syscall.Syscall(proc.GlobalLock.Addr(), 1,
		uintptr(hGlobal), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return unsafe.Pointer(ret)
}

// ⚠️ You must defer HGLOBAL.GlobalFree().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-globalrealloc
func (hGlobal HGLOBAL) GlobalReAlloc(dwBytes uint64, uFlags co.GMEM) HGLOBAL {
	ret, _, err := syscall.Syscall(proc.GlobalReAlloc.Addr(), 3,
		uintptr(hGlobal), uintptr(dwBytes), uintptr(uFlags))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HGLOBAL(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-globalsize
func (hGlobal HGLOBAL) GlobalSize() uint64 {
	ret, _, err := syscall.Syscall(proc.GlobalSize.Addr(), 1,
		uintptr(hGlobal), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return uint64(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-globalunlock
func (hGlobal HGLOBAL) GlobalUnlock() {
	ret, _, err := syscall.Syscall(proc.GlobalUnlock.Addr(), 1,
		uintptr(hGlobal), 0, 0)
	if wErr := errco.ERROR(err); ret == 0 && wErr != errco.SUCCESS {
		panic(wErr)
	}
}
