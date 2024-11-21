package main

import "fmt"

func main() {
	// Deklarasi chop sebagai slice kosong bertipe float64
	var chop []float64
	fmt.Println("Chop (kosong):", chop)

	// Deklarasi slice s101 dan inisialisai dengan elemen tertentu
	var s101 = []int{11, 2, 3, 5, 7, 13}
	fmt.Println("Isi slice s101:", s101)

	// Menggunakan fungsi built-in 'make' untuk membuat slice dengan kapasitas 5
	reazo := make([]int, 3, 5) // Slice dengan panjang 3 dan kapasitas 5
	fmt.Println("Isi slice reazo:", reazo)

	//len cap
	fmt.Println("Panjang slice reazo:", len(reazo), "Kapasitas slice reazo:", cap(reazo))

	// Menggunakan fungsi 'append' untuk menambahkan elemen ke dalam slice
	s101 = append(s101, 19, 23)
	fmt.Println("Isi slice s101 setelah append:", s101)

	// Membuat slice baru dengan mengambil 3 elemen pertama dari slice s101
	s104 := s101[:3] // Slice s101 dari indeks 0 sampai 3 (elemen ke-3 tidak termasuk)
	fmt.Println("Slice s104 (3 elemen pertama):", s104)

	// Mengambil beberapa elemen terakhir dari s101, dimulai dari indeks 5
	s105 := s101[5:]
	fmt.Println("Slice s105 (elemen terakhir setelah indeks 5):", s105)

	// Menyalin seluruh slice s105 ke slice baru s106
	s106 := s105[:]
	fmt.Println("Slice s106 (menyalin seluruh s105):", s106)

	// Menyalin elemen dari indeks 3 sampai 5 (tidak termsauk indeks 5) dari s106
	s107 := s106[3:5]
	fmt.Println("Slice s107 (elemen 3 sampai 5 dari s106):", s107)
}
