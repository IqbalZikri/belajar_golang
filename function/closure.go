// Closure adalah sebuah function berinteraksi dengan data-data disekitarnya dalam scope yang sama
// Harap gunakan fitur closure ini dengan bijak saat kita membuat aplikasi

package main

import "fmt"

func main(){
	counter := 0

	// Anonymous Function
	increment := func(){
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()
	fmt.Println(counter)
}