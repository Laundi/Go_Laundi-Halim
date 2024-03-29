package main

/*
- Penamaan variabel totalroda -> totalRoda
- Penamaan variabel kecepatanperjam -> kecepatanPerJam
- Penamaan struct umumnya kapital kendaraan -> Kendaraan
*/
type Kendaraan struct {
	totalRoda		int
	kecepatanPerJam int
}
/*
- Penamaan struct umumnya kapital kendaraan -> Kendaraan
- Penamaan struct umumnya kapital mobil -> Mobil
*/
type Mobil struct {
	Kendaraan
}
/*
- Penamaan method tambahkecepatan -> tambahKecepatan
- Penamaan struct umumnya kapital mobil -> Mobil
- Penamaan receiver sebaiknya menggunakan kata dibanding huruf m -> mobil
*/
func (mobil *Mobil) berjalan() {
	mobil.tambahKecepatan(10)
}

/*
- Penamaan method tambahKecepatan -> tambahKecepatan
- Penamaan variable kecepatanbaru -> kecepatanBaru'
- Penamaan variable kecepatanperjam -> kecepatanPerJam
- Penamaan receiver sebaiknya menggunakan kata dibanding huruf m -> mobil
*/
func (mobil *Mobil) tambahKecepatan(kecepatanBaru int) {
	mobil.kecepatanPerJam = mobil.kecepatanPerJam + kecepatanBaru
}

func main() {
	// Penamaan variable mobilcepat -> mobilCepat
	mobilCepat := Mobil {}
	mobilCepat.berjalan()
	mobilCepat.berjalan()
	mobilCepat.berjalan()
	// Penamaan variabel mobillamban -> mobilLamban
	mobilLamban := Mobil{}
	mobilLamban.berjalan()
}
