package data

import "fmt"

type DataMember struct {
	nama                string
	noAnggota           string
	judulBuku           string
	jumlahPinjam		int
	jumlahTelat    		int
	jumlahHari			int
	jumlahRusak			int
}

func Anggota(nama string, nonAnggota func(string) bool) {
	if nonAnggota(nama) {
		fmt.Printf("Mohon maaf, %s belum terdaftar sebagi anggota Perpustakaan Daerah. Harap mendaftar terlebih dahulu. \n", nama)
	}
}

func TampilData(){
	member1 := DataMember{
		nama: "Rizka",
		noAnggota: "1132",
		judulBuku: "Aroma Karsa, Supernova",
		jumlahPinjam: 2,
		jumlahTelat: 1,
		jumlahHari: 3,
		jumlahRusak: 1,
	}
	fmt.Println("Nama: ", member1.nama)
	fmt.Println("No. Anggota: ", member1.noAnggota)
	fmt.Println("Buku yang dipinjam: ", member1.judulBuku)
}

func Denda(jumlahTelat, jumlahHari, jumlahRusak int) (int, int) {
		const satuanDenda int = 10000
		const satuanRusak int = 50000
		bayarDenda := jumlahTelat * jumlahHari * satuanDenda
		bayarRusak := jumlahRusak * satuanRusak
		return bayarDenda, bayarRusak
}

func TotalDenda(bayarDenda, bayarRusak int) int {
	return bayarDenda + bayarRusak 
}
