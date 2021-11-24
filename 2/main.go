package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func Count(fileName, funcName string) (int64, error) {
	var count int64
	set := token.NewFileSet()
	astFile, err := parser.ParseFile(set, fileName, nil, 0)
	if err != nil {
		return 0, err
	}

	for _, d := range astFile.Decls {
		if fn, isFn := d.(*ast.FuncDecl); isFn {
			if fn.Name.String() != funcName {
				continue
			}
			count = countGoStmt(fn.Body.List)
			break
		}
	}
	return count, nil
}

func countGoStmt(stmts []ast.Stmt) int64 {
	var count int64
	for _, stmt := range stmts {
		switch v := stmt.(type) {
		case *ast.GoStmt:
			count++
		case *ast.IfStmt:
			count += countGoStmt(v.Body.List)
		case *ast.ForStmt:
			count += countGoStmt(v.Body.List)
		case *ast.SwitchStmt:
			count += countGoStmt(v.Body.List)
		case *ast.CaseClause:
			count += countGoStmt(v.Body)
		}
	}
	return count
}

func main() {
	count, err := Count("input.go", "ConcurrentWorker")
	if err != nil {
		fmt.Println("Error during async functions count:", err)
	} else {
		fmt.Println("Find async functions:", count)
	}

}
