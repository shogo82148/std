// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ir

import (
	"github.com/shogo82148/std/bufio"
	"github.com/shogo82148/std/io"
)

// An HTMLWriter dumps IR to multicolumn HTML, similar to what the
// ssa backend does for GOSSAFUNC.  This is not the format used for
// the ast column in GOSSAFUNC output.
type HTMLWriter struct {
	HTMLWriterBase
	Func *Func
}

type HTMLWriterBase struct {
	w             *BufferedWriterCloser
	canonIdMap    map[any]int
	prevCanonId   int
	path          string
	prevHash      []byte
	pendingPhases []string
	pendingTitles []string
	doDump        func(string) func()
}

func (h *HTMLWriterBase) Init(out io.WriteCloser, reportPath string, doDump func(string) func())

// BufferedWriterCloser is here to help avoid pre-buffering the whole
// rendered HTML in memory, which can cause problems for large inputs.
type BufferedWriterCloser struct {
	file io.Closer
	w    *bufio.Writer
}

func (b *BufferedWriterCloser) Write(p []byte) (n int, err error)

func (b *BufferedWriterCloser) Close() error

func NewBufferedWriterCloser(f io.WriteCloser) *BufferedWriterCloser

func NewHTMLWriter(path string, f *Func, cfgMask string) *HTMLWriter

func (h *HTMLWriterBase) Path() string

// CanonId assigns indices to nodes based on pointer identity.
// this helps ensure that output html files don't gratuitously
// differ from run to run.
func (h *HTMLWriterBase) CanonId(n any) int

// Fatalf reports an error and exits.
func (w *HTMLWriterBase) Fatalf(msg string, args ...any)

const (
	RightArrow = "►"
	DownArrow  = "▼"
)

func (w *HTMLWriterBase) Close(format string, args ...any)

// WritePhase writes f in a column headed by title.
// phase is used for collapsing columns and should be unique across the table.
func (w *HTMLWriterBase) WritePhase(phase, title string)

func (w *HTMLWriterBase) WriteMultiTitleColumn(phase string, titles []string, class string, writeContent func())

func (w *HTMLWriterBase) Printf(msg string, v ...any)

func (w *HTMLWriterBase) Print(s string)

func (w *HTMLWriter) FuncHTML(phase string) func()

const CSS = `<style>

body {
    font-size: 14px;
    font-family: Arial, sans-serif;
}

h1 {
    font-size: 18px;
    display: inline-block;
    margin: 0 1em .5em 0;
}

#helplink {
    display: inline-block;
}

#help {
    display: none;
}

table {
    border: 1px solid black;
    table-layout: fixed;
    width: 300px;
}

th, td {
    border: 1px solid black;
    overflow: hidden;
    width: 400px;
    vertical-align: top;
    padding: 5px;
    position: relative;
}

.resizer {
    display: inline-block;
    background: transparent;
    width: 10px;
    height: 100%;
    position: absolute;
    right: 0;
    top: 0;
    cursor: col-resize;
    z-index: 100;
}

td > h2 {
    cursor: pointer;
    font-size: 120%;
    margin: 5px 0px 5px 0px;
}

td.collapsed {
    font-size: 12px;
    width: 12px;
    border: 1px solid white;
    padding: 2px;
    cursor: pointer;
    background: #fafafa;
}

td.collapsed div {
    text-align: right;
    transform: rotate(180deg);
    writing-mode: vertical-lr;
    white-space: pre;
}

pre {
    font-family: Menlo, monospace;
    font-size: 12px;
}

pre {
    -moz-tab-size: 4;
    -o-tab-size:   4;
    tab-size:      4;
}

.allow-x-scroll {
    overflow-x: scroll;
}

.outline-node {
    cursor: cell;
}

.variable-name {
    cursor: crosshair;
}

.line-number {
    font-size: 11px;
    cursor: crosshair;
}

body.darkmode {
    background-color: rgb(21, 21, 21);
    color: rgb(230, 255, 255);
    opacity: 100%;
}

td.darkmode {
    background-color: rgb(21, 21, 21);
    border: 1px solid gray;
}

body.darkmode table, th {
    border: 1px solid gray;
}

body.darkmode text {
    fill: white;
}

.highlight-aquamarine     { background-color: aquamarine; color: black; }
.highlight-coral          { background-color: coral; color: black; }
.highlight-lightpink      { background-color: lightpink; color: black; }
.highlight-lightsteelblue { background-color: lightsteelblue; color: black; }
.highlight-palegreen      { background-color: palegreen; color: black; }
.highlight-skyblue        { background-color: skyblue; color: black; }
.highlight-lightgray      { background-color: lightgray; color: black; }
.highlight-yellow         { background-color: yellow; color: black; }
.highlight-lime           { background-color: lime; color: black; }
.highlight-khaki          { background-color: khaki; color: black; }
.highlight-aqua           { background-color: aqua; color: black; }
.highlight-salmon         { background-color: salmon; color: black; }


.outline-blue           { outline: #2893ff solid 2px; }
.outline-red            { outline: red solid 2px; }
.outline-blueviolet     { outline: blueviolet solid 2px; }
.outline-darkolivegreen { outline: darkolivegreen solid 2px; }
.outline-fuchsia        { outline: fuchsia solid 2px; }
.outline-sienna         { outline: sienna solid 2px; }
.outline-gold           { outline: gold solid 2px; }
.outline-orangered      { outline: orangered solid 2px; }
.outline-teal           { outline: teal solid 2px; }
.outline-maroon         { outline: maroon solid 2px; }
.outline-black          { outline: black solid 2px; }

/* Capture alternative for outline-black and ellipse.outline-black when in dark mode */
body.darkmode .outline-black        { outline: gray solid 2px; }

.toggle {
    cursor: pointer;
    display: inline-block;
    text-align: center;
    user-select: none;
    font-size: 12px; // hand-tweaked
}

</style>
`

func JS(opened ...string) string

const (
	JS1 = `<script type="text/javascript">

// Contains phase names which are expanded by default. Other columns are collapsed.
let expandedDefault = [
`
	JS2 = `];
if (history.state === null) {
    history.pushState({expandedDefault}, "", location.href);
}

// ordered list of all available highlight colors
var highlights = [
    "highlight-aquamarine",
    "highlight-coral",
    "highlight-lightpink",
    "highlight-lightsteelblue",
    "highlight-palegreen",
    "highlight-skyblue",
    "highlight-lightgray",
    "highlight-yellow",
    "highlight-lime",
    "highlight-khaki",
    "highlight-aqua",
    "highlight-salmon"
];

// state: which value is highlighted this color?
var highlighted = {};
for (var i = 0; i < highlights.length; i++) {
    highlighted[highlights[i]] = "";
}

// ordered list of all available outline colors
var outlines = [
    "outline-blue",
    "outline-red",
    "outline-blueviolet",
    "outline-darkolivegreen",
    "outline-fuchsia",
    "outline-sienna",
    "outline-gold",
    "outline-orangered",
    "outline-teal",
    "outline-maroon",
    "outline-black"
];

// state: which value is outlined this color?
var outlined = {};
for (var i = 0; i < outlines.length; i++) {
    outlined[outlines[i]] = "";
}

window.onload = function() {
    if (history.state !== null) {
        expandedDefault = history.state.expandedDefault;
    }
    if (window.matchMedia && window.matchMedia("(prefers-color-scheme: dark)").matches) {
        toggleDarkMode();
        document.getElementById("dark-mode-button").checked = true;
    }

    var irElemClicked = function(elem, event, selections, selected) {
        event.stopPropagation();

        // find all values with the same name
        var c = elem.classList.item(0);
        var x = document.getElementsByClassName(c);

        // if selected, remove selections from all of them
        // otherwise, attempt to add

        var remove = "";
        for (var i = 0; i < selections.length; i++) {
            var color = selections[i];
            if (selected[color] == c) {
                remove = color;
                break;
            }
        }

        if (remove != "") {
            for (var i = 0; i < x.length; i++) {
                x[i].classList.remove(remove);
            }
            selected[remove] = "";
            return;
        }

        // we're adding a selection
        // find first available color
        var avail = "";
        for (var i = 0; i < selections.length; i++) {
            var color = selections[i];
            if (selected[color] == "") {
                avail = color;
                break;
            }
        }
        if (avail == "") {
            alert("out of selection colors; go add more");
            return;
        }

        // set that as the selection
        for (var i = 0; i < x.length; i++) {
            x[i].classList.add(avail);
        }
        selected[avail] = c;
    };

    var irValueClicked = function(event) {
        irElemClicked(this, event, highlights, highlighted);
    };

    var irTreeClicked = function(event) {
        irElemClicked(this, event, outlines, outlined);
    };

    var irValues = document.getElementsByClassName("outline-node");
    for (var i = 0; i < irValues.length; i++) {
        irValues[i].addEventListener('click', irTreeClicked);
    }

    var lines = document.getElementsByClassName("line-number");
    for (var i = 0; i < lines.length; i++) {
        lines[i].addEventListener('click', irValueClicked);
    }

    var variableNames = document.getElementsByClassName("variable-name");
    for (var i = 0; i < variableNames.length; i++) {
        variableNames[i].addEventListener('click', irValueClicked);
    }

    function toggler(phase) {
        return function() {
            toggle_cell(phase+'-col');
            toggle_cell(phase+'-exp');
            const i = expandedDefault.indexOf(phase);
            if (i !== -1) {
                expandedDefault.splice(i, 1);
            } else {
                expandedDefault.push(phase);
            }
            history.pushState({expandedDefault}, "", location.href);
        };
    }

    function toggle_cell(id) {
        var e = document.getElementById(id);
        if (e.style.display == 'table-cell') {
            e.style.display = 'none';
        } else {
            e.style.display = 'table-cell';
        }
    }

    // Go through all columns and collapse needed phases.
    const td = document.getElementsByTagName("td");
    for (let i = 0; i < td.length; i++) {
        const id = td[i].id;
        const phase = id.substr(0, id.length-4);
        let show = expandedDefault.indexOf(phase) !== -1

        // If show == false, check to see if this is a combined column (multiple phases).
        // If combined, check each of the phases to see if they are in our expandedDefaults.
        // If any are found, that entire combined column gets shown.
        if (!show) {
            const combined = phase.split('--+--');
            const len = combined.length;
            if (len > 1) {
                for (let i = 0; i < len; i++) {
                    const num = expandedDefault.indexOf(combined[i]);
                    if (num !== -1) {
                        expandedDefault.splice(num, 1);
                        if (expandedDefault.indexOf(phase) === -1) {
                            expandedDefault.push(phase);
                            show = true;
                        }
                    }
                }
            }
        }
        if (id.endsWith("-exp")) {
            const h2Els = td[i].getElementsByTagName("h2");
            const len = h2Els.length;
            if (len > 0) {
                for (let i = 0; i < len; i++) {
                    h2Els[i].addEventListener('click', toggler(phase));
                }
            }
        } else {
            td[i].addEventListener('click', toggler(phase));
        }
        if (id.endsWith("-col") && show || id.endsWith("-exp") && !show) {
            td[i].style.display = 'none';
            continue;
        }
        td[i].style.display = 'table-cell';
    }

    var resizers = document.getElementsByClassName("resizer");
    for (var i = 0; i < resizers.length; i++) {
        var resizer = resizers[i];
        resizer.addEventListener('mousedown', initDrag, false);
    }
};

var startX, startWidth, resizableCol;

function initDrag(e) {
    resizableCol = this.parentElement;
    startX = e.clientX;
    startWidth = parseInt(document.defaultView.getComputedStyle(resizableCol).width, 10);
    document.documentElement.addEventListener('mousemove', doDrag, false);
    document.documentElement.addEventListener('mouseup', stopDrag, false);
}

function doDrag(e) {
    resizableCol.style.width = (startWidth + e.clientX - startX) + 'px';
}

function stopDrag(e) {
    document.documentElement.removeEventListener('mousemove', doDrag, false);
    document.documentElement.removeEventListener('mouseup', stopDrag, false);
}

function toggle_visibility(id) {
    var e = document.getElementById(id);
    if (e.style.display == 'block') {
        e.style.display = 'none';
    } else {
        e.style.display = 'block';
    }
}

function toggleDarkMode() {
    document.body.classList.toggle('darkmode');

    // Collect all of the "collapsed" elements and apply dark mode on each collapsed column
    const collapsedEls = document.getElementsByClassName('collapsed');
    const len = collapsedEls.length;

    for (let i = 0; i < len; i++) {
        collapsedEls[i].classList.toggle('darkmode');
    }
}

function toggle_node(e) {
    event.stopPropagation();
    var parent = e.parentNode;
    var children = parent.children;
    for (var i = 0; i < children.length; i++) {
        if (children[i].classList.contains("node-body")) {
            if (children[i].style.display == "none") {
                children[i].style.display = "";
            } else {
                children[i].style.display = "none";
            }
        }
    }
    if (e.innerText == "` + RightArrow + `") {
        e.innerText = "` + DownArrow + `";
    } else {
        e.innerText = "` + RightArrow + `";
    }
}

</script>
`
)
