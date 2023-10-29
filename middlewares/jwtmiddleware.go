package middlewares

import (
    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
    "net/http"
    "TaskBTPN/config" // Impor file konfigurasi
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Token JWT diperlukan"})
            c.Abort()
            return
        }

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            // Mengambil secret key dari konfigurasi
            return []byte(config.SecretKey), nil
        })

        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Token JWT tidak valid"})
            c.Abort()
            return
        }

        if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
            c.Set("user", claims)
        } else {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Token JWT tidak valid"})
            c.Abort()
            return
        }

        c.Next()
    }
}
