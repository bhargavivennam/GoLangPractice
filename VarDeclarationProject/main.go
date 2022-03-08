package main

import (
	"fmt"
	DataTypePackage "goLangAssignments/VarDeclarationProject/DataTypesModule"
	scope "goLangAssignments/VarDeclarationProject/ScopesDeclarationModule"
	VarDeclaration "goLangAssignments/VarDeclarationProject/VarDeclarationModule"
)

func main() {
	VarDeclaration.PrintDeclarationTypes() //calling the function using package name
	scope.ScopeDeclaration()
	scope.SampleFunc()
	fmt.Println("Global level var", scope.GlobalVar) //for global level scope it is packagename.variablename
	DataTypePackage.DataType()
	DataTypePackage.ArithmeticOperation()
}
