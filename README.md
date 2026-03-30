# Inventory Management Dashboard

Aplikasi inventory dashboard MVP untuk menampilkan data stok produk dari PostgreSQL atau Supabase menggunakan backend Go + Echo dan frontend plain HTML, CSS, dan JavaScript.

## Struktur Folder

```text
app         bootstrap aplikasi dan handler HTTP
repo        akses data ke PostgreSQL/Supabase
service     business logic
database    konfigurasi environment dan koneksi database
docs        dokumentasi Swagger
models      response dan model domain
frontend    halaman dashboard statis (HTML, CSS, JS)
sql         schema dan seed data
```

## Fitur

- `GET /api/v1/products` untuk mengambil data inventory
- `GET /healthz` untuk health check backend
- `GET /swagger/index.html` untuk dokumentasi API
- Dashboard frontend di `/` yang menampilkan product name, stock count, dan last updated

## Cara Menjalankan

1. Jalankan SQL di `sql/schema.sql` pada PostgreSQL atau Supabase.
2. Copy `.env.example` menjadi `.env` lalu isi koneksi database.
3. Install dependency Go:

```bash
go mod tidy
```

4. Jalankan server:

```bash
go run ./app
```

5. Buka dashboard di browser:

```text
http://localhost:8080
```

## Environment Variables

```env
PORT=8080
DATABASE_URL=postgres://...
SUPABASE_DB_URL=postgres://...
ALLOWED_ORIGIN=http://localhost:3000
```

## Endpoint

- `GET /api/v1/products`
- `GET /healthz`
- `GET /swagger/index.html`
