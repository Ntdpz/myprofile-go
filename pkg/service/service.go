package service

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

// GetImageBase64 แปลงไฟล์ภาพเป็น base64
func GetImageBase64() (string, error) {
	// ใช้เส้นทางเต็ม
	fullPath := "C:/Users/nnnon/Music/internship/git/myprofile-go/images/profile.jpg" // แก้ไขเส้นทางตรงนี้

	// เปิดไฟล์ภาพ
	file, err := os.Open(fullPath)
	if err != nil {
		return "", fmt.Errorf("could not read image file: %v", err)
	}
	defer file.Close()

	// อ่านข้อมูลไฟล์
	fileData, err := ioutil.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("could not read file data: %v", err)
	}

	// แปลงข้อมูลไฟล์เป็น base64
	return base64.StdEncoding.EncodeToString(fileData), nil
}
