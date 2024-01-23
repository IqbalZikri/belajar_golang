// Latihan 2
// Buatlah program yang meminta pengguna memasukkan usia. Program kemudian memberikan kategori usia, seperti "Balita", "Anak-anak", "Remaja", "Dewasa", atau "Lansia".
package main

import "fmt"

func main(){
	var usia int
	fmt.Print("Masukkan usia anda: ")
	fmt.Scan(&usia)
	if usia > 60 {
		fmt.Println("Anda sudah lansia")
	} else if usia >= 21 && usia <= 60 {
		fmt.Println("Anda sudah memasuki usia dewasa")
	}else if usia >= 12 && usia <= 20  {
		fmt.Println("Anda sudah memasuki usia remaja")
	}else if usia >= 6 && usia <= 11 {
		fmt.Println("Anda masih anak anak")
	}else if usia >= 0 && usia <= 5{
		fmt.Println("Anda balita1")
	}
}