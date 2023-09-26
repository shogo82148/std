// Code generated by cmd/cgo -godefs; DO NOT EDIT.
// cgo -godefs -- -fsigned-char types_openbsd.go

package syscall

type Timespec struct {
	Sec       int64
	Nsec      int32
	Pad_cgo_0 [4]byte
}

type Timeval struct {
	Sec       int64
	Usec      int32
	Pad_cgo_0 [4]byte
}

type Rusage struct {
	Utime    Timeval
	Stime    Timeval
	Maxrss   int32
	Ixrss    int32
	Idrss    int32
	Isrss    int32
	Minflt   int32
	Majflt   int32
	Nswap    int32
	Inblock  int32
	Oublock  int32
	Msgsnd   int32
	Msgrcv   int32
	Nsignals int32
	Nvcsw    int32
	Nivcsw   int32
}

type Rlimit struct {
	Cur uint64
	Max uint64
}

const (
	S_IFMT   = 0xf000
	S_IFIFO  = 0x1000
	S_IFCHR  = 0x2000
	S_IFDIR  = 0x4000
	S_IFBLK  = 0x6000
	S_IFREG  = 0x8000
	S_IFLNK  = 0xa000
	S_IFSOCK = 0xc000
	S_ISUID  = 0x800
	S_ISGID  = 0x400
	S_ISVTX  = 0x200
	S_IRUSR  = 0x100
	S_IWUSR  = 0x80
	S_IXUSR  = 0x40
	S_IRWXG  = 0x38
	S_IRWXO  = 0x7
)

type Stat_t struct {
	Mode           uint32
	Dev            int32
	Ino            uint64
	Nlink          uint32
	Uid            uint32
	Gid            uint32
	Rdev           int32
	Atim           Timespec
	Mtim           Timespec
	Ctim           Timespec
	Size           int64
	Blocks         int64
	Blksize        int32
	Flags          uint32
	Gen            uint32
	Pad_cgo_0      [4]byte
	X__st_birthtim Timespec
}

type Statfs_t struct {
	F_flags       uint32
	F_bsize       uint32
	F_iosize      uint32
	Pad_cgo_0     [4]byte
	F_blocks      uint64
	F_bfree       uint64
	F_bavail      int64
	F_files       uint64
	F_ffree       uint64
	F_favail      int64
	F_syncwrites  uint64
	F_syncreads   uint64
	F_asyncwrites uint64
	F_asyncreads  uint64
	F_fsid        Fsid
	F_namemax     uint32
	F_owner       uint32
	F_ctime       uint64
	F_fstypename  [16]int8
	F_mntonname   [90]int8
	F_mntfromname [90]int8
	F_mntfromspec [90]int8
	Pad_cgo_1     [2]byte
	Mount_info    [160]byte
}

type Flock_t struct {
	Start  int64
	Len    int64
	Pid    int32
	Type   int16
	Whence int16
}

type Dirent struct {
	Fileno       uint64
	Off          int64
	Reclen       uint16
	Type         uint8
	Namlen       uint8
	X__d_padding [4]uint8
	Name         [256]int8
}

type Fsid struct {
	Val [2]int32
}

type RawSockaddrInet4 struct {
	Len    uint8
	Family uint8
	Port   uint16
	Addr   [4]byte
	Zero   [8]int8
}

type RawSockaddrInet6 struct {
	Len      uint8
	Family   uint8
	Port     uint16
	Flowinfo uint32
	Addr     [16]byte
	Scope_id uint32
}

type RawSockaddrUnix struct {
	Len    uint8
	Family uint8
	Path   [104]int8
}

type RawSockaddrDatalink struct {
	Len    uint8
	Family uint8
	Index  uint16
	Type   uint8
	Nlen   uint8
	Alen   uint8
	Slen   uint8
	Data   [24]int8
}

type RawSockaddr struct {
	Len    uint8
	Family uint8
	Data   [14]int8
}

type RawSockaddrAny struct {
	Addr RawSockaddr
	Pad  [92]int8
}

type Linger struct {
	Onoff  int32
	Linger int32
}

type Iovec struct {
	Base *byte
	Len  uint32
}

type IPMreq struct {
	Multiaddr [4]byte
	Interface [4]byte
}

type IPv6Mreq struct {
	Multiaddr [16]byte
	Interface uint32
}

type Msghdr struct {
	Name       *byte
	Namelen    uint32
	Iov        *Iovec
	Iovlen     uint32
	Control    *byte
	Controllen uint32
	Flags      int32
}

type Cmsghdr struct {
	Len   uint32
	Level int32
	Type  int32
}

type Inet6Pktinfo struct {
	Addr    [16]byte
	Ifindex uint32
}

type IPv6MTUInfo struct {
	Addr RawSockaddrInet6
	Mtu  uint32
}

type ICMPv6Filter struct {
	Filt [8]uint32
}

const (
	SizeofSockaddrInet4    = 0x10
	SizeofSockaddrInet6    = 0x1c
	SizeofSockaddrAny      = 0x6c
	SizeofSockaddrUnix     = 0x6a
	SizeofSockaddrDatalink = 0x20
	SizeofLinger           = 0x8
	SizeofIPMreq           = 0x8
	SizeofIPv6Mreq         = 0x14
	SizeofMsghdr           = 0x1c
	SizeofCmsghdr          = 0xc
	SizeofInet6Pktinfo     = 0x14
	SizeofIPv6MTUInfo      = 0x20
	SizeofICMPv6Filter     = 0x20
)

const (
	PTRACE_TRACEME = 0x0
	PTRACE_CONT    = 0x7
	PTRACE_KILL    = 0x8
)

type Kevent_t struct {
	Ident     uint32
	Filter    int16
	Flags     uint16
	Fflags    uint32
	Pad_cgo_0 [4]byte
	Data      int64
	Udata     *byte
	Pad_cgo_1 [4]byte
}

type FdSet struct {
	Bits [32]uint32
}

const (
	SizeofIfMsghdr         = 0xa8
	SizeofIfData           = 0x90
	SizeofIfaMsghdr        = 0x18
	SizeofIfAnnounceMsghdr = 0x1a
	SizeofRtMsghdr         = 0x60
	SizeofRtMetrics        = 0x38
)

type IfMsghdr struct {
	Msglen  uint16
	Version uint8
	Type    uint8
	Hdrlen  uint16
	Index   uint16
	Tableid uint16
	Pad1    uint8
	Pad2    uint8
	Addrs   int32
	Flags   int32
	Xflags  int32
	Data    IfData
}

type IfData struct {
	Type         uint8
	Addrlen      uint8
	Hdrlen       uint8
	Link_state   uint8
	Mtu          uint32
	Metric       uint32
	Rdomain      uint32
	Baudrate     uint64
	Ipackets     uint64
	Ierrors      uint64
	Opackets     uint64
	Oerrors      uint64
	Collisions   uint64
	Ibytes       uint64
	Obytes       uint64
	Imcasts      uint64
	Omcasts      uint64
	Iqdrops      uint64
	Oqdrops      uint64
	Noproto      uint64
	Capabilities uint32
	Pad_cgo_0    [4]byte
	Lastchange   Timeval
}

type IfaMsghdr struct {
	Msglen  uint16
	Version uint8
	Type    uint8
	Hdrlen  uint16
	Index   uint16
	Tableid uint16
	Pad1    uint8
	Pad2    uint8
	Addrs   int32
	Flags   int32
	Metric  int32
}

type IfAnnounceMsghdr struct {
	Msglen  uint16
	Version uint8
	Type    uint8
	Hdrlen  uint16
	Index   uint16
	What    uint16
	Name    [16]int8
}

type RtMsghdr struct {
	Msglen   uint16
	Version  uint8
	Type     uint8
	Hdrlen   uint16
	Index    uint16
	Tableid  uint16
	Priority uint8
	Mpls     uint8
	Addrs    int32
	Flags    int32
	Fmask    int32
	Pid      int32
	Seq      int32
	Errno    int32
	Inits    uint32
	Rmx      RtMetrics
}

type RtMetrics struct {
	Pksent   uint64
	Expire   int64
	Locks    uint32
	Mtu      uint32
	Refcnt   uint32
	Hopcount uint32
	Recvpipe uint32
	Sendpipe uint32
	Ssthresh uint32
	Rtt      uint32
	Rttvar   uint32
	Pad      uint32
}

type Mclpool struct{}

const (
	SizeofBpfVersion = 0x4
	SizeofBpfStat    = 0x8
	SizeofBpfProgram = 0x8
	SizeofBpfInsn    = 0x8
	SizeofBpfHdr     = 0x14
)

type BpfVersion struct {
	Major uint16
	Minor uint16
}

type BpfStat struct {
	Recv uint32
	Drop uint32
}

type BpfProgram struct {
	Len   uint32
	Insns *BpfInsn
}

type BpfInsn struct {
	Code uint16
	Jt   uint8
	Jf   uint8
	K    uint32
}

type BpfHdr struct {
	Tstamp    BpfTimeval
	Caplen    uint32
	Datalen   uint32
	Hdrlen    uint16
	Pad_cgo_0 [2]byte
}

type BpfTimeval struct {
	Sec  uint32
	Usec uint32
}

type Termios struct {
	Iflag  uint32
	Oflag  uint32
	Cflag  uint32
	Lflag  uint32
	Cc     [20]uint8
	Ispeed int32
	Ospeed int32
}
