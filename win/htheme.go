package win

import (
	"syscall"
	"unsafe"

	"github.com/netfiredotnet/windigo/internal/proc"
	"github.com/netfiredotnet/windigo/win/co"
	"github.com/netfiredotnet/windigo/win/errco"
)

// Handle to a theme.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/uxtheme/
type HTHEME HANDLE

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-closethemedata
func (hTheme HTHEME) CloseThemeData() {
	if hTheme != 0 {
		syscall.Syscall(proc.CloseThemeData.Addr(), 1,
			uintptr(hTheme), 0, 0)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-drawthemebackground
func (hTheme HTHEME) DrawThemeBackground(
	hdc HDC, partStateId co.VS, rc *RECT, clipRc *RECT) {

	ret, _, _ := syscall.Syscall6(proc.DrawThemeBackground.Addr(), 6,
		uintptr(hTheme), uintptr(hdc),
		uintptr(partStateId.Part()), uintptr(partStateId.State()),
		uintptr(unsafe.Pointer(rc)), uintptr(unsafe.Pointer(clipRc)))
	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-getthemecolor
func (hTheme HTHEME) GetThemeColor(
	partStateId co.VS, propId co.TMT_COLOR) COLORREF {

	var color COLORREF
	ret, _, _ := syscall.Syscall6(proc.GetThemeColor.Addr(), 5,
		uintptr(hTheme), uintptr(partStateId.Part()), uintptr(partStateId.State()),
		uintptr(propId), uintptr(unsafe.Pointer(&color)), 0)
	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
	return color
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-getthemeint
func (hTheme HTHEME) GetThemeInt(partStateId co.VS, propId co.TMT_INT) int32 {
	var intVal int32
	ret, _, _ := syscall.Syscall6(proc.GetThemeInt.Addr(), 5,
		uintptr(hTheme), uintptr(partStateId.Part()), uintptr(partStateId.State()),
		uintptr(propId), uintptr(unsafe.Pointer(&intVal)), 0)
	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
	return intVal
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-getthememetric
func (hTheme HTHEME) GetThemeMetric(
	hdc HDC, partStateId co.VS, propId co.TMT_INT) int32 {

	var intVal int32
	ret, _, _ := syscall.Syscall6(proc.GetThemeMetric.Addr(), 6,
		uintptr(hTheme), uintptr(hdc),
		uintptr(partStateId.Part()), uintptr(partStateId.State()),
		uintptr(propId), uintptr(unsafe.Pointer(&intVal)))
	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
	return intVal
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-getthemeposition
func (hTheme HTHEME) GetThemePosition(
	partStateId co.VS, propId co.TMT_POSITION) POINT {

	var pt POINT
	ret, _, _ := syscall.Syscall6(proc.GetThemePosition.Addr(), 5,
		uintptr(hTheme), uintptr(partStateId.Part()), uintptr(partStateId.State()),
		uintptr(propId), uintptr(unsafe.Pointer(&pt)), 0)
	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
	return pt
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-getthemerect
func (hTheme HTHEME) GetThemeRect(partStateId co.VS, propId co.TMT_RECT) RECT {
	var rc RECT
	ret, _, _ := syscall.Syscall6(proc.GetThemeRect.Addr(), 5,
		uintptr(hTheme), uintptr(partStateId.Part()), uintptr(partStateId.State()),
		uintptr(propId), uintptr(unsafe.Pointer(&rc)), 0)
	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
	return rc
}

// ⚠️ You must defer HBRUSH.DeleteObject().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-getthemesyscolorbrush
func (hTheme HTHEME) GetThemeSysColorBrush(colorId co.TMT_COLOR) HBRUSH {
	ret, _, err := syscall.Syscall(proc.GetThemeSysColorBrush.Addr(), 2,
		uintptr(hTheme), uintptr(colorId), 0)
	if ret == 0 {
		panic(errco.ERROR(err)) // uncertain?
	}
	return HBRUSH(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-getthemesysfont
func (hTheme HTHEME) GetThemeSysFont(fontId co.TMT_FONT, lf *LOGFONT) {
	ret, _, _ := syscall.Syscall(proc.GetThemeSysFont.Addr(), 3,
		uintptr(hTheme), uintptr(fontId), uintptr(unsafe.Pointer(lf)))
	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-getthemetextmetrics
func (hTheme HTHEME) GetThemeTextMetrics(
	hdc HDC, partStateId co.VS, tm *TEXTMETRIC) {

	ret, _, _ := syscall.Syscall6(proc.GetThemeTextMetrics.Addr(), 5,
		uintptr(hTheme), uintptr(hdc),
		uintptr(partStateId.Part()), uintptr(partStateId.State()),
		uintptr(unsafe.Pointer(tm)), 0)
	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-isthemebackgroundpartiallytransparent
func (hTheme HTHEME) IsThemeBackgroundPartiallyTransparent(partStateId co.VS) bool {
	ret, _, _ := syscall.Syscall(proc.IsThemeBackgroundPartiallyTransparent.Addr(), 3,
		uintptr(hTheme), uintptr(partStateId.Part()), uintptr(partStateId.State()))
	return ret != 0
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/uxtheme/nf-uxtheme-isthemepartdefined
func (hTheme HTHEME) IsThemePartDefined(partStateId co.VS) bool {
	ret, _, _ := syscall.Syscall(proc.IsThemePartDefined.Addr(), 3,
		uintptr(hTheme), uintptr(partStateId.Part()), uintptr(partStateId.State()))
	return ret != 0
}
