// +build ignore

/*
 * Copyright 2020 The Goava authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"bufio"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
	"text/template"
)

const genCmt = "// Code generated. DO NOT EDIT."

func main() {
	// sanity check for ensuring that this program is invoked from the pkg directory
	dir, _ := os.Getwd()
	srcFile := path.Join(dir, "gen.go")
	if _, err := os.Stat(srcFile); err == nil {
		dir = path.Dir(path.Dir(dir))
	}

	srcFile = path.Join(dir, "base", "runematcher", "gen.go")
	if stat, err := os.Stat(srcFile); err != nil || stat.IsDir() {
		log.Fatalln("Cannot run Matcher generator in", dir)
	}

	// load function from interface
	srcFile = path.Join(dir, "base", "runematcher", "matcher.go")
	decls := parse(srcFile)["Matcher"]

	// load actual implementation
	srcFile = path.Join(dir, "base", "runematcher", strings.ToLower(os.Args[1])+".go")
	log.Println("Generating delegate methods in", srcFile)
	f, err := os.OpenFile(srcFile, os.O_RDWR, 0644)
	die(err)
	defer f.Close()

	fnNames := extractFnNames(f)

	var notFound []ast.Field
	for _, decl := range decls {
		if _, ok := fnNames[decl.Names[0].Name]; !ok {
			notFound = append(notFound, decl)
		}
	}

	for _, x := range []string{"\n", genCmt, "\n"} {
		_, err = f.WriteString(x)
		die(err)
	}

	for _, nf := range notFound {
		params := nf.Type.(*ast.FuncType).Params.List
		results := nf.Type.(*ast.FuncType).Results

		sig := strings.Builder{}
		args := strings.Builder{}
		for i, p := range params {
			if i != 0 {
				sig.WriteString(", ")
			}
			args.WriteString(", ")

			if t, ok := p.Type.(*ast.ArrayType); ok {
				sig.WriteString(p.Names[0].Name + " []" + t.Elt.(fmt.Stringer).String())
			} else {
				sig.WriteString(p.Names[0].Name + " " + p.Type.(fmt.Stringer).String())
			}
			args.WriteString(p.Names[0].Name)
		}

		doc := strings.TrimSuffix(nf.Doc.Text(), "\n")
		err = delegateTmpl.Execute(f, struct {
			Doc     string
			Matcher string
			Name    string
			Sig     string
			Ret     string
			Call    string
			Args    string
		}{
			strings.ReplaceAll(doc, "\n", "\n// "),
			os.Args[1],
			nf.Names[0].Name,
			sig.String(),
			results.List[0].Type.(fmt.Stringer).String(),
			strings.ToLower(nf.Names[0].Name[0:1]) + nf.Names[0].Name[1:],
			args.String(),
		})
	}

	die(err)
}

var delegateTmpl = template.Must(template.New("delegate").Parse(`
// {{.Doc}}
func (m {{.Matcher}}Matcher) {{.Name}}({{.Sig}}) {{.Ret}} {
	return {{.Call}}(m{{.Args}})
}
`))

func extractFnNames(f *os.File) map[string]struct{} {
	names := map[string]struct{}{}
	r := regexp.MustCompile(`.+\s(\S+)\(.+`)
	s := bufio.NewScanner(f)
	var n = 0
	for ok := s.Scan(); ok; ok = s.Scan() {
		if strings.Contains(s.Text(), genCmt) {
			_, err := f.Seek(int64(n-1), io.SeekStart)
			die(err)
			return names
		} else if strings.HasPrefix(s.Text(), "func (") {
			name := r.ReplaceAllString(s.Text(), "$1")
			names[name] = struct{}{}
		}
		n += len(s.Text()) + 1
	}

	return names
}

func die(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func parse(file string) map[string][]ast.Field {
	fset := token.NewFileSet()
	var fileAst *ast.File
	var err error
	if fileAst, err = parser.ParseFile(fset, file, nil, parser.ParseComments); err != nil || fileAst == nil {
		panic(err)
	}

	return scanAst(fileAst)
}

func scanAst(fileAst *ast.File) map[string][]ast.Field {
	fnsByIf := map[string][]ast.Field{}
	for _, decl := range fileAst.Decls {
		if decl, ok := decl.(*ast.GenDecl); ok {
			var fns []ast.Field
			for _, spec := range decl.Specs {
				if spec, ok := spec.(*ast.TypeSpec); ok {
					ms := spec.Type.(*ast.InterfaceType).Methods
					for _, m := range ms.List {
						if len(m.Names) > 0 {
							fns = append(fns, *m)
						}
					}
					fnsByIf[spec.Name.Name] = fns
				}
			}
		}
	}

	return fnsByIf
}
