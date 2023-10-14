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

	var packageName string
	if node.Name != nil {
		packageName = node.Name.Name
	}
	isTestPackage := strings.HasSuffix(packageName, "_test")
	isExample := isTestPackage && includesExampleTest(node)

	var decls []ast.Decl
	var comments []*ast.CommentGroup
	seen := map[*ast.CommentGroup]bool{}
	addComment := func(c *ast.CommentGroup) {
		if seen[c] {
			return
		}
		seen[c] = true
		comments = append(comments, c)
	}

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
						addComment(c)
					}
				}
			}
			if d.Doc != nil {
				addComment(d.Doc)
			}

		case *ast.GenDecl:
			var specs []ast.Spec
			for _, spec := range d.Specs {
				switch spec := spec.(type) {
				case *ast.ImportSpec:
					// 名前付きインポートは、ドキュメント中では使わないので除外する
					if spec.Name != nil {
						continue
					}
				case *ast.ValueSpec:
					var names []*ast.Ident
					for _, name := range spec.Names {
						// テスト中の変数は除外する
						if isTest && !isExample {
							continue
						}
						// 非公開の変数は除外する
						if name == nil || (!ast.IsExported(name.Name) && name.Name != "_") {
							continue
						}
						names = append(names, name)
					}
					if len(names) == 0 {
						continue
					}
					spec.Names = names

					if d.Doc != nil {
						addComment(d.Doc)
					}
					if spec.Doc != nil {
						addComment(spec.Doc)
					}
				case *ast.TypeSpec:
					if isTest || spec.Name == nil || !ast.IsExported(spec.Name.Name) {
						continue
					}

					if d.Doc != nil {
						addComment(d.Doc)
					}
					if spec.Doc != nil {
						addComment(spec.Doc)
					}
					if st, ok := spec.Type.(*ast.StructType); ok {
						for _, f := range st.Fields.List {
							if f.Doc != nil {
								comments = append(comments, f.Doc)
							}
						}
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
		if g.Pos() < node.Name.Pos() {
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
	for _, imp := range imports {
		if !astutil.UsesImport(node, imp) {
			// 未使用のimport文を削除
			astutil.DeleteImport(fset, node, imp)
		} else {
			// import文をgithub.com/shogo82148/stdに置き換える
			astutil.RewriteImport(fset, node, imp, "github.com/shogo82148/std/"+imp)
		}
	}

	var buf bytes.Buffer
	if err := format.Node(&buf, fset, node); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func isSpecialDir(s string) bool {
	return strings.HasPrefix(s, ".") || s == "testdata" || s == "vendor"
}

func includesExampleTest(node *ast.File) bool {
	for _, d := range node.Decls {
		if d, ok := d.(*ast.FuncDecl); ok && d.Name != nil && strings.HasPrefix(d.Name.Name, "Example") {
			return true
		}
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
			case *ast.IndexListExpr:
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
