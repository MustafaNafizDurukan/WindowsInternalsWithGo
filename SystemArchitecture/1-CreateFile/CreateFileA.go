func StringToCharPtr(str string) *uint8 {
	chars := append([]byte(str), 0) // null terminated
	return &chars[0]
}

const (
	InvalidHandle = ^Handle(0)
)

var (
	modkernel32     = windows.NewLazySystemDLL("kernel32.dll")
	procCreateFileA = modkernel32.NewProc("CreateFileA")
)

/*
HANDLE CreateFileA(
  [in]           LPCSTR                lpFileName,
  [in]           DWORD                 dwDesiredAccess,
  [in]           DWORD                 dwShareMode,
  [in, optional] LPSECURITY_ATTRIBUTES lpSecurityAttributes,
  [in]           DWORD                 dwCreationDisposition,
  [in]           DWORD                 dwFlagsAndAttributes,
  [in, optional] HANDLE                hTemplateFile
);
*/

func CreateFileA(name *uint8, access uint32, mode uint32, sa *windows.SecurityAttributes, createmode uint32, attrs uint32, templatefile Handle) (handle Handle, e1 error) {
	r0, _, e1 := syscall.Syscall9(procCreateFileA.Addr(), 7, uintptr(unsafe.Pointer(name)), uintptr(access), uintptr(mode), uintptr(unsafe.Pointer(sa)), uintptr(createmode), uintptr(attrs), uintptr(templatefile), 0, 0)
	handle = Handle(r0)
	return
}

type Handle uintptr

func main() {
	filePath := StringToCharPtr(`example.txt`)

	processMonitorHandle, err := CreateFileA(
		filePath,
		windows.GENERIC_WRITE,
		0,
		nil,
		windows.CREATE_ALWAYS,
		0,
		0)
	if err != windows.ERROR_SUCCESS {
		fmt.Printf("CreateFile failed: %v", err)
		return
	}

	fmt.Println(processMonitorHandle)
}