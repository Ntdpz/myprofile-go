package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"

	"myprofile-go/handler"
)

func main() {
	// สร้าง handler ของ CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},                   // อนุญาตให้เข้าถึงจาก localhost:3000
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // อนุญาต method ที่ต้องการ
		AllowedHeaders: []string{"Content-Type", "Authorization"},           // อนุญาต header ที่ต้องการ
	})

	// ตั้งค่า handler
	http.HandleFunc("/getMe", handler.GetMeHandler)
	http.HandleFunc("/", handler.StartHandler)

	// รันเซิร์ฟเวอร์
	log.Fatal(http.ListenAndServe(":8080", corsHandler.Handler(http.DefaultServeMux)))
}
