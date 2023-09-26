// Copyright 2013 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// A simulated network for use within NaCl.
// The simulation is not particularly tied to NaCl,
// but other systems have real networks.

package syscall

const (
	AF_UNSPEC = iota
	AF_UNIX
	AF_INET
	AF_INET6
)

const (
	SHUT_RD = iota
	SHUT_WR
	SHUT_RDWR
)

const (
	SOCK_STREAM = 1 + iota
	SOCK_DGRAM
	SOCK_RAW
	SOCK_SEQPACKET
)

const (
	IPPROTO_IP   = 0
	IPPROTO_IPV4 = 4
	IPPROTO_IPV6 = 0x29
	IPPROTO_TCP  = 6
	IPPROTO_UDP  = 0x11
)

// Misc constants expected by package net but not supported.
const (
	_ = iota
	SOL_SOCKET
	SO_TYPE
	NET_RT_IFLIST
	IFNAMSIZ
	IFF_UP
	IFF_BROADCAST
	IFF_LOOPBACK
	IFF_POINTOPOINT
	IFF_MULTICAST
	IPV6_V6ONLY
	SOMAXCONN
	F_DUPFD_CLOEXEC
	SO_BROADCAST
	SO_REUSEADDR
	SO_REUSEPORT
	SO_RCVBUF
	SO_SNDBUF
	SO_KEEPALIVE
	SO_LINGER
	SO_ERROR
	IP_PORTRANGE
	IP_PORTRANGE_DEFAULT
	IP_PORTRANGE_LOW
	IP_PORTRANGE_HIGH
	IP_MULTICAST_IF
	IP_MULTICAST_LOOP
	IP_ADD_MEMBERSHIP
	IPV6_PORTRANGE
	IPV6_PORTRANGE_DEFAULT
	IPV6_PORTRANGE_LOW
	IPV6_PORTRANGE_HIGH
	IPV6_MULTICAST_IF
	IPV6_MULTICAST_LOOP
	IPV6_JOIN_GROUP
	TCP_NODELAY
	TCP_KEEPINTVL
	TCP_KEEPIDLE

	SYS_FCNTL = 500
)

var SocketDisableIPv6 bool

// A Sockaddr is one of the SockaddrXxx structs.
type Sockaddr interface {
	copy() Sockaddr

	key() interface{}
}

type SockaddrInet4 struct {
	Port int
	Addr [4]byte
}

type SockaddrInet6 struct {
	Port   int
	ZoneId uint32
	Addr   [16]byte
}

type SockaddrUnix struct {
	Name string
}

type SockaddrDatalink struct {
	Len    uint8
	Family uint8
	Index  uint16
	Type   uint8
	Nlen   uint8
	Alen   uint8
	Slen   uint8
	Data   [12]int8
}

// RoutingMessage represents a routing message.
type RoutingMessage interface {
	unimplemented()
}

type IPMreq struct {
	Multiaddr [4]byte
	Interface [4]byte
}

type IPv6Mreq struct {
	Multiaddr [16]byte
	Interface uint32
}

type Linger struct {
	Onoff  int32
	Linger int32
}

type ICMPv6Filter struct {
	Filt [8]uint32
}

// A queue is the bookkeeping for a synchronized buffered queue.
// We do not use channels because we need to be able to handle
// writes after and during close, and because a chan byte would
// require too many send and receive operations in real use.

// A byteq is a byte queue.

// A msgq is a queue of messages.

// An addr is a sequence of bytes uniquely identifying a network address.
// It is not human-readable.

// A conn is one side of a stream-based network connection.
// That is, a stream-based network connection is a pair of cross-connected conns.

// A pktconn is one side of a packet-based network connection.
// That is, a packet-based network connection is a pair of cross-connected pktconns.

// A listener accepts incoming stream-based network connections.

// A netFile is an open network file.

// A netAddr is a network address in the global listener map.
// All the fields must have defined == operations.

// net records the state of the network.
// It maps a network address to the listener on that address.

// TODO(rsc): Some day, do a better job with port allocation.
// For playground programs, incrementing is fine.

// A netproto contains protocol-specific functionality
// (one for AF_INET, one for AF_INET6 and so on).
// It is a struct instead of an interface because the
// implementation needs no state, and I expect to
// add some data fields at some point.

func Socket(proto, sotype, unused int) (fd int, err error)

func Bind(fd int, sa Sockaddr) error

func StopIO(fd int) error

func Listen(fd int, backlog int) error

func Accept(fd int) (newfd int, sa Sockaddr, err error)

func Getsockname(fd int) (sa Sockaddr, err error)

func Getpeername(fd int) (sa Sockaddr, err error)

func Connect(fd int, sa Sockaddr) error

func Recvfrom(fd int, p []byte, flags int) (n int, from Sockaddr, err error)

func Sendto(fd int, p []byte, flags int, to Sockaddr) error

func Recvmsg(fd int, p, oob []byte, flags int) (n, oobn, recvflags int, from Sockaddr, err error)

func Sendmsg(fd int, p, oob []byte, to Sockaddr, flags int) error

func SendmsgN(fd int, p, oob []byte, to Sockaddr, flags int) (n int, err error)

func GetsockoptInt(fd, level, opt int) (value int, err error)

func SetsockoptInt(fd, level, opt int, value int) error

func SetsockoptByte(fd, level, opt int, value byte) error

func SetsockoptLinger(fd, level, opt int, l *Linger) error

func SetReadDeadline(fd int, t int64) error

func SetWriteDeadline(fd int, t int64) error

func Shutdown(fd int, how int) error

func SetsockoptICMPv6Filter(fd, level, opt int, filter *ICMPv6Filter) error
func SetsockoptIPMreq(fd, level, opt int, mreq *IPMreq) error
func SetsockoptIPv6Mreq(fd, level, opt int, mreq *IPv6Mreq) error
func SetsockoptInet4Addr(fd, level, opt int, value [4]byte) error
func SetsockoptString(fd, level, opt int, s string) error
func SetsockoptTimeval(fd, level, opt int, tv *Timeval) error
func Socketpair(domain, typ, proto int) (fd [2]int, err error)

func SetNonblock(fd int, nonblocking bool) error
