package http

import (
    "wfa-meetup-api/internal/domain"
    "github.com/gofiber/fiber/v2" // นำเข้า Fiber Framework
)

// เก็บ Usecase ไว้เรียกใช้งาน
type UserHandler struct {
    userUsecase domain.UserUsecase
}

// ฟังก์ชันสำหรับตั้งค่า API Route
func NewUserHandler(app *fiber.App, us domain.UserUsecase) {
    handler := &UserHandler{userUsecase: us}
    // สร้าง API เส้นทาง POST /users เมื่อเรียกมาให้ไปทำงานที่ฟังก์ชัน Register
    app.Post("/users", handler.Register)
}

// ฟังก์ชันรับ Request (Handler)
func (h *UserHandler) Register(c *fiber.Ctx) error {
    var user domain.User
    
    // 1. แปลงข้อมูล JSON จาก Client (Body) มาใส่ในตัวแปร user
    if err := c.BodyParser(&user); err != nil {
        // ถ้าแปลงไม่ได้ (Client ส่งข้อมูลผิด format) ให้ตอบกลับ Status 400 Bad Request
        return c.Status(400).JSON(fiber.Map{"error": "ข้อมูลไม่ถูกต้อง"})
    }
    
    // 2. ส่งข้อมูลไปให้ Usecase จัดการสมัครสมาชิก
    if err := h.userUsecase.Register(&user); err != nil {
        // ถ้าเกิด Error (เช่น Username ซ้ำ, DB ล่ม) ตอบกลับ Status 500
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
    
    // 3. ถ้าสำเร็จ ตอบกลับ Status 201 Created พร้อมข้อมูล User
    return c.Status(201).JSON(user)
}
