// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package url parses URLs and implements query escaping.
//
// See RFC 3986. This package generally follows RFC 3986, except where
// it deviates for compatibility reasons.
// RFC 6874 followed for IPv6 zone literals.
package url

// Error reports an error and the operation and URL that caused it.
type Error struct {
	Op  string
	URL string
	Err error
}

func (e *Error) Unwrap() error
func (e *Error) Error() string

func (e *Error) Timeout() bool

func (e *Error) Temporary() bool

type EscapeError string

func (e EscapeError) Error() string

type InvalidHostError string

func (e InvalidHostError) Error() string

// QueryUnescape does the inverse transformation of [QueryEscape],
// converting each 3-byte encoded substring of the form "%AB" into the
// hex-decoded byte 0xAB.
// It returns an error if any % is not followed by two hexadecimal
// digits.
func QueryUnescape(s string) (string, error)

// PathUnescape does the inverse transformation of [PathEscape],
// converting each 3-byte encoded substring of the form "%AB" into the
// hex-decoded byte 0xAB. It returns an error if any % is not followed
// by two hexadecimal digits.
//
// PathUnescape is identical to [QueryUnescape] except that it does not
// unescape '+' to ' ' (space).
func PathUnescape(s string) (string, error)

// QueryEscape escapes the string so it can be safely placed
// inside a [URL] query.
func QueryEscape(s string) string

// PathEscape escapes the string so it can be safely placed inside a [URL] path segment,
// replacing special characters (including /) with %XX sequences as needed.
func PathEscape(s string) string

// A URL represents a parsed URL (technically, a URI reference).
//
// The general form represented is:
//
//	[scheme:][//[userinfo@]host][/]path[?query][#fragment]
//
// URLs that do not start with a slash after the scheme are interpreted as:
//
//	scheme:opaque[?query][#fragment]
//
// The Host field contains the host and port subcomponents of the URL.
// When the port is present, it is separated from the host with a colon.
// When the host is an IPv6 address, it must be enclosed in square brackets:
// "[fe80::1]:80". The [net.JoinHostPort] function combines a host and port
// into a string suitable for the Host field, adding square brackets to
// the host when necessary.
//
// Note that the Path field is stored in decoded form: /%47%6f%2f becomes /Go/.
// A consequence is that it is impossible to tell which slashes in the Path were
// slashes in the raw URL and which were %2f. This distinction is rarely important,
// but when it is, the code should use the [URL.EscapedPath] method, which preserves
// the original encoding of Path. The Fragment field is also stored in decoded form,
// use [URL.EscapedFragment] to retrieve the original encoding.
//
// The [URL.String] method uses the [URL.EscapedPath] method to obtain the path.
type URL struct {
	Scheme   string
	Opaque   string
	User     *Userinfo
	Host     string
	Path     string
	Fragment string

	// RawQuery contains the encoded query values, without the initial '?'.
	// Use URL.Query to decode the query.
	RawQuery string

	// RawPath is an optional field containing an encoded path hint.
	// See the EscapedPath method for more details.
	//
	// In general, code should call EscapedPath instead of reading RawPath.
	RawPath string

	// RawFragment is an optional field containing an encoded fragment hint.
	// See the EscapedFragment method for more details.
	//
	// In general, code should call EscapedFragment instead of reading RawFragment.
	RawFragment string

	// ForceQuery indicates whether the original URL contained a query ('?') character.
	// When set, the String method will include a trailing '?', even when RawQuery is empty.
	ForceQuery bool

	// OmitHost indicates the URL has an empty host (authority).
	// When set, the String method will not include the host when it is empty.
	OmitHost bool
}

// User returns a [Userinfo] containing the provided username
// and no password set.
func User(username string) *Userinfo

// UserPassword returns a [Userinfo] containing the provided username
// and password.
//
// This functionality should only be used with legacy web sites.
// RFC 2396 warns that interpreting Userinfo this way
// “is NOT RECOMMENDED, because the passing of authentication
// information in clear text (such as URI) has proven to be a
// security risk in almost every case where it has been used.”
func UserPassword(username, password string) *Userinfo

// The Userinfo type is an immutable encapsulation of username and
// password details for a [URL]. An existing Userinfo value is guaranteed
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

// Parse parses a raw url into a [URL] structure.
//
// The url may be relative (a path, without a host) or absolute
// (starting with a scheme). Trying to parse a hostname and path
// without a scheme is invalid but may not necessarily return an
// error, due to parsing ambiguities.
func Parse(rawURL string) (*URL, error)

// ParseRequestURI parses a raw url into a [URL] structure. It assumes that
// url was received in an HTTP request, so the url is interpreted
// only as an absolute URI or an absolute path.
// The string url is assumed not to have a #fragment suffix.
// (Web browsers strip #fragment before sending the URL to a web server.)
func ParseRequestURI(rawURL string) (*URL, error)

// EscapedPath returns the escaped form of u.Path.
// In general there are multiple possible escaped forms of any path.
// EscapedPath returns u.RawPath when it is a valid escaping of u.Path.
// Otherwise EscapedPath ignores u.RawPath and computes an escaped
// form on its own.
// The [URL.String] and [URL.RequestURI] methods use EscapedPath to construct
// their results.
// In general, code should call EscapedPath instead of
// reading u.RawPath directly.
func (u *URL) EscapedPath() string

// EscapedFragment returns the escaped form of u.Fragment.
// In general there are multiple possible escaped forms of any fragment.
// EscapedFragment returns u.RawFragment when it is a valid escaping of u.Fragment.
// Otherwise EscapedFragment ignores u.RawFragment and computes an escaped
// form on its own.
// The [URL.String] method uses EscapedFragment to construct its result.
// In general, code should call EscapedFragment instead of
// reading u.RawFragment directly.
func (u *URL) EscapedFragment() string

// String reassembles the [URL] into a valid URL string.
// The general form of the result is one of:
//
//	scheme:opaque?query#fragment
//	scheme://userinfo@host/path?query#fragment
//
// If u.Opaque is non-empty, String uses the first form;
// otherwise it uses the second form.
// Any non-ASCII characters in host are escaped.
// To obtain the path, String uses u.EscapedPath().
//
// In the second form, the following rules apply:
//   - if u.Scheme is empty, scheme: is omitted.
//   - if u.User is nil, userinfo@ is omitted.
//   - if u.Host is empty, host/ is omitted.
//   - if u.Scheme and u.Host are empty and u.User is nil,
//     the entire scheme://userinfo@host/ is omitted.
//   - if u.Host is non-empty and u.Path begins with a /,
//     the form host/path does not add its own /.
//   - if u.RawQuery is empty, ?query is omitted.
//   - if u.Fragment is empty, #fragment is omitted.
func (u *URL) String() string

// Redacted is like [URL.String] but replaces any password with "xxxxx".
// Only the password in u.User is redacted.
func (u *URL) Redacted() string

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

// Has checks whether a given key is set.
func (v Values) Has(key string) bool

// ParseQuery parses the URL-encoded query string and returns
// a map listing the values specified for each key.
// ParseQuery always returns a non-nil map containing all the
// valid query parameters found; err describes the first decoding error
// encountered, if any.
//
// Query is expected to be a list of key=value settings separated by ampersands.
// A setting without an equals sign is interpreted as a key set to an empty
// value.
// Settings containing a non-URL-encoded semicolon are considered invalid.
func ParseQuery(query string) (Values, error)

// Encode encodes the values into “URL encoded” form
// ("bar=baz&foo=quux") sorted by key.
func (v Values) Encode() string

// IsAbs reports whether the [URL] is absolute.
// Absolute means that it has a non-empty scheme.
func (u *URL) IsAbs() bool

// Parse parses a [URL] in the context of the receiver. The provided URL
// may be relative or absolute. Parse returns nil, err on parse
// failure, otherwise its return value is the same as [URL.ResolveReference].
func (u *URL) Parse(ref string) (*URL, error)

// ResolveReference resolves a URI reference to an absolute URI from
// an absolute base URI u, per RFC 3986 Section 5.2. The URI reference
// may be relative or absolute. ResolveReference always returns a new
// [URL] instance, even if the returned URL is identical to either the
// base or reference. If ref is an absolute URL, then ResolveReference
// ignores base and returns a copy of ref.
func (u *URL) ResolveReference(ref *URL) *URL

// Query parses RawQuery and returns the corresponding values.
// It silently discards malformed value pairs.
// To check errors use [ParseQuery].
func (u *URL) Query() Values

// RequestURI returns the encoded path?query or opaque?query
// string that would be used in an HTTP request for u.
func (u *URL) RequestURI() string

// Hostname returns u.Host, stripping any valid port number if present.
//
// If the result is enclosed in square brackets, as literal IPv6 addresses are,
// the square brackets are removed from the result.
func (u *URL) Hostname() string

// Port returns the port part of u.Host, without the leading colon.
//
// If u.Host doesn't contain a valid numeric port, Port returns an empty string.
func (u *URL) Port() string

func (u *URL) MarshalBinary() (text []byte, err error)

func (u *URL) AppendBinary(b []byte) ([]byte, error)

func (u *URL) UnmarshalBinary(text []byte) error

// JoinPath returns a new [URL] with the provided path elements joined to
// any existing path and the resulting path cleaned of any ./ or ../ elements.
// Any sequences of multiple / characters will be reduced to a single /.
func (u *URL) JoinPath(elem ...string) *URL

// JoinPath returns a [URL] string with the provided path elements joined to
// the existing path of base and the resulting path cleaned of any ./ or ../ elements.
func JoinPath(base string, elem ...string) (result string, err error)
