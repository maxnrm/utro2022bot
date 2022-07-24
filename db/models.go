package db

// User is telegram user
type User struct {
	ID            int64  `json:"id" gorm:"primaryKey"`
	Group         string `json:"group"`
	JoinedChannel bool   `json:"joinedChannel"`
}
