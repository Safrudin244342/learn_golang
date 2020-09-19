package main

import "fmt"

func main() {
	var a = 10
	var b = 10

	var plus = a + b
	var minus = a - b
	var kali = a * b
	var bagi = a / b
	var sisa = a % b

	fmt.Println(plus)
	fmt.Println(minus)
	fmt.Println(kali)
	fmt.Println(bagi)
	fmt.Println(sisa)

	var i = 10
	i += 10 //augmented assignments

	fmt.Println(i)

	i++ //unary operator

	fmt.Println(i)

	i = -i

	fmt.Println(i)
}
