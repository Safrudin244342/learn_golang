package main

import "fmt"

func main() {
	var nama1 = "budi"
	var nama2 = "budi"
	var nama3 = "anduk"

	var result bool

	result = nama1 == nama2 && nama1 == nama3
	fmt.Println(result)

	result = nama1 == nama2 || nama1 == nama3
	fmt.Println(result)

	result = !(nama1 == nama2)
	fmt.Println(result)
}
