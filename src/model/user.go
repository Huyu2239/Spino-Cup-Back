package model

import "time"

type User struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	FirebaseUID string    `json:"firebase_uid" gorm:"unique"`
	Name        string    `json:"name" gorm:"not null"`
	Email       string    `json:"email" gorm:"unique"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
