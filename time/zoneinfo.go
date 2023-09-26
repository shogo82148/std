// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package time

// A Location maps time instants to the zone in use at that time.
// Typically, the Location represents the collection of time offsets
// in use in a geographical area, such as CEST and CET for central Europe.
type Location struct {
	name string
	zone []zone
	tx   []zoneTrans

	cacheStart int64
	cacheEnd   int64
	cacheZone  *zone
}

// A zone represents a single time zone such as CEST or CET.

// A zoneTrans represents a single time zone transition.

// alpha and omega are the beginning and end of time for zone
// transitions.

// UTC represents Universal Coordinated Time (UTC).
var UTC *Location = &utcLoc

// utcLoc is separate so that get can refer to &utcLoc
// and ensure that it never returns a nil *Location,
// even if a badly behaved client has changed UTC.

// Local represents the system's local time zone.
var Local *Location = &localLoc

// localLoc is separate so that initLocal can initialize
// it even if a client has changed Local.

// String returns a descriptive name for the time zone information,
// corresponding to the name argument to LoadLocation or FixedZone.
func (l *Location) String() string

// FixedZone returns a Location that always uses
// the given zone name and offset (seconds east of UTC).
func FixedZone(name string, offset int) *Location

// LoadLocation returns the Location with the given name.
//
// If the name is "" or "UTC", LoadLocation returns UTC.
// If the name is "Local", LoadLocation returns Local.
//
// Otherwise, the name is taken to be a location name corresponding to a file
// in the IANA Time Zone database, such as "America/New_York".
//
// The time zone database needed by LoadLocation may not be
// present on all systems, especially non-Unix systems.
// LoadLocation looks in the directory or uncompressed zip file
// named by the ZONEINFO environment variable, if any, then looks in
// known installation locations on Unix systems,
// and finally looks in $GOROOT/lib/time/zoneinfo.zip.
func LoadLocation(name string) (*Location, error)
