// mksyscall.pl -l32 -plan9 syscall_plan9.go
// MACHINE GENERATED BY THE COMMAND ABOVE; DO NOT EDIT

//go:build amd64 && plan9
// +build amd64,plan9

package syscall

func Dup(oldfd int, newfd int) (fd int, err error)

func Pread(fd int, p []byte, offset int64) (n int, err error)

func Pwrite(fd int, p []byte, offset int64) (n int, err error)

func Close(fd int) (err error)

func Fstat(fd int, edir []byte) (n int, err error)

func Fwstat(fd int, edir []byte) (err error)
