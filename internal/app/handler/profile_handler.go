package handler

import (
	"encoding/json"
	"net/http"

	"myprofile-go/pkg/service"

)

// GetMeHandler สำหรับการตอบกลับข้อมูลโปรไฟล์
func GetMeHandler(w http.ResponseWriter, r *http.Request) {
	// ดึงข้อมูลจาก service
	profile := service.GetProfile()

	// กำหนด headers เป็น JSON
	w.Header().Set("Content-Type", "application/json")

	// ส่งข้อมูล JSON กลับไปยัง client
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(profile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
