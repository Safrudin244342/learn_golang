package main

import "fmt"

func main() {
	type Married bool

	var marriedStatus Married = false
	fmt.Println(marriedStatus)
}
