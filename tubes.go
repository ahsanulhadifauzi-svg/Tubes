package main

import "fmt"

type pemilik struct {
	id, nama, alamat, nohp string
}

type kendaraan struct {
	plat, merk string
	tahunpro   int
}

type servis struct {
	blnservis, jumlahsrvsbln, jumlahker int
	jenisker string
}

type tabpemilik [100]pemilik
type tabkendaraan [100]kendaraan
type tabservis [100]servis

func tampemilik(p *tabpemilik, n *int) { // n * supaya nilai n bisa ke func global
	fmt.Println("Silahkan masukan data pemilik kendaraan")
	fmt.Print("ID Pemilik: ")
	fmt.Scan(&p[*n].id)
	fmt.Print("Nama: ")
	fmt.Scan(&p[*n].nama)
	fmt.Print("Alamat: ")
	fmt.Scan(&p[*n].alamat)
	fmt.Print("No HP: ")
	fmt.Scan(&p[*n].nohp)
	*n = *n + 1
}

func tamkendaraan(k *tabkendaraan, x *int, s *tabservis) { // sama kayak n
	fmt.Println("Silahkan masukan data kendaraan")
	fmt.Print("Plat Nomor: ")
	fmt.Scan(&k[*x].plat)
	fmt.Print("Merk: ")
	fmt.Scan(&k[*x].merk)
	fmt.Print("Tahun Produksi: ")
	fmt.Scan(&k[*x].tahunpro)
	fmt.Print("Bulan Servis: ")
	fmt.Scan(&s[*x].blnservis)
	s[*x].jenisker = "servis pertama" //  servis pertama untuk blm ada kerusakan
	*x = *x + 1
}

func tamservis(s *tabservis, y *int) { // sama kayak n
	var x int
	for x != -1 {
		if s[*y].blnservis == 0 {
			fmt.Print("Bulan Servis: ")
			fmt.Scan(&s[*y].blnservis)
			fmt.Print("Jenis Kerusakan: ")
			fmt.Scan(&s[*y].jenisker)
			x = -1
		}
		*y = *y + 1
	}
}

func selectionsort(s *tabservis, y int) {
	var i, j int
	var temp servis
	for i = 0; i < y-1; i++ {
		for j = i + 1; j < y; j++ {
			if s[i].blnservis > s[j].blnservis {
				temp = s[i]
				s[i] = s[j]
				s[j] = temp
			}
		}
	}
	fmt.Println("+---------------------------------------------+")
	for i = 0; i < y; i++ {
		fmt.Printf("|Bulan Servis: %d Jenis Kerusakan: %s|\n", s[i].blnservis, s[i].jenisker)
	}
	fmt.Println("+---------------------------------------------+")
}

func insertionsort(k *tabkendaraan, x int) {
	var i, pass int
	var temp kendaraan
	for pass <= x-1 {
		for i > 0 && k[i-1].tahunpro > k[i].tahunpro {
			temp = k[i]
			k[i] = k[i-1]
			k[i-1] = temp
			i = i - 1
		}
		pass = pass + 1
	}
	fmt.Println("+---------------------------------------------+")
	for i = 0; i < x; i++ {
		fmt.Printf("|Plat: %s Merk: %s Tahun Produksi: %d|\n", k[i].plat, k[i].merk, k[i].tahunpro)
	}
	fmt.Println("+---------------------------------------------+")
}

func binarysearch(k tabkendaraan, x int, cari string) int {
	var i, mid, ri, le, j int // ri untuk  kanan,le untuk  kiri
	var temp kendaraan

	for i = 0; i < x; i++ {
		for j = i + 1; j < x-1; j++ {
			if k[i].plat > k[j].plat {
				temp = k[i]
				k[i] = k[j]
				k[j] = temp
			}
		}
	}
	le = 0
	ri = x - 1
	mid = (le + ri) / 2
	for le <= ri {
		if k[mid].plat == cari {
			return mid
		} else if k[mid].plat < cari {
			le = mid + 1
		} else {
			ri = mid - 1
		}
		mid = (le + ri) / 2
	}
	return -1
}

func sequensialsearch(k tabkendaraan, x int, cari string) int {
	var i int
	for i = 0; i < x; i++ {
		if k[i].plat == cari {
			return i
		}
	}
	return -1
}

func mengubahdata(k *tabkendaraan, p *tabpemilik, n, x int) {
	var i, cari int         // z untuk menyimpan id yang dicari
	var q, z string         // q untuk menyimpan plat nomor yang dicari
	var ketemu bool = false // apakah data ditemukan atau gak
	fmt.Println("1.Ubah data pemilik")
	fmt.Println("2.Ubah data kendaraan")
	fmt.Print("pilih data yang ingin diubah: ")
	fmt.Scan(&cari)
	if cari == 1 {
		fmt.Print("Silahkan masukan ID pemilik yang ingin diubah: ")
		fmt.Scan(&z)
		for i = 0; i < n; i++ {
			if p[i].id == z {
				fmt.Print("Nama: ")
				fmt.Scan(&p[i].nama)
				fmt.Print("Alamat: ")
				fmt.Scan(&p[i].alamat)
				fmt.Print("No HP: ")
				fmt.Scan(&p[i].nohp)
				ketemu = true
			}
		}
		if ketemu == true {
			fmt.Println("Data pemilik berhasil diubah")
		} else {
			fmt.Println("Data pemilik tidak ditemukan")
		}
	} else if cari == 2 {
		fmt.Print("Silahkan masukan plat nomor kendaraan yang ingin diubah: ")
		fmt.Scan(&q)
		for i = 0; i < x; i++ { // bisa langsung diubah gak perlu di kurang atau ditambah lagi
			if k[i].plat == q {
				fmt.Print("Merk: ")
				fmt.Scan(&k[i].merk)
				fmt.Print("Tahun Produksi: ")
				fmt.Scan(&k[i].tahunpro)
				ketemu = true
			}
		}
		if ketemu == true {
			fmt.Println("Data kendaraan berhasil diubah")
		} else {
			fmt.Println("Data kendaraan tidak ditemukan")
		}
	}
}

func hapusdata(k *tabkendaraan, p *tabpemilik, n, x *int) {
	var i, cari, j int
	var q, z string
	var ketemu bool = false
	fmt.Println("1.Hapus data pemilik")
	fmt.Println("2.Hapus data kendaraan")
	fmt.Print("pilih data yang ingin dihapus: ")
	fmt.Scan(&cari)
	if cari == 1 {
		fmt.Print("Silahkan masukan ID pemilik yang ingin dihapus: ")
		fmt.Scan(&z)
		for i = 0; i < *n; i++ { // for untuk melihat yang dicari
			if p[i].id == z {
				for j = i; j < *n-1; j++ { // ini supaya saat menghapus data dapat digeser sehingga gak bolong datanya
					p[j] = p[j+1]
				}
				*n = *n - 1
				ketemu = true
			}
		}
		if ketemu == true {
			fmt.Println("Data pemilik berhasil dihapus")
		} else {
			fmt.Println("Data pemilik tidak ditemukan")
		}
	} else if cari == 2 {
		fmt.Print("Silahkan masukan kendaraan yang ingin dihapus dari data: ")
		fmt.Scan(&q)
		for i = 0; i < *x; i++ {
			if k[i].plat == q {
				for j = i; j < *x-1; j++ {
					k[j] = k[j+1]
				}
				*x = *x - 1
				ketemu = true
			}
		}
		if ketemu == true {
			fmt.Println("Data kendaraan berhasil dihapus")
		} else {
			fmt.Println("Data kendaraan tidak ditemukan")
		}
	}
}

func statistikservis(s tabservis, y int) {
	var i, j, max int
	var plser string // plser untuk nama kerusakan paling banyak
	for i = 0; i < y; i++ {
		switch s[i].blnservis {
		case 1:
			s[0].jumlahsrvsbln = s[0].jumlahsrvsbln + 1
		case 2:
			s[1].jumlahsrvsbln = s[1].jumlahsrvsbln + 1
		case 3:
			s[2].jumlahsrvsbln = s[2].jumlahsrvsbln + 1
		case 4:
			s[3].jumlahsrvsbln = s[3].jumlahsrvsbln + 1
		case 5:
			s[4].jumlahsrvsbln = s[4].jumlahsrvsbln + 1
		case 6:
			s[5].jumlahsrvsbln = s[5].jumlahsrvsbln + 1
		case 7:
			s[6].jumlahsrvsbln = s[6].jumlahsrvsbln + 1
		case 8:
			s[7].jumlahsrvsbln = s[7].jumlahsrvsbln + 1
		case 9:
			s[8].jumlahsrvsbln = s[8].jumlahsrvsbln + 1
		case 10:
			s[9].jumlahsrvsbln = s[9].jumlahsrvsbln + 1
		case 11:
			s[10].jumlahsrvsbln = s[10].jumlahsrvsbln + 1
		case 12:
			s[11].jumlahsrvsbln = s[11].jumlahsrvsbln + 1
		}
	}

	for i = 0; i <= 11; i++ {
		fmt.Printf("Bulan %d: %d servis\n", i+1, s[i].jumlahsrvsbln) // jika sama maka yang keluar yang terakhir
	}
	for i = 0; i < y; i++ {
		for j = 0; j < y; j++ {
			if s[i].jenisker == s[j].jenisker && s[i].jenisker != "servis pertama" { // supaya kerusakan terbanyak ngak keluar servis pertama
				s[i].jumlahker = s[i].jumlahker + 1
			}
		}
	}
	for i = 0; i < y; i++ {
		if s[i].jumlahker > max {
			max = s[i].jumlahker
			plser = s[i].jenisker
		}
	}
	fmt.Printf("%s dengan jumlah kerusakan sebanyak %d\n", plser, max) // untuk mengeliat jumlah terbanyak kerusakan 
}

func main() {
	var k tabkendaraan
	var p tabpemilik
	var s tabservis
	var x, y, z, a, tampung int // tampung adalah untuk menampung binary search dan sequential search
	var cari string
	a = -1
	fmt.Println()
	fmt.Println("Selamat datang di autocare")
	for a != 0 {
		fmt.Println()
		fmt.Println("Silahkan pilih menu yang tersedia")
		fmt.Println("1.Tambah data pemilik")
		fmt.Println("2.Tambah data kendaraan")
		fmt.Println("3.Tambah data servis")
		fmt.Println("4.Urutkan data sesuai bulan servis(selection sort)")
		fmt.Println("5.Urutkan data sesuai tahun produksi(insertion sort)")
		fmt.Println("6.Cari data kendaraan dengan plat nomor(binary search)")
		fmt.Println("7.Cari data kendaraan dengan plat nomor(sequential search)")
		fmt.Println("8.Ubah data")
		fmt.Println("9.Hapus data")
		fmt.Println("10.Statistik servis")
		fmt.Println("0.Keluar")
		fmt.Print("Silahkan masukan pilihan: ")
		fmt.Scan(&a)
		switch a {
		case 1:
			fmt.Println()
			tampemilik(&p, &x)
		case 2:
			fmt.Println()
			tamkendaraan(&k, &y, &s)
		case 3:
			fmt.Println()
			fmt.Println("Silahkan masukan data servis")
			tamservis(&s, &z)
		case 4:
			selectionsort(&s, z)
		case 5:
			insertionsort(&k, y)
		case 6:
			fmt.Println() // supaya ada spasi untuk ngeliat 
			fmt.Print("Plat yang dicari: ")
			fmt.Scan(&cari)
			tampung = binarysearch(k, y, cari)
			if tampung != -1 {
				fmt.Println("Data ditemukan")
				fmt.Printf("plat: %s merk: %s tahun produksi: %d\n", k[tampung].plat, k[tampung].merk, k[tampung].tahunpro)
			} else {
				fmt.Println("Data tidak ditemukan")
			}
		case 7:
			fmt.Println()
			fmt.Print("Plat yang dicari: ")
			fmt.Scan(&cari)
			tampung = sequensialsearch(k, y, cari)
			if tampung != -1 {
				fmt.Println("Data ditemukan")
				fmt.Printf("plat: %s merk: %s tahun produksi: %d\n", k[tampung].plat, k[tampung].merk, k[tampung].tahunpro)
			} else {
				fmt.Println("Data tidak ditemukan")
			}
		case 8:
			fmt.Println()
			mengubahdata(&k, &p, x, y)
		case 9:
			fmt.Println()
			hapusdata(&k, &p, &x, &y) // & supaya bisa mengubah semua data
		case 10:
			fmt.Println()
			fmt.Println("STATISTIK SERVIS AUTOCARE")
			statistikservis(s, z) 
		case 0:
			fmt.Println("Terima kasih telah menggunakan autocare")
		}
	}
}
