package controllers

import (
    "TaskBTPN/database"
    "TaskBTPN/models"
    "github.com/gin-gonic/gin"
    "net/http"
)

// CreatePhoto digunakan untuk membuat foto profil
func CreatePhoto(c *gin.Context) {
    var photo models.Photo
    if err := c.ShouldBindJSON(&photo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Periksa apakah pengguna telah masuk
    user, exists := c.Get("user")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Anda harus masuk terlebih dahulu"})
        return
    }

    // Tetapkan ID pengguna yang masuk ke foto profil
    photo.UserID = user.(*models.User).ID

    if err := database.DB.Create(&photo).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Foto profil berhasil diunggah"})
}

// GetPhotos digunakan untuk mengambil daftar foto profil
func GetPhotos(c *gin.Context) {
    var photos []models.Photo
    if err := database.DB.Find(&photos).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, photos)
}

// DeletePhoto digunakan untuk menghapus foto profil
func DeletePhoto(c *gin.Context) {
    photoID := c.Param("photoId")

    var photo models.Photo
    if err := database.DB.First(&photo, photoID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Foto tidak ditemukan"})
        return
    }

    // Periksa apakah pengguna yang meminta penghapusan adalah pemilik foto profil
    user, exists := c.Get("user")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Anda harus masuk terlebih dahulu"})
        return
    }
    if user.(*models.User).ID != photo.UserID {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Anda tidak memiliki izin untuk menghapus foto ini"})
        return
    }

    if err := database.DB.Delete(&photo).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Foto profil berhasil dihapus"})
}
