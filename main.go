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
	var b bool = s1 == s2

	fmt.Printf(
		"\nString:"+
			"\ns1 = %s"+
			"\ns2 = %s"+
			"\nb = %v",
		s1, s2, b)
}
