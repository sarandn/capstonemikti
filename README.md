# Capstone-Aplikasi 'Ticketing' menggunakan Golang

Proyek ini dibuat untuk untuk memenuhi syarat-syarat menyelesaikan kegiatan Study Independent MSIB batch 6 di mitra MIKTI. Selain itu, proyek ini juga bertujuan untuk memperdalam pemahaman menggunakan bahasa pemrograman Golang.

## Detail Proyek
Tema:

Aplikasi Ticketing Berbasis Web

Nama Aplikasi:

Depublic

Kelompok : 5

Deskripsi:

Platform ini merupakan tempat jual-beli tiket konser ataupun event. Platform ini membuka dan menyediakan berbagai jenis kategori kebutuhan. User yang mendaftarkan diri pada aplikasi ini dapat berperan buyer. Dalam hal ini, pengguna diharapkan dapat dengan mudah menemukan jadwal konser yang sesuai dengan kebutuhan mereka dan membayar tiket secara online. Selain itu, website ini juga diharapkan dapat menyediakan informasi yang akurat dan terkini tentang event yang sedang berlangsung.

Untuk detail lebih lanjut, dapat dilihat di [sini](https://docs.google.com/presentation/d/1Fg5eM2pDcXrN-cETuHndRRCExXjVBkF2KInIeNDFTqs/edit#slide=id.g248d5834739_0_11).
## Fitur-fitur 
- User-Service
- Event-Service
- Ticket-Service
- Order-Service
- Payment-Service

## Cara Menjalankan Proyek

- Unduh installer [Golang](https://golang.org/dl/) terlebih dahulu
- Setelah terunduh, jalankan installer, klik next hingga proses instalasi selesai. By default jika anda tidak merubah path pada saat instalasi, Go akan ter-install di ```C:\go```. 
- Path tersebut secara otomatis akan didaftarkan dalam ```PATH``` environment variable.
- Buka Command Prompt / CMD, eksekusi perintah berikut untuk mengecek versi Go.
```
go version
```
- Jika output adalah sama dengan versi Go yang ter-install, menandakan proses instalasi berhasil.
- Sering terjadi, command ```go version``` tidak bisa dijalankan meskipun instalasi sukses. Solusinya bisa dengan restart CMD (tutup CMD, kemudian buka lagi). Setelah itu coba jalankan ulang command di atas.
- Unduh dan install [PostgreSQL](https://www.postgresql.org/download/)
* Clone repository proyek ke lokal
```
git clone -b develop https://github.com/sarandn/capstonemikti.git
```
- Pindah ke direktori repositori lokal dengan command
```
cd capstonemikti
```

- Jalankan command ```go mod tidy``` untuk memvalidasi dependensi. Jika ada dependensi yang belum terunduh, maka akan otomatis diunduh.
- Unduh dan install [pgAdmin](https://www.pgadmin.org/download/)


## Langkah-langkah untuk Menjalankan Proyek

### 1. Persiapan Database

Buat database baru dengan nama `capstone_mikti`.

### 2. Konfigurasi File `.env`

Ganti password di file `.env` masing-masing layanan menggunakan password PostgreSQL pribadi.

### 3. Migrasi Data

Untuk setiap layanan, jalankan perintah migrasi data berikut:

#### Order Detail Service

```bash
cd order-detail-service
migrate -database "postgres://postgres:[PASSWORD]@localhost:5432/capstone_mikti?sslmode=disable" -path db/migrations up
```
Ganti `[PASSWORD]` dengan password PostgreSQL pribadi.

#### Order Service

```bash
cd order-service
migrate -database "postgres://postgres:[PASSWORD]@localhost:5432/capstone_mikti?sslmode=disable" -path db/migrations up
```
Ganti `[PASSWORD]` dengan password PostgreSQL pribadi.

#### Payment Service

```bash
cd payment-service
migrate -database "postgres://postgres:[PASSWORD]@localhost:5432/capstone_mikti?sslmode=disable" -path db/migrations up
```
Ganti `[PASSWORD]` dengan password PostgreSQL pribadi.

#### Ticket Service

```bash
cd ticket-service
migrate -database "postgres://postgres:[PASSWORD]@localhost:5432/capstone_mikti?sslmode=disable" -path db/migrations up
```
Ganti `[PASSWORD]` dengan password PostgreSQL pribadi.

#### Users Service

```bash
cd users-service
migrate -database "postgres://postgres:[PASSWORD]@localhost:5432/capstone_mikti?sslmode=disable" -path db/migrations up
```
Ganti `[PASSWORD]` dengan password PostgreSQL pribadi.

### 4. Menjalankan Layanan

Untuk setiap layanan, jalankan perintah berikut:

#### Order Detail Service

```bash
cd order-detail-service
Event-Service/cmd/go run main.go
```

#### Order Service

```bash
cd order-service
Order-Service/cmd/go run main.go
```

#### Payment Service

```bash
cd payment-service
Payment-Service/cmd/go run main.go
```

#### Ticket Service

```bash
cd ticket-service
Ticket-Service/cmd/go run main.go
```

#### Users Service

```bash
cd users-service
Users-Service/cmd/go run main.go
```

Gantilah `[PASSWORD]` dengan password PostgreSQL pribadi masing-masing untuk setiap layanan.

## Dokumentasi

- [Postman](https://depublic.postman.co/workspace/depublic-Workspace~ee94f749-a3c4-446c-8a87-b9f1b81e6d6b/collection/36476173-355ef619-c23c-4e1e-8e39-cd818f80304d?action=share&creator=36277884)
