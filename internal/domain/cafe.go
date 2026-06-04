// ไฟล์: internal/domain/cafe.go
package domain

import "context"

// Cafe คือหน้าตาของข้อมูลคาเฟ่ที่เราจะใช้งาน
type Cafe struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Address  string  `json:"address"`
	Rating   float64 `json:"rating"`
	Lat      float64 `json:"lat"`
	Lng      float64 `json:"lng"`
}

// CafeCacheRepository กำหนดหน้าที่ของ Redis
type CafeCacheRepository interface {
	GetNearbyCafes(ctx context.Context, lat, lng string) ([]Cafe, error)
	SetNearbyCafes(ctx context.Context, lat, lng string, cafes []Cafe) error
}

// CafeExternalService กำหนดหน้าที่ของ Google Maps API (หรือ API อื่นๆ ในอนาคต)
type CafeExternalService interface {
	FetchCafesFromAPI(ctx context.Context, lat, lng string) ([]Cafe, error)
}

// CafeUsecase คือ Business Logic หลักของเรา
type CafeUsecase interface {
	GetNearbyCafes(ctx context.Context, lat, lng string) ([]Cafe, error)
}