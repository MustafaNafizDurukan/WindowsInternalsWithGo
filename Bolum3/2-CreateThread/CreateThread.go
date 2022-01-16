package main

import (
	"fmt"
	"syscall"

	"golang.org/x/sys/windows"
)

func main() {
	r1, _, _ := procCreateThread.Call(0, 0, syscall.NewCallback(ThreadProc), 0, 0, 0)
	h := syscall.Handle(r1)
	syscall.WaitForSingleObject(h, syscall.INFINITE)
	syscall.CloseHandle(h)
}

func ThreadProc(p uintptr) uintptr {
	fmt.Println("FOO")
	return 0
}

var (
	modkernel32      = windows.NewLazySystemDLL("kernel32.dll")
	procCreateThread = modkernel32.NewProc("CreateThread")
)

type Handle uintptr

/*
HANDLE CreateThread(
  [in, optional]  LPSECURITY_ATTRIBUTES   lpThreadAttributes,
  [in]            SIZE_T                  dwStackSize,
  [in]            LPTHREAD_START_ROUTINE  lpStartAddress,
  [in, optional]  __drv_aliasesMem LPVOID lpParameter,
  [in]            DWORD                   dwCreationFlags,
  [out, optional] LPDWORD                 lpThreadId
);
*/

// func CreateThread(sa *windows.SecurityAttributes, st syscall.SIZE_T) (handle Handle, e1 error) {
// 	r0, _, e1 := procCreateThread.Call(uintptr(unsafe.Pointer(sa)))

// 	windows.
// 		handle = Handle(r0)
// 	return
// }
