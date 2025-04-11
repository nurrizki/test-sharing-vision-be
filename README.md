# 🧠 Test Sharing Vision BE

Project ini dibangun menggunakan **Golang**, dilengkapi dengan framework **Gin** untuk routing yang cepat dan efisien, serta **GORM** sebagai ORM untuk mempermudah interaksi dengan database.

---

## 🚀 Fitur

- Struktur project yang clean dan modular
- Menggunakan Gin sebagai web framework
- ORM GORM untuk pengelolaan database
- Konfigurasi berbasis file YAML
- Dukungan insert dummy data untuk setup awal

---

## 🛠️ Cara Menjalankan

### 1. Download Dependency

Jalankan perintah berikut untuk mendownload seluruh dependency yang diperlukan:

```bash
go mod tidy
```

### 2. Siapkan File Konfigurasi

Salin file konfigurasi contoh ke file yang akan digunakan:

```bash
cp ./config/config.example.yaml ./config/config.yaml
```

### 3. Jalankan Aplikasi (dengan Setup Dummy Data)

Untuk pertama kali menjalankan aplikasi dan mengisi dummy data ke database:

```bash
go run main.go --isSetupData true
```

### 4. Jalankan Aplikasi seperti Biasa

Setelah setup awal, kamu bisa menjalankan aplikasi tanpa flag:

```bash
go run main.go
```

## 📚 Dokumentasi API

Untuk dokumentasi lengkap, silakan kunjungi:  
[Dokumentasi API](https://documenter.getpostman.com/view/6711422/2sB2cYc141)
