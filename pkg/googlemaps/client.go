// ไฟล์: pkg/googlemaps/client.go
package googlemaps

import (
	"context"
	"wfa-meetup-api/internal/domain"
)

type googleMapsClient struct {
	apiKey string
}

func NewGoogleMapsClient(apiKey string) domain.CafeExternalService {
	return &googleMapsClient{
		apiKey: apiKey,
	}
}

// FetchCafesFromAPI จำลองการยิง API ไปหา Google Maps (Places API)
func (g *googleMapsClient) FetchCafesFromAPI(ctx context.Context, lat, lng string) ([]domain.Cafe, error) {
	// ในการทำงานจริง ตรงนี้คุณจะใช้ http.Client ยิง GET ไปที่:
	// [https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=lat,lng&radius=1500&type=cafe&key=YOUR_API_KEY](https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=lat,lng&radius=1500&type=cafe&key=YOUR_API_KEY)
	
	// สำหรับ Demo: เราจะจำลองการหน่วงเวลาของ Network 
	// time.Sleep(500 * time.Millisecond)

	// จำลองข้อมูลที่ได้กลับมาจาก Google Maps (Mock Data)
	mockCafes := []domain.Cafe{
		{ID: "place_1", Name: "Amazon Cafe", Address: "Bangkok", Rating: 4.5, Lat: 13.7563, Lng: 100.5018},
		{ID: "place_2", Name: "Starbucks", Address: "Bangkok", Rating: 4.8, Lat: 13.7565, Lng: 100.5020},
	}

	return mockCafes, nil
}
