// Latihan 6
// Buat program yang meminta pengguna memasukkan sebuah bilangan. Program kemudian menentukan apakah bilangan tersebut adalah bilangan prima atau bukan.

package main

import "fmt"

func main(){
	var angka int
	fmt.Print("Masukkan angka : ")
	fmt.Scan(&angka)
	if angka % 1 == 0 && angka % 2 == 0 && angka % 4 == 0 && angka % 6 == 0{
		fmt.Println("Angka ini termasuk angka prima")
	}else{
		fmt.Println("Tidak termasuk angka prima	")
	}
}