package usecase

import "wfa-meetup-api/internal/domain"

// เก็บ Repository ไว้ใช้งาน
type userUsecase struct {
    userRepo domain.UserRepository
}

// ฟังก์ชันสร้าง Usecase ตัวใหม่ โดยต้องรับ Repository เข้ามาด้วย
func NewUserUsecase(repo domain.UserRepository) domain.UserUsecase {
    return &userUsecase{userRepo: repo}
}

// ฟังก์ชัน Register ผูกกับ userUsecase
func (u *userUsecase) Register(user *domain.User) error {
    // ในอนาคตเราจะเขียนโค้ดตรวจสอบเงื่อนไขที่นี่ 
    // เช่น ตรวจสอบว่ารหัสผ่านยาวพอไหม, ทำการ Hash รหัสผ่าน
    // ตอนนี้ให้ส่งข้อมูลไปบันทึกที่ Repository เลย
    return u.userRepo.Create(user)
}