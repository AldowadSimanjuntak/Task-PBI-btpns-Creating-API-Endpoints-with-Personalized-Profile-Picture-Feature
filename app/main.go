package main

import (
    "TaskBTPN/database"
    "TaskBTPN/router"
)

func main() {
    // Inisialisasi database
    database.InitDB()

    // Membuat router dan menjalankan aplikasi
    r := router.SetupRouter()
    r.Run(":8080")
}
