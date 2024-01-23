// Panic
// Panic function adalh function yang bisa kita gunakan untuk menghentikan program
// Panic function biasanya dipanggil ketika terjadi panic pada saat program kita berjalan
// Saat panic function dipanggil, program akan terhenti, namun defer function akan tetap dieksekusi

// Recover
// Recover adalah function yang bisa kita gunakan untuk menangkap data panic
// Dengan recover proses panic akan terhenti, sehingga program akan tetap berjalan

package main

import "fmt"

func endApp(){
	fmt.Println("End App")
	message := recover()
	fmt.Println("Terjadi panic", message)
}

// Contoh penggunaan panic
func runApp(error bool){
	defer endApp()
	if error {
		panic("ERROR")
	}
}

// Kode program recover yang salah
// func runAppR(error bool){
// 	defer endApp()
// 	if error {
// 		panic("ERROR")
// 	}
// 	message := recover() // recover langsung mengambil isi atau data dari panic dan menyimpannya dalam message
// 	fmt.Println("Terjadi panic", message)

	// Penggunaang program recover di atas salah karena ketika program dieksekusi, program akan berhenti setelah mendeteksi panic sehingga program dibawah panic tidak akan dieksekusi
// }

// Kode program recover yang benar
func runAppRecover(error bool){
	defer endApp() // masukkan recover ke dalam function endApp()
	if error {
		panic("Error")
	}
}

func main(){
	runApp(true)
	fmt.Println("Tubagus Muhammad Iqbal Zikri")
}