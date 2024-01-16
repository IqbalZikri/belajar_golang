// constant
// constant adalah kata kunci seperti variable, bedanya ia tidak bisa diubah lagi setelah pertama kali diberi nilai
// kata kunci yang digunakan adalah const

package main

import "fmt"

func main(){
	const nama string = "Tubagus Muhammad Iqbal Zikri"
	const kampus = "Harvard University"

	// error
	// nama = "Bambang"
	// kampus = "Bimbing"
	// error karen constant tidak bisa diubah lagi

	fmt.Println(nama)
	fmt.Println(kampus)

	const(
		agama = "Islam"
		hobi = "apa ya" 
		target = "hmmm"
	)

	fmt.Println(agama)
	fmt.Println(hobi)
	fmt.Println(target)
}