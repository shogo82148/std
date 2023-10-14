// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unix

import (
	"github.com/shogo82148/std/syscall"
)

const (
	AI_CANONNAME = 0x2
	AI_ALL       = 0x100
	AI_V4MAPPED  = 0x800
	AI_MASK      = 0x1407

	EAI_AGAIN    = 2
	EAI_NODATA   = 7
	EAI_NONAME   = 8
	EAI_SERVICE  = 9
	EAI_SYSTEM   = 11
	EAI_OVERFLOW = 14

	NI_NAMEREQD = 4
)

type Addrinfo struct {
	Flags     int32
	Family    int32
	Socktype  int32
	Protocol  int32
	Addrlen   uint32
	Canonname *byte
	Addr      *syscall.RawSockaddr
	Next      *Addrinfo
}

func Getaddrinfo(hostname, servname *byte, hints *Addrinfo, res **Addrinfo) (int, error)

func Freeaddrinfo(ai *Addrinfo)

func Getnameinfo(sa *syscall.RawSockaddr, salen int, host *byte, hostlen int, serv *byte, servlen int, flags int) (int, error)

func GaiStrerror(ecode int) string

func GoString(p *byte) string

type ResState struct {
	unexported [69]uintptr
}

func ResNinit(state *ResState) error

func ResNclose(state *ResState)

func ResNsearch(state *ResState, dname *byte, class, typ int, ans *byte, anslen int) (int, error)
