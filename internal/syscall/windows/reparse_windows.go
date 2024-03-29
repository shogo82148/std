// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package windows

// Reparse tag values are taken from
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-fscc/c8e77b37-3909-4fe6-a4ea-2b9d423b1ee4
const (
	FSCTL_SET_REPARSE_POINT    = 0x000900A4
	IO_REPARSE_TAG_MOUNT_POINT = 0xA0000003
	IO_REPARSE_TAG_DEDUP       = 0x80000013
	IO_REPARSE_TAG_AF_UNIX     = 0x80000023

	SYMLINK_FLAG_RELATIVE = 1
)

type REPARSE_DATA_BUFFER struct {
	ReparseTag        uint32
	ReparseDataLength uint16
	Reserved          uint16
	DUMMYUNIONNAME    byte
}

// REPARSE_DATA_BUFFER_HEADER is a common part of REPARSE_DATA_BUFFER structure.
type REPARSE_DATA_BUFFER_HEADER struct {
	ReparseTag uint32
	// The size, in bytes, of the reparse data that follows
	// the common portion of the REPARSE_DATA_BUFFER element.
	// This value is the length of the data starting at the
	// SubstituteNameOffset field.
	ReparseDataLength uint16
	Reserved          uint16
}

type SymbolicLinkReparseBuffer struct {
	// The integer that contains the offset, in bytes,
	// of the substitute name string in the PathBuffer array,
	// computed as an offset from byte 0 of PathBuffer. Note that
	// this offset must be divided by 2 to get the array index.
	SubstituteNameOffset uint16
	// The integer that contains the length, in bytes, of the
	// substitute name string. If this string is null-terminated,
	// SubstituteNameLength does not include the Unicode null character.
	SubstituteNameLength uint16
	// PrintNameOffset is similar to SubstituteNameOffset.
	PrintNameOffset uint16
	// PrintNameLength is similar to SubstituteNameLength.
	PrintNameLength uint16
	// Flags specifies whether the substitute name is a full path name or
	// a path name relative to the directory containing the symbolic link.
	Flags      uint32
	PathBuffer [1]uint16
}

// Path returns path stored in rb.
func (rb *SymbolicLinkReparseBuffer) Path() string

type MountPointReparseBuffer struct {
	// The integer that contains the offset, in bytes,
	// of the substitute name string in the PathBuffer array,
	// computed as an offset from byte 0 of PathBuffer. Note that
	// this offset must be divided by 2 to get the array index.
	SubstituteNameOffset uint16
	// The integer that contains the length, in bytes, of the
	// substitute name string. If this string is null-terminated,
	// SubstituteNameLength does not include the Unicode null character.
	SubstituteNameLength uint16
	// PrintNameOffset is similar to SubstituteNameOffset.
	PrintNameOffset uint16
	// PrintNameLength is similar to SubstituteNameLength.
	PrintNameLength uint16
	PathBuffer      [1]uint16
}

// Path returns path stored in rb.
func (rb *MountPointReparseBuffer) Path() string
