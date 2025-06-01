package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
	SISTEM PENERIMAAN MAHASISWA BUKIT DURI

# Sistem bedasarkan seleksi nilai UTBK apabila nilai lebih dari 600 maka diterima

======== SPESIFIKASI UMUM
✅ Program harus dibuat secara modular dengan menggunakan subprogram.
✅ Struktur subprogram ditentukan oleh masing-masing kelompok. Setiap subprogram harus dilengkapi dengan parameter dan spesifikasinya.
✅ Program harus mengimplementasikan Array dan Tipe Bentukan. Array statis harus digunakan, bukan array dinamis atau slice.
✅ Program harus mengimplementasikan algoritma Sequential dan Binary Search untuk pencarian data, pengubahan (edit), atau penghapusan data tertentu.
✅ Program harus mengimplementasikan algoritma Selection Sort dan Insertion Sort untuk mengurutkan data dalam kategori yang berbeda saat menampilkan hasil. Setiap kategori harus bisa diurutkan dalam urutan naik (ascending) maupun turun (descending).
✅ Penggunaan statement break (kecuali dalam repeat-until) dan continue tidak diperbolehkan.
✅ Variabel global hanya diperbolehkan untuk menyimpan array utama yang akan diproses.
*/
type dataPendaftar struct {
	nisn              int
	nama              string
	tempatLahir       string
	tanggalLahir      string
	jenisKelamin      string
	agama             string
	email             string
	jurusan           string
	asalSekolah       string
	tahunLulus        int
	jurusanYangDituju string
	nilaiUTBK         float64
	status            string
}
type dataMahasigma struct {
	nim     int
	nama    string
	jurusan string
}

const max = 10

type dataSiswa = [max]dataPendaftar

func main() {
	var calonMahasiswa dataSiswa
	// var mahasigma [10]dataMahasigma

	//Data dummy calon mahasiswa
	//Nilai UTBK Min 0
	//Nilai UTBK Max 1000

	calonMahasiswa[0] = dataPendaftar{nisn: 1234567890, nama: "Andi Pratama", tempatLahir: "Surabaya", tanggalLahir: "2000-01-01", jenisKelamin: "Laki-laki", agama: "Islam", email: "andi@email.com", jurusan: "IPA", asalSekolah: "SMA 1 Surabaya", tahunLulus: 2018, jurusanYangDituju: "Teknik Informatika", nilaiUTBK: 0}
	calonMahasiswa[1] = dataPendaftar{nisn: 1234567891, nama: "Siti Nurjanah", tempatLahir: "Malang", tanggalLahir: "2001-02-02", jenisKelamin: "Perempuan", agama: "Islam", email: "siti@email.com", jurusan: "IPS", asalSekolah: "SMA 2 Malang", tahunLulus: 2019, jurusanYangDituju: "Manajemen", nilaiUTBK: 1}
	calonMahasiswa[2] = dataPendaftar{nisn: 1234567892, nama: "Budi Santoso", tempatLahir: "Surabaya", tanggalLahir: "2000-03-03", jenisKelamin: "Laki-laki", agama: "Kristen", email: "budi@email.com", jurusan: "IPA", asalSekolah: "SMA 3 Surabaya", tahunLulus: 2019, jurusanYangDituju: "Teknik Mesin", nilaiUTBK: 710}
	calonMahasiswa[3] = dataPendaftar{nisn: 1234567893, nama: "Dewi Lestari", tempatLahir: "Bandung", tanggalLahir: "1999-04-04", jenisKelamin: "Perempuan", agama: "Hindu", email: "dewi@email.com", jurusan: "IPA", asalSekolah: "SMA 4 Bandung", tahunLulus: 2018, jurusanYangDituju: "Teknik Elektro", nilaiUTBK: 720}
	calonMahasiswa[4] = dataPendaftar{nisn: 1234567894, nama: "Fajar Ramadhan", tempatLahir: "Yogyakarta", tanggalLahir: "2000-05-05", jenisKelamin: "Laki-laki", agama: "Islam", email: "fajar@email.com", jurusan: "IPS", asalSekolah: "SMA 5 Yogyakarta", tahunLulus: 2019, jurusanYangDituju: "Ekonomi", nilaiUTBK: 680}
	calonMahasiswa[5] = dataPendaftar{nisn: 1234567895, nama: "Hanafi Rizky", tempatLahir: "Jakarta", tanggalLahir: "1999-06-06", jenisKelamin: "Laki-laki", agama: "Islam", email: "hanafi@email.com", jurusan: "IPA", asalSekolah: "SMA 6 Jakarta", tahunLulus: 2018, jurusanYangDituju: "Fisika", nilaiUTBK: 1000}
	calonMahasiswa[6] = dataPendaftar{nisn: 1234567896, nama: "Indah Purnama", tempatLahir: "Semarang", tanggalLahir: "2001-07-07", jenisKelamin: "Perempuan", agama: "Islam", email: "indah@email.com", jurusan: "IPS", asalSekolah: "SMA 7 Semarang", tahunLulus: 2019, jurusanYangDituju: "Psikologi", nilaiUTBK: 630}
	calonMahasiswa[7] = dataPendaftar{nisn: 1234567897, nama: "Joko Susanto", tempatLahir: "Surakarta", tanggalLahir: "2000-08-08", jenisKelamin: "Laki-laki", agama: "Kristen", email: "joko@email.com", jurusan: "IPA", asalSekolah: "SMA 8 Surakarta", tahunLulus: 2019, jurusanYangDituju: "Teknik Industri", nilaiUTBK: 715}
	calonMahasiswa[8] = dataPendaftar{nisn: 1234567898, nama: "Kartika Sari", tempatLahir: "Solo", tanggalLahir: "1999-09-09", jenisKelamin: "Perempuan", agama: "Islam", email: "kartika@email.com", jurusan: "IPA", asalSekolah: "SMA 9 Solo", tahunLulus: 2018, jurusanYangDituju: "Teknik Kimia", nilaiUTBK: 705}
	calonMahasiswa[9] = dataPendaftar{nisn: 1234567899, nama: "Lutfi Fauzi", tempatLahir: "Surabaya", tanggalLahir: "2001-10-10", jenisKelamin: "Laki-laki", agama: "Islam", email: "lutfi@email.com", jurusan: "IPS", asalSekolah: "SMA 10 Surabaya", tahunLulus: 2019, jurusanYangDituju: "Ilmu Komunikasi", nilaiUTBK: 690}

	// calonMahasiswa[10] = dataPendaftar{nisn: 1234567900, nama: "Rina Marlina", tempatLahir: "Medan", tanggalLahir: "2001-11-11", jenisKelamin: "Perempuan", agama: "Kristen", email: "rina@email.com", jurusan: "IPA", asalSekolah: "SMA 11 Medan", tahunLulus: 2019, jurusanYangDituju: "Biologi", nilaiUTBK: 680}
	// calonMahasiswa[11] = dataPendaftar{nisn: 1234567901, nama: "Toni Suryadi", tempatLahir: "Jakarta", tanggalLahir: "1999-12-12", jenisKelamin: "Laki-laki", agama: "Islam", email: "toni@email.com", jurusan: "IPS", asalSekolah: "SMA 12 Jakarta", tahunLulus: 2018, jurusanYangDituju: "Sosiologi", nilaiUTBK: 675}
	// calonMahasiswa[12] = dataPendaftar{nisn: 1234567902, nama: "Siti Aisyah", tempatLahir: "Surabaya", tanggalLahir: "2000-01-13", jenisKelamin: "Perempuan", agama: "Islam", email: "aisyah@email.com", jurusan: "IPA", asalSekolah: "SMA 13 Surabaya", tahunLulus: 2019, jurusanYangDituju: "Teknik Lingkungan", nilaiUTBK: 690}
	// calonMahasiswa[13] = dataPendaftar{nisn: 1234567903, nama: "Wawan Prasetyo", tempatLahir: "Semarang", tanggalLahir: "1999-02-14", jenisKelamin: "Laki-laki", agama: "Buddha", email: "wawan@email.com", jurusan: "IPA", asalSekolah: "SMA 14 Semarang", tahunLulus: 2018, jurusanYangDituju: "Arsitektur", nilaiUTBK: 725}
	// calonMahasiswa[14] = dataPendaftar{nisn: 1234567904, nama: "Eva Nabila", tempatLahir: "Bandung", tanggalLahir: "2001-03-15", jenisKelamin: "Perempuan", agama: "Islam", email: "eva@email.com", jurusan: "IPS", asalSekolah: "SMA 15 Bandung", tahunLulus: 2019, jurusanYangDituju: "Hukum", nilaiUTBK: 660}
	// calonMahasiswa[15] = dataPendaftar{nisn: 1234567905, nama: "Aditya Putra", tempatLahir: "Yogyakarta", tanggalLahir: "2000-04-16", jenisKelamin: "Laki-laki", agama: "Islam", email: "aditya@email.com", jurusan: "IPA", asalSekolah: "SMA 16 Yogyakarta", tahunLulus: 2018, jurusanYangDituju: "Informatika", nilaiUTBK: 700}
	// calonMahasiswa[16] = dataPendaftar{nisn: 1234567906, nama: "Nurul Azzahra", tempatLahir: "Surakarta", tanggalLahir: "1999-05-17", jenisKelamin: "Perempuan", agama: "Islam", email: "nurul@email.com", jurusan: "IPS", asalSekolah: "SMA 17 Surakarta", tahunLulus: 2019, jurusanYangDituju: "Akuntansi", nilaiUTBK: 680}
	// calonMahasiswa[17] = dataPendaftar{nisn: 1234567907, nama: "Rudi Kurniawan", tempatLahir: "Surabaya", tanggalLahir: "2001-06-18", jenisKelamin: "Laki-laki", agama: "Kristen", email: "rudi@email.com", jurusan: "IPA", asalSekolah: "SMA 18 Surabaya", tahunLulus: 2019, jurusanYangDituju: "Bioteknologi", nilaiUTBK: 710}
	// calonMahasiswa[18] = dataPendaftar{nisn: 1234567908, nama: "Hendri Susilo", tempatLahir: "Jakarta", tanggalLahir: "2000-07-19", jenisKelamin: "Laki-laki", agama: "Islam", email: "hendri@email.com", jurusan: "IPA", asalSekolah: "SMA 19 Jakarta", tahunLulus: 2018, jurusanYangDituju: "Teknik Komputer", nilaiUTBK: 690}
	// calonMahasiswa[19] = dataPendaftar{nisn: 1234567909, nama: "Siti Mulyani", tempatLahir: "Semarang", tanggalLahir: "2001-08-20", jenisKelamin: "Perempuan", agama: "Hindu", email: "siti@email.com", jurusan: "IPS", asalSekolah: "SMA 20 Semarang", tahunLulus: 2019, jurusanYangDituju: "Manajemen", nilaiUTBK: 660}

	status(&calonMahasiswa) //Hapus jika menggunakan input manual

	fmt.Println("🏫 SELAMAT DATANG 🏫")
	fmt.Println("Penerimaan Mahasiswa Bukit Duri")
selesai:
	for {
		input := mainMenu()
		switch input {
		case 1:
			fmt.Println("------- Tampilkan semua data -------")
			tampilkanIsi(calonMahasiswa)

		case 2:
			fmt.Println("------- Urutkan Data -------")
			for {
				var metode int
				fmt.Println("Pilh metode pengurutan:")
				fmt.Println("1. Selection Sort (Ascending)")
				fmt.Println("2. Selection Sort (Descending)")
				fmt.Println("3. Insertion Sort (Ascending)")
				fmt.Println("4. Insertion Sort (Descending)")
				fmt.Println("5. Kembali")
				fmt.Print("Pilih metode pengurutan: ")
				fmt.Scan(&metode)

				switch metode {
				case 1:
					fmt.Println("Pengurutan dengan Selection Sort (Ascending)")
					selectionSortAscending(&calonMahasiswa)
				case 2:
					fmt.Println("Pengurutan dengan Selection Sort (Descending)")
					selectionSortDescending(&calonMahasiswa)

				case 3:
					fmt.Println("Pengurutan dengan Insertion Sort (Ascending)")
					insertionSort_Asc(&calonMahasiswa)
				case 4:
					fmt.Println("Pengurutan dengan Insertion Sort (Descending)")
					insertionSort_Desc(&calonMahasiswa)
				default:
					break
				}
				if metode == 5 {
					break
				}
			}

		case 3:
			fmt.Println("------- Tambah data -------")
			var banyakData int
			fmt.Print("Masukkan banyak data yang akan ditambahkan: ")
			fmt.Scan(&banyakData)
			tambahData(&calonMahasiswa, banyakData)

			for i := 0; i < banyakData; i++ {
				fmt.Println(i+1, ". NISN: ", calonMahasiswa[i].nisn)
				fmt.Println("Nama: ", calonMahasiswa[i].nama)
				fmt.Println("Tempat Lahir: ", calonMahasiswa[i].tempatLahir)
				fmt.Println("Tanggal Lahir: ", calonMahasiswa[i].tanggalLahir)
				fmt.Println("Jenis Kelamin: ", calonMahasiswa[i].jenisKelamin)
				fmt.Println("Agama: ", calonMahasiswa[i].agama)
				fmt.Println("Email: ", calonMahasiswa[i].email)
				fmt.Println("Jurusan: ", calonMahasiswa[i].jurusan)
				fmt.Println("Asal Sekolah: ", calonMahasiswa[i].asalSekolah)
				fmt.Println("Tahun Lulus: ", calonMahasiswa[i].tahunLulus)
				fmt.Println("Jurusan yang Dituju: ", calonMahasiswa[i].jurusanYangDituju)
				fmt.Println("Nilai UTBK: ", calonMahasiswa[i].nilaiUTBK)
				fmt.Println("Status: ", calonMahasiswa[i].status)
				fmt.Println()
			}
		case 4:
			//Melakukan pencarian data dengan sequential search dan binary search
			fmt.Println("------- Cari Data -------")
			var cariNISN int
		menuCari:
			for {
				var metode int
				fmt.Println("Pilih metode pencarian:")
				fmt.Println("1. Sequential Search")
				fmt.Println("2. Binary Search")
				fmt.Println("3. Kembali")
				fmt.Print("Pilih metode pengurutan: ")
				fmt.Scan(&metode)
				switch metode {
				case 1:
					fmt.Println("📊 Pencarian dengan Sequential Search 📊")
					fmt.Print("Masukkan NISN yang dicari: ")
					fmt.Scan(&cariNISN)
					sequentianSearchNISN(&calonMahasiswa, cariNISN)

					switch menuEditable() {
					case 1:
						// Ubah
						ubahDataCalon(&calonMahasiswa, cariNISN)
					case 2:
						//Hapus
			
					case 3:
						break menuCari
					default:
						break
					}
				case 2:
					fmt.Println("🔍 Pencarian dengan Binary Search 🔍")
					//Urutkan data
					urutkanNISN(&calonMahasiswa)
					//Input NISN
					fmt.Print("Masukkan NISN yang dicari: ")
					fmt.Scan(&cariNISN)

					binarySearchNISN(&calonMahasiswa, cariNISN)
					switch menuEditable() {
					case 1:
						//Ubah
						ubahDataCalon(&calonMahasiswa, cariNISN)
					case 2:
						calonMahasiswa = hapusDataBsByNISN(&calonMahasiswa, cariNISN)
					case 3:
						break menuCari
					default:
						break
					}

				default:
					break
				}
				if metode == 3 {
					break
				}
			}

		case 5:
			fmt.Println("Terima kasih telah menggunakan fitur ini")
			break selesai
		}
	}
}

// Fungsi mengurutkan ASC bedasarkan NISN
func urutkanNISN(calonMahasiswa *dataSiswa) {
	for i := 1; i < len(calonMahasiswa); i++ {
		temp := calonMahasiswa[i]
		// j penunjuk kedua (berada di posisi belakang i)
		j := i
		for j > 0 && float64(temp.nisn) < float64(calonMahasiswa[j-1].nisn) {
			calonMahasiswa[j] = calonMahasiswa[j-1]
			j--
		}
		calonMahasiswa[j] = temp
	}
}

// Fungsi menampilkan semua isi array
func tampilkanIsi(calonMahasiswa dataSiswa) {
	for i, p := range calonMahasiswa {
		fmt.Println(i+1, ". NISN: ", p.nisn)
		fmt.Println("Nama: ", p.nama)
		fmt.Println("Tempat Lahir: ", p.tempatLahir)
		fmt.Println("Tanggal Lahir: ", p.tanggalLahir)
		fmt.Println("Jenis Kelamin: ", p.jenisKelamin)
		fmt.Println("Agama: ", p.agama)
		fmt.Println("Email: ", p.email)
		fmt.Println("Jurusan: ", p.jurusan)
		fmt.Println("Asal Sekolah: ", p.asalSekolah)
		fmt.Println("Tahun Lulus: ", p.tahunLulus)
		fmt.Println("Jurusan yang Dituju: ", p.jurusanYangDituju)
		fmt.Println("Nilai UTBK: ", p.nilaiUTBK)
		fmt.Println("Status: ", p.status)
		fmt.Println()
	}
}

func ubahDataCalon(calonMahasiswa *dataSiswa, cariNISN int) {
	var gantiData int
	fmt.Println("-------------------------------------")
	fmt.Println("Pilih data apa yang ingin di ubah")
	fmt.Println("1. Nama")
	fmt.Println("2. Tempat Lahir")
	fmt.Println("3. Tanggal Lahir")
	fmt.Println("4. Jenis Kelamin")
	fmt.Println("5. Agama")
	fmt.Println("6. Email")
	fmt.Println("7. Jurusan")
	fmt.Println("8. Asal Sekolah")
	fmt.Println("9. Tahun Lulus")
	fmt.Println("10. Jurusan yang Dituju")
	fmt.Print("Masukkan data yang ingin diganti: ")
	fmt.Scan(&gantiData)

	for i := 0; i < len(calonMahasiswa); i++ {
		if calonMahasiswa[i].nisn == cariNISN {
			switch gantiData {
			case 1:
				fmt.Print("Masukkan nama baru: ")
				reader := bufio.NewReader(os.Stdin)
				reader.ReadString('\n')
				inputNamaBaru, _ := reader.ReadString('\n')
				calonMahasiswa[i].nama = strings.TrimSpace(inputNamaBaru)
				fmt.Println("🎉 Nama berhasil diganti 🎉")
				fmt.Println("---------------------------------")
			case 2:
				fmt.Print("Masukkan tempat lahir baru: ")
				reader := bufio.NewReader(os.Stdin)
				reader.ReadString('\n')
				inputTempatLahirBaru, _ := reader.ReadString('\n')
				calonMahasiswa[i].tempatLahir = strings.TrimSpace(inputTempatLahirBaru)
				fmt.Println("🎉 Tempat lahir berhasil diganti 🎉")
				fmt.Println("-----------------------------------------")
			case 3:
				fmt.Print("Masukkan tanggal lahir baru: ")
				reader := bufio.NewReader(os.Stdin)
				reader.ReadString('\n')
				inputTanggalLahirBaru, _ := reader.ReadString('\n')
				calonMahasiswa[i].tanggalLahir = strings.TrimSpace(inputTanggalLahirBaru)
				fmt.Println("🎉 Tanggal lahir berhasil diganti 🎉")
				fmt.Println("------------------------------------------")
			case 4:
				fmt.Print("Masukkan jenis kelamin baru: ")
				reader := bufio.NewReader(os.Stdin)
				reader.ReadString('\n')
				inputJenisKelaminBaru, _ := reader.ReadString('\n')
				calonMahasiswa[i].jenisKelamin = strings.TrimSpace(inputJenisKelaminBaru)
				fmt.Println("🎉 Jenis Kelamin berhasil diganti 🎉")
				fmt.Println("-------------------------------------------")
			case 5:
				fmt.Print("Masukkan agama baru: ")
				reader := bufio.NewReader(os.Stdin)
				reader.ReadString('\n')
				inputAgamaBaru, _ := reader.ReadString('\n')
				calonMahasiswa[i].agama = strings.TrimSpace(inputAgamaBaru)
				fmt.Println("🎉 Agama berhasil diganti 🎉")
				fmt.Println("----------------------------------")
			case 6:
				fmt.Print("Masukkan email baru: ")
				reader := bufio.NewReader(os.Stdin)
				reader.ReadString('\n')
				inputEmailBaru, _ := reader.ReadString('\n')
				calonMahasiswa[i].email = strings.TrimSpace(inputEmailBaru)
				fmt.Println("🎉 Email berhasil diganti 🎉")
				fmt.Println("----------------------------------")
			case 7:
				fmt.Print("Masukkan jurusan baru: ")
				reader := bufio.NewReader(os.Stdin)
				reader.ReadString('\n')
				inputJurusanBaru, _ := reader.ReadString('\n')
				calonMahasiswa[i].jurusan = strings.TrimSpace(inputJurusanBaru)
				fmt.Println("🎉 Jurusan berhasil diganti 🎉")
				fmt.Println("-----------------------------------")
			case 8:
				fmt.Print("Masukkan asal sekolah baru: ")
				reader := bufio.NewReader(os.Stdin)
				reader.ReadString('\n')
				inputAsalSekolahBaru, _ := reader.ReadString('\n')
				calonMahasiswa[i].asalSekolah = strings.TrimSpace(inputAsalSekolahBaru)
				fmt.Println("🎉 Asal Sekolah berhasil diganti 🎉")
				fmt.Println("-----------------------------------------")
			case 9:
				fmt.Print("Masukkan tahun lulus baru: ")
				fmt.Scan(&calonMahasiswa[i].tahunLulus)
				fmt.Println("🎉 Tahun Lulus berhasil diganti 🎉")
				fmt.Println("----------------------------------------")
			case 10:
				fmt.Print("Masukkan jurusan yang dituju baru: ")
				reader := bufio.NewReader(os.Stdin)
				reader.ReadString('\n')
				inputJurusanYangDitujuBaru, _ := reader.ReadString('\n')
				calonMahasiswa[i].jurusanYangDituju = strings.TrimSpace(inputJurusanYangDitujuBaru)
				fmt.Println("🎉 Jurusan yang dituju berhasil diganti 🎉")
				fmt.Println("------------------------------------------------")
			}
		}
	}
}

// Fungsi untuk pencarian (sequential search)
func sequentianSearchNISN(calonMahasiswa *dataSiswa, cariNISN int) {
	found := false
	for _, j := range calonMahasiswa {
		if j.nisn == cariNISN {
			fmt.Printf("🕵️  Ditemukan Data Calon Mahasiswa dengan NISN %d 🕵️\n", cariNISN)
			fmt.Println("NISN: ", j.nisn)
			fmt.Println("Nama: ", j.nama)
			fmt.Println("Tempat Lahir: ", j.tempatLahir)
			fmt.Println("Tanggal Lahir: ", j.tanggalLahir)
			fmt.Println("Jenis Kelamin: ", j.jenisKelamin)
			fmt.Println("Agama: ", j.agama)
			fmt.Println("Email: ", j.email)
			fmt.Println("Jurusan: ", j.jurusan)
			fmt.Println("Asal Sekolah: ", j.asalSekolah)
			fmt.Println("Tahun Lulus: ", j.tahunLulus)
			fmt.Println("Jurusan yang Dituju: ", j.jurusanYangDituju)
			fmt.Println("Nilai UTBK: ", j.nilaiUTBK)
			fmt.Println("Status: ", j.status)
			found = true
			break
		}
		if !found {
			fmt.Printf("🧐  Tidak ditemukan data untuk calon mahasiswa dengan NISN %d 🧐\n", cariNISN)
			break
		}
	}
}

// Fungsi untuk penarian (binary search)
func binarySearchNISN(calonMahasiswa *dataSiswa, cariNISN int) int {
	low := 0
	high := len(calonMahasiswa) - 1

	for low <= high {
		mid := (low + high) / 2
		if calonMahasiswa[mid].nisn == cariNISN {
			return mid
		} else if calonMahasiswa[mid].nisn < cariNISN {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// Fungsi untuk menu ubah dan hapus
func hapusDataBsByNISN(calonMahasiswa *dataSiswa, cariNisn int) dataSiswa {
	idx := binarySearchNISN(calonMahasiswa, cariNisn)
	if idx == -1 {
		fmt.Println("404 data not found")
		return *calonMahasiswa
	}

	var newData [max]dataPendaftar
	j := 0
	for i := 0; i < len(calonMahasiswa); i++ {
		if i != idx {
			newData[j] = calonMahasiswa[i]
			j++
		}
	}

	// Bersihkan sisanya
	for k := j; k < max; k++ {
		newData[k] = dataPendaftar{}
	}

	return newData
}

func menuEditable() int {
	var aksi int
	fmt.Println("-------------")
	fmt.Println("1. Ubah Data")
	fmt.Println("2. Hapus Data")
	fmt.Println("3. Kembali ke menu utama")
	fmt.Print("Aksi: ")
	fmt.Scan(&aksi)
	return aksi
}

// Fungsi Insertion Sort ASC
func insertionSort_Asc(calonMahasiswa *dataSiswa) {
	// i penunjuk pertama
	for i := 1; i < len(calonMahasiswa); i++ {
		temp := calonMahasiswa[i]
		// j penunjuk kedua (berada di posisi belakang i)
		j := i
		for j > 0 && temp.nilaiUTBK < calonMahasiswa[j-1].nilaiUTBK {
			calonMahasiswa[j] = calonMahasiswa[j-1]
			j--
		}
		calonMahasiswa[j] = temp
	}
}

// Fungsi Insertion Sort DESC
func insertionSort_Desc(calonMahasiswa *dataSiswa) {
	// i penunjuk pertama
	for i := 1; i < len(calonMahasiswa); i++ {
		temp := calonMahasiswa[i]
		// j penunjuk kedua (berada di posisi belakang i)
		j := i
		for j > 0 && temp.nilaiUTBK > calonMahasiswa[j-1].nilaiUTBK {
			calonMahasiswa[j] = calonMahasiswa[j-1]
			j--
		}
		calonMahasiswa[j] = temp
	}
}

// Fungsi untuk mengurutkan data menggunakan selection sort ascending
func selectionSortAscending(calonMahasiswa *dataSiswa) {
	for i := 0; i < len(calonMahasiswa)-1; i++ {
		idx_min := i
		for j := i + 1; j < len(calonMahasiswa); j++ {
			if calonMahasiswa[j].nilaiUTBK < calonMahasiswa[idx_min].nilaiUTBK {
				idx_min = j
			}
		}
		calonMahasiswa[i], calonMahasiswa[idx_min] = calonMahasiswa[idx_min], calonMahasiswa[i]
	}
}

// Fungsi untuk mengurutkan data menggunakan selection sort ascending
func selectionSortDescending(calonMahasiswa *dataSiswa) {
	for i := 0; i < len(calonMahasiswa)-1; i++ {
		idx_min := i
		for j := i + 1; j < len(calonMahasiswa); j++ {
			if calonMahasiswa[j].nilaiUTBK > calonMahasiswa[idx_min].nilaiUTBK {
				idx_min = j
			}
		}
		calonMahasiswa[i], calonMahasiswa[idx_min] = calonMahasiswa[idx_min], calonMahasiswa[i]
	}
}

// Fungsi untuk menentukan status penerimaan
func status(calonMahasiswa *dataSiswa) {
	for i, k := range calonMahasiswa {
		if k.nilaiUTBK >= 600 {
			calonMahasiswa[i].status = "Diterima"
		} else {
			calonMahasiswa[i].status = "Tidak Diterima"
		}
	}
}

// Fungsi untuk menambah data
func tambahData(calonMahasiswa *dataSiswa, banyakData int) {

	for i := 0; i < banyakData; i++ {
		fmt.Print("Masukkan NISN: ")
		fmt.Scan(&calonMahasiswa[i].nisn)
		fmt.Print("Masukkan Nama: ")
		fmt.Scan(&calonMahasiswa[i].nama)
		fmt.Print("Masukkan Tempat Lahir: ")
		fmt.Scan(&calonMahasiswa[i].tempatLahir)
		fmt.Print("Masukkan Tanggal Lahir: ")
		fmt.Scan(&calonMahasiswa[i].tanggalLahir)
		fmt.Print("Masukkan Jenis Kelamin: ")
		fmt.Scan(&calonMahasiswa[i].jenisKelamin)
		fmt.Print("Masukkan Agama: ")
		fmt.Scan(&calonMahasiswa[i].agama)
		fmt.Print("Masukkan Email: ")
		fmt.Scan(&calonMahasiswa[i].email)
		fmt.Print("Masukkan Jurusan: ")
		fmt.Scan(&calonMahasiswa[i].jurusan)
		fmt.Print("Masukkan Asal Sekolah: ")
		fmt.Scan(&calonMahasiswa[i].asalSekolah)
		fmt.Print("Masukkan Tahun Lulus: ")
		fmt.Scan(&calonMahasiswa[i].tahunLulus)
		fmt.Print("Masukkan Jurusan yang Dituju: ")
		fmt.Scan(&calonMahasiswa[i].jurusanYangDituju)
		fmt.Print("Masukkan Nilai UTBK: ")
		fmt.Scan(&calonMahasiswa[i].nilaiUTBK)
		status(calonMahasiswa)
		fmt.Println()
	}
	fmt.Println("✅ Data berhasil ditambahkan")
}

// Prosedur untuk menampilkan menu utama
func mainMenu() int {
	var pilihan int
	fmt.Println("-------- MENU UTAMA --------")
	fmt.Println("1. Tampilkan Semua Data")
	fmt.Println("2. Urutkan Data")
	fmt.Println("3. Tambah Data")
	fmt.Println("4. Cari Data")
	fmt.Println("5. Keluar")
	fmt.Print("Pilih menu: ")
	fmt.Scan(&pilihan)
	return pilihan

}

// ============ *Kalau ada fitur login beb, jangan panik duluw
/*
func verifikasiLogin() {
	var nama, password string
	fmt.Println("======= SELAMAT DATANG =======")
	fmt.Println("Penerimaan Mahasiswa Bukit Duri")
	//Input nama
	fmt.Print("Masukkan nama : ")
	fmt.Scanln(&nama)
	//Input password
	fmt.Print("Masukkan password : ")
	fmt.Scanln(&password)

	switch {
	case nama == "admin" && password == "admin":
		fmt.Println("Hello admin")
		//Masuk ke menu admin
	default:
	}
}
*/
