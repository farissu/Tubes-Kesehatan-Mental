package main

import (
	"fmt"
)

func menu(pilihan int) {
	fmt.Println("---------------------------------------------------")
	fmt.Println("        Aplikasi Manajemen Kesehatan Mental        ")
	fmt.Println("---------------------------------------------------")
	fmt.Println("Silakan pilih opsi di bawah ini:")
	fmt.Println("1. Isi Kuesioner")
	fmt.Println("2. Edit Kuesioner")
	fmt.Println("3. Hapus Kuesioner")
	fmt.Println("4. Cari Kuesioner")
	fmt.Println("5. Cek Laporan")
	fmt.Println("6. Keluar")
	fmt.Println("---------------------------------------------------")
	fmt.Print("Pilih opsi (1-6): ")
	fmt.Scanln(&pilihan)

}
func main() {
	var pilihan int
	menu(pilihan)
}
