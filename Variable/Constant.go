package main

import "fmt"

func main() {
	const NAME = "safrudin"
	const COUNTRY string = "Indonesia"

	fmt.Println(NAME)
	fmt.Println(COUNTRY)

	const (
		SATU = "satu"
		DUA  = "dua"
	)

	fmt.Println(SATU)
	fmt.Println(DUA)
}
