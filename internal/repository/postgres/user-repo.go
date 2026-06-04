package postgres

import (
    // นำเข้า package domain เพื่อใช้ Struct และ Interface
    "wfa-meetup-api/internal/domain"
    // นำเข้า gorm สำหรับต่อ Database
    "gorm.io/gorm"
)

// สร้าง struct ซ่อนไว้ใช้งานภายใน (สังเกตตัวพิมพ์เล็ก) เพื่อเก็บการเชื่อมต่อ DB
type userRepo struct {
    db *gorm.DB
}

// ฟังก์ชันสร้าง Repository ตัวใหม่ ส่งคืนค่าเป็น Interface ของ Domain
// เสมือนโรงงานผลิตเครื่องมือคุยกับ DB
func NewUserRepository(db *gorm.DB) domain.UserRepository {
    return &userRepo{db: db} // ส่งเครื่องมือกลับไปให้ระบบใช้
}

// นำฟังก์ชัน Create มาผูกกับ userRepo เพื่อให้ตรงตามสัญญาใน Interface
func (r *userRepo) Create(user *domain.User) error {
    // ใช้คำสั่ง r.db.Create ของ GORM เพื่อบันทึกข้อมูลลง Database
    // .Error จะส่งค่า error กลับไปถ้าบันทึกไม่สำเร็จ (เช่น database ล่ม)
    return r.db.Create(user).Error
}