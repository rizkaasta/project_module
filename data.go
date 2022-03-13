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


func (member dataMember) Ucapan() {
	fmt.Println("Selamat datang kembali di Prpustakaan Daerah, ", member.nama)
}

func TampilData(){
	member1 := dataMember{
		nama: "Rizka",
		noAnggota: "1132",
		judulBuku: "Aroma Karsa, Supernova",
		jumlahPinjam: 2,
		jumlahTelat: 1,
		jumlahHari: 3,
		jumlahRusak: 1,
	}

	fmt.Println("Nama: ", member1.nama)
	
}
func Denda(jumlahTelat, jumlahHari, jumlahRusak int) (int, int) {
		const satuanDenda int = 10000
		const satuanRusak int = 50000
		bayarDenda := jumlahTelat * jumlahHari * satuanDenda
		bayarRusak := jumlahRusak * satuanRusak
		return bayarDenda, bayarRusak
}
