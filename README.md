# Simple Payment RestAPI

Ini adalah proyek Simple Payment RestAPI yang dibangun menggunakan Gin di Golang.

## Deskripsi

Proyek ini adalah implementasi sederhana dari sistem pembayaran dengan menggunakan RESTful API yang dibangun menggunakan Gin, sebuah framework web yang ringan dan cepat untuk Golang.

## Fitur-fitur utama dari proyek ini:

Pembuatan dan manajemen akun pengguna
Pembuatan dan manajemen transaksi pembayaran
Otorisasi pengguna menggunakan token JWT
Validasi data input menggunakan middleware
Penggunaan database PostgreSQL untuk menyimpan data
Instalasi

## Setup postgres database

1. Create database with name simple_payment in postgres
2. Run query in ddl.sql from line 1 to line 41

## Instalasi

1. Pastikan Anda memiliki Go terinstal di sistem Anda. Untuk informasi lebih lanjut, kunjungi [dokumentasi](https://golang.org/doc/install).

2. Clone repositori ini ke dalam direktori lokal Anda:

```bash
git clone https://github.com/Rahmattullah13/Simple-Payment-RestAPI.git
```

3. Masuk ke direktori proyek:

```bash
cd Simple-Payment-RestAPI
```

4. Instal dependensi menggunakan Go modules:

```bash
go mod tidy
```

5. Konfigurasi file .env sesuai dengan preferensi Anda.

6. Jalankan aplikasi

```bash
go run main.go
```

Aplikasi akan dijalankan pada http://localhost:8080.

## Dokumentasi

- Unutk melihat dokumentasi bisa buka:

```bash
http://localhost:8080/swagger/index.html#
```

## Testing

- Untuk pengujian bisa Import simple-payment-go.postman_collection.json di postman
