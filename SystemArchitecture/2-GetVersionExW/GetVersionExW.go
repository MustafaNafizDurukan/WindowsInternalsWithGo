package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

var (
	modkernel32       = syscall.NewLazyDLL("kernel32.dll")
	procGetVersionExW = modkernel32.NewProc("GetVersionExW")
)

func main() {
	var osvi OSVERSIONINFOEX
	osvi.OSVersionInfoSize = uint32(unsafe.Sizeof(osvi))
	if err := GetVersionEx(&osvi); err != nil {
		return
	}
	fmt.Println(osvi)
}

type OSVERSIONINFO struct {
	OSVersionInfoSize uint32
	MajorVersion      uint32
	MinorVersion      uint32
	BuildNumber       uint32
	PlatformId        uint32
	CSDVersion        [128]uint16
}

type OSVERSIONINFOEX struct {
	OSVERSIONINFO
	ServicePackMajor uint16
	ServicePackMinor uint16
	SuiteMask        uint16
	ProductType      uint8
	Reserved         uint8
}

func GetVersionEx(osvi *OSVERSIONINFOEX) error {
	r1, _, e1 := syscall.Syscall(procGetVersionExW.Addr(), 1, uintptr(unsafe.Pointer(osvi)), 0, 0)
	if r1 == 0 {
		return e1
	}
	return nil
}
