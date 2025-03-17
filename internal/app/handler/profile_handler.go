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

	// สร้าง Response struct โดยใช้ข้อมูลจาก profile
	// หาก profile เป็น struct และต้องการนับข้อมูลในฟิลด์ต่างๆ ให้เปลี่ยนเป็น slice หรือ array ที่เหมาะสม
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
