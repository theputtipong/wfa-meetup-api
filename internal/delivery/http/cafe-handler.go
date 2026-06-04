// ไฟล์: internal/delivery/http/cafe_handler.go
package http

import (
	"wfa-meetup-api/internal/domain"

	"github.com/gofiber/fiber/v2"
)

type CafeHandler struct {
	CafeUsecase domain.CafeUsecase
}

func NewCafeHandler(app *fiber.App, uc domain.CafeUsecase) {
	handler := &CafeHandler{
		CafeUsecase: uc,
	}
	
	// ตั้ง Route สำหรับดึงข้อมูลคาเฟ่
	app.Get("/cafes/nearby", handler.GetNearbyCafes)
}

func (h *CafeHandler) GetNearbyCafes(c *fiber.Ctx) error {
	// รับ Query Parameters เช่น ?lat=13.75&lng=100.50
	lat := c.Query("lat")
	lng := c.Query("lng")

	if lat == "" || lng == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "กรุณาระบุ lat และ lng ให้ครบถ้วน",
		})
	}

	cafes, err := h.CafeUsecase.GetNearbyCafes(c.Context(), lat, lng)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": cafes,
	})
}
