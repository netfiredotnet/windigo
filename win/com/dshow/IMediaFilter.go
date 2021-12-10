package dshow

import (
	"syscall"
	"time"
	"unsafe"

	"github.com/netfiredotnet/windigo/internal/util"
	"github.com/netfiredotnet/windigo/win"
	"github.com/netfiredotnet/windigo/win/com/dshow/dshowco"
	"github.com/netfiredotnet/windigo/win/com/dshow/dshowvt"
	"github.com/netfiredotnet/windigo/win/com/idl"
	"github.com/netfiredotnet/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-imediafilter
type IMediaFilter struct{ idl.IPersist }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IMediaFilter.Release().
func NewIMediaFilter(base win.IUnknown) IMediaFilter {
	return IMediaFilter{IPersist: idl.NewIPersist(base)}
}

// Pass -1 for infinite timeout.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediafilter-getstate
func (me *IMediaFilter) GetState(msTimeout int) (dshowco.FILTER_STATE, error) {
	var state dshowco.FILTER_STATE
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IMediaFilter)(unsafe.Pointer(*me.Ptr())).GetState, 3,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(int32(msTimeout)), uintptr(unsafe.Pointer(&state)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return state, nil
	} else if hr == errco.VFW_S_STATE_INTERMEDIATE {
		return state, hr
	} else {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediafilter-pause
func (me *IMediaFilter) Pause() bool {
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IMediaFilter)(unsafe.Pointer(*me.Ptr())).Pause, 1,
		uintptr(unsafe.Pointer(me.Ptr())),
		0, 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return true
	} else if hr == errco.S_FALSE {
		return false
	} else {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediafilter-run
func (me *IMediaFilter) Run(start time.Duration) bool {
	iStart := util.DurationToNano100(start)
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IMediaFilter)(unsafe.Pointer(*me.Ptr())).Run, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(iStart), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return true
	} else if hr == errco.S_FALSE {
		return false
	} else {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediafilter-stop
func (me *IMediaFilter) Stop() bool {
	ret, _, _ := syscall.Syscall(
		(*dshowvt.IMediaFilter)(unsafe.Pointer(*me.Ptr())).Stop, 1,
		uintptr(unsafe.Pointer(me.Ptr())),
		0, 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return true
	} else if hr == errco.S_FALSE {
		return false
	} else {
		panic(hr)
	}
}
