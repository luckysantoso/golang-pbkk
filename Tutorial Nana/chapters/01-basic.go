package main

import "fmt"

func main() {
	fmt.Print("Hello World\n")

	var a int = 10   
	var b int = 20
	var max int
	fmt.Println("Nilai a:", a)
	fmt.Println("Nilai b:", b)

	
	max := (a > b) ? a : b

	fmt.Println("Nilai maksimum adalah:", max)
}
