package utils

import (
	"image"
	"mime/multipart"

	"github.com/SA-TailorStore/Kanok-API/domain/exceptions"
	"github.com/gofiber/fiber/v2"
	"github.com/liyue201/goqr"
)

func OpenFile(c *fiber.Ctx) (multipart.File, error) {

	// Pull form file
	fileHeader, err := c.FormFile("image")
	if err != nil {
		return nil, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Failed to get file",
			"status":  "400",
			"message": err.Error(),
		})
	}

	// Open File
	file, err := fileHeader.Open()
	if err != nil {
		return nil, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Failed to open file",
			"status":  "400",
			"message": err.Error(),
		})
	}
	defer file.Close()

	return file, nil
}

func DecodeImage(file multipart.File) (image.Image, error) {
	if file != nil {
		img, _, err := image.Decode(file)
		if err != nil {
			return nil, exceptions.ErrInvalidImage
		}

		return img, nil
	} else {
		return nil, exceptions.ErrNoImage
	}
}

func ReadQRCode(img image.Image) ([]*goqr.QRData, error) {
	codes, err := goqr.Recognize(img)
	if err != nil {
		return nil, err
	}
	return codes, nil
}
