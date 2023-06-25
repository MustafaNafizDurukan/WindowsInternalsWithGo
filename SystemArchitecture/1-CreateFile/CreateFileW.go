package main

import (
	"fmt"

	"golang.org/x/sys/windows"
)

func main() {
	filePath, err := windows.UTF16PtrFromString(`example.txt`)
	if nil != err {
		fmt.Printf("init failed: %v", err)
		return
	}

	processMonitorHandle, err := windows.CreateFile(
		filePath,
		windows.GENERIC_WRITE,
		0,
		nil,
		windows.CREATE_ALWAYS,
		0,
		0)
	if nil != err {
		fmt.Printf("CreateFile failed: %v", err)
		return
	}

	fmt.Println(processMonitorHandle)
}
