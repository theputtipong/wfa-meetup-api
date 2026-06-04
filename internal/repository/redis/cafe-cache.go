// ไฟล์: internal/repository/redis/cafe_cache.go
package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"wfa-meetup-api/internal/domain"

	"github.com/redis/go-redis/v9"
)

type cafeCacheRepo struct {
	client *redis.Client
}

func NewCafeCacheRepository(client *redis.Client) domain.CafeCacheRepository {
	return &cafeCacheRepo{client: client}
}

// สร้าง Key สำหรับ Redis เช่น cafes:13.75:100.50
func generateKey(lat, lng string) string {
	return fmt.Sprintf("cafes:%s:%s", lat, lng)
}

func (r *cafeCacheRepo) GetNearbyCafes(ctx context.Context, lat, lng string) ([]domain.Cafe, error) {
	key := generateKey(lat, lng)
	
	// ดึงข้อมูลจาก Redis
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil // ไม่เจอใน Cache (Cache Miss)
	} else if err != nil {
		return nil, err // เกิด Error อื่นๆ กับ Redis
	}

	// ถ้าเจอ แปลง JSON string กลับเป็น Struct (Cache Hit)
	var cafes []domain.Cafe
	err = json.Unmarshal([]byte(val), &cafes)
	if err != nil {
		return nil, err
	}

	return cafes, nil
}

func (r *cafeCacheRepo) SetNearbyCafes(ctx context.Context, lat, lng string, cafes []domain.Cafe) error {
	key := generateKey(lat, lng)
	
	// แปลง Struct เป็น JSON string
	data, err := json.Marshal(cafes)
	if err != nil {
		return err
	}

	// เซฟลง Redis และตั้งเวลาหมดอายุ (TTL) 10 นาที
	return r.client.Set(ctx, key, data, 10*time.Minute).Err()
}