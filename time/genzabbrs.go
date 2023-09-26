// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore

//
// usage:
//
// go run genzabbrs.go -output zoneinfo_abbrs_windows.go
//

package main

type MapZone struct {
	Other     string `xml:"other,attr"`
	Territory string `xml:"territory,attr"`
	Type      string `xml:"type,attr"`
}

type SupplementalData struct {
	Zones []MapZone `xml:"windowsZones>mapTimezones>mapZone"`
}
