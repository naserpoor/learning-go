package main

import (
	"fmt"
)

func main() {
	// Salam
	fmt.Println("Salam!!!\n")

	// base types
	var t1 int = 100
	var t2 uint32 = 120
	var t3 byte = 100
	var t4 rune = 'A'
	var t5 uintptr = 0x1121

	fmt.Printf("\nIntegers:"+
		"\nt1 = %d"+
		"\nt2 = %d"+
		"\nt3 = %d"+
		"\nt4 = %d"+
		"\nt5 = %d",
		t1, t2, t3, t4, t5)

	var f1 float32 = 12.5
	var f2 float64 = 11.24

	fmt.Printf("\nFloat:"+
		"\nf1 = %f"+
		"\nf2 = %f",
		f1, f2)

	var c1 complex64 = complex(1.2, 2.3)
	var c2 complex128 = complex(1.3, 2.4)

	fmt.Printf("\nComplex:"+
		"\nc1 = %v"+
		"\nc2 = %v",
		c1, c2)

	var s1 string = "Alirexa"
	var s2 string = "Alireza"
	var b1 bool = s1 == s2

	fmt.Printf(
		"\nString:"+
			"\ns1 = %s"+
			"\ns2 = %s"+
			"\nb1 = %v",
		s1, s2, b1)

	// variables
	var var1 string = "var1"
	var var2 = "var2"
	var var3 string
	var var4, var5 string = "var4", "var5"
	var var6, var7 = "var6", 12

	var8 := "var8"
	var9, var7 := "var9", 13

	fmt.Printf("\nVariables:"+
		"\n var1 = %v"+
		"\n var2 = %v"+
		"\n var3 = %v"+
		"\n var4 = %v"+
		"\n var5 = %v"+
		"\n var6 = %v"+
		"\n var7 = %v"+
		"\n var8 = %v"+
		"\n var9 = %v",
		var1, var2, var3, var4, var5,
		var6, var7, var8, var9)

	// constants
	const typed = 12
	const untyped float64 = 1.12

	// var temp2 = var7 + int(f1)
	// var temp1 = typed + f1

	// integer const
	// 85
	// 0213
	// 0x4b
	// 30
	// 30u
	// 30l
	// 30ul
	// 212
	// 215u
	// 0xFeeL

	// complex const
	// (0.0, 0.0)
	// (-123.456E+30, 987.654E-29)

	// float const
	// 3.14159
	// 314159E-5L

	// string const
	// 	"hello, geeksforgeeks"
	// "hello, \
	// geeksforgeeks"
	// "hello, " "geeks" "forgeeks"
	// var sd =  `salam
	// aleik`

	// condition
	if t1 > 99 {

	} else if t1 > 200 {

	} else {

	}

	// for loop
	for i := 0; i < 1000; i++ {

	}

	for {
		break
	}

	for t1 > 100 {

	}

	array := []string{"1", "2", "3"}
	for _, item := range array {
		fmt.Print(item)
	}

	mmap := map[int]string{
		22: "S1",
		33: "S2",
	}
	for key, value := range mmap {
		fmt.Println(key, value)
	}

	chnl := make(chan int)
	go func() {
		chnl <- 100
		chnl <- 1000
		chnl <- 10000
		chnl <- 100000
		close(chnl)
	}()
	for i := range chnl {
		fmt.Println(i)
	}

	// switch

	switch variable := 4; variable {
	case 1:
		fmt.Println("Case 1")
	case 2:
		fmt.Println("Case 2")
	case 3:
		fmt.Println("Case 3")
	case 4:
		fmt.Println("Case 4")
	default:
		fmt.Println("Invalid")
	}
	value := 1
	switch {
	case value == 1:
		fmt.Println("Hello")
	case value == 2:
		fmt.Println("Hello 2")
	case value == 3:
		fmt.Println("Hello 3")
	default:
		fmt.Println("Hello 4")
	}


	// select
	c := make(chan int)
	c <- 10
	select {
		case x1 := <- c:
			fmt.Println("Value: ", x1)
		default:
			fmt.Println("Default case..!")
    }

}
