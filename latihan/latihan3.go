// latihan 3
/* Buatlah sebuah fungsi yang menerima dua paramter, nama dan umur. Fungsi ini harus mencetak pesan
"Halo [nama]! kamu masih muda!" jika umur kurang dari 25 tahun, atau "Halo [nama]! semoga hari kamu menyenangkan!"
jika umur 25 tahun atau lebih. */

package main

import "fmt"

func dataDiri(nama string, umur int)(pesan1 string, pesan2 string){
	pesan1 = fmt.Sprintf("Halo %s! kamu masih muda", nama)
	pesan2 = fmt.Sprintf("Halo %s! semoga hari kamu menyenangkan", nama)
	return pesan1, pesan2
}

func main(){
	var nama string
	var umur int
	fmt.Print("Masukkan nama : ")
	fmt.Scan(&nama)
	fmt.Print("Masukkan umur : ")
	fmt.Scan(&umur)
	if umur < 25 {
		pesan1, _ := dataDiri(nama,umur)
		fmt.Println(pesan1)
	}else if umur >= 25 {
		_, pesan2 := dataDiri(nama,umur)
		fmt.Println(pesan2)
	}
}