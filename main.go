package main

import (
	"fmt"
	"time"
)

type Kuesioner struct {
	IDKuesioner  int
	IDPengguna   int
	Nama         string
	Usia         int
	JenisKelamin string
	Tanggal      string
	Jawaban      [5]int
}

var dataKuesioner [5]Kuesioner
var jumlahData int

func IsiKuesioner() {
	if jumlahData >= len(dataKuesioner) {
		fmt.Println("Maaf, tidak bisa menambah kuesioner baru.")
		return
	}
	var input Kuesioner
	var save string
	input.Tanggal = time.Now().Format("2006-01-02")
	input.IDKuesioner = jumlahData + 1

	fmt.Println("-------------------------------------------")
	fmt.Println("Anda memilih \"Isi Kuesioner\"")
	fmt.Println("-------------------------------------------")
	fmt.Println("ID Kuesioner: ", input.IDKuesioner)
	fmt.Print("ID Pengguna: ")
	fmt.Scanln(&input.IDPengguna)
	fmt.Print("Nama: ")
	fmt.Scanln(&input.Nama)
	fmt.Print("Usia: ")
	fmt.Scanln(&input.Usia)
	fmt.Print("Jenis Kelamin: ")
	fmt.Scanln(&input.JenisKelamin)
	fmt.Println("Tanggal Pengisian:", input.Tanggal)

	fmt.Println("-------------------------------------------")
	fmt.Println("Isilah jawaban kuesioner berikut (1-5):")
	fmt.Println("1 = tidak sama sekali")
	fmt.Println("2 = beberapa hari terakhir (1-2 hari)")
	fmt.Println("3 = kadang-kadang (3-5 hari)")
	fmt.Println("4 = terlalu sering (6-10 hari)")
	fmt.Println("5 = hampir setiap hari (11-14 hari)")
	fmt.Println("-------------------------------------------")

	pertanyaan := [5]string{
		"1. Apakah kamu merasa sedih, tertekan atau putus asa akhir-akhir ini?",
		"2. Apakah kamu sulit untuk merasa senang atau menikmati hal-hal yang biasanya disukai?",
		"3. Apakah kamu merasa khawatir atau tegang secara terus-menerus?",
		"4. Apakah kamu kesulitan tidur atau tidur terlalu banyak?",
		"5. Apakah kamu cemas di situasi sosial atau merasa tidak fokus?",
	}

	for i := 0; i < 5; i++ {
		fmt.Println(pertanyaan[i])
		fmt.Print("Jawaban (1-5): ")
		fmt.Scanln(&input.Jawaban[i])
	}

	fmt.Print("Apakah Anda ingin menyimpan kuesioner ini? (Y/N): ")
	fmt.Scanln(&save)

	if save == "Y" || save == "y" {
		dataKuesioner[jumlahData] = input
		jumlahData++
		fmt.Println("Kuesioner berhasil disimpan.")
	} else {
		fmt.Println("Kuesioner tidak disimpan.")
	}
}

func EditKuesioner() {
	var id int
	fmt.Print("Masukkan ID kuesioner yang ingin diubah: ")
	fmt.Scanln(&id)
	if id < 1 || id > jumlahData {
		fmt.Println("ID kuesioner tidak valid.")
		return
	}
	var input Kuesioner
	var save string
	input.Tanggal = time.Now().Format("2006-01-02")
	input.IDKuesioner = id
	fmt.Println("-------------------------------------------")
	fmt.Println("Anda memilih \"Edit Kuesioner\"")
	fmt.Println("-------------------------------------------")
	fmt.Println("ID Kuesioner: ", input.IDKuesioner)
	fmt.Print("ID Pengguna: ")
	fmt.Scanln(&input.IDPengguna)
	fmt.Print("Nama: ")
	fmt.Scanln(&input.Nama)
	fmt.Print("Usia: ")
	fmt.Scanln(&input.Usia)
	fmt.Print("Jenis Kelamin: ")
	fmt.Scanln(&input.JenisKelamin)
	fmt.Println("Tanggal Pengisian:", input.Tanggal)
	fmt.Println("-------------------------------------------")
	fmt.Println("Isilah jawaban kuesioner berikut (1-5):")
	fmt.Println("1 = tidak sama sekali")
	fmt.Println("2 = beberapa hari terakhir (1-2 hari)")
	fmt.Println("3 = kadang-kadang (3-5 hari)")
	fmt.Println("4 = terlalu sering (6-10 hari)")
	fmt.Println("5 = hampir setiap hari (11-14 hari)")
	fmt.Println("-------------------------------------------")
	pertanyaan := [5]string{
		"1. Apakah kamu merasa sedih, tertekan atau putus asa akhir-akhir ini?",
		"2. Apakah kamu sulit untuk merasa senang atau menikmati hal-hal yang biasanya disukai?",
		"3. Apakah kamu merasa khawatir atau tegang secara terus-menerus?",
		"4. Apakah kamu kesulitan tidur atau tidur terlalu banyak?",
		"5. Apakah kamu cemas di situasi sosial atau merasa tidak fokus?",
	}
	for i := 0; i < 5; i++ {
		fmt.Println(pertanyaan[i])
		fmt.Print("Jawaban (1-5): ")
		fmt.Scanln(&input.Jawaban[i])
	}
	fmt.Print("Apakah Anda ingin menyimpan perubahan kuesioner ini? (Y/N): ")
	fmt.Scanln(&save)
	if save == "Y" || save == "y" {
		dataKuesioner[id-1] = input
		fmt.Println("Kuesioner berhasil diubah.")
	} else {
		fmt.Println("Perubahan kuesioner tidak disimpan.")
	}
}
func HapusKuesioner() {
	var id int
	fmt.Print("Masukkan ID kuesioner yang ingin dihapus: ")
	fmt.Scanln(&id)
	if id < 1 || id > jumlahData {
		fmt.Println("ID kuesioner tidak valid.")
		return
	}
	for i := id - 1; i < jumlahData-1; i++ {
		dataKuesioner[i] = dataKuesioner[i+1]
		dataKuesioner[i].IDKuesioner = i + 1
	}
	jumlahData--
	fmt.Println("Kuesioner berhasil dihapus.")
}

func menu() {
	var pilihan int
	for {
		fmt.Println("---------------------------------------------------")
		fmt.Println("        Aplikasi Manajemen Kesehatan Mental        ")
		fmt.Println("---------------------------------------------------")
		fmt.Println("Silakan pilih opsi di bawah ini:")
		fmt.Println("1. Isi Kuesioner")
		fmt.Println("2. Edit Kuesioner")
		fmt.Println("3. Hapus Kuesioner (belum tersedia)")
		fmt.Println("4. Cari Kuesioner (belum tersedia)")
		fmt.Println("5. Cek Laporan (belum tersedia)")
		fmt.Println("6. Keluar")
		fmt.Println("---------------------------------------------------")
		fmt.Print("Pilih opsi (1-6): ")
		fmt.Scanln(&pilihan)

		if pilihan == 1 {
			IsiKuesioner()
		} else if pilihan == 2 {
			EditKuesioner()
		} else if pilihan == 3 {
			HapusKuesioner()
		} else if pilihan == 6 {
			fmt.Println("Terima kasih telah menggunakan aplikasi ini.")
			return
		} else {
			fmt.Println("input tidak valid.")
		}
	}
}

func main() {
	menu()
}
