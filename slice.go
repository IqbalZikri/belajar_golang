// Tipe Data Slice
// Tipe Data Slice adalah potongan dari array
// SLice mirip dengan array, tetapi ukuran data Slice bisa diubah
// Slice dan Array saling terkoneksi, dimana Slice adalah data yang mengakses sebagian atau seluruh data di Array
// Tipe data Slice memiliki 3 data, yaitu pointer, length dan capacity
// Pointer adalah penunjuk data pertama di array para slice
// Length adalah Panjang dari slice
// capacity adalah kapasitas dari SLice, dimana length tidak boleh lebih dari capacity
// array[low:high] = mengambil data / membuat slice dari array dari data low hingga data high
// array[low:] = mengambil data / membuat slice dari  array dari data low hingga akhir array
// array[:high] = mengambil data / membuat slice dari array dari awal data array hingga akhir array
package main

import "fmt"

func main(){
	names := [...]string{"tb","Muhammad","Iqbal","Zikri","Budi","Joko"}
	slice1 := names[4:6]
	slice2 := names[2:]
	slice3 := names[:5]
	slice4 := names[:]

	fmt.Println(slice1[0])
	fmt.Println(slice1[1])
	fmt.Println(slice1)

	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)

	// Function Slice
	// len(slice) = untuk mendapatkan panjang slice
	// cap(slice) = untuk mendapatkan kapasitas slice
	// append(slice,data) = Membuat slice baru dengan menambah data ke posisi terakhir slice, jika kapasitas sudah penuh, maka akan membuat array baru
	// make([]TypeData,length,Capacity) membuat slice baru
	// copy(destination,source) menyalin slice dari source ke destination
	fmt.Println(len(slice1))

	days := [...]string{"Senin","Selasa","Rabu","Kamis","Jum'at","Sabtu","Minggu"}
	sliceDays := days[5:]
	sliceDays[0] = "Sabtu Baru"
	sliceDays[1] = "Minggu Baru"
	fmt.Println(days)
	fmt.Println(sliceDays)

	daySlice := append(sliceDays,"Libur Baru") // Menambahkan data "Libur Baru" dalam slice sliceDays
	fmt.Println(daySlice)
	fmt.Println(days) // tidak akan ditambahkan data "Libur Baru" Karena array tidak bisa ditambahkan, dengan kata lain daySLice membuat array baru

	newSlice := make([]string, 2, 5) // 2 = panjang slice, 5 = kapasitas slice
	newSlice[0] = "Iqbal"
	newSlice[1] = "TB"
	// newSlice[3] = "TB Iqbal" error karena membuat slice dengan panjang 2, jika ingin menambahkan data dalam array maka gunakan append

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
	
	newSliceName := append(newSlice, "Bambang")
	fmt.Println(newSliceName)
	fmt.Println(len(newSliceName))
	fmt.Println(cap(newSliceName))

	newSliceName[0] = "Bimbing"
	fmt.Println(newSliceName)
	fmt.Println(newSlice)

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(toSlice)

	// Note (Hati hati saat membuat array) karena bisa jadi jika salah kita bukan membuat array tapi slice, dan juga kebalikannya
	iniArray := [...]int{1,2,3,4,5}
	iniSlice := []int{1,2,3,4,5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

	// rata-rata di golang lebih banyak menggunakan slice daripada array
}