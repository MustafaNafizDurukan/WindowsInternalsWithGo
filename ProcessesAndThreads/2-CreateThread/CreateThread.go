package main

import (
	"fmt"
	"math"
	"syscall"
	"time"
	"unsafe"

	"golang.org/x/sys/windows"
)

const (
	NUM_THREADS = 8
)

type PrimeData struct {
	first, last int
	count       int
}

func main() {
	first := 3
	last := 20000000

	data := make([]PrimeData, NUM_THREADS)
	delta := (last - first + 1) / NUM_THREADS

	hThread := make([]windows.Handle, NUM_THREADS)

	fmt.Printf("Working... %d\n", windows.GetCurrentThreadId())
	start := time.Now()

	var id uint32
	for i := 0; i < NUM_THREADS; i++ {
		data[i].first = first + i*delta

		if i == NUM_THREADS-1 {
			data[i].last = last
		} else {
			data[i].last = first + (i+1)*delta - 1
		}

		r1, _, _ := procCreateThread.Call(0, 0, syscall.NewCallback(CalcPrimes), uintptr(unsafe.Pointer(&data[i])), 0, uintptr(unsafe.Pointer(&id)))
		hThread[i] = windows.Handle(syscall.Handle(r1))
	}

	windows.WaitForMultipleObjects(hThread, true, windows.INFINITE)

	t := time.Now()
	elapsed := t.Sub(start)

	total := 0
	for i := 0; i < NUM_THREADS; i++ {
		fmt.Printf("Thread %d result: %d \n", i, data[i].count)
		total += data[i].count
	}
	fmt.Printf("\nTotal: %d\n", total)

	for i := 0; i < NUM_THREADS; i++ {
		windows.CloseHandle(hThread[i])
	}

	fmt.Println(elapsed)
}

func CalcPrimes(p uintptr) uintptr {
	data := (*PrimeData)(unsafe.Pointer(p))

	count := 0

	for i := data.first; i <= data.last; i++ {
		limit := (int)(math.Sqrt((float64)(i)))
		var j int
		for j = 2; j <= limit; j++ {
			if i%j == 0 {
				break
			}
		}

		if j == limit+1 {
			count++
		}
	}

	data.count = count
	return 0
}

var (
	modkernel32      = windows.NewLazySystemDLL("kernel32.dll")
	procCreateThread = modkernel32.NewProc("CreateThread")
)
