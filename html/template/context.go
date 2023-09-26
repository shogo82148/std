// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package template

// context describes the state an HTML parser must be in when it reaches the
// portion of HTML produced by evaluating a particular template node.
//
// The zero value of type context is the start context for a template that
// produces an HTML fragment as defined at
// https://www.w3.org/TR/html5/syntax.html#the-end
// where the context element is null.

// state describes a high-level HTML parser state.
//
// It bounds the top of the element stack, and by extension the HTML insertion
// mode, but also contains state that does not correspond to anything in the
// HTML5 parsing algorithm because a single token production in the HTML
// grammar may contain embedded actions in a template. For instance, the quoted
// HTML attribute produced by
//
//	<div title="Hello {{.World}}">
//
// is a single token in HTML's grammar but in a template spans several nodes.

// delim is the delimiter that will end the current HTML attribute.

// urlPart identifies a part in an RFC 3986 hierarchical URL to allow different
// encoding strategies.

// jsCtx determines whether a '/' starts a regular expression literal or a
// division operator.

// element identifies the HTML element when inside a start tag or special body.
// Certain HTML element (for example <script> and <style>) have bodies that are
// treated differently from stateText so the element type is necessary to
// transition into the correct context at the end of a tag and to identify the
// end delimiter for the body.

// attr identifies the current HTML attribute when inside the attribute,
// that is, starting from stateAttrName until stateTag/stateText (exclusive).
