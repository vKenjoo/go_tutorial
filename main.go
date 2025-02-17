package main

import "fmt"

func main()  {
	/* Ints
	the key word int automatically sets it to either int32/int64 
	(depende ong your architecture)
	Signed Bits
	int8 = 1 byte with a range of [-128 | 127]
	int16 = 2 bytes with a range of [-32,768 | 32,767]
	int32 = 3 bytes with a range of [-2,147,483,648 | 2,147,483,647]
	int64 = 4 bytes with a range of [-9,223,372,036,854,775,808 | 9,223,372,036,854,775,807]
	
	Unsigned Bits (positive integers ONLY)
	uint8 = 1 byte with a range of [0 | 255]
	uint16 = 2 bytes with a range of [0 | 65,535]
	uint32 = 3 bytes with a range of [0 | 4,294,967,295]
	uint64 = 4 bytes with a range of [0 | 18,446,744,073,709,551,615]
	*/
	var intNum int // since it was not defined, it is always equal to 0
	fmt.Println(intNum)


	/* Floats
	float32 is single precision floating point
	float64 is double precision floating point

	float64 collects more memory than float32 but is more precise
	Always consider the purpose and range of the value you are using
	*/
	var floatNum1 float32 = 12345678.9
	var floatNum2 float64 = 12345678.9
	fmt.Println("here is float32: ", floatNum1)
	fmt.Println("here is float64: ", floatNum2)

	//var floatError2 float64 = floatNum1 + floatNum2
	/* invalid operation: floatNum1 + floatNum2 \
	(mismatched types float32 and float64)compilerMismatchedTypes
	*/

	// and you cant add mixed types DIRECTLY 
	// you need to type cast
	var floatError1 int = int(floatNum1) + int(floatNum2)
	fmt.Println(floatError1)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1/intNum2) // output = 1
	fmt.Println(intNum1%intNum2) // output = 1.5

	/* SUMMARY OF INTS AND FLOATS
		it is always best to considere where you will be using your value.
		for example, if you would be using x in the an rgb code 
		Ex: rgb(x, 255, 130)
		
		You might probably consider using int16 for this. And that is by no means wrong
		but there is a more better approach to it by using uint8
		
		uint8 = 1 byte 
		int16 = 2 bytes.

		that is one byte saved had you thought about it a litte more efficient
		TLDR: BE MINDFUL OF WHAT YOU USE AND WHERE YOU'll BE USING IT <3
	*/

	// Strings
	var myString1 string = "hello world"
	var myString2 string = `hello world (in backticks)` //backticks DIRECTLY format strings
	fmt.Println(myString1 + "\n" + myString2)
}
