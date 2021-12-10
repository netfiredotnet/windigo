package dshow

import (
	"syscall"
	"unsafe"

	"github.com/netfiredotnet/windigo/win"
	"github.com/netfiredotnet/windigo/win/com/dshow/dshowvt"
	"github.com/netfiredotnet/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ienummediatypes
type IEnumMediaTypes struct{ win.IUnknown }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IEnumMediaTypes.Release().
func NewIEnumMediaTypes(base win.IUnknown) IEnumMediaTypes {
	return IEnumMediaTypes{IUnknown: base}
}

// ⚠️ You must defer IEnumMediaTypes.Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ienummediatypes-clone
func (me *IEnumMediaTypes) Clone() IEnumMediaTypes {
	var ppQueried win.IUnknown
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IEnumMediaTypes)(unsafe.Pointer(*me.Ptr())).Clone, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&ppQueried)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewIEnumMediaTypes(ppQueried)
	} else {
		panic(hr)
	}
}

// Calls Skip() until the end of the enum to retrieve the actual number of media
// types, then calls Reset().
func (me *IEnumMediaTypes) Count() int {
	count := int(0)
	for {
		gotOne := me.Skip(1)
		if gotOne {
			count++
		} else {
			me.Reset()
			return count
		}
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ienummediatypes-next
func (me *IEnumMediaTypes) Next(mt *AM_MEDIA_TYPE) bool {
	ret, _, _ := syscall.Syscall6(
		(*dshowvt.IEnumMediaTypes)(unsafe.Pointer(*me.Ptr())).Next, 4,
		uintptr(unsafe.Pointer(me.Ptr())),
		1, uintptr(unsafe.Pointer(&mt)), 0, 0, 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return true
	} else if hr == errco.S_FALSE {
		return false
	} else {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ienummediatypes-reset
func (me *IEnumMediaTypes) Reset() {
	syscall.Syscall(
		(*dshowvt.IEnumMediaTypes)(unsafe.Pointer(*me.Ptr())).Reset, 1,
		uintptr(unsafe.Pointer(me.Ptr())),
		0, 0)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ienummediatypes-skip
func (me *IEnumMediaTypes) Skip(numMediaTypes int) bool {
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IEnumMediaTypes)(unsafe.Pointer(*me.Ptr())).Skip, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(uint32(numMediaTypes)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return true
	} else if hr == errco.S_FALSE {
		return false
	} else {
		panic(hr)
	}
}
