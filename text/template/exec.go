// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package template

import (
	"github.com/shogo82148/std/io"
)

// state represents the state of an execution. It's not part of the
// template so that multiple executions of the same template
// can execute in parallel.

// variable holds the dynamic value of a variable such as $, $x etc.

// ExecuteTemplate applies the template associated with t that has the given name
// to the specified data object and writes the output to wr.
func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error

// Execute applies a parsed template to the specified data object,
// and writes the output to wr.
func (t *Template) Execute(wr io.Writer, data interface{}) (err error)
