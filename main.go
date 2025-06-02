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

	fmt.Println("---------------------------------------------------")
	fmt.Println("Anda memilih \"Isi Kuesioner\"")
	fmt.Println("---------------------------------------------------")
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

	fmt.Println("---------------------------------------------------")
	fmt.Println("Isilah jawaban kuesioner berikut (1-5):")
	fmt.Println("1 = tidak sama sekali")
	fmt.Println("2 = beberapa hari terakhir (1-2 hari)")
	fmt.Println("3 = kadang-kadang (3-5 hari)")
	fmt.Println("4 = terlalu sering (6-10 hari)")
	fmt.Println("5 = hampir setiap hari (11-14 hari)")
	fmt.Println("---------------------------------------------------")

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
	fmt.Println("---------------------------------------------------")
	fmt.Println("Anda memilih \"Edit Kuesioner\"")
	fmt.Println("---------------------------------------------------")
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
	fmt.Println("---------------------------------------------------")
	fmt.Println("Isilah jawaban kuesioner berikut (1-5):")
	fmt.Println("1 = tidak sama sekali")
	fmt.Println("2 = beberapa hari terakhir (1-2 hari)")
	fmt.Println("3 = kadang-kadang (3-5 hari)")
	fmt.Println("4 = terlalu sering (6-10 hari)")
	fmt.Println("5 = hampir setiap hari (11-14 hari)")
	fmt.Println("---------------------------------------------------")
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

func CariKuesioner() {
	var id int
	fmt.Print("Masukkan ID kuesioner yang ingin dicari: ")
	fmt.Scanln(&id)
	if id < 1 || id > jumlahData {
		fmt.Println("ID kuesioner tidak valid.")
		return
	}
	fmt.Println("---------------------------------------------------")
	fmt.Println("Anda memilih \"Cari Kuesioner\"")
	fmt.Println("---------------------------------------------------")
	fmt.Println("ID Kuesioner: ", dataKuesioner[id-1].IDKuesioner)
	fmt.Println("ID Pengguna: ", dataKuesioner[id-1].IDPengguna)
	fmt.Println("Nama: ", dataKuesioner[id-1].Nama)
	fmt.Println("Usia: ", dataKuesioner[id-1].Usia)
	fmt.Println("Jenis Kelamin: ", dataKuesioner[id-1].JenisKelamin)
	fmt.Println("Tanggal Pengisian: ", dataKuesioner[id-1].Tanggal)
	fmt.Println("---------------------------------------------------")
	fmt.Println("Jawaban kuesioner:")
	fmt.Println("1. Apakah kamu merasa sedih, tertekan atau putus asa akhir-akhir ini?")
	fmt.Println("Jawaban: ", dataKuesioner[id-1].Jawaban[0])
	fmt.Println("2. Apakah kamu sulit untuk merasa senang atau menikmati hal-hal yang biasanya disukai?")
	fmt.Println("Jawaban: ", dataKuesioner[id-1].Jawaban[1])
	fmt.Println("3. Apakah kamu merasa khawatir atau tegang secara terus-menerus?")
	fmt.Println("Jawaban: ", dataKuesioner[id-1].Jawaban[2])
	fmt.Println("4. Apakah kamu kesulitan tidur atau tidur terlalu banyak?")
	fmt.Println("Jawaban: ", dataKuesioner[id-1].Jawaban[3])
	fmt.Println("5. Apakah kamu cemas di situasi sosial atau merasa tidak fokus?")
	fmt.Println("Jawaban: ", dataKuesioner[id-1].Jawaban[4])
	fmt.Println("---------------------------------------------------")
}
func CekLaporan() {
	fmt.Println("---------------------------------------------------")
	fmt.Println("Anda memilih \"Cek Laporan\"")
	fmt.Println("---------------------------------------------------")
	fmt.Println("Jumlah data kuesioner: ", jumlahData)
	fmt.Println("---------------------------------------------------")
	fmt.Println("Data kuesioner:")
	fmt.Println("---------------------------------------------------")
	for i := 0; i < jumlahData; i++ {
		fmt.Println(i+1, ".", dataKuesioner[i].Tanggal, " | ", dataKuesioner[i].Jawaban[0]+dataKuesioner[i].Jawaban[1]+dataKuesioner[i].Jawaban[2]+dataKuesioner[i].Jawaban[3]+dataKuesioner[i].Jawaban[4])
	}
	fmt.Println("---------------------------------------------------")
	rata, hasil, deskripsi := rata_rata()
	fmt.Printf("Rata-rata skor anda : %.1f\n", rata)
	fmt.Println("Hasil: ", hasil)
	fmt.Println("Deskripsi: ", deskripsi)

}

func rata_rata() (float64, string, string) {
	var total int
	for i := 0; i < jumlahData; i++ {
		total += dataKuesioner[i].Jawaban[0] + dataKuesioner[i].Jawaban[1] + dataKuesioner[i].Jawaban[2] + dataKuesioner[i].Jawaban[3] + dataKuesioner[i].Jawaban[4]
	}

	var hasil, deskripsi string
	total_rata_rata := float64(total) / float64(jumlahData)

	if total_rata_rata <= 5 {
		hasil = "Kesehatan mentalmu sehat (Normal)."
		deskripsi = "Kondisi kesehatan kamu stabil dan pertahankan kebiasaan baik serta kamu bisa melanjutkan aktivitas yang mendukung kesehatan mental kamu"
	} else if total_rata_rata <= 10 {
		hasil = "Kesehatan mentalmu sedikit stres."
		deskripsi = "Kamu memiliki beberapa gejala karena kelelahan yang diakibatkan tekanan ke diri kamu. Kamu harus istirahat dan melakukan aktivitas yang menyenangkan (AIA)"
	} else if total_rata_rata <= 15 {
		hasil = "Kesehatan mentalmu stres sedang."
		deskripsi = "Kamu memiliki tanda tanda stress yang cukup sedang. Kamu bisa menggunakan cara seperti menggunakan teknik coping olahraga, memasak, journaling, melakukan hobi, dan meditasi (Siloam, 2024)"
	} else if total_rata_rata <= 20 {
		hasil = "Kecemasan atau Beban Emosional Tinggi"
		deskripsi = "Kamu memiliki gejala kecemasan dan bisa mengarah ke gangguan psikologis jika tidak ditangani. Kamu bisa menerapkan Metode 5-4-3-2-1 dengan teknik mengurangi distraksi yang dilakukan dengan fokus pada fungsi indra tubuh dan efektif untuk mengatasi cemas berlebihan dan kamu bisa bercerita ke orang terdekat tentang rasa cemas yang dirasakan (Enesis, 2024)"
	} else {
		hasil = "Gejala Depresi Berat"
		deskripsi = "Kamu mengalami gejala depresi berat yang mengarah pada kondisi serius. Kamu harus segera menghubungi dokter untuk mengatasi depresi berat yang kamu alami (Halodoc, 2024)"
	}

	fmt.Println(hasil)
	fmt.Println(deskripsi)

	return total_rata_rata, hasil, deskripsi
}

func Keluar() {
	fmt.Println("---------------------------------------------------")
	fmt.Println("Anda memilih \"Keluar\"")
	fmt.Println("---------------------------------------------------")
	fmt.Println("Terima kasih telah menggunakan Aplikasi Manajemen Kesehatan Mental.")
	fmt.Println()
	fmt.Println("Jaga kesehatan mentalmu, karena kamu berharga.")
	fmt.Println("Sampai jumpa!")
	fmt.Println("---------------------------------------------------")
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
		fmt.Println("3. Hapus Kuesioner")
		fmt.Println("4. Cari Kuesioner")
		fmt.Println("5. Cek Laporan")
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
		} else if pilihan == 4 {
			CariKuesioner()
		} else if pilihan == 5 {
			CekLaporan()
		} else if pilihan == 6 {
			Keluar()
		} else {
			fmt.Println("input tidak valid.")
		}
	}
}

func main() {
	menu()
}
