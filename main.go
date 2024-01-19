package main

import (
	"fmt"
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/analysis/singlechecker"
	"golang.org/x/tools/go/ast/inspector"

	. "github.com/dave/jennifer/jen"
)

// surrealTagAnalyzer は構造体のフィールドタグを解析するためのAnalyzerを定義します。
var surrealTagAnalyzer = &analysis.Analyzer{
	Name: "surrealtag",
	Doc:  "generates code for surreal struct tags",
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.TypeSpec)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		typeStmt := n.(*ast.TypeSpec)
		// return if not struct
		structType, ok := typeStmt.Type.(*ast.StructType)
		if !ok {
			return
		}
		structName := typeStmt.Name.Name
		fields := []*ast.Field{}

		for _, field := range structType.Fields.List {
			if field.Tag != nil {
				tag := strings.Trim(field.Tag.Value, "`")
				if strings.Contains(tag, `surreal:`) {
					fields = append(fields, field)
				}
			}
		}
		if 0 < len(fields) {
			generateStructByFields(structName, fields)
		}
	})

	return nil, nil
}

func generateStructByFields(structName string, fields []*ast.Field) {
	f := NewFile(strings.ToLower(structName) + "_surreal")
	fieldDecls := []Code{}
	for _, field := range fields {
		fieldName := field.Names[0].Name
		// get surreal tag value
		tag := strings.Trim(field.Tag.Value, "`")
		tag = strings.Replace(tag, `surreal:`, "", 1)
		tag = strings.Replace(tag, `"`, "", -1)
		jsonName := strings.Split(tag, ",")[0]
		fieldDecls = append(fieldDecls, Id(fieldName).Id("StrandJSON").Tag(map[string]string{"json": jsonName}))
	}
	f.Type().Id(structName + "JSON").Struct(fieldDecls...)
	f.Type().Id("StrandJSON").Struct(Id("Strand").String().Tag(map[string]string{"json": "strand"}))

	block := []Code{Id("s").Op(":=").Id(structName).Values()}
	for _, field := range fields {
		fieldName := field.Names[0].Name
		block = append(block, Id("s").Dot(fieldName).Op("=").Id("j").Dot(fieldName).Dot("Strand"))
	}
	block = append(block, Return(Id("s")))

	f.Func().Id("From" + structName + "JSON").Params(Id("j").Id(structName + "JSON")).Id(structName).Block(
		block...,
	)
	fmt.Printf("%#v", f)
}

func main() {
	singlechecker.Main(surrealTagAnalyzer)
}
