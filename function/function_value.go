// Function Value
// Function adalah first class citizen
// function bisa kita simpan di dalam variable atau dijadikan sebuah value atau nilai
// function juga bisa di bilang sebagai tipe data jika di golang

package main

import "fmt"

func getGoodBye(name string)string{
	return "Good Bye " + name
}

func main(){
	goodbye := getGoodBye
	fmt.Println(goodbye("Ball"))

	// bisa juga masuk ke lebih dari 1 variable
	contoh := getGoodBye
	fmt.Println(contoh("Budi"))

}