# Golang Backend

Ini adalah backend server yang dikembangkan menggunakan bahasa pemrograman **Golang**, dengan pendekatan **Clean Architecture** dan penyimpanan data **in-memory**.

Aplikasi ini digunakan untuk:
- Mengelola autentikasi pengguna (register, login, logout, update)
- Mengelola data item (CRUD item)
- Melindungi route menggunakan JWT Middleware
- Menyediakan endpoint REST API yang terintegrasi dengan frontend berbasis Next.js

🔗 Repo frontend: [frontend-app](https://github.com/dickysetiawan031000/frontend-app)

---

## 🔧 Fitur Utama

- ✅ Clean Architecture (Modular & scalable)
- ✅ RESTful API
- ✅ Register, Login, Logout, Update Profile
- ✅ Middleware JWT (Token-based Authentication)
- ✅ Proteksi token aktif via in-memory store (blacklist token saat logout)
- ✅ CRUD Item
- ✅ Validasi menggunakan DTO
- ✅ Penyimpanan in-memory (tanpa database)

---

## 🗂️ Struktur Folder

```
go-backend/
├── dto/             # Data Transfer Object
├── handler/         # HTTP handlers
├── mapper/          # Konversi antar model dan DTO
├── middleware/      # JWT Middleware dan Auth
├── model/           # Model data
├── repository/      # In-memory data store
├── usecase/         # Business logic
├── utils/           # Helper
└── main.go          # Entry point aplikasi
```

---

## 🚀 Cara Menjalankan

### 1. Clone Repository

```bash
git clone https://github.com/dickysetiawan031000/go-backend.git

cd go-backend
```

### 2. Run

```bash
go run main.go
```

Aplikasi akan berjalan pada: `http://localhost:8080`

---

## 🔐 Endpoint Autentikasi

- `POST /register` – Register pengguna baru
- `POST /login` – Login dan mendapatkan JWT
- `POST /logout` – Logout dan blacklist token
- `GET /profile` – Mendapatkan data pengguna (dengan auth)
- `PUT /profile` – Mengubah data pengguna (dengan auth)

---

## 📦 Endpoint Item

- `GET /items` – List semua item
- `GET /items/{id}` – Detail item
- `POST /items` – Tambah item
- `PUT /items/{id}` – Edit item
- `DELETE /items/{id}` – Hapus item

> Semua endpoint item dilindungi oleh JWT

📄 [Download Postman Collection (JSON)](https://pioneertech.postman.co/workspace/TechnicalTest-JDI~e3783dde-90f7-4699-8417-2e9e76e22f4c/collection/18467327-920d24e7-0e38-4609-b9fd-eccaeb742284?action=share&creator=18467327&active-environment=18467327-3c3f7638-13a2-4b03-b1ff-cdd729dbb549)


---

## ⚙️ Teknologi

- Golang
- Pendekatan Clean Architecture
- JWT
- In-Memory Storage

---


© 2025 - Dicky Setiawan
