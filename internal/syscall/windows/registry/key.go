// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows

// Package registry provides access to the Windows registry.
//
// Here is a simple example, opening a registry key and reading a string value from it.
//
//	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer k.Close()
//
//	s, _, err := k.GetStringValue("SystemRoot")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("Windows system root is %q\n", s)
//
// NOTE: This package is a copy of golang.org/x/sys/windows/registry
// with KeyInfo.ModTime removed to prevent dependency cycles.
package registry

import (
	"github.com/shogo82148/std/syscall"
)

const (
	// Registry key security and access rights.
	// See https://msdn.microsoft.com/en-us/library/windows/desktop/ms724878.aspx
	// for details.
	ALL_ACCESS         = 0xf003f
	CREATE_LINK        = 0x00020
	CREATE_SUB_KEY     = 0x00004
	ENUMERATE_SUB_KEYS = 0x00008
	EXECUTE            = 0x20019
	NOTIFY             = 0x00010
	QUERY_VALUE        = 0x00001
	READ               = 0x20019
	SET_VALUE          = 0x00002
	WOW64_32KEY        = 0x00200
	WOW64_64KEY        = 0x00100
	WRITE              = 0x20006
)

// Key is a handle to an open Windows registry key.
// Keys can be obtained by calling OpenKey; there are
// also some predefined root keys such as CURRENT_USER.
// Keys can be used directly in the Windows API.
type Key syscall.Handle

const (
	// Windows defines some predefined root keys that are always open.
	// An application can use these keys as entry points to the registry.
	// Normally these keys are used in OpenKey to open new keys,
	// but they can also be used anywhere a Key is required.
	CLASSES_ROOT   = Key(syscall.HKEY_CLASSES_ROOT)
	CURRENT_USER   = Key(syscall.HKEY_CURRENT_USER)
	LOCAL_MACHINE  = Key(syscall.HKEY_LOCAL_MACHINE)
	USERS          = Key(syscall.HKEY_USERS)
	CURRENT_CONFIG = Key(syscall.HKEY_CURRENT_CONFIG)
)

// Close closes open key k.
func (k Key) Close() error

// OpenKey opens a new key with path name relative to key k.
// It accepts any open key, including CURRENT_USER and others,
// and returns the new key and an error.
// The access parameter specifies desired access rights to the
// key to be opened.
func OpenKey(k Key, path string, access uint32) (Key, error)

// ReadSubKeyNames returns the names of subkeys of key k.
func (k Key) ReadSubKeyNames() ([]string, error)

// CreateKey creates a key named path under open key k.
// CreateKey returns the new key and a boolean flag that reports
// whether the key already existed.
// The access parameter specifies the access rights for the key
// to be created.
func CreateKey(k Key, path string, access uint32) (newk Key, openedExisting bool, err error)

// DeleteKey deletes the subkey path of key k and its values.
func DeleteKey(k Key, path string) error

// A KeyInfo describes the statistics of a key. It is returned by Stat.
type KeyInfo struct {
	SubKeyCount     uint32
	MaxSubKeyLen    uint32
	ValueCount      uint32
	MaxValueNameLen uint32
	MaxValueLen     uint32
	lastWriteTime   syscall.Filetime
}

// Stat retrieves information about the open key k.
func (k Key) Stat() (*KeyInfo, error)
