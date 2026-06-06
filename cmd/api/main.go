package main

import (
    // 1. Import Package ของ Redis
	"github.com/redis/go-redis/v9"
    "github.com/gofiber/fiber/v2"
    "gorm.io/gorm"
    "log"

    // 2. Import โฟลเดอร์ต่างๆ ในโปรเจกต์ของเรา
	"wfa-meetup-api/internal/usecase"
	"wfa-meetup-api/pkg/googlemaps"
    "wfa-meetup-api/internal/delivery/http"
    "wfa-meetup-api/internal/domain"
    "wfa-meetup-api/internal/repository/postgres"

    // ตั้งชื่อเล่น (Alias) ให้บาง Package เพื่อไม่ให้ชื่อซ้ำกับระบบ หรือจำง่ายขึ้น
    gorm_postgres "gorm.io/driver/postgres"
	http_delivery "wfa-meetup-api/internal/delivery/http"
	redis_repo "wfa-meetup-api/internal/repository/redis"
    ws_delivery "wfa-meetup-api/internal/delivery/ws"
)

func main() {
    // 1. ตั้งค่าการเชื่อมต่อ PostgreSQL 
    // (ตอนนี้เราใช้รหัสผ่านตามที่เราตั้งไว้ใน docker-compose.yml ก่อนหน้านี้)
    dsn := "host=localhost user=user password=password dbname=wfa_meetup port=5432 sslmode=disable"
    
    // เปิดการเชื่อมต่อ Database
    db, err := gorm.Open(gorm_postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("เชื่อมต่อ Database ไม่สำเร็จ:", err)
    }

    // 2. สั่ง Auto Migrate เพื่อสร้างตารางใน DB ตาม Struct ของเราอัตโนมัติ
    db.AutoMigrate(&domain.User{})

    // 3. สร้าง Fiber App (จำลอง Web Server)
    app := fiber.New()

    // 1. เชื่อมต่อ Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // ไม่มีรหัสผ่านใน default config
		DB:       0,
	})

	// 2. สร้าง Components ของ Cafe
	cafeCacheRepo := redis_repo.NewCafeCacheRepository(redisClient)
	googleMapsSvc := googlemaps.NewGoogleMapsClient("MOCK_API_KEY") 
    
    // 4. ประกอบร่าง (Dependency Injection)
    // สร้าง Repository
    userRepo := postgres.NewUserRepository(db)
    // เอา Repository ไปใส่ใน Usecase
    // เอา Usecase ไปผูกกับ Handler (Fiber Route)
    userUsecase := usecase.NewUserUsecase(userRepo)
	// ตรงนี้แหละครับที่เราเรียกใช้ usecase
	cafeUsecase := usecase.NewCafeUsecase(cafeCacheRepo, googleMapsSvc)

	// 3. ผูกเข้ากับ Fiber App (Handler)
    http.NewUserHandler(app, userUsecase)
	http_delivery.NewCafeHandler(app, cafeUsecase)
    // --- Phase 3: Setup WebSockets ---
	chatHub := ws_delivery.NewChatHub()
	ws_delivery.NewChatHandler(app, chatHub)

    
    // 5. สั่งให้ Server เริ่มทำงานที่ Port 3000
    log.Println("Server is running on port 3000...")
    app.Listen(":3000")
}