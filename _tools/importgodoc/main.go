package main

import (
	"bytes"
	"flag"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"golang.org/x/tools/go/ast/astutil"
)

func main() {
	flag.Parse()
	if flag.NArg() != 2 {
		flag.Usage()
		os.Exit(2)
	}

	src, err := filepath.Abs(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	dst, err := filepath.Abs(flag.Arg(1))
	if err != nil {
		log.Fatal(err)
	}
	err = filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			if isSpecialDir(info.Name()) {
				return filepath.SkipDir
			}
			return nil
		}
		if !strings.HasSuffix(info.Name(), ".go") {
			return nil
		}

		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		outputPath := filepath.Join(dst, rel)

		doc, err := godoc(path)
		if err != nil {
			return err
		}
		if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
			return err
		}
		if err := os.WriteFile(outputPath, doc, 0644); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func godoc(path string) ([]byte, error) {
	isTest := strings.HasSuffix(path, "_test.go")

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	var decls []ast.Decl
	var comments []*ast.CommentGroup

	for _, d := range node.Decls {
		switch d := d.(type) {
		case *ast.FuncDecl:
			if d.Name == nil || !ast.IsExported(d.Name.Name) {
				continue
			}
			if !recvExported(d) {
				continue
			}

			isExample := strings.HasPrefix(d.Name.Name, "Example")
			if isTest && !isExample {
				continue
			}
			if !isTest || !isExample {
				d.Body = nil
			} else {
				for _, c := range node.Comments {
					if c.End() > d.Body.Pos() && c.End() < d.Body.End() {
						comments = append(comments, c)
					}
				}
			}
			if d.Doc != nil {
				comments = append(comments, d.Doc)
			}

		case *ast.GenDecl:
			if d.Doc != nil {
				comments = append(comments, d.Doc)
			}

			var specs []ast.Spec
			for _, spec := range d.Specs {
				switch spec := spec.(type) {
				case *ast.ValueSpec:
					var names []*ast.Ident
					for _, name := range spec.Names {
						if name == nil || (!ast.IsExported(name.Name) && name.Name != "_") {
							continue
						}
						names = append(names, name)
					}
					if len(names) == 0 {
						continue
					}
					spec.Names = names

					if spec.Doc != nil {
						comments = append(comments, spec.Doc)
					}
				case *ast.TypeSpec:
					if spec.Name == nil || !ast.IsExported(spec.Name.Name) {
						continue
					}
					if spec.Doc != nil {
						comments = append(comments, spec.Doc)
					}
				}
				specs = append(specs, spec)
			}
			if len(specs) == 0 {
				continue
			}
			d.Specs = specs
		}
		decls = append(decls, d)
	}
	node.Decls = decls

	// remove comments
	for _, g := range node.Comments {
		group := g.List[:0]
		for _, c := range g.List {
			if c.Pos() < node.Name.Pos() {
				group = append(group, c)
				continue
			}
			if isSpecialComment(c) {
				group = append(group, c)
			}
		}
		if len(group) != 0 {
			g.List = group
			comments = append(comments, g)
		}
	}
	slices.SortFunc(comments, func(a, b *ast.CommentGroup) int {
		return int(a.Pos() - b.Pos())
	})
	node.Comments = comments

	// re-write import path and shrink
	var imports []string
	for _, imp := range node.Imports {
		v, err := strconv.Unquote(imp.Path.Value)
		if err != nil {
			return nil, err
		}
		imports = append(imports, v)
	}
	if isTest {
		for _, imp := range imports {
			if isInternal(imp) || !astutil.UsesImport(node, imp) {
				astutil.DeleteImport(fset, node, imp)
			}
		}
	} else {
		for _, imp := range imports {
			if isInternal(imp) || !astutil.UsesImport(node, imp) {
				astutil.DeleteImport(fset, node, imp)
			} else {
				astutil.RewriteImport(fset, node, imp, "github.com/shogo82148/std/"+imp)
			}
		}
	}

	var buf bytes.Buffer
	if err := format.Node(&buf, fset, node); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func isSpecialDir(s string) bool {
	return strings.HasPrefix(s, ".") || s == "testdata" || s == "internal" || s == "vendor"
}

func isSpecialComment(c *ast.Comment) bool {
	return strings.HasPrefix(c.Text, "//go:build ") || strings.HasPrefix(c.Text, "// +build ")
}

func isInternal(path string) bool {
	if strings.HasPrefix(path, "internal/") {
		return true
	}
	if strings.HasSuffix(path, "/internal") {
		return true
	}
	if strings.Contains(path, "/internal/") {
		return true
	}
	return false
}

func recvExported(d *ast.FuncDecl) bool {
	if d.Recv == nil {
		return true
	}
	for _, recv := range d.Recv.List {
		switch typ := recv.Type.(type) {
		case *ast.StarExpr:
			switch typ := typ.X.(type) {
			case *ast.IndexExpr:
				if ident, ok := typ.X.(*ast.Ident); !ok || !ast.IsExported(ident.Name) {
					return false
				}
			case *ast.Ident:
				if !ast.IsExported(typ.Name) {
					return false
				}
			default:
				log.Fatalf("unknown type: %T", typ)
			}
		case *ast.Ident:
			if !ast.IsExported(typ.Name) {
				return false
			}
		}
	}
	return true
}
