// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package template

import (
	"github.com/shogo82148/std/text/template/parse"
)

// Template is the representation of a parsed template. The *parse.Tree
// field is exported only for use by [html/template] and should be treated
// as unexported by all other clients.
type Template struct {
	name string
	*parse.Tree
	*common
	leftDelim  string
	rightDelim string
}

// New allocates a new, undefined template with the given name.
func New(name string) *Template

// Name returns the name of the template.
func (t *Template) Name() string

// New allocates a new, undefined template associated with the given one and with the same
// delimiters. The association, which is transitive, allows one template to
// invoke another with a {{template}} action.
//
// Because associated templates share underlying data, template construction
// cannot be done safely in parallel. Once the templates are constructed, they
// can be executed in parallel.
func (t *Template) New(name string) *Template

// Clone returns a duplicate of the template, including all associated
// templates. The actual representation is not copied, but the name space of
// associated templates is, so further calls to [Template.Parse] in the copy will add
// templates to the copy but not to the original. Clone can be used to prepare
// common templates and use them with variant definitions for other templates
// by adding the variants after the clone is made.
func (t *Template) Clone() (*Template, error)

// AddParseTree associates the argument parse tree with the template t, giving
// it the specified name. If the template has not been defined, this tree becomes
// its definition. If it has been defined and already has that name, the existing
// definition is replaced; otherwise a new template is created, defined, and returned.
func (t *Template) AddParseTree(name string, tree *parse.Tree) (*Template, error)

// Templates returns a slice of defined templates associated with t.
func (t *Template) Templates() []*Template

// Delims sets the action delimiters to the specified strings, to be used in
// subsequent calls to [Template.Parse], [Template.ParseFiles], or [Template.ParseGlob]. Nested template
// definitions will inherit the settings. An empty delimiter stands for the
// corresponding default: {{ or }}.
// The return value is the template, so calls can be chained.
func (t *Template) Delims(left, right string) *Template

// Funcs adds the elements of the argument map to the template's function map.
// It must be called before the template is parsed.
// It panics if a value in the map is not a function with appropriate return
// type or if the name cannot be used syntactically as a function in a template.
// It is legal to overwrite elements of the map. The return value is the template,
// so calls can be chained.
func (t *Template) Funcs(funcMap FuncMap) *Template

// Lookup returns the template with the given name that is associated with t.
// It returns nil if there is no such template or the template has no definition.
func (t *Template) Lookup(name string) *Template

// Parse parses text as a template body for t.
// Named template definitions ({{define ...}} or {{block ...}} statements) in text
// define additional templates associated with t and are removed from the
// definition of t itself.
//
// Templates can be redefined in successive calls to Parse.
// A template definition with a body containing only white space and comments
// is considered empty and will not replace an existing template's body.
// This allows using Parse to add new named template definitions without
// overwriting the main template body.
func (t *Template) Parse(text string) (*Template, error)
