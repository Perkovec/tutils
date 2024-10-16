//go:build darwin || dragonfly || freebsd || linux || nacl || netbsd || openbsd || solaris

package tutils

import (
	"syscall"

	"golang.org/x/sys/unix"
)

type winsize struct {
	row uint16
	col uint16
	x   uint16
	y   uint16
}

func getSize() (Size, error) {
	size, err := unix.IoctlGetWinsize(syscall.Stdout, unix.TIOCGWINSZ)
	if err != nil {
		return Size{}, err
	}

	return Size{
		Rows:    int(size.Row),
		Columns: int(size.Col),
	}, nil
}
