// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package list

type Context struct {
	GOARCH        string   `json:",omitempty"`
	GOOS          string   `json:",omitempty"`
	GOROOT        string   `json:",omitempty"`
	GOPATH        string   `json:",omitempty"`
	CgoEnabled    bool     `json:",omitempty"`
	UseAllFiles   bool     `json:",omitempty"`
	Compiler      string   `json:",omitempty"`
	BuildTags     []string `json:",omitempty"`
	ToolTags      []string `json:",omitempty"`
	ReleaseTags   []string `json:",omitempty"`
	InstallSuffix string   `json:",omitempty"`
}
