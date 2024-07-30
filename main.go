package main

import "fmt"

func main()  {
	var intNum int // since it was not defined, it is always equal to 0
	fmt.Println(intNum)

	/* the key word int automatically sets it to either int32/int64 
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
}
