# Inventory Management Dashboard Backend

Backend API menggunakan Go, Echo, PostgreSQL/Supabase, `godotenv`, dan Swagger UI dari Swaggo.

## Struktur Folder

```text
app         bootstrap aplikasi, handler HTTP, main.go
repo        akses data ke PostgreSQL/Supabase
service     business logic
database    konfigurasi environment dan koneksi database
docs        dokumen Swagger untuk Swaggo
models      response dan model domain
sql         schema dan seed data
```

## Fitur

- `GET /api/v1/products` untuk mengambil data inventory
- `GET /healthz` untuk health check
- `GET /swagger/index.html` untuk dokumentasi API

## Cara Menjalankan

1. Jalankan SQL di `sql/schema.sql` pada PostgreSQL atau Supabase.
2. Copy `.env.example` menjadi `.env` lalu isi koneksi database.
3. Jalankan dependency install:

```bash
go mod tidy
```

4. Jalankan server:

```bash
go run ./app
```

## Environment Variables

```env
PORT=8080
DATABASE_URL=postgres://...
SUPABASE_DB_URL=postgres://...
ALLOWED_ORIGIN=http://localhost:3000
```
