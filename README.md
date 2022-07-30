# Dependensi Aplikasi

- Go 1.17
- Gin Gonic (Golang Web Framework)[Kunjungi](https://gin-gonic.com/).
- GORM (ORM Library for golang)[Kunjungi](https://gorm.io/index.html).
- Postgres
- golang-jwt [Kunjungi](https://github.com/golang-jwt/jwt)

# Memulai
###### Instalasi

1. Rename .env.example menjadi .env
```
cp .env.example .env
```
2. Jalankan postgres database
3. Kemudian isilah variable env sesuai dengan environment kamu
- Lalu jalankan perintah berikut pada terminal kamu di direktori aplikasi
1. untuk mengintall dependensi aplikasi
```
make install
```
2. make run untuk menjalankan aplikasi
```
make run
```

###### Instalasi Menggunakan Docker

1. Pull postgres image terlebih dahulu, bisa skip step ini kalau sudah
```
docker pull postgres
```
2. Build docker image project ini
```
sudo docker build -t inventory-server:latest .
```
3. Rename .env.example menjadi .env
```
cp .env.example .env
```
4. Kemudian isilah variable env sesuai dengan environment kamu
5. Jalankan aplikasi dan database dengan docker compose
```
docker-compose up
```