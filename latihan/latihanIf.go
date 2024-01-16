package main

import (
	"fmt"
)

func stringToList(s string) []rune {
	// konversi string ke slice rune
	chars := []rune(s)
	return chars
}

func main() {
	// Cek Bilangan Ganjil dan Genap
	// angka := 5

	// if angka % 2 == 0 {
	// 	fmt.Println("Ini adalah angka genap")
	// }else{
	// 	fmt.Println("Ini adalah angka ganjil")
	// }

	// Hitung Diskon Pembelian
	// Buat program yang meminta pengguna memasukkan jumlah total pembelian. Jika total pembelian lebih dari 100, berikan diskon sebesar 10%. Tampilkan total yang harus dibayarkan setelah diskon.
	// hargaBarang := 20000.0
	// totalBarang := 120
	// totalHargaBarang := hargaBarang * float64(totalBarang)
	// fmt.Println(totalHargaBarang)
	// persentaseDiskon := 10.0

	// if totalBarang >= 100 {
	// 	fmt.Println("Anda mendapatkan diskon sebesar 10%")
	// 	potonganHarga := totalHargaBarang * (persentaseDiskon / 100)
	// 	fmt.Println(potonganHarga)
	// 	hargaSetelahDiskon := totalHargaBarang - potonganHarga
	// 	fmt.Println("Harga barang setelah diskon:",hargaSetelahDiskon)
	// }else{
	// 	fmt.Println("Anda tidak mendapatkan diskon")
	// }

	// Kategori Usia
	// Buat program yang meminta pengguna memasukkan usia mereka. Tampilkan kategori usia berdasarkan aturan berikut:
	// 0-12 tahun: Anak-anak
	// 13-17 tahun: Remaja
	// 18-59 tahun: Dewasa
	// 60 tahun ke atas: Lansia

	// var usia int
	// fmt.Print("Masukkan usia anda: ")
	// fmt.Scan(&usia)

	// if usia >= 0 && usia <= 12{
	// 	fmt.Println("Anda masih dalam usia anak-anak")
	// }else if usia >= 13 && usia <= 17 {
	// 	fmt.Println("Anda dalam usia remaja")
	// }else if usia >= 18 && usia <= 59 {
	// 	fmt.Println("Anda dalam usia Dewasa")
	// }else if usia >= 60 {
	// 	fmt.Println("Anda lansia")
	// }else{
	// 	fmt.Println("Silahkan masukkan usia anda")
	// }
	// fmt.Println("Anda Berumur:", usia, "tahun")

	// Cek Kata Palindrom:
	// Buat program yang memeriksa apakah suatu kata yang dimasukkan pengguna adalah palindrom (	 yang sama jika dibaca dari depan maupun dari belakang).

	var kata string
	fmt.Print("Masukkan kata :")
	fmt.Scanln(&kata)

	// Mencetak setiap karakter dalam list
	runes := []rune(kata)

	start, end := 0, len(runes)-1

	for start < end {
		runes[start], runes[end] = runes[end], runes[start]
		start++
		end--
	}

	reversedStr := string(runes)
	if kata == reversedStr {
		fmt.Printf("Kata ini adalah polindrom %v", reversedStr)
		return
	}

	fmt.Printf("Bukan polindrom : %v", reversedStr)

	// for kata := 0; kata < count; kata++ {

	// }
	// if kata == reverseKata {
	// 	fmt.Println("Kata ini termasuk kata polindrom")
	// }else{
	// 	fmt.Println("Kata ini bukan termasuk kata polindrom")
	// }

	// Buat sebuah program dari 1 sampai 100, dimana setiap bilangan yang habis dibagi 3 akan muncul Fizz di sebelahnya, dan jika angkanya habis dibagi 5 tamoilkan Buzz disebelahnya. Bila angka tersebut habis dibagi 3 dan habis dibagi 5, tampilkan kata FizzBuzz pada angka tersebut
	// gunakan fitur perulangan (for loops) dan if else
	// angka := 1
	// for angka <= 100{
	// 		if angka % 3 == 0 && angka % 5 == 0 {
	// 			fmt.Println(angka, "FizzBuzz")
	// 		}else if angka % 3 == 0 {
	// 			fmt.Println(angka,"Fizz")
	// 		}else if angka % 5 == 0 {
	// 			fmt.Println(angka, "Buzz")
	// 		}else{
	// 			fmt.Println(angka)
	// 		}
	// 		angka++
	// 	}

	// Goroutine 
}
