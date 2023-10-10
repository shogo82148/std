// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syntax

// Progはコンパイルされた正規表現プログラムです。
type Prog struct {
	Inst   []Inst
	Start  int
	NumCap int
}

// InstOpは命令のオペコードです。
type InstOp uint8

const (
	InstAlt InstOp = iota
	InstAltMatch
	InstCapture
	InstEmptyWidth
	InstMatch
	InstFail
	InstNop
	InstRune
	InstRune1
	InstRuneAny
	InstRuneAnyNotNL
)

func (i InstOp) String() string

// EmptyOpは、ゼロ幅アサーションの種類または混合を指定します。
type EmptyOp uint8

const (
	EmptyBeginLine EmptyOp = 1 << iota
	EmptyEndLine
	EmptyBeginText
	EmptyEndText
	EmptyWordBoundary
	EmptyNoWordBoundary
)

// EmptyOpContextは、r1とr2のルーンの間の位置で満たされる
// ゼロ幅のアサーションを返します。
// r1 == -1を渡すと、位置がテキストの先頭にあることを示します。
// r2 == -1を渡すと、位置がテキストの末尾にあることを示します。
func EmptyOpContext(r1, r2 rune) EmptyOp

// IsWordCharは、\bおよび\Bゼロ幅のアサーションの評価中にrが「単語文字」と見なされるかどうかを報告します。
// これらのアサーションはASCIIのみです：単語文字は[A-Za-z0-9_]です。
func IsWordChar(r rune) bool

// Instは正規表現プログラム内の単一の命令です。
type Inst struct {
	Op   InstOp
	Out  uint32
	Arg  uint32
	Rune []rune
}

func (p *Prog) String() string

// Prefix は正規表現のすべての一致した結果が始まるリテラル文字列を返します。もし Prefix が完全な一致である場合、Complete は true になります。
func (p *Prog) Prefix() (prefix string, complete bool)

// StartCondは、どのマッチにおいても真である必要がある先頭の空幅条件を返します。
// マッチが不可能な場合は、^EmptyOp(0)を返します。
func (p *Prog) StartCond() EmptyOp

// MatchRune は指定した r に instruction が一致し、それを消費するかどうかを報告します。
// i.Op == InstRune の場合にのみ呼び出すべきです。
func (i *Inst) MatchRune(r rune) bool

// MatchRunePosは、命令がrと一致しているかどうか（そして消費するかどうか）を確認します。
// そうであれば、MatchRunePosは一致するルーンのペアのインデックスを返します
// （または、len(i.Rune) == 1の場合、ルーンの単一要素）。
// 一致しない場合、MatchRunePosは-1を返します。
// MatchRunePosは、i.Op == InstRuneの場合のみ呼び出す必要があります。
func (i *Inst) MatchRunePos(r rune) int

// MatchEmptyWidthは、runesの前と後の間に空の文字列が
// マッチしているかどうかを報告します。
// i.Op == InstEmptyWidthの場合にのみ呼び出すべきです。
func (i *Inst) MatchEmptyWidth(before rune, after rune) bool

func (i *Inst) String() string
