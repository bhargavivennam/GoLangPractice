package scope

import "fmt"

var pkgVariable int = 200 // declaring the variable in package level scope
var GlobalVar int = 22

func ScopeDeclaration() {
	var locVariable int
	locVariable = 5 //a is local variable it is only initialised for this function
	fmt.Println("local variable or functional level scope value: ", locVariable)

	fmt.Println("Pkg levelvariable", pkgVariable) //Calling the package level variable to the function
}
func SampleFunc() {
	var locVariable int
	locVariable = 10 //as this is different function the value is changed to 10
	fmt.Println("local variable or functional level scope value: ", locVariable)

	fmt.Println("Pkg level variable", pkgVariable) //Calling the package level variable to the function
}
