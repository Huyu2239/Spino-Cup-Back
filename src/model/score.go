package model

import "time"

type Score struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	Point      uint      `json:"point"`
	Time       uint      `json:"time" gorm:"not null"`
	MissCount  uint      `json:"miss_count"`
	Mode       string    `json:"mode" gorm:"not null"`
	Difficulty string    `json:"difficulty" gorm:"not null"`
	UserID     uint      `json:"user_id" gorm:"not null"`
	User       User      `json:"user" gorm:"foreignKey:UserID; constraint:OnDelete:CASCADE"`
	QuizID     uint      `json:"quiz_id" gorm:"not null"`
	Quiz       Quiz      `json:"quiz" gorm:"foreignKey:QuizID; constraint:OnDelete:CASCADE"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type ScoreResponse struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Point     uint      `json:"point"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
