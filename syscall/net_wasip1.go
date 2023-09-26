// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build wasip1

package syscall

const (
	SHUT_RD   = 0x1
	SHUT_WR   = 0x2
	SHUT_RDWR = SHUT_RD | SHUT_WR
)

func Socket(proto, sotype, unused int) (fd int, err error)

func Bind(fd int, sa Sockaddr) error

func StopIO(fd int) error

func Listen(fd int, backlog int) error

func Accept(fd int) (int, Sockaddr, error)

func Connect(fd int, sa Sockaddr) error

func Recvfrom(fd int, p []byte, flags int) (n int, from Sockaddr, err error)

func Sendto(fd int, p []byte, flags int, to Sockaddr) error

func Recvmsg(fd int, p, oob []byte, flags int) (n, oobn, recvflags int, from Sockaddr, err error)

func SendmsgN(fd int, p, oob []byte, to Sockaddr, flags int) (n int, err error)

func GetsockoptInt(fd, level, opt int) (value int, err error)

func SetsockoptInt(fd, level, opt int, value int) error

func SetReadDeadline(fd int, t int64) error

func SetWriteDeadline(fd int, t int64) error

func Shutdown(fd int, how int) error
