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
	var9, var7 := "var9", 12

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

}
