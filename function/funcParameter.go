// Function Parameter
// kadang-kadang kita membutuhkan data dari luar, atau kita sebut parameter
// kita bisa menambahkan parameter di function lebih dari satu
// parameter tidak lah wajib
// jika kita memasukkan parameter ke function, maka wajib memasukkan data ke parameternya saat memanggil function tersebut


package main

import "fmt"

func sayHelloTo(firstName string, lastName string){
	fmt.Println("Hello", firstName, lastName)
}

func main(){
	sayHelloTo("Iqbal", "TB") // Iqbal adalah parameter pertama (firstName), TB adalah parameter kedua (lastName)
	// hasil yang akan keluar adalah Hello Iqbal TB
	sayHelloTo("Zikri", "Tubagus")
	// output == Hello Zikri Tubagus
}