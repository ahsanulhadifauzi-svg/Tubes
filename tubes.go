package main

import "fmt"

const NMAX int = 100

type pemilik struct {
	idPemilik string
	nama      string
	alamat    string
	noHP      string
}

type kendaraan struct {
	platNomor     string
	merk          string
	tahunProduksi int
	bulanServis   string
	idPemilik     string
}

type riwayatServis struct {
	platNomor      string
	bulanServis    string
	jenisKerusakan string
}

type tabKendaraan [NMAX]kendaraan
type tabPemilik [NMAX]pemilik
type tabServis [NMAX]riwayatServis

func tambahPemilik(P *tabPemilik, nP *int) {
	fmt.Print("ID Pemilik: ")
	fmt.Scan(&P[*nP].idPemilik)

	fmt.Print("Nama: ")
	fmt.Scan(&P[*nP].nama)

	fmt.Print("Alamat: ")
	fmt.Scan(&P[*nP].alamat)

	fmt.Print("No HP: ")
	fmt.Scan(&P[*nP].noHP)

	*nP++
}

func tambahKendaraan(K *tabKendaraan, nK *int) {
	fmt.Print("Plat Nomor: ")
	fmt.Scan(&K[*nK].platNomor)

	fmt.Print("Merk: ")
	fmt.Scan(&K[*nK].merk)

	fmt.Print("Tahun Produksi: ")
	fmt.Scan(&K[*nK].tahunProduksi)

	fmt.Print("bulan Servis Terakhir: ")
	fmt.Scan(&K[*nK].bulanServis)

	fmt.Print("ID Pemilik: ")
	fmt.Scan(&K[*nK].idPemilik)

	*nK++
}

func tampilKendaraan(K tabKendaraan, nK int) {
	var i int

	fmt.Println()
	fmt.Println("DATA KENDARAAN")

	for i = 0; i < nK; i++ {
		fmt.Println("Plat :", K[i].platNomor)
		fmt.Println("Merk :", K[i].merk)
		fmt.Println("Tahun:", K[i].tahunProduksi)
		fmt.Println("bulan Servis:", K[i].bulanServis)
		fmt.Println()
	}
}

func sequentialSearch(K tabKendaraan, nK int, plat string) int {
	var i int

	for i = 0; i < nK; i++ {
		if K[i].platNomor == plat {
			return i
		}
	}
	return -1
}

func selectionSortTahun(K *tabKendaraan, nK int) {
	var pass, idx, i int
	var temp kendaraan

	for pass = 0; pass < nK-1; pass++ {
		idx = pass

		for i = pass + 1; i < nK; i++ {
			if K[i].tahunProduksi < K[idx].tahunProduksi {
				idx = i
			}
		}

		temp = K[pass]
		K[pass] = K[idx]
		K[idx] = temp
	}
}

func insertionSortPlat(K *tabKendaraan, nK int) {
	var pass, i int
	var temp kendaraan

	for pass = 1; pass < nK; pass++ {
		temp = K[pass]
		i = pass

		for i > 0 && temp.platNomor < K[i-1].platNomor {
			K[i] = K[i-1]
			i--
		}

		K[i] = temp
	}
}

// Fungsi untuk melakukan binary search pada data kendaraan yang sudah diurutkan berdasarkan plat nomor
func binarySearch(K tabKendaraan, nK int, plat string) int {
	var left, right, mid int

	left = 0
	right = nK - 1

	for left <= right {
		mid = (left + right) / 2

		if K[mid].platNomor == plat {
			return mid
		} else if K[mid].platNomor < plat {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

// Fungsi untuk menambah riwayat servis
func tambahServis(S *tabServis, nS *int) {
	fmt.Print("Plat Nomor: ")
	fmt.Scan(&S[*nS].platNomor)

	fmt.Print("bulan Servis: ")
	fmt.Scan(&S[*nS].bulanServis)

	fmt.Print("Jenis Kerusakan: ")
	fmt.Scan(&S[*nS].jenisKerusakan)

	*nS++
}

// Fungsi untuk mengubah data pemilik atau kendaraan
func ubahData(P *tabPemilik, nP int, K *tabKendaraan, nK int) {
	var pilih int
	var id, plat string
	var i int
	var ketemu bool

	fmt.Println("1. Ubah Pemilik")
	fmt.Println("2. Ubah Kendaraan")
	fmt.Print("Pilihan: ")
	fmt.Scan(&pilih)

	if pilih == 1 {

		fmt.Print("ID Pemilik: ")
		fmt.Scan(&id)

		ketemu = false

		for i = 0; i < nP; i++ {
			if P[i].idPemilik == id {

				fmt.Print("Nama baru: ")
				fmt.Scan(&P[i].nama)

				fmt.Print("Alamat baru: ")
				fmt.Scan(&P[i].alamat)

				fmt.Print("No HP baru: ")
				fmt.Scan(&P[i].noHP)

				fmt.Println("Data pemilik berhasil diubah")
				ketemu = true
			}
		}

		if !ketemu {
			fmt.Println("Data tidak ditemukan")
		}

	} else if pilih == 2 {

		fmt.Print("Plat Nomor: ")
		fmt.Scan(&plat)

		ketemu = false

		for i = 0; i < nK; i++ {
			if K[i].platNomor == plat {

				fmt.Print("Merk baru: ")
				fmt.Scan(&K[i].merk)

				fmt.Print("Tahun produksi baru: ")
				fmt.Scan(&K[i].tahunProduksi)

				fmt.Print("Bulan servis baru: ")
				fmt.Scan(&K[i].bulanServis)

				fmt.Println("Data kendaraan berhasil diubah")
				ketemu = true
			}
		}

		if !ketemu {
			fmt.Println("Data tidak ditemukan")
		}
	}
}

// Fungsi untuk menghapus data pemilik atau kendaraan
func hapusData(P *tabPemilik, nP *int, K *tabKendaraan, nK *int) {
	var pilih int
	var id, plat string
	var i, idx int

	fmt.Println("1. Hapus Pemilik")
	fmt.Println("2. Hapus Kendaraan")
	fmt.Print("Pilihan: ")
	fmt.Scan(&pilih)

	if pilih == 1 {

		fmt.Print("ID Pemilik: ")
		fmt.Scan(&id)

		idx = -1

		for i = 0; i < *nP; i++ {
			if P[i].idPemilik == id {
				idx = i
			}
		}

		if idx != -1 {

			for i = idx; i < *nP-1; i++ {
				P[i] = P[i+1]
			}

			*nP--

			fmt.Println("Data pemilik berhasil dihapus")
		} else {
			fmt.Println("Data tidak ditemukan")
		}

	} else if pilih == 2 {

		fmt.Print("Plat Nomor: ")
		fmt.Scan(&plat)

		idx = sequentialSearch(*K, *nK, plat)

		if idx != -1 {

			for i = idx; i < *nK-1; i++ {
				K[i] = K[i+1]
			}

			*nK--

			fmt.Println("Data kendaraan berhasil dihapus")
		} else {
			fmt.Println("Data tidak ditemukan")
		}
	}
}

// Fungsi untuk menghitung statistik bulanan servis
func statistikBulanan(S tabServis, nS int) {
	var bulan [12]string
	var jumlah [12]int
	var nBulan int
	var i, j int
	var ketemu bool

	for i = 0; i < nS; i++ {

		ketemu = false

		for j = 0; j < nBulan; j++ {
			if bulan[j] == S[i].bulanServis {
				jumlah[j]++
				ketemu = true
			}
		}

		if !ketemu {
			bulan[nBulan] = S[i].bulanServis
			jumlah[nBulan] = 1
			nBulan++
		}
	}

	fmt.Println("+++ STATISTIK BULANAN +++")

	for i = 0; i < nBulan; i++ {
		fmt.Println("bulan ke-", bulan[i], ":", jumlah[i], "servis")
	}
}

// Fungsi untuk menghitung jenis kerusakan yang paling sering terjadi
func kerusakanTerbanyak(S tabServis, nS int) {
	var jenis [100]string
	var jumlah [100]int
	var nJenis int
	var i, j int
	var ketemu bool
	var idxMax int

	for i = 0; i < nS; i++ {

		ketemu = false

		for j = 0; j < nJenis; j++ {
			if jenis[j] == S[i].jenisKerusakan {
				jumlah[j]++
				ketemu = true
			}
		}

		if !ketemu {
			jenis[nJenis] = S[i].jenisKerusakan
			jumlah[nJenis] = 1
			nJenis++
		}
	}

	if nJenis > 0 {
		idxMax = 0

		for i = 1; i < nJenis; i++ {
			if jumlah[i] > jumlah[idxMax] {
				idxMax = i
			}
		}

		fmt.Println("Kerusakan terbanyak :", jenis[idxMax])
		fmt.Println("Jumlah :", jumlah[idxMax])
	}
}
func main() {
	var K tabKendaraan
	var P tabPemilik
	var S tabServis

	var nK, nP, nS int
	var pilih int
	var plat string
	var idx int

	pilih = -1

	for pilih != 0 {
		fmt.Println()
		fmt.Println("+++ AUTOCARE +++")
		fmt.Println("1. Tambah Pemilik")
		fmt.Println("2. Tambah Kendaraan")
		fmt.Println("3. Tambah Riwayat Servis")
		fmt.Println("4. Tampil Kendaraan")
		fmt.Println("5. Sequential Search")
		fmt.Println("6. Selection Sort Tahun")
		fmt.Println("7. Binary Search")
		fmt.Println("8. Ubah Data")
		fmt.Println("9. Hapus Data")
		fmt.Println("10. Statistik Servis")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilih)

		switch pilih {

		case 1:
			tambahPemilik(&P, &nP)

		case 2:
			tambahKendaraan(&K, &nK)

		case 3:
			tambahServis(&S, &nS)

		case 4:
			tampilKendaraan(K, nK)

		case 5:
			fmt.Print("Plat Nomor: ")
			fmt.Scan(&plat)

			idx = sequentialSearch(K, nK, plat)

			if idx != -1 {
				fmt.Println("Data ditemukan pada indeks", idx)
			} else {
				fmt.Println("Data tidak ditemukan")
			}

		case 6:
			selectionSortTahun(&K, nK)
			fmt.Println("Data berhasil diurutkan")

		case 7:
			insertionSortPlat(&K, nK)

			fmt.Print("Plat Nomor: ")
			fmt.Scan(&plat)

			idx = binarySearch(K, nK, plat)

			if idx != -1 {
				fmt.Println("Data ditemukan pada indeks", idx)
			} else {
				fmt.Println("Data tidak ditemukan")
			}

		case 8:
			ubahData(&P, nP, &K, nK)

		case 9:
			hapusData(&P, &nP, &K, &nK)

		case 10:
			fmt.Println("----- AUTO CARE -----")
			statistikBulanan(S, nS)
			kerusakanTerbanyak(S, nS)

		case 0:
			fmt.Println("Terima kasih telah menggunakan AutoCare!")
		}
	}
}
