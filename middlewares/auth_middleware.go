package middlewares

import (
	"net/http"
	"strings"

	"github.com/aris4p/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(config.GetEnv("JWT_SECRET", "secret_key"))

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil header Authorization dari request
		tokenString := c.GetHeader("Authorization")
		// jika token kosong, kembalikan respons 401 Unauthorized
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token is required",
			})
			c.Abort()
			return
		}
		// Hapus prefix "Bearer" dari token
		// Header biasanya berbentuk: "Bearer <token>"
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		// Buat struct untuk menampung klaim token
		claims := &jwt.RegisteredClaims{}
		// parse token dan verifikasi tanda tangan dengan jwtkey
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			// Kembalikan kunci rahasia untuk memverifikasi token
			return jwtKey, nil
		})

		// jika token tidak valid atau terjadi error saat parsing
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token",
			})
			c.Abort()
			return
		}

		// simpan klaim "sub" (Username) ke dalam context
		c.Set("username", claims.Subject)

		// lanjut ke handler berikutnya
		c.Next()

	}
}
