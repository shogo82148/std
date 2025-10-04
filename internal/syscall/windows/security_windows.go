// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package windows

import (
	"github.com/shogo82148/std/syscall"
)

const (
	SecurityAnonymous      = 0
	SecurityIdentification = 1
	SecurityImpersonation  = 2
	SecurityDelegation     = 3
)

const (
	TOKEN_ADJUST_PRIVILEGES = 0x0020
	SE_PRIVILEGE_ENABLED    = 0x00000002
)

type LUID struct {
	LowPart  uint32
	HighPart int32
}

type LUID_AND_ATTRIBUTES struct {
	Luid       LUID
	Attributes uint32
}

type TOKEN_PRIVILEGES struct {
	PrivilegeCount uint32
	Privileges     [1]LUID_AND_ATTRIBUTES
}

func AdjustTokenPrivileges(token syscall.Token, disableAllPrivileges bool, newstate *TOKEN_PRIVILEGES, buflen uint32, prevstate *TOKEN_PRIVILEGES, returnlen *uint32) error

type SID_AND_ATTRIBUTES struct {
	Sid        *syscall.SID
	Attributes uint32
}

type TOKEN_MANDATORY_LABEL struct {
	Label SID_AND_ATTRIBUTES
}

func (tml *TOKEN_MANDATORY_LABEL) Size() uint32

const SE_GROUP_INTEGRITY = 0x00000020

type TokenType uint32

const (
	TokenPrimary       TokenType = 1
	TokenImpersonation TokenType = 2
)

const (
	LG_INCLUDE_INDIRECT  = 0x1
	MAX_PREFERRED_LENGTH = 0xFFFFFFFF
)

type LocalGroupUserInfo0 struct {
	Name *uint16
}

type UserInfo4 struct {
	Name            *uint16
	Password        *uint16
	PasswordAge     uint32
	Priv            uint32
	HomeDir         *uint16
	Comment         *uint16
	Flags           uint32
	ScriptPath      *uint16
	AuthFlags       uint32
	FullName        *uint16
	UsrComment      *uint16
	Parms           *uint16
	Workstations    *uint16
	LastLogon       uint32
	LastLogoff      uint32
	AcctExpires     uint32
	MaxStorage      uint32
	UnitsPerWeek    uint32
	LogonHours      *byte
	BadPwCount      uint32
	NumLogons       uint32
	LogonServer     *uint16
	CountryCode     uint32
	CodePage        uint32
	UserSid         *syscall.SID
	PrimaryGroupID  uint32
	Profile         *uint16
	HomeDirDrive    *uint16
	PasswordExpired uint32
}

// GetSystemDirectory retrieves the path to current location of the system
// directory, which is typically, though not always, `C:\Windows\System32`.
//
//go:linkname GetSystemDirectory
func GetSystemDirectory() string

// GetUserName retrieves the user name of the current thread
// in the specified format.
func GetUserName(format uint32) (string, error)

type TOKEN_GROUPS struct {
	GroupCount uint32
	Groups     [1]SID_AND_ATTRIBUTES
}

func (g *TOKEN_GROUPS) AllGroups() []SID_AND_ATTRIBUTES

func GetTokenGroups(t syscall.Token) (*TOKEN_GROUPS, error)

// https://learn.microsoft.com/en-us/windows/win32/api/winnt/ns-winnt-sid_identifier_authority
type SID_IDENTIFIER_AUTHORITY struct {
	Value [6]byte
}

const (
	SID_REVISION = 1
	// https://learn.microsoft.com/en-us/windows/win32/services/localsystem-account
	SECURITY_LOCAL_SYSTEM_RID = 18
	// https://learn.microsoft.com/en-us/windows/win32/services/localservice-account
	SECURITY_LOCAL_SERVICE_RID = 19
	// https://learn.microsoft.com/en-us/windows/win32/services/networkservice-account
	SECURITY_NETWORK_SERVICE_RID = 20
)

var SECURITY_NT_AUTHORITY = SID_IDENTIFIER_AUTHORITY{
	Value: [6]byte{0, 0, 0, 0, 0, 5},
}

//go:nocheckptr
func GetSidIdentifierAuthority(sid *syscall.SID) SID_IDENTIFIER_AUTHORITY

//go:nocheckptr
func GetSidSubAuthority(sid *syscall.SID, subAuthorityIdx uint32) uint32

//go:nocheckptr
func GetSidSubAuthorityCount(sid *syscall.SID) uint8
