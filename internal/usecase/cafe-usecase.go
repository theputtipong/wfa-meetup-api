// ไฟล์: internal/usecase/cafe_usecase.go
package usecase

import (
	"context"
	"fmt"
	"wfa-meetup-api/internal/domain"
)

type cafeUsecase struct {
	cacheRepo   domain.CafeCacheRepository
	externalSvc domain.CafeExternalService
}

func NewCafeUsecase(cacheRepo domain.CafeCacheRepository, externalSvc domain.CafeExternalService) domain.CafeUsecase {
	return &cafeUsecase{
		cacheRepo:   cacheRepo,
		externalSvc: externalSvc,
	}
}

func (u *cafeUsecase) GetNearbyCafes(ctx context.Context, lat, lng string) ([]domain.Cafe, error) {
	// 1. ลองดึงจาก Redis ก่อน (Cache-Aside Pattern)
	cachedCafes, err := u.cacheRepo.GetNearbyCafes(ctx, lat, lng)
	if err == nil && cachedCafes != nil {
		// Cache Hit!
		fmt.Println("🚀 [Cache Hit] ดึงข้อมูลคาเฟ่จาก Redis เร็วปรื๊ด!")
		return cachedCafes, nil
	}

	// 2. ถ้าไม่เจอใน Cache (Cache Miss) ให้ไปเรียก Google Maps API
	fmt.Println("🐢 [Cache Miss] ไม่เจอใน Redis กำลังไปดึงจาก Google Maps API...")
	cafes, err := u.externalSvc.FetchCafesFromAPI(ctx, lat, lng)
	if err != nil {
		return nil, err
	}

	// 3. เอาข้อมูลใหม่ที่เพิ่งได้มา เซฟลง Redis ซะ รอบหน้าจะได้เร็ว
	err = u.cacheRepo.SetNearbyCafes(ctx, lat, lng, cafes)
	if err != nil {
		// Log error ไว้เฉยๆ ไม่ต้อง return error ให้ user เพราะถึง cache เซฟพัง เราก็ยังมีข้อมูลตอบกลับ
		fmt.Println("⚠️ [Redis Warning] ไม่สามารถเซฟ Cache ได้:", err)
	}

	return cafes, nil
}
