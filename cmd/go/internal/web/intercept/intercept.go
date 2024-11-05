package intercept

import (
	"github.com/shogo82148/std/net/http"
	"github.com/shogo82148/std/net/url"
)

// Interceptor is used to change the host, and maybe the client,
// for a request to point to a test host.
type Interceptor struct {
	Scheme   string
	FromHost string
	ToHost   string
	Client   *http.Client
}

// EnableTestHooks installs the given interceptors to be used by URL and Request.
func EnableTestHooks(interceptors []Interceptor) error

// DisableTestHooks disables the installed interceptors.
func DisableTestHooks()

var (
	// TestHooksEnabled is true if interceptors are installed
	TestHooksEnabled = false
)

// URL returns the Interceptor to be used for a given URL.
func URL(u *url.URL) (*Interceptor, bool)

// Request updates the host to actually use for the request, if it is to be intercepted.
func Request(req *http.Request)
