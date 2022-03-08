package VarDeclaration //creating custom package

import "fmt" //format

func PrintDeclarationTypes() {
	p := 25 //Direct declaration of p value as 25
	q := 20 //Direct declaration of q value as 20

	fmt.Println(p + q)

	var a int //Declaring variable a as type int
	var b int //Declaring variable b as type int
	a = 10    //Initializing a value as 10
	b = 12    //Initializing b value as 12
	fmt.Println(a + b)

	var x int = 20 //Initializing and declaring x value as 20
	var y int = 22 //Initializing and declaring y value as 22
	fmt.Println(x + y)

	str1 := "Hello\n"
	str2 := "Welcome to GO Language"
	fmt.Println(str1 + str2) //Concatenating two strings

	var str3, str4 string //Declaring as type string
	str3 = "Bhargavi"     //Giving the input to string
	str4 = "Vennam"
	fmt.Println("My full name is :", str3+" "+str4)

	var str5 string = "This program is related to variable declaration "
	var str6 string = "using GO programming language"
	fmt.Println(str5 + str6)
}
