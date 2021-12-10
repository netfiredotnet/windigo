package win

import (
	"syscall"
	"unsafe"

	"github.com/netfiredotnet/windigo/internal/proc"
	"github.com/netfiredotnet/windigo/win/errco"
)

// A handle returned by FindFirstFile() function.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-findfirstfilew
type HFIND HANDLE

// Returns true if a file was found.
//
// ⚠️ You must defer HFIND.FindClose().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-findfirstfilew
func FindFirstFile(fileName string,
	findFileData *WIN32_FIND_DATA) (HFIND, bool, error) {

	ret, _, err := syscall.Syscall(proc.FindFirstFile.Addr(), 2,
		uintptr(unsafe.Pointer(Str.ToNativePtr(fileName))),
		uintptr(unsafe.Pointer(findFileData)), 0)

	if int(ret) == _INVALID_HANDLE_VALUE {
		if wErr := errco.ERROR(err); wErr == errco.FILE_NOT_FOUND || wErr == errco.PATH_NOT_FOUND {
			return HFIND(0), false, nil // no matching files, not an error
		} else {
			return HFIND(0), false, wErr
		}
	}
	return HFIND(ret), true, nil // a file was found
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-findclose
func (hFind HFIND) FindClose() {
	ret, _, err := syscall.Syscall(proc.FindClose.Addr(), 1,
		uintptr(hFind), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// Returns true if a file was found.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-findnextfilew
func (hFind HFIND) FindNextFile(findFileData *WIN32_FIND_DATA) (bool, error) {
	ret, _, err := syscall.Syscall(proc.FindNextFile.Addr(), 2,
		uintptr(hFind), uintptr(unsafe.Pointer(findFileData)), 0)

	if ret == 0 {
		if wErr := errco.ERROR(err); wErr == errco.NO_MORE_FILES {
			return false, nil // not an error, search ended
		} else {
			return false, wErr
		}
	}
	return true, nil // a file was found
}
