package ui

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/netfiredotnet/windigo/win"
	"github.com/netfiredotnet/windigo/win/co"
)

// Base to all native controls.
type _NativeControlBase struct {
	hWnd         win.HWND
	ctrlId       int
	parent       AnyParent
	eventsSubcl  _EventsWm // subclass events
	subclassId   uint32
	subclassProc uintptr // necessary to circumvent InvalidInitCycle error
}

func (me *_NativeControlBase) new(parent AnyParent, ctrlId int) {
	me.hWnd = win.HWND(0)
	me.ctrlId = ctrlId
	me.parent = parent
	me.eventsSubcl.new()
	me.subclassId = 0
}

// Implements AnyWindow.
func (me *_NativeControlBase) Hwnd() win.HWND {
	return me.hWnd
}

// Implements AnyControl.
func (me *_NativeControlBase) CtrlId() int {
	return me.ctrlId
}

// Implements AnyControl.
func (me *_NativeControlBase) Parent() AnyParent {
	return me.parent
}

// Implements AnyNativeControl.
func (me *_NativeControlBase) OnSubclass() *_EventsWm {
	if me.Hwnd() != 0 {
		panic("Cannot add event handling after the control is created.")
	}
	return &me.eventsSubcl
}

// Calls CreateWindowEx() and installs subclass.
func (me *_NativeControlBase) createWindow(
	exStyle co.WS_EX, className win.ClassName, title win.StrOrNil, style co.WS,
	pos win.POINT, size win.SIZE, hMenu win.HMENU) {

	if me.hWnd != 0 {
		panic(fmt.Sprintf("Control already created: \"%s\".", className))
	}

	me.hWnd = win.CreateWindowEx(exStyle, className, title, style,
		pos.X, pos.Y, size.Cx, size.Cy, me.parent.Hwnd(), hMenu,
		me.parent.Hwnd().Hinstance(), 0)

	me.installSubclassIfNeeded()
}

// Calls GetDlgItem() and installs subclass.
func (me *_NativeControlBase) assignDlgItem() {
	if me.hWnd != 0 {
		panic(fmt.Sprintf("Dialog control already assigned: \"%d\".", me.ctrlId))
	}
	if !me.parent.isDialog() {
		panic(fmt.Sprintf("Parent is not dialog, cannot assign control: \"%d\".", me.ctrlId))
	}

	me.hWnd = me.parent.Hwnd().GetDlgItem(int32(me.ctrlId))
	me.installSubclassIfNeeded()
}

func (me *_NativeControlBase) installSubclassIfNeeded() {
	if me.eventsSubcl.hasMessages() {
		_globalSubclassId++
		me.subclassId = _globalSubclassId
		me.subclassProc = _globalSubclassProc

		_globalNativeControlBasePtrs[me] = struct{}{} // store pointer in the set

		// Subclass is installed after window creation, thus WM_CREATE can never
		// be handled for a subclassed control.
		me.hWnd.SetWindowSubclass(me.subclassProc,
			me.subclassId, unsafe.Pointer(me)) // pass pointer to object itself
	}
}

func (me *_NativeControlBase) focus() {
	// Will draw the borders correctly in a button.
	me.hWnd.GetParent().SendMessage(co.WM_NEXTDLGCTL, win.WPARAM(me.hWnd), 1)
}

var (
	// incremented at each subclass installed
	_globalSubclassId = uint32(0)

	// A set keeping all *_NativeControlBase that were retrieved in _SubclassProc.
	_globalNativeControlBasePtrs = make(map[*_NativeControlBase]struct{}, 10)

	// Default window procedure for subclassed child controls.
	_globalSubclassProc uintptr = syscall.NewCallback(_SubclassProc)
)

func _SubclassProc(
	hWnd win.HWND, uMsg co.WM, wParam win.WPARAM, lParam win.LPARAM,
	uIdSubclass, dwRefData uintptr) uintptr {

	pMe := (*_NativeControlBase)(unsafe.Pointer(dwRefData)) // retrieve passed pointer

	// If object pointer is not stored, then no processing is done.
	// Prevents processing after WM_NCDESTROY.
	if _, isStored := _globalNativeControlBasePtrs[pMe]; isStored {
		// Try to process the message with an user handler.
		retVal, meaningfulRet, wasHandled :=
			pMe.eventsSubcl.processMessage(uMsg, wParam, lParam)

		if uMsg == co.WM_NCDESTROY { // even if the user handles WM_NCDESTROY, we must ensure cleanup
			delete(_globalNativeControlBasePtrs, pMe) // remove from set
			hWnd.RemoveWindowSubclass(pMe.subclassProc, pMe.subclassId)
		}

		if wasHandled {
			if meaningfulRet {
				return retVal
			}
			return 0 // message processed, default return value
		}
	} else if uMsg == co.WM_NCDESTROY {
		// https://devblogs.microsoft.com/oldnewthing/20031111-00/?p=41883
		delete(_globalNativeControlBasePtrs, pMe) // remove from set
		hWnd.RemoveWindowSubclass(pMe.subclassProc, pMe.subclassId)
	}

	return hWnd.DefSubclassProc(uMsg, wParam, lParam) // message was not processed
}
