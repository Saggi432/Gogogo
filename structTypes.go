
package main

import "fmt"

// Readability is important
// Not important to re-structure for padding although go
// does not do it explicitly.

type example struct {
	 num int16
	 name string
	 flag bool
}



func main() {


var b bool = true //Normal declaration
c := false // short hand declaration

fmt.Printf("%+v\t%+v\n", b, c)


var new example;

fmt.Printf("%+v\n", new)
fmt.Printf("%v\n", new)

new.num=10
new.name="hello"
new.flag=true

one :=example{num:420, name:"Satish", flag:true,} //Shorthand declaration for variables
fmt.Printf("%+v\n", one)

two := example{540,"World",false}

fmt.Printf("The assigned variable declaration is %+v\n", two)


}
