package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	connectToDatabase()
	defer db.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n===== MENU =====")
		fmt.Println("1. Tambah Buku")
		fmt.Println("2. Tampilkan Daftar Buku")
		fmt.Println("3. Ubah Buku")
		fmt.Println("4. Hapus Buku")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih menu (1-5): ")

		input, _ := reader.ReadString('\n')
		menu := strings.TrimSpace(input)

		switch menu {
		case "1":
			tambahBuku(reader)
		case "2":
			tampilkanDaftarBuku()
		case "3":
			ubahBuku(reader)
		case "4":
			hapusBuku(reader)
		case "5":
			fmt.Println("Terima kasih!")
			return
		default:
			fmt.Println("Menu tidak valid!")
		}
	}
}

func tambahBuku(reader *bufio.Reader) {
	fmt.Println("Jenis Buku:")
	fmt.Println("1. Non Fiksi")
	fmt.Println("2. Fiksi")
	fmt.Print("Pilihan: ")
	jenisInput, _ := reader.ReadString('\n')
	jenis := strings.TrimSpace(jenisInput)

	fmt.Print("ID Buku: ")
	id, _ := reader.ReadString('\n')
	fmt.Print("Judul: ")
	judul, _ := reader.ReadString('\n')
	fmt.Print("Pengarang: ")
	pengarang, _ := reader.ReadString('\n')
	fmt.Print("Tahun Terbit: ")
	tahunInput, _ := reader.ReadString('\n')
	tahunTerbit, _ := strconv.Atoi(strings.TrimSpace(tahunInput))
	fmt.Print("Harga: Rp")
	hargaInput, _ := reader.ReadString('\n')
	harga, _ := strconv.ParseFloat(strings.TrimSpace(hargaInput), 64)

	var subJenis, genre, subjek string
	if jenis == "1" {
		fmt.Print("Subjek: ")
		subjek, _ = reader.ReadString('\n')
		jenis = "Non Fiksi"
	} else if jenis == "2" {
		fmt.Print("Sub Jenis: ")
		subJenis, _ = reader.ReadString('\n')
		fmt.Print("Genre: ")
		genre, _ = reader.ReadString('\n')
		jenis = "Fiksi"
	} else {
		fmt.Println("Jenis tidak valid!")
		return
	}

	sql := "INSERT INTO buku (id_buku, judul, pengarang, tahun_terbit, harga, jenis, sub_jenis, genre, subjek) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, err := db.Exec(sql, strings.TrimSpace(id), strings.TrimSpace(judul), strings.TrimSpace(pengarang), tahunTerbit, harga, jenis, strings.TrimSpace(subJenis), strings.TrimSpace(genre), strings.TrimSpace(subjek))
	if err != nil {
		fmt.Printf("Error adding book: %v\n", err)
		return
	}
	fmt.Println("Buku berhasil ditambahkan.")
}

func tampilkanDaftarBuku() {
	sql := "SELECT * FROM buku"
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Printf("Error fetching books: %v\n", err)
		return
	}
	defer rows.Close()

	fmt.Printf("\n%-10s | %-40s | %-20s | %-10s | %-15s | %-20s\n", "ID", "Judul", "Pengarang", "Tahun", "Harga", "Detail")
	fmt.Println(strings.Repeat("=", 110))

	for rows.Next() {
		var b Buku
		err := rows.Scan(&b.ID, &b.Judul, &b.Pengarang, &b.TahunTerbit, &b.Harga, &b.Jenis, &b.SubJenis, &b.Genre, &b.Subjek)
		if err != nil {
			fmt.Printf("Error reading row: %v\n", err)
			return
		}
		if b.Jenis == "Non Fiksi" {
			fmt.Printf("%-10s | %-40s | %-20s | %-10d | Rp%-15.2f | %-20s\n", b.ID, b.Judul, b.Pengarang, b.TahunTerbit, b.Harga, b.Subjek)
		} else {
			fmt.Printf("%-10s | %-40s | %-20s | %-10d | Rp%-15.2f | %-20s: %s\n", b.ID, b.Judul, b.Pengarang, b.TahunTerbit, b.Harga, b.SubJenis, b.Genre)
		}
	}
}

func ubahBuku(reader *bufio.Reader) {
	fmt.Print("Masukkan ID buku yang ingin diubah: ")
	id, _ := reader.ReadString('\n')
	id = strings.TrimSpace(id)

	sql := "SELECT * FROM buku WHERE id_buku = ?"
	row := db.QueryRow(sql, id)

	var b Buku
	err := row.Scan(&b.ID, &b.Judul, &b.Pengarang, &b.TahunTerbit, &b.Harga, &b.Jenis, &b.SubJenis, &b.Genre, &b.Subjek)
	if err != nil {
		fmt.Println("Buku tidak ditemukan.")
		return
	}

	fmt.Printf("Data lama: %s - %s - %s\n", b.ID, b.Judul, b.Pengarang)
	fmt.Print("Judul baru (kosong untuk tidak mengubah): ")
	judul, _ := reader.ReadString('\n')
	if strings.TrimSpace(judul) != "" {
		b.Judul = strings.TrimSpace(judul)
	}

	sqlUpdate := "UPDATE buku SET judul = ? WHERE id_buku = ?"
	_, err = db.Exec(sqlUpdate, b.Judul, b.ID)
	if err != nil {
		fmt.Printf("Error updating book: %v\n", err)
		return
	}
	fmt.Println("Buku berhasil diubah.")
}

func hapusBuku(reader *bufio.Reader) {
	fmt.Print("Masukkan ID buku yang ingin dihapus: ")
	id, _ := reader.ReadString('\n')
	id = strings.TrimSpace(id)

	sql := "DELETE FROM buku WHERE id_buku = ?"
	_, err := db.Exec(sql, id)
	if err != nil {
		fmt.Printf("Error deleting book: %v\n", err)
		return
	}
	fmt.Println("Buku berhasil dihapus.")
}
