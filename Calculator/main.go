package main

import (
	Operations "GOMODULE/calculator/Operations"
	"fmt"
)

//Calling all the functions using switchcase
func main() {
	fmt.Println("Choose one of the following options:\n1. Addition\n2.Subtraction\n3.Multiplication\n4.Division")
	var option int
	fmt.Scan(&option)

	switch option {
	case 1:
		fmt.Println("Addition operation")
		Operations.AddInt8()
		Operations.AddInt16()
		Operations.AddInt32()
		Operations.AddInt64()

		Operations.AddUint8()
		Operations.AddUint16()
		Operations.AddUint32()
		Operations.AddUint64()

		Operations.AddFloat32()
		Operations.AddFloat64()

	case 2:
		fmt.Println("Subtraction Operation")
		Operations.SubInt8()
		Operations.SubInt16()
		Operations.SubInt32()
		Operations.SubInt64()

		Operations.SubUint8()
		Operations.SubUint16()
		Operations.SubUint32()
		Operations.SubUint64()

		Operations.SubFloat32()
		Operations.SubFloat64()

	case 3:
		fmt.Println("Multiplication operation")
		Operations.MultiplyInt8()
		Operations.MultiplyInt16()
		Operations.MultiplyInt32()
		Operations.MultiplyInt64()

		Operations.MultiplyUint8()
		Operations.MultiplyUint16()
		Operations.MultiplyUint32()
		Operations.MultiplyUint64()

		Operations.MultiplyFloat32()
		Operations.MultiplyFloat64()

	case 4:
		fmt.Println("Division operation")
		Operations.DivideInt8()
		Operations.DivideInt16()
		Operations.DivideInt32()
		Operations.DivideInt64()

		Operations.DivideUint8()
		Operations.DivideUint16()
		Operations.DivideUint32()
		Operations.DivideUint64()

		Operations.DivideFloat32()
		Operations.DivideFloat64()
	}

}
