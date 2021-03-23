package proc

import (
	"syscall"
)

var (
	kernel32 = syscall.NewLazyDLL("kernel32.dll")

	CloseHandle                     = kernel32.NewProc("CloseHandle")
	CreateDirectory                 = kernel32.NewProc("CreateDirectoryW")
	CreateFile                      = kernel32.NewProc("CreateFileW")
	CreateFileMapping               = kernel32.NewProc("CreateFileMappingW")
	DeleteFile                      = kernel32.NewProc("DeleteFileW")
	FileTimeToSystemTime            = kernel32.NewProc("FileTimeToSystemTime")
	FindClose                       = kernel32.NewProc("FindClose")
	FindFirstFile                   = kernel32.NewProc("FindFirstFileW")
	FindNextFile                    = kernel32.NewProc("FindNextFileW")
	GetCurrentProcessId             = kernel32.NewProc("GetCurrentProcessId")
	GetCurrentThreadId              = kernel32.NewProc("GetCurrentThreadId")
	GetFileAttributes               = kernel32.NewProc("GetFileAttributesW")
	GetFileSizeEx                   = kernel32.NewProc("GetFileSizeEx")
	GetModuleHandle                 = kernel32.NewProc("GetModuleHandleW")
	GetSystemTime                   = kernel32.NewProc("GetSystemTime")
	MapViewOfFile                   = kernel32.NewProc("MapViewOfFile")
	MulDiv                          = kernel32.NewProc("MulDiv")
	ReadFile                        = kernel32.NewProc("ReadFile")
	SetEndOfFile                    = kernel32.NewProc("SetEndOfFile")
	SetFilePointerEx                = kernel32.NewProc("SetFilePointerEx")
	Sleep                           = kernel32.NewProc("Sleep")
	SystemTimeToFileTime            = kernel32.NewProc("SystemTimeToFileTime")
	SystemTimeToTzSpecificLocalTime = kernel32.NewProc("SystemTimeToTzSpecificLocalTime")
	TzSpecificLocalTimeToSystemTime = kernel32.NewProc("TzSpecificLocalTimeToSystemTime")
	UnmapViewOfFile                 = kernel32.NewProc("UnmapViewOfFile")
	VerifyVersionInfo               = kernel32.NewProc("VerifyVersionInfoW")
	VerSetConditionMask             = kernel32.NewProc("VerSetConditionMask")
	WriteFile                       = kernel32.NewProc("WriteFile")
)
