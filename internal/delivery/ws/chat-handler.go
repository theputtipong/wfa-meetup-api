
package ws

import (
	"encoding/json"
	"fmt"
	"log"
	"wfa-meetup-api/internal/domain"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/contrib/websocket"
)

type ChatHandler struct {
	Hub *ChatHub
}

func NewChatHandler(app *fiber.App, hub *ChatHub) {
	handler := &ChatHandler{
		Hub: hub,
	}

	// Middleware เช็คว่า Request นี้ต้องการอัปเกรดเป็น WebSocket หรือไม่
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	// Endpoint สำหรับเข้าห้องแชท (เช่น ws://localhost:3000/ws/meetup/123)
	app.Get("/ws/meetup/:meetup_id", websocket.New(handler.HandleChat))
}

func (h *ChatHandler) HandleChat(c *websocket.Conn) {
	meetupID := c.Params("meetup_id")
	
	// 1. ลงทะเบียนเข้าห้อง
	h.Hub.Register(meetupID, c)
	fmt.Printf("ผู้ใช้เชื่อมต่อเข้าห้อง: %s\n", meetupID)

	// 2. เมื่อฟังก์ชันนี้จบ (คนปิดหน้าเว็บ) ให้ลบชื่อออกจากห้อง
	defer func() {
		h.Hub.Unregister(meetupID, c)
		c.Close()
		fmt.Printf("ผู้ใช้ออกจากห้อง: %s\n", meetupID)
	}()

	// 3. วนลูปอ่านข้อความที่มีคนพิมพ์เข้ามาเรื่อยๆ
	for {
		messageType, payload, err := c.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break // ถ้ามี Error (เช่น สายหลุด) ให้ออกจากลูป
		}

		// แปลงข้อความ JSON ที่รับมา ให้อยู่ในรูปแบบ Struct
		var msg domain.ChatMessage
		if err := json.Unmarshal(payload, &msg); err != nil {
			log.Println("Invalid message format:", err)
			continue
		}

		// บังคับให้ MeetupID ตรงกับห้องที่เข้ามา
		msg.MeetupID = meetupID 
		fmt.Printf("[Room %s] %s: %s\n", msg.MeetupID, msg.UserID, msg.Text)

		// ส่งข้อความนี้กลับไปให้ "ทุกคน" ที่อยู่ในห้องเดียวกัน (Broadcast)
		h.Hub.mu.RLock()
		for conn := range h.Hub.Rooms[meetupID] {
			// พ่น JSON กลับไปให้ Client
			err := conn.WriteMessage(messageType, payload)
			if err != nil {
				log.Println("Error writing message:", err)
			}
		}
		h.Hub.mu.RUnlock()
	}
}