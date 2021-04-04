//+build !windows

package cmd

import (
	"syscall"
	"unsafe"
)

// inject injects the string into the parent processes
// input queue (effectively typing it out into the tty)
//
// used to inject `cd` commands into the parent shell
func inject(str string) {
	for _, char := range str {
		syscall.Syscall(
			syscall.SYS_IOCTL,
			0,
			syscall.TIOCSTI,
			uintptr(unsafe.Pointer(&char)),
		)
	}
}
