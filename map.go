// Tipe Data Map
// Pada Array atau slice, untuk mengakses data, kita gunakan index number yang dimulai dari 0
// Map adalah tipe data lain yang berisikan kumpulan data yang sama, namun kita bisa menentukan jenis tipe data index yang akan kita gunakan
// Sederhananya, Map adalah tipe data kumpulan key-value (kata kunci-nilai), dimana kata kuncinya bersifat unik, tidak boleh sama
// Berbeda dengan array dan Slice, kita bisa memasukkan data di dalam map sebanyak-banyaknya, asalkan kata kuncinya berbeda, jika kita gunakan kata kunci yang sama, otomatis data sebelumnya akan tergantikan dengan yang baru

package main

import "fmt"

func main(){
	person := map[string]string{
		"name" : "Iqbal",
		"email" : "tbmiqbalz@gmail.com",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["email"])
	fmt.Println(len(person)) // untuk mendapatkan jumlah data pada map person

	// cara menambhkan data ke map
	fakultas := map[string]string{

	}

	fakultas["teknik"] = "Informatika, Elektro"
	fakultas["Ekonomi dan Bisnis"] = "Ekonomi"
	fakultas["teknik"] = "industri" // data pada key teknik akan berubah valuenya menjadi industri
	fmt.Println(fakultas)
	delete(fakultas,"teknik") // menghapus data key teknik pada map fakultas
	fmt.Println(fakultas)

	// membuat map baru
	kendaraan := make(map[string]string)
	kendaraan["motor"] = "KLX"
	kendaraan["mobil"] = "Toyota"
	kendaraan["Perahu"] = "Apa aja lah"

	fmt.Println(kendaraan)
	delete(kendaraan,"Perahu")
	fmt.Println(kendaraan)	

	// len(map) = mendapatkan jumlah data pada map
	// Delete(Map, Key) = menghapus data pada key berdasarkan kata kuncinya
	// make(map[TypeData]TypeData) = salah satu cara untuk membuat map baru
	// namaMap := map[TypeData]TypeData = membuat map
	// map[key] = mengambil data / value pada map berdasarkan key / kata kuncinya
}