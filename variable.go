// Variable
// Variable digunakan agar kita bisa mengakses data yang sama dimanapun kita mau
// Variable dalam golang hanya bisa menyimpan tipe data yang sama, jika ingin berbeda maka harus membuat beberapa variable
// untuk membuat variable gunakan kata kunci var, lalu diikuti dengan nama variable dan tipe datanya

package main

import "fmt"

func main(){
	var name string
	// name adalah nama dari variablenya
	// string adalah tipe datanya

	name = "Tubagus Muhammad Iqbal Zikri"
	fmt.Println(name)

	name = "Cid Kagenou"
	fmt.Println(name)

	// namun jika kita langsung menginisialisasikan data pada variablenya, maka tidak wajib menyebutkan tipe data variablenya
	var kampus = "Universitas Cendekia Abditama"
	fmt.Println(kampus)

	// jika tidak ingin menggunakan kata kunci var gunakan :=
	nim := 12110045
	fmt.Println(nim)

	// Multiple Variable
	var (
		firstName = "Cid"
		lastName = "Kagenou"
		kodePos = 15143
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(kodePos)
}