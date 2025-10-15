package main

import "fmt"

const NMAX int = 100

type SparePart struct {
	ID        int
	Nama      string
	HargaPart int
	Freq      int
}

type Service struct {
	Pelanggan  string
	SparePart  string
	HargaPart  int
	HargaJasa  int
	TotalHarga int
}

type Transaksi struct {
	Nama       string
	IDSpare    int
	HargaJasa  int
	TotalHarga int
	Pelanggan  int
}

type Pelanggan struct {
	ID        int
	Nama      string
	NoHp      string
	HargaJasa int
}

var (
	dataSparePart [NMAX]SparePart
	nSparePart    int

	dataPelanggan [NMAX]Pelanggan
	nPelanggan    int

	dataService [NMAX]Service
	nService    int

	dataTransaksi [NMAX]Transaksi
	nTransaksi    int
)

func main() {
	var klik int
	for klik != 5 {
		dashboardBengkel()
		fmt.Scan(&klik)
		if klik == 1 {
			menuSparePart(&dataSparePart, &nSparePart)
		} else if klik == 2 {
			menuPelanggan(&dataPelanggan, &nPelanggan)
		} else if klik == 3 {
			menuService(&dataService, &nService)
		} else if klik == 4 {
			menuTransaksi(&dataTransaksi, &nTransaksi)
		} else if klik == 5 {
			fmt.Println("Keluar program.")
		} else {
			fmt.Println("Menu tidak tersedia.")
		}
	}
}

func dashboardBengkel() {
	fmt.Println("======================")
	fmt.Println(" Selamat datang di bengkel kami ")
	fmt.Println("=====menu utama=====")
	fmt.Println("1. Menu Sparepart ")
	fmt.Println("2. Menu Pelanggan")
	fmt.Println("3. Menu Service")
	fmt.Println("4. Menu Transaksi ")
	fmt.Println("5. Keluar ")
	fmt.Println("====================")
	fmt.Print("Silahkan masukkan pilihan : ")
}

// =================== MENU SPAREPART =========================
func menuSparePart(data *[NMAX]SparePart, n *int) {
	var pilih int
	for pilih != 5 {
		fmt.Println("===== Menu Sparepart Motor =====")
		fmt.Println("1. Tambah Data Sparepart      ")
		fmt.Println("2. Edit Sparepart             ")
		fmt.Println("3. Hapus Sparepart            ")
		fmt.Println("4. Lihat semua Sparepart      ")
		fmt.Println("5. Kembali                    ")
		fmt.Println("==============================")
		fmt.Print("Masukan pilihan anda(1-5): ")
		fmt.Scan(&pilih)
		if pilih == 1 {
			tambahDataSparePart(data, n)
		} else if pilih == 2 {
			editDataSparePart(data, n)
		} else if pilih == 3 {
			hapusDataSparePart(data, n)
		} else if pilih == 4 {
			lihatDataSparePart(data, n)
		} else if pilih == 5 {
			fmt.Println("Kembali ke menu utama.")
			return
		} else {
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

// Tambah Data SparePart
func tambahDataSparePart(data *[NMAX]SparePart, n *int) {
	if *n > NMAX {
		fmt.Println("Data sparepart sudah penuh!")
		return
	}
	var s SparePart
	fmt.Print("Masukkan ID sparepart: ")
	fmt.Scan(&s.ID)
	fmt.Scanln()
	fmt.Print("Masukkan nama sparepart: ")
	fmt.Scan(&s.Nama)
	fmt.Scanln()

	s.HargaPart = 0
	for s.HargaPart <= 0 {
		fmt.Print("Masukkan harga sparepart: ")
		fmt.Scan(&s.HargaPart)
		if s.HargaPart <= 0 {
			fmt.Println("Harga tidak vali ")
		}
	}
	s.Freq = 0
	data[*n] = s
	*n++
	fmt.Println("Data sparepart berhasil ditambah!")
}

// Edit Data SparePart
func editDataSparePart(data *[NMAX]SparePart, n *int) {
	if *n == 0 {
		fmt.Println("Belum ada data sparepart.")
		return
	}
	var id int
	fmt.Print("Masukkan ID sparepart yang ingin diubah: ")
	fmt.Scan(&id)
	fmt.Scanln()
	idx := sequentialSearchSparePart(data, *n, id)
	if idx == -1 {
		fmt.Println("Data tidak ditemukan.")
		return
	}
	fmt.Print("Masukkan nama sparepart baru: ")
	fmt.Scan(&data[idx].Nama)
	fmt.Print("Masukkan harga sparepart baru: ")
	fmt.Scan(&data[idx].HargaPart)
	fmt.Println("Data berhasil diubah.")
}

// Hapus Data SparePart
func hapusDataSparePart(data *[NMAX]SparePart, n *int) {
	if *n == 0 {
		fmt.Println("Belum ada data sparepart.")
		return
	}
	var id int
	fmt.Print("Masukkan ID sparepart yang ingin dihapus: ")
	fmt.Scan(&id)
	idx := sequentialSearchSparePart(data, *n, id)
	if idx == -1 {
		fmt.Println("Data tidak ditemukan.")
		return
	}
	for i := idx; i < *n-1; i++ {
		data[i] = data[i+1]
	}
	*n--
	fmt.Println("Data berhasil dihapus.")
}

// Lihat Data SparePart (urut Freq terbanyak, Selection Sort)
func lihatDataSparePart(data *[NMAX]SparePart, n *int) {
	if *n == 0 {
		fmt.Println("Belum ada data sparepart.")
		return
	}
	selectionSortSparePartByFreq(data, *n)
	fmt.Println("\n--- Daftar Sparepart (Top Freq First) ---")
	fmt.Printf("%-5s %-20s %-10s %-10s\n", "ID", "Nama", "Harga", "Freq")
	for i := 0; i < *n; i++ {
		fmt.Printf("%-5d %-20s %-10d %-10d\n", data[i].ID, data[i].Nama, data[i].HargaPart, data[i].Freq)
	}
	fmt.Print("Tekan ENTER untuk kembali ke menu ")
	fmt.Scanln()
	fmt.Scanln()
}

// Selection Sort SparePart by Freq Descending
func selectionSortSparePartByFreq(data *[NMAX]SparePart, n int) {
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if data[j].Freq > data[maxIdx].Freq {
				maxIdx = j
			}
		}
		data[i], data[maxIdx] = data[maxIdx], data[i]
	}
}

// Sequential Search SparePart by ID
func sequentialSearchSparePart(data *[NMAX]SparePart, n int, id int) int {
	for i := 0; i < n; i++ {
		if data[i].ID == id {
			return i
		}
	}
	return -1
}

func tampilkanListSparepart(data *[NMAX]SparePart, n int) {
	fmt.Println("===== DAFTAR SPAREPART =====")
	fmt.Printf("%-5s %-20s %-10s\n", "No", "Nama", "Harga")
	for i := 0; i < n; i++ {
		fmt.Printf("%-5d %-20s %-10d\n", i+1, data[i].Nama, data[i].HargaPart)
	}
}

// =================== MENU PELANGGAN =========================
func menuPelanggan(data *[NMAX]Pelanggan, n *int) {
	var klik int
	for klik != 5 {
		fmt.Println("=====   Menu Pelanggan  =====")
		fmt.Println("1. Tambah pelanggan      ")
		fmt.Println("2. Edit pelanggan        ")
		fmt.Println("3. Hapus pelanggan       ")
		fmt.Println("4. Lihat semua pelanggan ")
		fmt.Println("5. Kembali               ")
		fmt.Println("==============================")
		fmt.Print("Masukan pilihan anda: ")
		fmt.Scan(&klik)
		if klik == 1 {
			tambahPelanggan(data, n)
		} else if klik == 2 {
			editPelanggan(data, n)
		} else if klik == 3 {
			hapusPelanggan(data, n)
		} else if klik == 4 {
			lihatPelanggan(data, n)
		} else if klik == 5 {
			fmt.Println("Kembali ke menu utama.")
			return
		} else {
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

// Tambah Pelanggan
func tambahPelanggan(data *[NMAX]Pelanggan, n *int) {
	if *n >= NMAX {
		fmt.Println("Data pelanggan sudah penuh!")
		return
	}
	var p Pelanggan
	fmt.Print("Masukan ID Pelanggan: ")
	fmt.Scan(&p.ID)
	fmt.Print("Masukan Nama pelanggan: ")
	fmt.Scan(&p.Nama)
	fmt.Print("Masukan No Hp pelanggan: ")
	fmt.Scan(&p.NoHp)
	fmt.Print("Masukan harga jasa: ")
	fmt.Scan(&p.HargaJasa)
	data[*n] = p
	*n++
	fmt.Println("Data pelanggan berhasil ditambah!")
}

// Edit Pelanggan
func editPelanggan(data *[NMAX]Pelanggan, n *int) {
	if *n == 0 {
		fmt.Println("Belum ada data pelanggan.")
		return
	}
	var id int
	fmt.Print("Masukkan ID pelanggan yang ingin diubah: ")
	fmt.Scan(&id)
	idx := sequentialSearchPelanggan(data, *n, id)
	if idx == -1 {
		fmt.Println("Data tidak ditemukan.")
		return
	}
	fmt.Print("Nama baru: ")
	fmt.Scan(&data[idx].Nama)
	fmt.Print("No HP baru: ")
	fmt.Scan(&data[idx].NoHp)
	fmt.Print("Harga jasa baru: ")
	fmt.Scan(&data[idx].HargaJasa)
	fmt.Println("Data pelanggan berhasil diubah!")
}

// Hapus Pelanggan
func hapusPelanggan(data *[NMAX]Pelanggan, n *int) {
	if *n == 0 {
		fmt.Println("Belum ada data pelanggan.")
		return
	}
	var id int
	fmt.Print("Masukkan ID pelanggan yang ingin dihapus: ")
	fmt.Scan(&id)
	idx := sequentialSearchPelanggan(data, *n, id)
	if idx == -1 {
		fmt.Println("Data tidak ditemukan.")
		return
	}
	for i := idx; i < *n-1; i++ {
		data[i] = data[i+1]
	}
	*n--
	fmt.Println("Data pelanggan berhasil dihapus.")
}

// Lihat Pelanggan
func lihatPelanggan(data *[NMAX]Pelanggan, n *int) {
	if *n == 0 {
		fmt.Println("Belum ada data pelanggan.")
		return
	}
	fmt.Printf("%-5s %-20s %-15s %-10s\n", "ID", "Nama", "No HP", "HargaJasa")
	for i := 0; i < *n; i++ {
		fmt.Printf("%-5d %-20s %-15s %-10d\n", data[i].ID, data[i].Nama, data[i].NoHp, data[i].HargaJasa)
	}
}

// Sequential Search Pelanggan by ID
func sequentialSearchPelanggan(data *[NMAX]Pelanggan, n int, id int) int {
	for i := 0; i < n; i++ {
		if data[i].ID == id {
			return i
		}
	}
	return -1
}

// =================== MENU SERVICE =========================
func menuService(data *[NMAX]Service, n *int) {
	var pilih int
	for pilih != 4 {
		fmt.Println("==== Menu Service ====")
		fmt.Println("1. Tambah Data service")
		fmt.Println("2. Lihat semua service")
		fmt.Println("3. Edit data service")
		fmt.Println("4. Kembali")
		fmt.Print("Masukan Pilihan anda (1-4): ")
		fmt.Scan(&pilih)
		if pilih == 1 {
			tambahDataService(data, n)
		} else if pilih == 2 {
			lihatDataService(data, n)
		} else if pilih == 3 {
			editDataService(data, n)
		} else if pilih == 4 {
			fmt.Println("Kembali ke menu utama.")
			return
		} else {
			fmt.Println("Menu tidak ada")
		}
	}
}

// Tambah Data Service
func tambahDataService(data *[NMAX]Service, n *int) {
	if *n > NMAX {
		fmt.Println("Data service sudah penuh!")
		return
	}
	var s Service
	fmt.Print("Masukkan nama pelanggan: ")
	fmt.Scan(&s.Pelanggan)
	tampilkanListSparepart(&dataSparePart, nSparePart)
	var pilih int
	fmt.Print("Pilih nomor sparepart yang diinginkan:")
	fmt.Scan(&pilih)
	if pilih < 1 || pilih > nSparePart {
		fmt.Println("pilihan tidak valid.")
		return
	}
	s.SparePart = dataSparePart[pilih-1].Nama
	s.HargaPart = dataSparePart[pilih-1].HargaPart
	fmt.Print("Masukkan harga jasa: ")
	fmt.Scan(&s.HargaJasa)
	s.TotalHarga = s.HargaPart + s.HargaJasa
	partIdx := sequentialSearchSparePart(&dataSparePart, nSparePart, cariIDSparePartByNama(s.SparePart))
	if partIdx != -1 {
		dataSparePart[partIdx].Freq++
	}
	data[*n] = s
	*n++
	fmt.Println("Data service berhasil ditambah!")
}

// Lihat Data Service
func lihatDataService(data *[NMAX]Service, n *int) {
	if *n == 0 {
		fmt.Println("Belum ada data service.")
		return
	}
	fmt.Printf("%-3s %-20s %-20s %-12s %-12s %-12s\n", "No", "Pelanggan", "SparePart", "HargaPart", "HargaJasa", "Total")
	for i := 0; i < *n; i++ {
		fmt.Printf("%-3d %-20s %-20s %-12d %-12d %-12d\n", i+1, data[i].Pelanggan, data[i].SparePart, data[i].HargaPart, data[i].HargaJasa, data[i].TotalHarga)
	}
}

// Edit Data Service
func editDataService(data *[NMAX]Service, n *int) {
	if *n == 0 {
		fmt.Println("Belum ada data service.")
		return
	}
	var namaCari string
	fmt.Print("Masukkan nama pelanggan yang ingin diubah servicenya: ")
	fmt.Scan(&namaCari)
	idx := sequentialSearchService(data, *n, namaCari)
	if idx == -1 {
		fmt.Println("Data tidak ditemukan.")
		return
	}
	fmt.Print("Nama pelanggan baru: ")
	fmt.Scan(&data[idx].Pelanggan)
	fmt.Print("Nama sparepart baru: ")
	fmt.Scan(&data[idx].SparePart)
	fmt.Print("Harga part baru: ")
	fmt.Scan(&data[idx].HargaPart)
	fmt.Print("Harga jasa baru: ")
	fmt.Scan(&data[idx].HargaJasa)
	data[idx].TotalHarga = data[idx].HargaPart + data[idx].HargaJasa
	fmt.Println("Data service berhasil diubah!")
}

// Sequential Search Service by Nama Pelanggan
func sequentialSearchService(data *[NMAX]Service, n int, nama string) int {
	for i := 0; i < n; i++ {
		if data[i].Pelanggan == nama {
			return i
		}
	}
	return -1
}

// Cari ID SparePart by Nama
func cariIDSparePartByNama(nama string) int {
	for i := 0; i < nSparePart; i++ {
		if dataSparePart[i].Nama == nama {
			return dataSparePart[i].ID
		}
	}
	return -1
}

// =================== MENU TRANSAKSI =========================
func menuTransaksi(data *[NMAX]Transaksi, n *int) {
	var klik int
	for klik != 3 {
		fmt.Println("================")
		fmt.Println("1. Tambah Transaksi")
		fmt.Println("2. Lihat riwayat Transaksi ")
		fmt.Println("3. Kembali")
		fmt.Print("Masukan pilihan : ")
		fmt.Scan(&klik)
		if klik == 1 {
			tambahDataTransaksi(data, n)
		} else if klik == 2 {
			lihatRiwayatTransaksi(data, n)
		} else if klik == 3 {
			fmt.Println("Kembali ke menu utama.")
			return
		} else {
			fmt.Println("Menu tidak ada")
		}
	}
}

// Tambah Data Transaksi
func tambahDataTransaksi(data *[NMAX]Transaksi, n *int) {
	if *n >= NMAX {
		fmt.Println("Data transaksi sudah penuh!")
		return
	}
	var t Transaksi
	fmt.Print("Masukkan nama pelanggan: ")
	fmt.Scan(&t.Nama)
	fmt.Print("Masukkan ID spare part: ")
	fmt.Scan(&t.IDSpare)
	fmt.Print("Masukkan harga jasa: ")
	fmt.Scan(&t.HargaJasa)
	var HargaPart int
	for i := 0; i < nSparePart; i++ {
		if dataSparePart[i].ID == t.IDSpare {
			HargaPart = dataSparePart[i].HargaPart
		}
	}
	t.TotalHarga = t.HargaJasa + HargaPart
	fmt.Printf("total harga (otomatis dihitung): %d\n", t.TotalHarga)
	data[*n] = t
	*n++
	fmt.Println("Data transaksi berhasil ditambah!")

}

// Lihat Riwayat Transaksi
func lihatRiwayatTransaksi(data *[NMAX]Transaksi, n *int) {
	if *n == 0 {
		fmt.Println("Belum ada data transaksi.")
		return
	}
	fmt.Printf("%-3s %-20s %-12s %-12s %-12s\n", "No", "Nama", "IDSpare", "HargaJasa", "Total")
	for i := 0; i < *n; i++ {
		fmt.Printf("%-3d %-20s %-12d %-12d %-12d\n", i+1, data[i].Nama, data[i].IDSpare, data[i].HargaJasa, data[i].TotalHarga)

	}

}
