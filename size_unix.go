//go:build darwin || dragonfly || freebsd || linux || nacl || netbsd || openbsd || solaris

package tutils

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix"
)

type winsize struct {
	row uint16
	col uint16
	x   uint16
	y   uint16
}

func getSize() (Size, error) {
	size := winsize{}

	_, _, err := syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(unix.TIOCGWINSZ),
		uintptr(unsafe.Pointer(&size)),
	)

	if err != 0 {
		return Size{}, err
	}

	return Size{
		Rows:    int(size.row),
		Columns: int(size.col),
	}, nil
}
