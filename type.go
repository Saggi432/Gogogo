package main

import "fmt"



func main() {
	//Declare variables that are set to their zero value
	var a int
	var b string //Assigns 2 words for a string ??
	var c float64
	var d bool
	fmt.Printf("var a is int\t %T [%v]\n", a, a)
	fmt.Printf("var b string \t %T [%v]\n", b, b)
	fmt.Printf("var c float64 \t %T [%v]\n", c, c)
	fmt.Printf("var b bool \t %T [%v]\n", d, d)

	//Short variable declaration operator
	
	aa := 10
	bb := "String"
	cc := 3.4090
	dd := true
	
	
	fmt.Printf("var aa is int\t %T [%v]\n", aa, aa)
	fmt.Printf("var bb string \t %T [%v]\n", bb, bb)
	fmt.Printf("var cc float64 \t %T [%v]\n", cc, cc)
	fmt.Printf("var dd bool \t %T [%v]\n", dd, dd)
}
