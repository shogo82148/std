// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package url parses URLs and implements query escaping.
// See RFC 3986.
package url

// Error reports an error and the operation and URL that caused it.
type Error struct {
	Op  string
	URL string
	Err error
}

func (e *Error) Error() string

type EscapeError string

func (e EscapeError) Error() string

// QueryUnescape does the inverse transformation of QueryEscape, converting
// %AB into the byte 0xAB and '+' into ' ' (space). It returns an error if
// any % is not followed by two hexadecimal digits.
func QueryUnescape(s string) (string, error)

// QueryEscape escapes the string so it can be safely placed
// inside a URL query.
func QueryEscape(s string) string

// A URL represents a parsed URL (technically, a URI reference).
// The general form represented is:
//
//	scheme://[userinfo@]host/path[?query][#fragment]
//
// URLs that do not start with a slash after the scheme are interpreted as:
//
//	scheme:opaque[?query][#fragment]
//
// Note that the Path field is stored in decoded form: /%47%6f%2f becomes /Go/.
// A consequence is that it is impossible to tell which slashes in the Path were
// slashes in the raw URL and which were %2f. This distinction is rarely important,
// but when it is a client must use other routines to parse the raw URL or construct
// the parsed URL. For example, an HTTP server can consult req.RequestURI, and
// an HTTP client can use URL{Host: "example.com", Opaque: "//example.com/Go%2f"}
// instead of URL{Host: "example.com", Path: "/Go/"}.
type URL struct {
	Scheme   string
	Opaque   string
	User     *Userinfo
	Host     string
	Path     string
	RawQuery string
	Fragment string
}

// User returns a Userinfo containing the provided username
// and no password set.
func User(username string) *Userinfo

// UserPassword returns a Userinfo containing the provided username
// and password.
// This functionality should only be used with legacy web sites.
// RFC 2396 warns that interpreting Userinfo this way
// “is NOT RECOMMENDED, because the passing of authentication
// information in clear text (such as URI) has proven to be a
// security risk in almost every case where it has been used.”
func UserPassword(username, password string) *Userinfo

// The Userinfo type is an immutable encapsulation of username and
// password details for a URL. An existing Userinfo value is guaranteed
// to have a username set (potentially empty, as allowed by RFC 2396),
// and optionally a password.
type Userinfo struct {
	username    string
	password    string
	passwordSet bool
}

// Username returns the username.
func (u *Userinfo) Username() string

// Password returns the password in case it is set, and whether it is set.
func (u *Userinfo) Password() (string, bool)

// String returns the encoded userinfo information in the standard form
// of "username[:password]".
func (u *Userinfo) String() string

// Parse parses rawurl into a URL structure.
// The rawurl may be relative or absolute.
func Parse(rawurl string) (url *URL, err error)

// ParseRequestURI parses rawurl into a URL structure.  It assumes that
// rawurl was received in an HTTP request, so the rawurl is interpreted
// only as an absolute URI or an absolute path.
// The string rawurl is assumed not to have a #fragment suffix.
// (Web browsers strip #fragment before sending the URL to a web server.)
func ParseRequestURI(rawurl string) (url *URL, err error)

// String reassembles the URL into a valid URL string.
func (u *URL) String() string

// Values maps a string key to a list of values.
// It is typically used for query parameters and form values.
// Unlike in the http.Header map, the keys in a Values map
// are case-sensitive.
type Values map[string][]string

// Get gets the first value associated with the given key.
// If there are no values associated with the key, Get returns
// the empty string. To access multiple values, use the map
// directly.
func (v Values) Get(key string) string

// Set sets the key to value. It replaces any existing
// values.
func (v Values) Set(key, value string)

// Add adds the value to key. It appends to any existing
// values associated with key.
func (v Values) Add(key, value string)

// Del deletes the values associated with key.
func (v Values) Del(key string)

// ParseQuery parses the URL-encoded query string and returns
// a map listing the values specified for each key.
// ParseQuery always returns a non-nil map containing all the
// valid query parameters found; err describes the first decoding error
// encountered, if any.
func ParseQuery(query string) (m Values, err error)

// Encode encodes the values into “URL encoded” form
// ("bar=baz&foo=quux") sorted by key.
func (v Values) Encode() string

// IsAbs returns true if the URL is absolute.
func (u *URL) IsAbs() bool

// Parse parses a URL in the context of the receiver.  The provided URL
// may be relative or absolute.  Parse returns nil, err on parse
// failure, otherwise its return value is the same as ResolveReference.
func (u *URL) Parse(ref string) (*URL, error)

// ResolveReference resolves a URI reference to an absolute URI from
// an absolute base URI, per RFC 3986 Section 5.2.  The URI reference
// may be relative or absolute.  ResolveReference always returns a new
// URL instance, even if the returned URL is identical to either the
// base or reference. If ref is an absolute URL, then ResolveReference
// ignores base and returns a copy of ref.
func (u *URL) ResolveReference(ref *URL) *URL

// Query parses RawQuery and returns the corresponding values.
func (u *URL) Query() Values

// RequestURI returns the encoded path?query or opaque?query
// string that would be used in an HTTP request for u.
func (u *URL) RequestURI() string
