package domain

// ChatMessage คือหน้าตาของข้อความที่เราจะส่งหากันไปมา
type ChatMessage struct {
	MeetupID string `json:"meetup_id"`
	UserID   string `json:"user_id"`
	Text     string `json:"text"`
}