package dshow

import (
	"syscall"
	"unsafe"

	"github.com/netfiredotnet/windigo/win"
	"github.com/netfiredotnet/windigo/win/com/dshow/dshowvt"
	"github.com/netfiredotnet/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ifilesourcefilter
type IFileSourceFilter struct{ win.IUnknown }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IFileSourceFilter.Release().
func NewIFileSourceFilter(base win.IUnknown) IFileSourceFilter {
	return IFileSourceFilter{IUnknown: base}
}

// Returns false if no file is opened.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ifilesourcefilter-getcurfile
func (me *IFileSourceFilter) GetCurFile(mt *AM_MEDIA_TYPE) (string, bool) {
	var pv uintptr
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IFileSourceFilter)(unsafe.Pointer(*me.Ptr())).GetCurFile, 3,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&pv)), uintptr(unsafe.Pointer(mt)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		name := win.Str.FromNativePtr((*uint16)(unsafe.Pointer(pv)))
		win.CoTaskMemFree(pv)
		return name, true
	} else if hr == errco.E_FAIL {
		return "", false
	} else {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ifilesourcefilter-load
func (me *IFileSourceFilter) Load(fileName string, mt *AM_MEDIA_TYPE) {
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IFileSourceFilter)(unsafe.Pointer(*me.Ptr())).Load, 3,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(fileName))),
		uintptr(unsafe.Pointer(mt)))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}
