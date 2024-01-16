// Operasi Perbandingan
// Operasi yang digunakan saat ingin membandingkan 2 buah data
// Operasi perbandingan menghasilkan nilai true dan false (boolean)
// > = lebih dari , < = Kurang dari , >= adalah lebih dari sama dengan , <= adalah kurang dari sama dengan , == adalah Sama dengan , != adalah tidak sama dengan

package main

import "fmt"

func main(){
	var name1 = "tb"
	var name2 = "tb"
	var result bool = name1 == name2
	fmt.Println(result)

	var abjad1 = "a"
	var abjad2 = "b"
	var hasil = abjad1 > abjad2 // bisa juga digunakan untuk membandingkan abjad
	fmt.Println(hasil)
}