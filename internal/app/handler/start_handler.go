package handler

import (
	"fmt"
	"net/http"
)

// StartHandler สำหรับการตอบกลับข้อความ "Go my profile service start"
func StartHandler(w http.ResponseWriter, r *http.Request) {
	// กำหนด headers เป็นข้อความธรรมดา
	w.Header().Set("Content-Type", "text/plain")

	// ส่งข้อความกลับไปยัง client
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Go my profile service start")
}
