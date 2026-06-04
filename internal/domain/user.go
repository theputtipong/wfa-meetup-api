// ประกาศว่าไฟล์นี้อยู่ใน package ชื่อ domain
package domain

// สร้างโครงสร้างข้อมูล User (เสมือนตารางใน Database)
type User struct {
    // ID เป็นตัวเลข (uint) เป็น Primary Key ของตาราง
    ID       uint   `json:"id" gorm:"primaryKey"`
    // Username เป็นข้อความ (string) และห้ามซ้ำกัน (unique)
    Username string `json:"username" gorm:"unique"`
    // Email ของผู้ใช้
    Email    string `json:"email"`
}

// UserRepository คือข้อตกลง (Interface) สำหรับคุยกับ Database
// ระบบไม่สนว่าจะใช้ Postgres หรือ Mongo สนแค่ว่าต้องมีฟังก์ชัน Create รับข้อมูล User เข้าไป
type UserRepository interface {
    Create(user *User) error
}

// UserUsecase คือข้อตกลงสำหรับ Business Logic (กฎของธุรกิจ)
type UserUsecase interface {
    Register(user *User) error
}
