// Operasi Boolean
// Digunakan untuk mengetahui dari sebuah nilai (true atau false)
// && adalah and/dan , || adalah or/atau , ! adalah kebalikan
// Untuk menggunakan && kedua nilai harus true jika ingin hasilnya true selain itu false
// Untuk menggunakan || salah satu nilai cukup true jika ingin hasilnya true, selain itu false
// ! true maka akan menjadi false
// ! false maka akan menjadi true

package main

import "fmt"

func main(){
	var nilaiAkhir int = 90
	var nilaiKehadiran int = 80

	var lulusNilaiAkhir = nilaiAkhir >= 80
	var lulusNilaiKehadiran = nilaiKehadiran >= 80

	var lulus bool = lulusNilaiAkhir && lulusNilaiKehadiran
	fmt.Println(lulus)

	var absensi = 70
	var tugas = 50

	var nilaiAbsensi = absensi >= 70
	var nilaiTugas = tugas >= 80

	var lanjut = nilaiAbsensi && nilaiTugas
	fmt.Println(lanjut)

}