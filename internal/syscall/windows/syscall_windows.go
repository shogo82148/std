// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package windows

import (
	"github.com/shogo82148/std/syscall"
	"github.com/shogo82148/std/unsafe"
)

// CanUseLongPaths is true when the OS supports opting into
// proper long path handling without the need for fixups.
//
//go:linkname CanUseLongPaths
var CanUseLongPaths bool

// UTF16PtrToString is like UTF16ToString, but takes *uint16
// as a parameter instead of []uint16.
func UTF16PtrToString(p *uint16) string

const (
	ERROR_INVALID_HANDLE         syscall.Errno = 6
	ERROR_BAD_LENGTH             syscall.Errno = 24
	ERROR_SHARING_VIOLATION      syscall.Errno = 32
	ERROR_LOCK_VIOLATION         syscall.Errno = 33
	ERROR_NOT_SUPPORTED          syscall.Errno = 50
	ERROR_CALL_NOT_IMPLEMENTED   syscall.Errno = 120
	ERROR_INVALID_NAME           syscall.Errno = 123
	ERROR_LOCK_FAILED            syscall.Errno = 167
<<<<<<< HEAD
=======
	ERROR_IO_INCOMPLETE          syscall.Errno = 996
	ERROR_NO_TOKEN               syscall.Errno = 1008
>>>>>>> upstream/release-branch.go1.25
	ERROR_NO_UNICODE_TRANSLATION syscall.Errno = 1113
	ERROR_CANT_ACCESS_FILE       syscall.Errno = 1920
)

const (
	GAA_FLAG_INCLUDE_PREFIX   = 0x00000010
	GAA_FLAG_INCLUDE_GATEWAYS = 0x0080
)

const (
	IF_TYPE_OTHER              = 1
	IF_TYPE_ETHERNET_CSMACD    = 6
	IF_TYPE_ISO88025_TOKENRING = 9
	IF_TYPE_PPP                = 23
	IF_TYPE_SOFTWARE_LOOPBACK  = 24
	IF_TYPE_ATM                = 37
	IF_TYPE_IEEE80211          = 71
	IF_TYPE_TUNNEL             = 131
	IF_TYPE_IEEE1394           = 144
)

type SocketAddress struct {
	Sockaddr       *syscall.RawSockaddrAny
	SockaddrLength int32
}

type IpAdapterUnicastAddress struct {
	Length             uint32
	Flags              uint32
	Next               *IpAdapterUnicastAddress
	Address            SocketAddress
	PrefixOrigin       int32
	SuffixOrigin       int32
	DadState           int32
	ValidLifetime      uint32
	PreferredLifetime  uint32
	LeaseLifetime      uint32
	OnLinkPrefixLength uint8
}

type IpAdapterAnycastAddress struct {
	Length  uint32
	Flags   uint32
	Next    *IpAdapterAnycastAddress
	Address SocketAddress
}

type IpAdapterMulticastAddress struct {
	Length  uint32
	Flags   uint32
	Next    *IpAdapterMulticastAddress
	Address SocketAddress
}

type IpAdapterDnsServerAdapter struct {
	Length   uint32
	Reserved uint32
	Next     *IpAdapterDnsServerAdapter
	Address  SocketAddress
}

type IpAdapterPrefix struct {
	Length       uint32
	Flags        uint32
	Next         *IpAdapterPrefix
	Address      SocketAddress
	PrefixLength uint32
}

type IpAdapterWinsServerAddress struct {
	Length   uint32
	Reserved uint32
	Next     *IpAdapterWinsServerAddress
	Address  SocketAddress
}

type IpAdapterGatewayAddress struct {
	Length   uint32
	Reserved uint32
	Next     *IpAdapterGatewayAddress
	Address  SocketAddress
}

type IpAdapterAddresses struct {
	Length                 uint32
	IfIndex                uint32
	Next                   *IpAdapterAddresses
	AdapterName            *byte
	FirstUnicastAddress    *IpAdapterUnicastAddress
	FirstAnycastAddress    *IpAdapterAnycastAddress
	FirstMulticastAddress  *IpAdapterMulticastAddress
	FirstDnsServerAddress  *IpAdapterDnsServerAdapter
	DnsSuffix              *uint16
	Description            *uint16
	FriendlyName           *uint16
	PhysicalAddress        [syscall.MAX_ADAPTER_ADDRESS_LENGTH]byte
	PhysicalAddressLength  uint32
	Flags                  uint32
	Mtu                    uint32
	IfType                 uint32
	OperStatus             uint32
	Ipv6IfIndex            uint32
	ZoneIndices            [16]uint32
	FirstPrefix            *IpAdapterPrefix
	TransmitLinkSpeed      uint64
	ReceiveLinkSpeed       uint64
	FirstWinsServerAddress *IpAdapterWinsServerAddress
	FirstGatewayAddress    *IpAdapterGatewayAddress
}

type SecurityAttributes struct {
	Length             uint16
	SecurityDescriptor uintptr
	InheritHandle      bool
}

type FILE_BASIC_INFO struct {
	CreationTime   int64
	LastAccessTime int64
	LastWriteTime  int64
	ChangedTime    int64
	FileAttributes uint32

	// Pad out to 8-byte alignment.
	//
	// Without this padding, TestChmod fails due to an argument validation error
	// in SetFileInformationByHandle on windows/386.
	//
	// https://learn.microsoft.com/en-us/cpp/build/reference/zp-struct-member-alignment?view=msvc-170
	// says that “The C/C++ headers in the Windows SDK assume the platform's
	// default alignment is used.” What we see here is padding rather than
	// alignment, but maybe it is related.
	_ uint32
}

const (
	IfOperStatusUp             = 1
	IfOperStatusDown           = 2
	IfOperStatusTesting        = 3
	IfOperStatusUnknown        = 4
	IfOperStatusDormant        = 5
	IfOperStatusNotPresent     = 6
	IfOperStatusLowerLayerDown = 7
)

const (
	// flags for CreateToolhelp32Snapshot
	TH32CS_SNAPMODULE   = 0x08
	TH32CS_SNAPMODULE32 = 0x10
)

const MAX_MODULE_NAME32 = 255

type ModuleEntry32 struct {
	Size         uint32
	ModuleID     uint32
	ProcessID    uint32
	GlblcntUsage uint32
	ProccntUsage uint32
	ModBaseAddr  uintptr
	ModBaseSize  uint32
	ModuleHandle syscall.Handle
	Module       [MAX_MODULE_NAME32 + 1]uint16
	ExePath      [syscall.MAX_PATH]uint16
}

const SizeofModuleEntry32 = unsafe.Sizeof(ModuleEntry32{})

const (
	WSA_FLAG_OVERLAPPED        = 0x01
	WSA_FLAG_NO_HANDLE_INHERIT = 0x80

	WSAEINVAL       syscall.Errno = 10022
	WSAEMSGSIZE     syscall.Errno = 10040
	WSAEAFNOSUPPORT syscall.Errno = 10047

	MSG_PEEK   = 0x2
	MSG_TRUNC  = 0x0100
	MSG_CTRUNC = 0x0200
)

var WSAID_WSASENDMSG = syscall.GUID{
	Data1: 0xa441e712,
	Data2: 0x754f,
	Data3: 0x43ca,
	Data4: [8]byte{0x84, 0xa7, 0x0d, 0xee, 0x44, 0xcf, 0x60, 0x6d},
}

var WSAID_WSARECVMSG = syscall.GUID{
	Data1: 0xf689d7c8,
	Data2: 0x6f1f,
	Data3: 0x436b,
	Data4: [8]byte{0x8a, 0x53, 0xe5, 0x4f, 0xe3, 0x51, 0xc3, 0x22},
}

type WSAMsg struct {
	Name        syscall.Pointer
	Namelen     int32
	Buffers     *syscall.WSABuf
	BufferCount uint32
	Control     syscall.WSABuf
	Flags       uint32
}

func WSASendMsg(fd syscall.Handle, msg *WSAMsg, flags uint32, bytesSent *uint32, overlapped *syscall.Overlapped, croutine *byte) error

func WSARecvMsg(fd syscall.Handle, msg *WSAMsg, bytesReceived *uint32, overlapped *syscall.Overlapped, croutine *byte) error

const (
	ComputerNameNetBIOS                   = 0
	ComputerNameDnsHostname               = 1
	ComputerNameDnsDomain                 = 2
	ComputerNameDnsFullyQualified         = 3
	ComputerNamePhysicalNetBIOS           = 4
	ComputerNamePhysicalDnsHostname       = 5
	ComputerNamePhysicalDnsDomain         = 6
	ComputerNamePhysicalDnsFullyQualified = 7
	ComputerNameMax                       = 8

	MOVEFILE_REPLACE_EXISTING      = 0x1
	MOVEFILE_COPY_ALLOWED          = 0x2
	MOVEFILE_DELAY_UNTIL_REBOOT    = 0x4
	MOVEFILE_WRITE_THROUGH         = 0x8
	MOVEFILE_CREATE_HARDLINK       = 0x10
	MOVEFILE_FAIL_IF_NOT_TRACKABLE = 0x20
)

func Rename(oldpath, newpath string) error

const (
	LOCKFILE_FAIL_IMMEDIATELY = 0x00000001
	LOCKFILE_EXCLUSIVE_LOCK   = 0x00000002
)

const MB_ERR_INVALID_CHARS = 8

// Constants from lmshare.h
const (
	STYPE_DISKTREE  = 0x00
	STYPE_TEMPORARY = 0x40000000
)

type SHARE_INFO_2 struct {
	Netname     *uint16
	Type        uint32
	Remark      *uint16
	Permissions uint32
	MaxUses     uint32
	CurrentUses uint32
	Path        *uint16
	Passwd      *uint16
}

const (
	FILE_NAME_NORMALIZED = 0x0
	FILE_NAME_OPENED     = 0x8

	VOLUME_NAME_DOS  = 0x0
	VOLUME_NAME_GUID = 0x1
	VOLUME_NAME_NONE = 0x4
	VOLUME_NAME_NT   = 0x2
)

func ErrorLoadingGetTempPath2() error

type FILE_ID_BOTH_DIR_INFO struct {
	NextEntryOffset uint32
	FileIndex       uint32
	CreationTime    syscall.Filetime
	LastAccessTime  syscall.Filetime
	LastWriteTime   syscall.Filetime
	ChangeTime      syscall.Filetime
	EndOfFile       uint64
	AllocationSize  uint64
	FileAttributes  uint32
	FileNameLength  uint32
	EaSize          uint32
	ShortNameLength uint32
	ShortName       [12]uint16
	FileID          uint64
	FileName        [1]uint16
}

type FILE_FULL_DIR_INFO struct {
	NextEntryOffset uint32
	FileIndex       uint32
	CreationTime    syscall.Filetime
	LastAccessTime  syscall.Filetime
	LastWriteTime   syscall.Filetime
	ChangeTime      syscall.Filetime
	EndOfFile       uint64
	AllocationSize  uint64
	FileAttributes  uint32
	FileNameLength  uint32
	EaSize          uint32
	FileName        [1]uint16
}

type RUNTIME_FUNCTION struct {
	BeginAddress uint32
	EndAddress   uint32
	UnwindData   uint32
}

type SERVICE_STATUS struct {
	ServiceType             uint32
	CurrentState            uint32
	ControlsAccepted        uint32
	Win32ExitCode           uint32
	ServiceSpecificExitCode uint32
	CheckPoint              uint32
	WaitHint                uint32
}

const (
	SERVICE_RUNNING      = 4
	SERVICE_QUERY_STATUS = 4
)

func FinalPath(h syscall.Handle, flags uint32) (string, error)

// QueryPerformanceCounter retrieves the current value of performance counter.
//
//go:linkname QueryPerformanceCounter
func QueryPerformanceCounter() int64

// QueryPerformanceFrequency retrieves the frequency of the performance counter.
// The returned value is represented as counts per second.
//
//go:linkname QueryPerformanceFrequency
func QueryPerformanceFrequency() int64

const (
	PIPE_ACCESS_INBOUND  = 0x00000001
	PIPE_ACCESS_OUTBOUND = 0x00000002
	PIPE_ACCESS_DUPLEX   = 0x00000003

	PIPE_TYPE_BYTE    = 0x00000000
	PIPE_TYPE_MESSAGE = 0x00000004

	PIPE_READMODE_BYTE    = 0x00000000
	PIPE_READMODE_MESSAGE = 0x00000002
)

// NTStatus corresponds with NTSTATUS, error values returned by ntdll.dll and
// other native functions.
type NTStatus uint32

func (s NTStatus) Errno() syscall.Errno

func (s NTStatus) Error() string

// x/sys/windows/mkerrors.bash can generate a complete list of NTStatus codes.
//
// At the moment, we only need a couple, so just put them here manually.
// If this list starts getting long, we should consider generating the full set.
const (
	STATUS_OBJECT_NAME_COLLISION     NTStatus = 0xC0000035
	STATUS_FILE_IS_A_DIRECTORY       NTStatus = 0xC00000BA
	STATUS_DIRECTORY_NOT_EMPTY       NTStatus = 0xC0000101
	STATUS_NOT_A_DIRECTORY           NTStatus = 0xC0000103
	STATUS_CANNOT_DELETE             NTStatus = 0xC0000121
	STATUS_REPARSE_POINT_ENCOUNTERED NTStatus = 0xC000050B
)

const (
	FileModeInformation = 16
)

// https://learn.microsoft.com/en-us/windows-hardware/drivers/ddi/ntifs/ns-ntifs-_file_mode_information
type FILE_MODE_INFORMATION struct {
	Mode uint32
}
