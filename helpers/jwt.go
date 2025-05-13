package helpers

import (
	"time"

	"github.com/aris4p/config"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(config.GetEnv("JWT_SECRET", "secret_key"))

func GenerateToken(username string) string {
	// mengatur waktu kadaluwarsa token, disini kita set 60 menit dari waktu sekarang
	expirationTime := time.Now().Add(60 * time.Minute)

	// Membuat klaim (claims) JWT
	// Subject berisi username, dan ExpiresAt menentukan waktu expired token
	claims := &jwt.RegisteredClaims{
		Subject:   username,
		ExpiresAt: jwt.NewNumericDate((expirationTime)),
	}
	// membuat token baru dengan klaim yang telah dibuat
	// menggunakan algoritma HS256 untuk menandatangani token
	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtKey)

	// mengembalikan token dalam bentuk string
	return token
}
