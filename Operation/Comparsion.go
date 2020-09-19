package main

import "fmt"

func main() {
	var nama1 = "budi"
	var nama2 = "budi"
	var nama3 = "anduk"

	var result = nama1 == nama2

	fmt.Println(result)

	result = nama2 == nama3

	fmt.Println(result)
}
