package main

import (
	"fmt"
	"time"
)

func IsiKuesioner() {
	var nama string
	var ID_pengguna, usia int
	var jenisKelamin string
	var jawaban1, jawaban2, jawaban3, jawaban4, jawaban5 int
	var save string
	currentTime := time.Now()
	ID_kuesioner := 1
	fmt.Println("-------------------------------------------")
	fmt.Println("Anda memilih \"Isi Kuesioner\"")
	fmt.Println("-------------------------------------------")
	fmt.Println("Silahkan isi kuesioner berikut :")
	fmt.Println("-------------------------------------------")
	fmt.Println("ID Kuisoner: ", ID_kuesioner)
	fmt.Print("ID Pengguna: ")
	fmt.Scanln(&ID_pengguna)
	fmt.Print("Nama: ")
	fmt.Scanln(&nama)
	fmt.Print("Usia: ")
	fmt.Scanln(&usia)
	fmt.Print("Jenis Kelamin: ")
	fmt.Scanln(&jenisKelamin)
	fmt.Println("Tanggal Pengisian:", currentTime.Format("2006-01-02"))
	fmt.Println("-------------------------------------------")
	fmt.Println("Isilah jawaban kuesioner berikut dengan memilih angka dari 1 hingga 5, sesuai dengan tingkat depresi Anda:")
	fmt.Println("1 = tidak sama sekali")
	fmt.Println("2 = beberapa hari terakhir (1-2 hari)")
	fmt.Println("3 = kadang-kadang (3-5 hari)")
	fmt.Println("4 = terlalu sering (6-10 hari)")
	fmt.Println("5 = hampir setiap hari (11-14 hari)")
	fmt.Println("-------------------------------------------")

	fmt.Println("1. Apakah kamu merasakan kelelahan emosional, sedih, tertekan atau putus asa akhir akhir ini?")
	fmt.Print("Jawaban (1-5): ")
	fmt.Scanln(&jawaban1)

	fmt.Println("\n2. Apakah kamu sulit untuk merasa senang atau tidak menikmati hal-hal yang biasanya disukai?")
	fmt.Print("Jawaban (1-5): ")
	fmt.Scanln(&jawaban2)

	fmt.Println("\n3. Apakah kamu merasa khawatir atau tegang secara terus-menerus?")
	fmt.Print("Jawaban (1-5): ")
	fmt.Scanln(&jawaban3)

	fmt.Println("\n4. Apakah kamu merasakan kesulitan untuk tertidur atau tidur terlalu banyak?")
	fmt.Print("Jawaban (1-5): ")
	fmt.Scanln(&jawaban4)

	fmt.Println("\n5. Apakah kamu merasa cemas ketika berada di situasi sosial atau berbicara dengan orang lain dan merasa tidak fokus?")
	fmt.Print("Jawaban (1-5): ")
	fmt.Scanln(&jawaban5)
	fmt.Println("-------------------------------------------")
	fmt.Println("Apakah Anda ingin menyimpan kuesioner ini? (Y/N):")
	fmt.Scanln(&save)
	if save == "Y" || save == "y" {
		main()
	} else if save == "N" || save == "n" {
		main()
	} else {
		fmt.Println("Input tidak valid. Mohon masukkan Y/N")
		fmt.Print("Apakah Anda ingin menyimpan kuesioner ini? (Y/N): ")
		fmt.Scanln(&save)
	}
	fmt.Println("-------------------------------------------")

}

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
	if pilihan == 1 {
		IsiKuesioner()
	}

}
func main() {
	var pilihan int
	menu(pilihan)
}
