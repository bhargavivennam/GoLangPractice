package DataTypePackage

import (
	"fmt"
)

// Perform all possible arithmetic operations on integer, float variables
// Calculate the min value and maximum value that can be saved in float64 variable -----> maximum value = -(2^31).((2^32)-1)
// Try type conversions and observe behavior of those conversations(compiler error)

func ArithmeticOperation() {
	var num int8 ////variable of type int8
	num = 3
	fmt.Println("int8 variable ", num)
	var num1 int16 //variable of type int16
	num1 = 3076
	fmt.Println("int16 variable ", num1)
	result := int16(num) + num1 // typeconversion by converting integer 8bit to 16bit and performing addition operation
	fmt.Println("Addition operation by typeconversion of int8 to int16 ", result)
	result1 := int16(num) - num1 // typeconversion by converting integer 8bit to 16bit and performing subtraction operation
	fmt.Println("Subtraction operation ", result1)
	result2 := int16(num) * num1 // typeconversion by converting integer 8bit to 16bit and performing multiplication operation
	fmt.Println("Multiplication operation ", result2)
	result7 := int16(num) / num1 // typeconversion by converting integer 8bit to 16bit and performing division operation
	fmt.Println("Division operation ", result7)
	//result10:=num+ num1 //gives error because they are of two different size variables.
	//fmt.Println(result10)
	var num3 float32
	num3 = 22.22
	fmt.Println("float32 value ", num3)

	var num4 float64
	num4 = 6589.38
	fmt.Println("float64 value ", num4)

	result3 := float64(num3) + num4 // typeconversion by converting float 32bit to 64bit and performing addition operation
	fmt.Println("Addition operation by typeconversion of float32 to float64 ", result3)
	result4 := float64(num3) - num4 // typeconversion by converting float 32bit to 64bit and performing subtraction operation
	fmt.Println("Subtraction operation ", result4)
	result5 := float64(num3) * num4 // typeconversion by converting float 32bit to 64bit and performing multiplication operation
	fmt.Println("Multiplication operation ", result5)

	var num5 int64
	num5 = 300009
	fmt.Println("int64 datatype ", num5)

	result6 := num5 + int64(num4) // Adding float value and integer value by converting float to integer
	fmt.Println("Addition of float and integer value ", result6)

}
