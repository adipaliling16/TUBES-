package main

import (
	"fmt"
)

type User struct {
	ID       int
	Username string
	Password string
	Role     string
	Approved bool
}

type Product struct {
	ID    int
	Name  string
	Price float64
}

type Transaction struct {
	ID       int
	Buyer    User
	Product  Product
	Quantity int
	Total    float64
}

type Store struct {
	Owners       []User
	Products     []Product
	Users        []User
	Transactions []Transaction
}

func main() {
	var pilihan int
	var store Store

	// digunakan untuk mengecek siapa yang login
	store.Users = append(store.Users, User{ID: 1, Username: "admin", Password: "123", Role: "admin", Approved: true})
	for {
		fmt.Println("Menu :")
		fmt.Println("1. Daftar Akun")
		fmt.Println("2. Masuk")
		fmt.Println("3. Keluar")
		fmt.Print("Masukkan Pilihan: ")

		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			// Register user
			var username, pass, role string
			fmt.Print("Masukkan Username: ")
			fmt.Scan(&username)
			fmt.Print("Masukkan Password: ")
			fmt.Scan(&pass)
			fmt.Print("Sebagai? (pembeli/pemilik): ")
			fmt.Scan(&role)
			store.RegisterUser(username, pass, role)
		case 2:
			// Login user
			var username, pass string
			fmt.Print("Masukkan Username: ")
			fmt.Scan(&username)
			fmt.Print("Masukkan Password: ")
			fmt.Scan(&pass)
			user, found := store.LoginUser(username, pass)
			if found {
				store.UserMenu(user)
			} else {
				fmt.Println("Akun belom disetujui")
			}
		case 3:
			fmt.Println("Terimakasih!")
			return
		default:
			fmt.Println("Salah masukkan, Silahkan input ulang!")
		}
	}
}

func (s *Store) RegisterUser(username, password, role string) {
	id := len(s.Users) + 1
	user := User{ID: id, Username: username, Password: password, Role: role, Approved: false}
	s.Users = append(s.Users, user)
	fmt.Println("Pengguna", username, "terdaftar.")
	fmt.Println("Menunggu Admin Untuk Mensetujui")
}

func (s *Store) LoginUser(username, password string) (User, bool) {
	for _, user := range s.Users {
		if user.Username == username && user.Password == password && user.Approved {
			return user, true
		}
	}
	return User{}, false
}

func (s *Store) UserMenu(user User) {
	switch user.Role {
	case "admin":
		s.AdminMenu()
	case "pemilik":
		s.OwnerMenu(user)
	case "pembeli":
		s.BuyerMenu(user)
	}
}

func (s *Store) AdminMenu() {
	for {
		fmt.Println("Admin Menu :")
		fmt.Println("1. List Daftar Akun")
		fmt.Println("2. Persetujuan Akun")
		fmt.Println("3. Keluar")
		fmt.Print("Masukkan Pilihan: ")
		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			s.AdminListUsers()
		case 2:
			fmt.Print("Masukkan ID Pengguna yang Ingin disetujui : ")
			var userID int
			fmt.Scan(&userID)
			s.ApproveUser(userID)
		case 3:
			fmt.Println("Terimakasih!")
			return
		default:
			fmt.Println("Salah masukkan, Silahkan input ulang!")
		}
	}
}

func (s *Store) OwnerMenu(owner User) {
	for {
		fmt.Println("Owner Menu :")
		fmt.Println("1. Tambahkan Produk")
		fmt.Println("2. Edit Produk")
		fmt.Println("3. Hapus Produk")
		fmt.Println("4. Lihat Produk")
		fmt.Println("5. Lihat Transaksi")
		fmt.Println("6. Keluar")
		fmt.Print("Masukkan Pilihan: ")
		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			var namaProduk string
			var harga float64
			fmt.Print("Masukkan Nama produk: ")
			fmt.Scan(&namaProduk)
			fmt.Print("Masukkan Harga produk: ")
			fmt.Scan(&harga)
			s.AddProduct(namaProduk, harga)
		case 2:
			var productID int
			var harga float64
			var namaProduk string
			fmt.Print("Masukkan ID produk yang akan di edit: ")
			fmt.Scan(&productID)
			fmt.Println("Masukkan nama produk baru: ")
			fmt.Scan(&namaProduk)
			fmt.Print("Masukkan harga produk baru: ")
			fmt.Scan(&harga)
			s.EditProduct(productID, namaProduk, harga)
		case 3:
			fmt.Print("Masukkan ID produk yang ingin dihapus: ")
			var productID int
			fmt.Scan(&productID)
			s.DeleteProduct(productID)
		case 4:
			s.ListProducts()
		case 5:
			s.ListTransactions()
		case 6:
			fmt.Println("TERIMAKASIH!")
			return
		default:
			fmt.Println("Salah masukkan, Silahkan input ulang!")
		}
	}
}

func (s *Store) BuyerMenu(buyer User) {
	for {
		fmt.Println("Menu Pembeli: ")
		fmt.Println("1. Lihat produk")
		fmt.Println("2. Beli produk")
		fmt.Println("3. Keluar")
		fmt.Print("Masukkan Pilihan: ")
		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			s.ListProducts()
		case 2:
			var productID, jumlah int
			fmt.Print("Masukkan ID produk yang ingin dibeli: ")
			fmt.Scan(&productID)
			fmt.Print("Masukkan Jumlah Produk yang ingin dibeli: ")
			fmt.Scan(&jumlah)
			s.BuyProduct(buyer, productID, jumlah)
		case 3:
			fmt.Println("Terimakasih!")
			return
		default:
			fmt.Println("Salah masukkan, Silahkan input ulang!")
		}
	}
}

func (s *Store) ListProducts() {
	fmt.Println("Produk Tersedia:")
	for _, product := range s.Products {
		fmt.Printf("ID: %d, Nama: %s, Harga: Rp.%.2f\n", product.ID, product.Name, product.Price)
	}
}

func (s *Store) AddProduct(nama string, price float64) {
	id := len(s.Products) + 1
	product := Product{ID: id, Name: nama, Price: price}
	s.Products = append(s.Products, product)
	fmt.Println("Produk", nama, "ditambahkan.")
}

func (s *Store) EditProduct(productID int, nama string, harga float64) {
	for i, product := range s.Products {
		if product.ID == productID {
			s.Products[i].Name = nama
			s.Products[i].Price = harga
			fmt.Println("Produk", nama, "diedit.")
			return
		}
	}
	fmt.Println("Produk Tidak Ditemukan.")
}

func (s *Store) DeleteProduct(productID int) {
	for i, product := range s.Products {
		if product.ID == productID {
			s.Products = append(s.Products[:i], s.Products[i+1:]...)
			fmt.Println("Produk", product.Name, "dihapus.")

			// Agar ID selanjut nya terupdate ke ID baru
			for j := i; j < len(s.Products); j++ {
				s.Products[j].ID--
			}

			return
		}
	}
	fmt.Println("Produk Tidak ditemukan.")
}

func (s *Store) BuyProduct(buyer User, productID, quantity int) {
	for _, product := range s.Products {
		if product.ID == productID {
			total := product.Price * float64(quantity)
			transactionID := len(s.Transactions) + 1
			transaction := Transaction{
				ID:       transactionID,
				Buyer:    buyer,
				Product:  product,
				Quantity: quantity,
				Total:    total,
			}
			s.Transactions = append(s.Transactions, transaction)
			fmt.Println("Transaksi Berhasil.")
			return
		}
	}
	fmt.Println("Produk Tidak Ditemukan.")
}

func (s *Store) ListTransactions() {
	fmt.Println("Transaksi: ")
	for _, transaction := range s.Transactions {
		fmt.Printf("ID: %d, Pembeli: %s, Produk: %s, Jumlah: %d, Total: %.2f\n",
			transaction.ID, transaction.Buyer.Username, transaction.Product.Name,
			transaction.Quantity, transaction.Total)
	}
}

func (s *Store) AdminListUsers() {
	fmt.Println("List nama Pengguna:")
	for _, user := range s.Users {
		fmt.Printf("ID: %d, Username: %s, Role: %s, Status: %t\n",
			user.ID, user.Username, user.Role, user.Approved)
	}
}

func (s *Store) ApproveUser(userID int) {
	for i, user := range s.Users {
		if user.ID == userID {
			s.Users[i].Approved = true
			fmt.Println("Pengguna", user.Username, "Disetujui.")
			return
		}
	}
	fmt.Println("Pengguna tidak Ditemukan.")
}
