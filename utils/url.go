package utils

import (
	"fmt"
	"net/url"
	"path"
	"strings"
)

func ExtractPublicID(imageURL string) (string, error) {
	// Parse
	parsedURL, err := url.Parse(imageURL)
	if err != nil {
		return "", err
	}

	// ดึง path จาก URL เช่น /demo/image/upload/v1234567890/sample.jpg
	imagePath := parsedURL.Path

	// Split path
	parts := strings.Split(imagePath, "/")

	// ตรวจสอบว่า URL มีรูปแบบถูกต้องหรือไม่
	if len(parts) < 4 || parts[len(parts)-1] == "" {
		return "", fmt.Errorf("Invalid URL format")
	}

	// ดึงชื่อไฟล์ตัวสุดท้ายออกจาก path เช่น sample.jpg
	fileName := parts[len(parts)-1]

	// ลบส่วนที่เป็น extension เช่น .jpg ออกไป
	publicID := strings.TrimSuffix(fileName, path.Ext(fileName))

	return publicID, nil
}
