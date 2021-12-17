package main

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

func main() {
	notepad, err := windows.UTF16PtrFromString(`notepad.exe`)
	if nil != err {
		fmt.Printf("init failed: %v", err)
		return
	}

	var si windows.StartupInfo
	si.Cb = uint32(unsafe.Sizeof(si))
	var pi windows.ProcessInformation

	success := windows.CreateProcess(nil, notepad, nil, nil, false, 0, nil, nil, &si, &pi)
	if success != nil {
		fmt.Println("Creation failed: ", windows.GetLastError())
		return
	}
	fmt.Println("PID: ", pi.ProcessId)
	fmt.Println("TID: ", pi.ThreadId)
	fmt.Println("Process handle: ", pi.Process)

	windows.WaitForSingleObject(pi.Process, windows.INFINITE)

	var code uint32
	windows.GetExitCodeProcess(pi.Process, &code)
	fmt.Println("Notepad has exited. Exit code: ", code)

}
