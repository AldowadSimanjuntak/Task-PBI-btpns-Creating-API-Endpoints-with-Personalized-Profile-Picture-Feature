package controllers

import (
    "TaskBTPN/database"
    "TaskBTPN/models"
    "github.com/gin-gonic/gin"
    "net/http"
)

// Register digunakan untuk mendaftarkan pengguna
func Register(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Hash kata sandi sebelum menyimpan ke database
    err := database.DB.Create(&user).Error
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Pengguna berhasil terdaftar"})
}

// Login digunakan untuk login pengguna
func Login(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var existingUser models.User
    if err := database.DB.Where("email = ?", user.Email).First(&existingUser).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Email atau kata sandi salah"})
        return
    }

    // Periksa apakah kata sandi cocok
    if !existingUser.CheckPassword(user.Password) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Email atau kata sandi salah"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Berhasil login"})
}
