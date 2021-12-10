package win

import (
	"syscall"
	"unsafe"

	"github.com/netfiredotnet/windigo/internal/proc"
	"github.com/netfiredotnet/windigo/win/co"
	"github.com/netfiredotnet/windigo/win/errco"
)

// A handle to the clipboard. Actually this handle does not exist, it only
// serves the purpose of logically group the clipboard functions.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/dataxchg/clipboard
type HCLIPBOARD struct{}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-closeclipboard
func (HCLIPBOARD) CloseClipboard() {
	ret, _, err := syscall.Syscall(proc.CloseClipboard.Addr(), 0,
		0, 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-emptyclipboard
func (HCLIPBOARD) EmptyClipboard() {
	ret, _, err := syscall.Syscall(proc.EmptyClipboard.Addr(), 0,
		0, 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enumclipboardformats
func (HCLIPBOARD) EnumClipboardFormats() []co.CF {
	formats := make([]co.CF, 0, 30) // arbitrary
	thisFormat := co.CF(0)

	for {
		ret, _, err := syscall.Syscall(proc.EnumClipboardFormats.Addr(), 1,
			uintptr(thisFormat), 0, 0)

		if ret == 0 {
			if wErr := errco.ERROR(err); wErr == errco.SUCCESS {
				return formats
			} else {
				panic(wErr)
			}
		} else {
			thisFormat = co.CF(ret)
			formats = append(formats, thisFormat)
		}
	}
}

// ⚠️ hMem will be owned by the clipboard, do not call HGLOBAL.Free() anymore.
//
// Unless you're doing something specific, prefer HCLIPBOARD.WriteBitmap() or
// HCLIPBOARD.WriteString().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setclipboarddata
func (HCLIPBOARD) SetClipboardData(format co.CF, hMem HGLOBAL) {
	ret, _, err := syscall.Syscall(proc.SetClipboardData.Addr(), 2,
		uintptr(format), uintptr(hMem), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// ⚠️ hBmp will be owned by the clipboard, do not call HBITMAP.DeleteObject() anymore.
//
// Writes a bitmap to the clipboard with HCLIPBOARD.SetClipboardData().
func (hClip HCLIPBOARD) WriteBitmap(hBmp HBITMAP) {
	hClip.SetClipboardData(co.CF_BITMAP, HGLOBAL(hBmp))
}

// Writes a string to the clipboard with HCLIPBOARD.SetClipboardData().
func (hClip HCLIPBOARD) WriteString(text string) {
	text16 := Str.ToNativeSlice(text)
	text8 := unsafe.Slice((*byte)(unsafe.Pointer(&text16[0])), len(text16)*2) // direct pointer conversion

	hGlob := GlobalAlloc(co.GMEM_MOVEABLE, uint64(len(text8)))
	pMem := hGlob.GlobalLock()
	pMemSlice := unsafe.Slice((*byte)(pMem), len(text8))
	copy(pMemSlice, text8)
	hGlob.GlobalUnlock()

	hClip.SetClipboardData(co.CF_UNICODETEXT, hGlob)
}
