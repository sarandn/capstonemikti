package main

import (
    "log"
    "payment-service/internal/domain/model"
    "payment-service/internal/infra/repository"
    "payment-service/interfaces/api"

    "github.com/gin-gonic/gin"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func main() {
    // Ganti dengan DSN PostgreSQL Anda
    dsn := "host=localhost user=youruser password=yourpassword dbname=yourdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Gagal terhubung ke database: ", err)
    }

    // Migrate schema
    db.AutoMigrate(&model.Payment{})

    // Inisialisasi repository dan handler
    paymentRepo := repository.NewPaymentRepository(db)
    paymentHandler := api.NewPaymentHandler(paymentRepo)

    // Setup router
    r := gin.Default()

    // Rute dengan middleware
    r.POST("/payments", paymentHandler.CreatePayment)
    r.GET("/payments/:id", paymentHandler.GetPayment)
    r.PUT("/payments/:id", paymentHandler.UpdatePayment)
    r.DELETE("/payments/:id", paymentHandler.DeletePayment)

    // Jalankan server
    r.Run(":8080")
}