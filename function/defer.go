// Defer
// Defer function adalah function yang bisa kita jadwalkan untuk dieksekusi setelah sebuah function selesai dieksekusi
// Defer function akan selalu dieksekusi walaupun terjadi error di function yang akan dieksekusi

package main

import "fmt"

func logging(){
	fmt.Println("Selesai memanggil function")
}

func runApplication(){
	fmt.Println("Run Application")
	// Jika terjadi error maka func logging dibawah ini tidak akan dieksekusi karena tidak menggunakan defer
	logging()

	// defer logging()
	// fmt.Println("Run Application")
}

func main(){
	runApplication()
}