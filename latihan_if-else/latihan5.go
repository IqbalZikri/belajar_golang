// Latihan 5
// Buatlah program yang meminta pengguna memasukkan berat (dalam kilogram) dan tinggi (dalam meter). Program kemudian menghitung IMT dan memberikan kategori berat badan, seperti "Kurus", "Normal", "Gemuk", atau "Obesitas".

package main

import (
	"fmt"
)

func main(){
	var beratBadan float64
	var tinggiBadan float64
	fmt.Print("Masukkan berat badan anda: ")
	fmt.Scan(&beratBadan)
	fmt.Print("Masukkan Tinggi anda: ")
	fmt.Scan(&tinggiBadan)
	imt := beratBadan / (tinggiBadan * tinggiBadan)
	fmt.Printf("IMT anda adalah %g\n", imt)
}