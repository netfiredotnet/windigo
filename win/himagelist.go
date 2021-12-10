package win

import (
	"syscall"
	"unsafe"

	"github.com/netfiredotnet/windigo/internal/proc"
	"github.com/netfiredotnet/windigo/internal/util"
	"github.com/netfiredotnet/windigo/win/co"
	"github.com/netfiredotnet/windigo/win/errco"
)

// A handle to an image list.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/image-lists
type HIMAGELIST HANDLE

// Usually flags is ILC_COLOR32.
//
// ⚠️ You must defer HIMAGELIST.Destroy().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_create
func ImageListCreate(
	cx, cy uint32, flags co.ILC, szInitial, szGrow uint32) HIMAGELIST {

	ret, _, err := syscall.Syscall6(proc.ImageList_Create.Addr(), 5,
		uintptr(cx), uintptr(cy), uintptr(flags),
		uintptr(szInitial), uintptr(szGrow), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HIMAGELIST(ret)
}

// If icon was loaded from resource with LoadIcon(), it doesn't need to be
// destroyed, because all icon resources are automatically freed.
// Otherwise, if loaded with CreateIcon(), it must be destroyed.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_addicon
func (hImg HIMAGELIST) AddIcon(hIcons ...HICON) {
	for _, hIco := range hIcons {
		hImg.ReplaceIcon(-1, hIco)
	}
}

// Loads icons from the shell, used by Windows Explorer to represent the given
// file extensions, like "mp3".
func (hImg HIMAGELIST) AddIconFromShell(fileExtensions ...string) {
	sz := hImg.GetIconSize()
	isIco16 := sz.Cx == 16 && sz.Cy == 16
	isIco32 := sz.Cx == 32 && sz.Cy == 32
	if !isIco16 && !isIco32 {
		panic("AddIconFromShell can load only 16x16 or 32x32 icons.")
	}

	shgfi := co.SHGFI_USEFILEATTRIBUTES | co.SHGFI_ICON |
		util.Iif(isIco32, co.SHGFI_LARGEICON, co.SHGFI_SMALLICON).(co.SHGFI)

	fi := SHFILEINFO{}
	for _, fileExt := range fileExtensions {
		SHGetFileInfo("*."+fileExt, co.FILE_ATTRIBUTE_NORMAL, &fi, shgfi)
		hImg.AddIcon(fi.HIcon)
		fi.HIcon.DestroyIcon()
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_destroy
func (hImg HIMAGELIST) Destroy() {
	// http://www.catch22.net/tuts/win32/system-image-list
	// https://www.autohotkey.com/docs/commands/ListView.htm
	ret, _, err := syscall.Syscall(proc.ImageList_Destroy.Addr(), 1,
		uintptr(hImg), 0, 0)
	if ret == 0 && errco.ERROR(err) != errco.SUCCESS {
		panic(errco.ERROR(err))
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_geticonsize
func (hImg HIMAGELIST) GetIconSize() SIZE {
	var sz SIZE
	ret, _, err := syscall.Syscall(proc.ImageList_GetIconSize.Addr(), 3,
		uintptr(hImg),
		uintptr(unsafe.Pointer(&sz.Cx)), uintptr(unsafe.Pointer(&sz.Cy)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return sz
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_getimagecount
func (hImg HIMAGELIST) GetImageCount() uint32 {
	ret, _, _ := syscall.Syscall(proc.ImageList_GetImageCount.Addr(), 1,
		uintptr(hImg), 0, 0)
	return uint32(ret)
}

// If icon was loaded from resource with LoadIcon(), it doesn't need to be
// destroyed, because all icon resources are automatically freed.
// Otherwise, if loaded with CreateIcon(), it must be destroyed.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_replaceicon
func (hImg HIMAGELIST) ReplaceIcon(i int32, hIcon HICON) int32 {
	ret, _, err := syscall.Syscall(proc.ImageList_ReplaceIcon.Addr(), 3,
		uintptr(hImg), uintptr(i), uintptr(hIcon))
	if int(ret) == -1 {
		panic(errco.ERROR(err))
	}
	return int32(ret)
}
