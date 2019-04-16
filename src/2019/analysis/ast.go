package main

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"
)

func main() {
	// 解析するソースコード
	const src = `
package main

import (
	"fmt"
)

func main() {
	greeting := "Hello, world"
	fmt.Println(greeting)
}`

	// 構文解析
	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, "my.go", src, 0)
	if err != nil {
		log.Fatal(err.Error())
	}

	// AST出力
	ast.Print(fs, f)

	// 型情報解析
	info := types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
		Defs:  make(map[*ast.Ident]types.Object),
		Uses:  make(map[*ast.Ident]types.Object),
	}

	conf := &types.Config{
		Importer: importer.Default(),
	}
	_, err = conf.Check("fib", fs, []*ast.File{f}, &info)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n\nTypes")
	for k, v := range info.Types {
		fmt.Printf("%v : %v\n", k, v)
	}

	fmt.Println("\n\nDefs")
	for k, v := range info.Defs {
		fmt.Printf("%v : %v\n", k, v)
	}

	fmt.Println("\n\nUses")
	for k, v := range info.Uses {
		fmt.Printf("%v : %v\n", k, v)
	}

	// ASTをトラバースして型チェック
	fmt.Println("\n\n#### Inspect!!")
	ast.Inspect(f, func(n ast.Node) bool {

		// 識別子ではない場合は無視
		ident, ok := n.(*ast.Ident)
		if !ok {
			return true
		}

		// 識別子が定義または利用されている部分の情報を取得
		obj := info.Defs[ident]
		if obj == nil {
			obj = info.Uses[ident]
			if obj == nil {
				return true
			}
		}

		typ := obj.Type()

		fmt.Printf("%v %s %T %v\n", fs.Position(ident.Pos()), ident.Name, typ, typ)

		return true
	})

}
