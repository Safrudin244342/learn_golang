package main

import "fmt"

func main() {
	var nilai32 int32 = 23434
	var nilai64 int64 = int64(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)

	var nama string = "azis"
	var huruf byte = nama[0]
	var Shuruf = string(huruf)

	fmt.Println(nama)
	fmt.Println(huruf)
	fmt.Println(Shuruf)
}
