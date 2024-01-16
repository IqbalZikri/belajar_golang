// function juga bisa multiple value
// untuk melakukan hal ini, kita harus memberikan atau menulis semua tipe data pada return pada valuenya di function

package main

import "fmt"

func getFullName()(string, string){
	return "Iqbal" , "Zikri"
}

func campuran()(string, string, bool, int){
	return "Tubagus", "Muhammad", true, 99
}

func mobil()(string, string, bool){
	return "Toyota", "Honda", false
}

func main(){
	namaPertama, namaTerakhir := getFullName()
	fmt.Println(namaPertama, namaTerakhir)
	// output Iqbal Zikri

	namaDepan, namaTengah, benar, angka := campuran()
	fmt.Println(namaDepan, namaTengah, benar, angka)
	// output Tubagus, Muhammad, true, 99

	fmt.Println(campuran())

	// Menghiraukan return value
	// Multiple return value wajib ditangkap semua valuenya
	// Jika kita ingin menghiraukan return value tersebut, kita bisa menggunakan tanda_ (garis bawah)
	mobilPertama, mobilKedua, _ := mobil()
	fmt.Println(mobilPertama, mobilKedua)

}