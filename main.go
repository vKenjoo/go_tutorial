package main

import "fmt"

func MoveLastToFirst(arr []int) []int {
    if len(arr) == 0 {
        return arr
    }
    
    // Get the last element
    lastElement := arr[len(arr)-1]
    
    // Create a new slice with the last element at the beginning
    newArr := append([]int{lastElement}, arr[:len(arr)-1]...)
    
    return newArr
}

func main()  {
	var intNum int // since it was not defined, it is always equal to 0
	fmt.Println(intNum)

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

	var floatNum1 float32 = 12345678.9
	var floatNum2 float64 = 12345678.9
	fmt.Println("here is float32: ", floatNum1)
	fmt.Println("here is float64: ", floatNum2)

	/* Floats
	float32 is single precision floating point
	float64 is double precision floating point

	float64 collects more memory than float32 but is more precise
	Always consider the purpose and range of the value you are using
	*/

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
	
	// hello!

	for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

	arr := []int{1, 2, 3, 4, 5}
    fmt.Println("Original array:", arr)
    arr = MoveLastToFirst(arr)
    fmt.Println("Modified array:", arr)
}
