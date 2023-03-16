package main

import "fmt"

type user struct {
	ID			int
	Username	string // Mengubah tipe data dari int menjadi string
	Paswoord	string // Mengubah tipe data dari int menjadi string
}

type userService struct {
	t []user
}

func (u *userService) GetAllUsers() []user { // Mengubah parameter dan pointer receiver
	return u.t
}

func (u *userService) GetUserByID(id int) *user {
	for i, r := range u.t {
		if id == r.ID {
			return &u.t[i] // Mengembalikan pointer ke alamat memori variabel user pada slice u.t
		}
	}

	return nil // Mengembalikan nil jika tidak ditemukan user dengan id tertentu
}

func main() {
	// Inisialisasi objek userService
	user := userService{
		t: []user{
			{ID: 1, Username: "user1", Password: "password1"},
			{ID: 2, Username: "user2", Password: "password2"},
			{ID: 3, Username: "user3", Password: "password3"},
		},
	}

	// Memanggil method GetAllUsers() untuk mendapatkan semua data user
	allUsers := users.GetAllUsers()
	fmt.Println("All users:")
	for _, u:= range allUsers {
		fmt.Println(u)
	}

	// Memanggil method GetUserByID() untuk mendapatkan data user dengan ID tertentu
	id := 2
	userByID := users.GetUserByID(id)
	if userByID != nil {
		fmt.Printf("User dengan ID %d:\n%v\n", id, *userByID)
	} else {
		fmt.Printf(User dengan ID %d tidak ditemukan\n", id)
	}
}

// Kekurangan terjadi pada bagian-bagian berikut:
// Tidak mendeklarasikan func main sehingga tidak memiliki output
// Tipe data field username dan password seharusnya berupa string, bukan int.
// Nama field username dan password seharusnya diawali huruf besar agar bisa diakses dari luar paket.
// Parameter u pada method getallusers dan getuserbyid seharusnya bertipe pointer (*userservice) agar dapat mengubah
// Method getallusers dan getuserbyid seharusnya ditulis dengan menggunakan pointer receiver (func (u *userservice).
// Nilai balik pada method getuserbyid seharusnya bertipe pointer {*user) agar bisa diketahui apakah nilai baliknya