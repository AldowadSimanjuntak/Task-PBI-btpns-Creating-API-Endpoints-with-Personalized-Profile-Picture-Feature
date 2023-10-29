package database

import (
    "TaskBTPN/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "TaskBTPN/config" // Impor file konfigurasi
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
    db, err := gorm.Open(postgres.Open(config.DatabaseURL), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    // Auto Migrate akan membuat tabel berdasarkan model struktur
    db.AutoMigrate(&models.User{}, &models.Photo{})

    DB = db

    return db, nil
}
