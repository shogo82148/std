// codecmp compares the code of two Go source trees, ignoring comments
// (except //go: directives) and blank lines. It is used to check that the
// translated branch has the same declarations as the upstream branch.
//
// Usage:
//
//	git worktree add /tmp/upstream origin/upstream/master
//	cd _tools && go run ./codecmp [-out dir] .. /tmp/upstream
//
// Each line of the output is one of:
//
//	DIFF path            the code differs
//	ONLY-MASTER path     the file exists only in the first tree
//	ONLY-UPSTREAM path   the file exists only in the second tree
//
// With -out, the stripped versions of differing files are written to
// dir/a and dir/b so that they can be compared with diff -r.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func strip(path string) ([]byte, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, parser.ParseComments|parser.SkipObjectResolution)
	if err != nil {
		return nil, err
	}
	var keep []*ast.CommentGroup
	for _, cg := range f.Comments {
		var list []*ast.Comment
		for _, c := range cg.List {
			if strings.HasPrefix(c.Text, "//go:") {
				list = append(list, c)
			}
		}
		if len(list) > 0 {
			keep = append(keep, &ast.CommentGroup{List: list})
		}
	}
	f.Comments = keep
	ast.Inspect(f, func(n ast.Node) bool {
		switch n := n.(type) {
		case *ast.File:
			n.Doc = nil
		case *ast.GenDecl:
			n.Doc = nil
		case *ast.FuncDecl:
			n.Doc = nil
		case *ast.Field:
			n.Doc, n.Comment = nil, nil
		case *ast.ValueSpec:
			n.Doc, n.Comment = nil, nil
		case *ast.TypeSpec:
			n.Doc, n.Comment = nil, nil
		case *ast.ImportSpec:
			n.Doc, n.Comment = nil, nil
		}
		return true
	})
	var buf bytes.Buffer
	cfg := printer.Config{Mode: printer.UseSpaces | printer.TabIndent, Tabwidth: 8}
	if err := cfg.Fprint(&buf, fset, f); err != nil {
		return nil, err
	}
	// drop blank lines
	var out []string
	for _, l := range strings.Split(buf.String(), "\n") {
		if strings.TrimSpace(l) == "" {
			continue
		}
		out = append(out, l)
	}
	return []byte(strings.Join(out, "\n") + "\n"), nil
}

func list(root string) map[string]bool {
	m := map[string]bool{}
	filepath.WalkDir(root, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, _ := filepath.Rel(root, p)
		if d.IsDir() {
			if rel == "_tools" || (rel != "." && strings.HasPrefix(d.Name(), ".")) {
				return filepath.SkipDir
			}
			return nil
		}
		if strings.HasSuffix(p, ".go") {
			m[rel] = true
		}
		return nil
	})
	return m
}

func main() {
	outDir := flag.String("out", "", "write stripped versions to out/{a,b}")
	flag.Parse()
	a, b := flag.Arg(0), flag.Arg(1)
	ma, mb := list(a), list(b)
	var all []string
	for k := range ma {
		all = append(all, k)
	}
	for k := range mb {
		if !ma[k] {
			all = append(all, k)
		}
	}
	sort.Strings(all)
	for _, k := range all {
		switch {
		case !mb[k]:
			fmt.Println("ONLY-MASTER", k)
		case !ma[k]:
			fmt.Println("ONLY-UPSTREAM", k)
		default:
			sa, ea := strip(filepath.Join(a, k))
			sb, eb := strip(filepath.Join(b, k))
			if ea != nil || eb != nil {
				fmt.Println("ERROR", k, ea, eb)
				continue
			}
			if !bytes.Equal(sa, sb) {
				fmt.Println("DIFF", k)
				if *outDir != "" {
					for _, x := range []struct {
						d string
						s []byte
					}{{"a", sa}, {"b", sb}} {
						p := filepath.Join(*outDir, x.d, k)
						os.MkdirAll(filepath.Dir(p), 0o755)
						os.WriteFile(p, x.s, 0o644)
					}
				}
			}
		}
	}
}
