// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package template

import (
	"github.com/shogo82148/std/text/template/parse"
)

// common holds the information shared by related templates.

// Template is the representation of a parsed template. The *parse.Tree
// field is exported only for use by html/template and should be treated
// as unexported by all other clients.
type Template struct {
	name string
	*parse.Tree
	*common
	leftDelim  string
	rightDelim string
}

// New allocates a new template with the given name.
func New(name string) *Template

// Name returns the name of the template.
func (t *Template) Name() string

// New allocates a new template associated with the given one and with the same
// delimiters. The association, which is transitive, allows one template to
// invoke another with a {{template}} action.
func (t *Template) New(name string) *Template

// Clone returns a duplicate of the template, including all associated
// templates. The actual representation is not copied, but the name space of
// associated templates is, so further calls to Parse in the copy will add
// templates to the copy but not to the original. Clone can be used to prepare
// common templates and use them with variant definitions for other templates
// by adding the variants after the clone is made.
func (t *Template) Clone() (*Template, error)

// AddParseTree creates a new template with the name and parse tree
// and associates it with t.
func (t *Template) AddParseTree(name string, tree *parse.Tree) (*Template, error)

// Templates returns a slice of the templates associated with t, including t
// itself.
func (t *Template) Templates() []*Template

// Delims sets the action delimiters to the specified strings, to be used in
// subsequent calls to Parse, ParseFiles, or ParseGlob. Nested template
// definitions will inherit the settings. An empty delimiter stands for the
// corresponding default: {{ or }}.
// The return value is the template, so calls can be chained.
func (t *Template) Delims(left, right string) *Template

// Funcs adds the elements of the argument map to the template's function map.
// It panics if a value in the map is not a function with appropriate return
// type. However, it is legal to overwrite elements of the map. The return
// value is the template, so calls can be chained.
func (t *Template) Funcs(funcMap FuncMap) *Template

// Lookup returns the template with the given name that is associated with t,
// or nil if there is no such template.
func (t *Template) Lookup(name string) *Template

// Parse parses a string into a template. Nested template definitions will be
// associated with the top-level template t. Parse may be called multiple times
// to parse definitions of templates to associate with t. It is an error if a
// resulting template is non-empty (contains content other than template
// definitions) and would replace a non-empty template with the same name.
// (In multiple calls to Parse with the same receiver template, only one call
// can contain text other than space, comments, and template definitions.)
func (t *Template) Parse(text string) (*Template, error)
