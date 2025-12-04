# Go Clean Architecture & DDD Boilerplate

Boilerplate ini adalah fondasi kokoh untuk membangun aplikasi backend berbasis Go dengan praktik Clean Architecture dan Domain Driven Design (DDD) yang baik dan benar. Struktur proyek ini dirancang agar modular, scalable, dan mudah di-maintain, serta siap untuk kebutuhan produksi.

## ✨ Fitur

- **Arsitektur Bersih**: Clean Architecture & DDD, pemisahan domain, usecase, adapter, dan infrastruktur.
- **Observability**: Integrasi OpenTelemetry untuk tracing dan metrics.
- **Error Handling**: Penanganan error yang konsisten dan terstruktur.
- **Structured Logging**: Logging terstruktur menggunakan zerolog.
- **Form Validation**: Validasi input dengan go-playground/validator.
- **Unit Test Ready**: Mudah diintegrasikan dengan unit test dan mocking.
- **Environment Management**: Konfigurasi environment dengan godotenv.
- **Database & Migration**: PostgreSQL driver dan Goose untuk migrasi database.
- **Caching**: Redis integration untuk caching dan message broker.

## 📂 Struktur Proyek

```
/
├── cmd/
│   └── api/                # Entry point aplikasi (HTTP server, CLI, dsb)
├── internal/
│   ├── adapter/            # Adapter komunikasi eksternal (DB, API, dsb)
│   ├── domain/             # Entitas, value object, dan logika bisnis inti
│   ├── infra/              # Implementasi infrastruktur (DB, cache, dsb)
│   ├── usecase/            # Use case / application service (orchestrator bisnis)
│   └── app.go              # Inisialisasi aplikasi
├── pkg/
│   ├── utils/              # Utility/helper reusable lintas modul
│   └── .gitkeep            # Penanda folder kosong
├── .gitignore
├── Dockerfile
├── go.mod
├── go.sum
```

### Penjelasan Folder

- **cmd/**  
  Berisi entry point aplikasi, misal HTTP server atau CLI.
- **internal/adapter/**  
  Adapter untuk komunikasi dengan dunia luar seperti database, API eksternal, dsb.
- **internal/domain/**  
  Berisi entitas, value object, dan logika bisnis inti (core domain).
- **internal/infra/**  
  Implementasi infrastruktur seperti database, cache, dsb.
- **internal/usecase/**  
  Use case atau application service yang mengorkestrasi logika bisnis.
- **internal/app.go**  
  Inisialisasi dan konfigurasi utama aplikasi.
- **pkg/utils/**  
  Utility/helper yang dapat digunakan lintas modul.
- **pkg/.gitkeep**  
  Penanda agar folder tetap ada di repository meski kosong.

## ⚙️ Dependency Utama

- **OpenTelemetry**  
  Observability, tracing, dan metrics (`go.opentelemetry.io/otel`)
- **Zerolog**  
  Structured logging (`github.com/rs/zerolog`)
- **Validator**  
  Validasi input/form (`github.com/go-playground/validator/v10`)
- **UUID**  
  Unique identifier (`github.com/google/uuid`)
- **Godotenv**  
  Manajemen environment variable (`github.com/joho/godotenv`)
- **PostgreSQL Driver**  
  Database connection (`github.com/lib/pq`)
- **Goose**  
  Database migration (`github.com/pressly/goose/v3`)
- **Redis**  
  Caching/message broker (`github.com/redis/go-redis/v9`)

## 🚀 Memulai

### Prasyarat

- [Go](https://golang.org/) versi terbaru
- Database PostgreSQL (atau sesuaikan driver di `internal/infra`)
- Redis (optional, untuk caching/message broker)

### Instalasi

1. **Clone repository ini:**
    ```bash
    git clone <URL_REPO_KAMU>
    cd goboilerplate-domain-driven
    ```

2. **Install dependency:**
    ```bash
    go mod tidy
    ```

3. **Siapkan Environment Variables:**
    - Buat file `.env` sesuai kebutuhan, contoh bisa lihat di dokumentasi masing-masing dependency.

---

## ⚙️ Konfigurasi Environment (.env)

Aplikasi ini menggunakan file `.env` untuk konfigurasi environment. Berikut contoh dan penjelasan tiap variabel:

```env
APP_MODE="release"                # Mode aplikasi: "release" atau "debug"
APP_HOST="localhost"              # Host aplikasi
APP_PORT=8086                     # Port aplikasi

DB_HOST="localhost"               # Host database
DB_USER="postgres"                # Username database
DB_PASSWORD="1sampai8"            # Password database
DB_NAMES="pos"                    # Nama database
DB_SSL_MODE="disable"             # SSL mode database (biasanya "disable" untuk lokal)
DB_DRIVER="postgres"              # Driver database yang digunakan
DB_MIGRATION_PATH="./internal/infra/migrations"  # Path migrasi database

# Observability (OpenTelemetry)
# Kosongkan untuk menonaktifkan OTLP, atau set "stdout" untuk menggunakan stdout exporter.
OTEL_TRACER_MODE="otlp"           # Mode tracer: "otlp", "stdout", atau kosong
OTEL_TRACER_OTLP_ENDPOINT="localhost:4319"   # Endpoint OTLP untuk tracing
OTEL_MATRIC_MODE="otlp"           # Mode metric: "otlp", "stdout", atau kosong
OTEL_MATRIC_OTLP_ENDPOINT="localhost:4319"   # Endpoint OTLP untuk metrics
```

**Keterangan:**
- `APP_MODE`, `APP_HOST`, `APP_PORT`: Konfigurasi mode dan alamat aplikasi.
- `DB_*`: Konfigurasi koneksi database PostgreSQL.
- `DB_MIGRATION_PATH`: Path folder migrasi untuk Goose.
- `OTEL_*`: Konfigurasi observability dengan OpenTelemetry.  
  - Kosongkan untuk menonaktifkan OTLP.
  - Set ke "stdout" untuk menggunakan stdout exporter.
  - Endpoint biasanya mengarah ke OTEL Collector.

Pastikan semua variabel sudah diisi sesuai kebutuhan environment kamu (development, staging, production).

4. **Migrasi akan dijalankan otomatis pada startup time:**

### Menjalankan Aplikasi

```bash
go run cmd/api/main.go
```

---

## 🗄️ Migrasi Database dengan Goose

Untuk melakukan migrasi database, kamu wajib menginstall Goose terlebih dahulu:

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

Setelah Goose terinstall, kamu bisa menggunakan command `goose` untuk membuat dan menjalankan migrasi.

### Membuat File Migrasi Baru

```bash
goose -dir internal/infra/migrations/postgres create goods_table sql
```

Perintah di atas akan membuat file migrasi SQL baru untuk tabel `goods_table` di folder `internal/infra/migrations/postgres`.

### Menjalankan Migrasi

setelah migrasi dibuat dan di config, maka ketika menjalankan startup migrasi otomatis akan migrate

## 🧪 Testing

Struktur kode mendukung unit test dan mocking, sehingga kamu dapat menulis dan menjalankan test dengan mudah.

### Menjalankan Unit Test

Untuk menjalankan seluruh unit test di proyek, gunakan perintah berikut:

```bash
go test ./...
```

Perintah ini akan menjalankan semua test di seluruh folder.

### Melihat Test Coverage

Untuk mengetahui seberapa banyak kode yang sudah tercover oleh test, jalankan:

```bash
go test ./... -cover
```

Jika ingin melihat laporan coverage dalam bentuk file HTML yang lebih detail, gunakan:

```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

Laporan coverage akan terbuka di browser dan menampilkan bagian mana saja dari kode yang sudah teruji.

### Best Practice

- Simpan file test di folder yang sama dengan kode yang diuji, dengan akhiran `_test.go`.
- Gunakan mocking untuk dependency eksternal agar test lebih reliable.
- Pastikan setiap fitur utama memiliki unit test dan, jika perlu, integration test.


## 🔒 Praktik Keamanan

- Validasi input dan error handling yang aman.
- Logging terstruktur tanpa membocorkan data sensitif.
- Konfigurasi environment yang terpisah untuk development, testing, dan production.

## 🤝 Kontribusi

Silakan kontribusi dengan menambah fitur, memperbaiki bug, atau meningkatkan dokumentasi. Pastikan mengikuti prinsip Clean Architecture dan DDD.

## 📝 Lisensi

Proyek ini dilisensikan di bawah MIT License.
