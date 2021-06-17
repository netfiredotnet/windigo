package dshow

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/com/dshow/dshowco"
	"github.com/rodrigocfd/windigo/win/errco"
)

type _IMFGetServiceVtbl struct {
	win.IUnknownVtbl
	GetService uintptr
}

//------------------------------------------------------------------------------

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/mfidl/nn-mfidl-imfgetservice
type IMFGetService struct {
	win.IUnknown // Base IUnknown.
}

// ⚠️ You must defer Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/mfidl/nf-mfidl-imfgetservice-getservice
func (me *IMFGetService) GetService(guidService, riid *win.GUID) win.IUnknown {
	var ppQueried **win.IUnknownVtbl
	ret, _, _ := syscall.Syscall6(
		(*_IMFGetServiceVtbl)(unsafe.Pointer(*me.Ppv)).GetService, 4,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(guidService)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(&ppQueried)), 0, 0)

	if err := errco.ERROR(ret); err != errco.S_OK {
		panic(err)
	}
	return win.IUnknown{Ppv: ppQueried}
}

// Calls IMFGetService.GetService() to return IMFVideoDisplayControl.
//
// ⚠️ You must defer Release().
func (me *IMFGetService) GetServiceIMFVideoDisplayControl() IMFVideoDisplayControl {
	iUnk := me.GetService(
		win.NewGuidFromClsid(dshowco.CLSID_MR_VideoRenderService),
		win.NewGuidFromIid(dshowco.IID_IMFVideoDisplayControl))
	return IMFVideoDisplayControl{IUnknown: iUnk}
}
