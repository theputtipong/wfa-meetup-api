package ws

import (
	"sync"

	"github.com/gofiber/contrib/websocket"
)

// ChatHub เป็นตัวเก็บข้อมูลว่าแต่ละห้อง (MeetupID) มีใครเชื่อมต่ออยู่บ้าง
type ChatHub struct {
	// Rooms รูปแบบคือ map[meetupID]map[connection]boolean
	Rooms map[string]map[*websocket.Conn]bool
	mu    sync.RWMutex // ใช้ป้องกันปัญหาเวลาคนเชื่อมต่อเข้ามาพร้อมๆ กัน (Concurrency)
}

func NewChatHub() *ChatHub {
	return &ChatHub{
		Rooms: make(map[string]map[*websocket.Conn]bool),
	}
}

// Register เอาไว้ลงทะเบียนเวลาคนเข้าห้องแชท
func (h *ChatHub) Register(meetupID string, conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.Rooms[meetupID] == nil {
		h.Rooms[meetupID] = make(map[*websocket.Conn]bool)
	}
	h.Rooms[meetupID][conn] = true
}

// Unregister เอาไว้ลบข้อมูลเวลาคนออกห้องแชท (ปิดเน็ต/ปิดหน้าเว็บ)
func (h *ChatHub) Unregister(meetupID string, conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if _, ok := h.Rooms[meetupID]; ok {
		delete(h.Rooms[meetupID], conn)
		// ถ้าห้องว่างแล้ว ก็ลบห้องทิ้งไปเลย คืน Memory ให้ระบบ
		if len(h.Rooms[meetupID]) == 0 {
			delete(h.Rooms, meetupID)
		}
	}
}