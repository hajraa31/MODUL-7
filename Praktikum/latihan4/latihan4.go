package main

import "fmt"

func main() {
	// Deklarasi variabel 'det' sebagai map dengan kunci string dan nilai
	var det map[string]int
	// Pencetakan map 'det' yang nilainya masih kosong (nil) karena belum di inisialisai
	fmt.Println("Map det(kosong):", det)

	// Deklarasi dan inisialisasi map 'det1' dengan kunci dan nilai bertipe string
	var det1 = map[string]string{
		"john":  "developer",
		"anne":  "designer",
		"sarah": "manager",
	}

	// Menampilkan isi map 'det1'
	fmt.Println("Isi map det1:", det1)

	// Deklarasi dan prealokasi map 'det2' dengan kapasitas 10
	var det2 map[float64]int = make(map[float64]int, 10)

	// Menampilkan map 'det2' setelah diinisialisasi
	fmt.Println("Isi map det2 setelah inisialisai:", det2)

	// Mengakses nilai dengan kunci "john" dari map 'det1'
	fmt.Println("Nilai untuk kunci 'john' di det1:", det1["john"])

	//Mengganti nilai yang tersimpan pada kunci "anne" dan menambahkan entri baru
	det1["anne"] = "lovely"    // Mengubah nilai untuk kunci "anne"
	det1["boy"] = " runaround" // Menambahkan entri baru dengan kunci "boy"
	fmt.Println("Isi map det1 setelah perubahan:", det1)

	// Menghapus entri dengan kunci "john" dari map 'det1'
	delete(det1, "john")
	fmt.Println("Isi map det1 setelah penghapusan kunci 'john':", det1)
}
