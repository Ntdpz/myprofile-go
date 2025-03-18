package handler

import (
	"encoding/json"
	"net/http"

	"myprofile-go/pkg/models"
	"myprofile-go/pkg/service"
)

// Response สำหรับจัดการรูปแบบของ response ที่มี Count และ Data
type Response struct {
	Count int         `json:"count"`
	Data  interface{} `json:"data"`
}

// GetMeHandler สำหรับการตอบกลับข้อมูลโปรไฟล์
func GetMeHandler(w http.ResponseWriter, r *http.Request) {
	// ดึงข้อมูลจาก service
	profile := service.GetProfile()

	// แปลงภาพเป็น base64
	imageBase64, err := service.GetImageBase64() // ใช้ฟังก์ชันที่เราได้แก้ไข
	if err != nil {
		http.Error(w, "Failed to convert image to base64", http.StatusInternalServerError)
		return
	}

	// เพิ่มข้อมูล base64 ไปใน profile
	profile.ImageBase64 = imageBase64

	// สร้าง Response struct โดยใช้ข้อมูลจาก profile
	profileData := []models.Profile{profile} // แปลง profile เป็น slice
	response := Response{
		Count: len(profileData), // นับจำนวนข้อมูลใน slice
		Data:  profileData,      // ข้อมูลโปรไฟล์
	}

	// กำหนด headers เป็น JSON
	w.Header().Set("Content-Type", "application/json")

	// ส่งข้อมูล JSON กลับไปยัง client
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
