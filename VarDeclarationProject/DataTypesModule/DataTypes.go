package DataTypePackage

import "fmt"

func DataType() {
	//Boolean datatype
	var a bool
	a = true
	fmt.Println(a)

	//a="Bhargavi" -- cannot declare as string because variable a is of boolean datatype

	var b string
	b = "Hello this is string datatype"
	fmt.Println(b)

	var num int //we know that 1byte=8bits
	num = 20
	fmt.Println("Integer datatype", num)
	fmt.Printf("%v, %T\n", num, num) //%v is for the value of the variable and %T for the type of varaiable

	var num1 int8 //8 bits = 1 byte, 2^8=256 combination
	num1 = 99     //range is -128 to 127 ->-(2^7)to 2^7-1
	fmt.Println("This is int8 datatype", num1)

	var num2 int16 //16 bits = 2 bytes, 2^16 combinations
	num2 = 10000   //range is -32,768 to 32,767 ->-(2^15) to 2^15-1
	fmt.Println("This is int16 datatype", num2)

	var num3 int32   //32 bits = 4 bytes, 2^32 combinations
	num3 = 548921345 //range is -2,147,483,648 to 2,147,483,647 -> -(2^31) to 2^31-1
	fmt.Println("This is int32 datatype", num3)

	var num4 int64             //64 bits = 8 bytes, 2^64 combinations
	num4 = 9182949948895939392 //range is -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807 -> -(2^63) to 2^63-1
	fmt.Println("This is int64 datatype", num4)

}
