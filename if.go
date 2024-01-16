// if expression / pengkondisian
// kata kunci yang digunakan untuk percabangan

package main

import "fmt"

func main(){
	// nama := "Bambang"

	// if (nama == "iqbal") {
	// 	fmt.Println("Hallo iqbal")
	// }else{
	// 	fmt.Println("Kamu Siapaaaaa")
	// }

	// angka1 := 2012
	// angka2 := 2022

	// if angka1 > angka2 {
	// 	fmt.Println("Angka 1 lebih besar")
	// }else if angka2 > angka1 {
	// 	fmt.Println("Angka 2 lebih besar")
	// }

	// nilai := 20

	// if nilai >= 90 && nilai <= 100 {
	// 	fmt.Println("A")
	// }else if nilai >= 80 && nilai <= 89 {
	// 	fmt.Println("B")
	// }else if nilai >= 70 && nilai <= 79  {
	// 	fmt.Println("C")
	// }else{
	// 	fmt.Println("Kamu belum ikut ujian")
	// }

	// beratPaket := 5

	// if beratPaket < 1 {
	// 	fmt.Println("Biaya ongkir 5 dollar")
	// }else if beratPaket > 1 && beratPaket < 5 {
	// 	fmt.Println("Biaya ongkir 10 dollar")
	// }else if beratPaket >= 5 {
	// 	fmt.Println("Biaya ongkir adalah 20 dollar")
	// }


	// angka := 0

	// if angka >= 1 {
	// 	fmt.Println("Ini adalah angka positif")
	// }else if angka <= -1 {
	// 	fmt.Println("Ini adalah angka negatif")
	// }else{
	// 	fmt.Println("ini adalah angka 0")
	// }
	

	angka := 8

	if angka % 2 == 0 {
		fmt.Println("ini adalah angka genap")
	}else{
		fmt.Println("ini adalah angka ganjil")
	}

	name := "Tubagus Muhammad"
	panjangNama := len(name)

	if panjangNama >= 20 {
		fmt.Println("Namanya Kepanjangan woiii")
	}else{
		fmt.Println("Nama diterima")
	}

	if length := len(name); length > 5{
		fmt.Println("Namanya kepanjangan")
	}else{
		fmt.Println("Nama diterima")
	}
}