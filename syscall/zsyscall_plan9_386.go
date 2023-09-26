// mksyscall.pl -l32 -plan9 -tags plan9,386 syscall_plan9.go
// Code generated by the command above; DO NOT EDIT.

//go:build plan9 && 386
// +build plan9,386

package syscall

func Dup(oldfd int, newfd int) (fd int, err error)

func Pread(fd int, p []byte, offset int64) (n int, err error)

func Pwrite(fd int, p []byte, offset int64) (n int, err error)

func Close(fd int) (err error)

func Fstat(fd int, edir []byte) (n int, err error)

func Fwstat(fd int, edir []byte) (err error)
