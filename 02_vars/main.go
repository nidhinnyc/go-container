package main

import "fmt"

//name := "Nidhin" //this type of declaration is not permitted

func main() {
	//var name = "Nidhin"
	name := "Nidhin S"
	//var age int = 35
	//var age = 35 //default integer type is infered as int
	var age int64 = 35
	//var isPresent = false
	const isPresent = false
	// isPresent = true //const can be changed
	pie := 3.14  //by default infered float type is float64
	a, b := 1, 2 // multiple variables can be declared and assigned in a single line
	fmt.Println(name, age, isPresent, pie, a, b)
	fmt.Printf("%T\n", pie)
}
