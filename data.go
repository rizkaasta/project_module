package data

import "fmt"

type dataMember struct {
	nama                 string
	no_anggota           int
	judul_buku           string
	jumlah_buku_dipinjam int
	jumlah_buku_telat    int
}

const satuanDenda int = 10000

func GetData(newname string) {
	fmt.Printf("Selamat bergabung di Perpustakaan, %s", newname)
}

func Denda(jumlah_buku_telat int, satuanDenda int) int {
		bayarDenda := jumlah_buku_telat * satuanDenda
		return bayarDenda
}
