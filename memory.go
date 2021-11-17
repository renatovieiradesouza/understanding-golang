package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

/*
Elements position is very important here
Minimum unit = 1 bit
8 bits = 1 byte
unsafe: not recommend for use in prodution projects
*/

func main() {

	//Int 32/64 automatic - will define by your machine
	//int8
	//int16
	//int32

	var valueInt int8 = 1
	fmt.Println("Working Int")
	fmt.Println(unsafe.Sizeof(valueInt))
	fmt.Println(reflect.TypeOf(valueInt))
	fmt.Println("")

	//Array
	fmt.Println("Working Array")
	var array [3]int
	fmt.Println(unsafe.Sizeof(array))
	fmt.Println(reflect.TypeOf(array))
	fmt.Println("")

	//Slice
	//Slice not need define length
	fmt.Println("Working Slice")
	var slice []int
	fmt.Println(unsafe.Sizeof(slice))
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println("")

	fmt.Println("Understanding how work the bytes calcule")
	var data1 struct {
		sexo   bool  // 1Bytes
		idade  int32 // 4Bytes
		idade2 int32 // 4Bytes
		//Expected return = 9Bytes
		//Real return = 12Bytes
		//The bytes work with multiples of 4.
		// int32 -> [ x ][ x ][ x ][ x ] int32 -> [ x ][ x ][ x ][ x ] bool -> [ x ] | [ ][ ][ ] 3 not used
	}

	fmt.Println(unsafe.Sizeof(data1))
	fmt.Println(reflect.TypeOf(data1))
	fmt.Println("")

	var data2 struct {
		sexo   bool  // 1Bytes
		idade  int8  // 1Bytes
		idade2 int32 // 4Bytes
		//Expected return = 6Bytes
		//Real return = 8Bytes
		//The bytes work with multiples of 4.
		// int8 -> [ x ] int32 -> [ x ][ x ][ x ][ x ] bool -> [ x ] | [ ] [ ] 2 not used
	}

	fmt.Println(unsafe.Sizeof(data2))
	fmt.Println(reflect.TypeOf(data2))
	fmt.Println("")

	var data3 struct {
		sexo   bool  // 1Bytes
		idade  int8  // 1Bytes
		idade2 int64 // 8Bytes
		//Expected return = 10Bytes
		//Real return = 16Bytes
		//The bytes work with multiples of 8.
		// int64 -> [ x ][ x ][ x ][ x ][ x ][ x ][ x ][ x ] int8 -> [ x ] bool -> [ x ] | [ ][ ][ ][ ][ ][ ] 6 not used
	}

	fmt.Println(unsafe.Sizeof(data3))
	fmt.Println(reflect.TypeOf(data3))
	fmt.Println("")

	var data4 struct {
		sexo   bool  // 1Bytes
		idade  int32 // 4Bytes
		idade2 int32 // 4Bytes
		//Expected return = 9Bytes
		//Real return = 12Bytes
		//The bytes work with multiples of 4.
		// int32 -> [ x ][ x ][ x ][ x ] int32 -> [ x ][ x ][ x ][ x ] bool -> [ x ] | [ ][ ][ ] 3 not used
	}

	fmt.Println(unsafe.Sizeof(data4))
	fmt.Println(reflect.TypeOf(data4))
	fmt.Println("")

	var data5 struct {
		//The variable bool be distant of the other. However, this calcule is equal 16
		sexo   bool  // 1Bytes
		idade  int32 // 4Bytes
		idade2 int32 // 4Bytes
		nerd   bool  // 1Bytes
		//Expected return = 10Bytes
		//Real return = 16Bytes
		//The bytes work with multiples of 4.
		// bool -> [ x ][ ][ ][ ] int32 -> [ x ][ x ][ x ][ x ] int32 -> [ x ][ x ][ x ][ x ] bool -> [ x ][ ][ ][ ] |  6 not used
	}

	fmt.Println(unsafe.Sizeof(data5))
	fmt.Println(reflect.TypeOf(data5))
	fmt.Println("")

	var data6 struct {
		//The variable bool be side of the other. However, this calcule is equal 12
		idade  int32 // 4Bytes
		idade2 int32 // 4Bytes
		nerd   bool  // 1Bytes
		sexo   bool  // 1Bytes
		//Expected return = 10Bytes
		//Real return = 12Bytes
		//The bytes work with multiples of 4.
		// int32 -> [ x ][ x ][ x ][ x ] int32 -> [ x ][ x ][ x ][ x ] bool -> [ x ][ x ][ ][ ] |  2 not used
	}

	fmt.Println(unsafe.Sizeof(data6))
	fmt.Println(reflect.TypeOf(data6))
	fmt.Println("")

	var data7 struct {
		idade  int32 // 4Bytes
		idade2 int64 // 8Bytes
		nerd   bool  // 1Bytes
		//Expected return = 13Bytes
		//Real return = 24Bytes
		//The bytes work with multiples of 8.
		/* int32 -> [ x ][ x ][ x ][ x ] [ ][ ][ ][ ] int64 -> [ x ][ x ][ x ][ x ][ x ][ x ][ x ][ x ]
		bool -> [ x ][ ][ ][ ][ ][ ][ ][ ] |  11 not used
		*/
	}

	fmt.Println(unsafe.Sizeof(data7))
	fmt.Println(reflect.TypeOf(data7))
	fmt.Println("")
}
