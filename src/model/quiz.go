package model

import (
	"time"
)

type Quiz2 struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Question     string    `json:"question" gorm:"not null"` //内容
	Code         string    `json:"code" gorm:"not null"`
	InputSample  string    `json:"input_sample"`
	OutputSample string    `json:"output_sample"`
	InputSecret  string    `json:"input_secret"`
	OutputSecret string    `json:"output_secret"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type QuizResponse struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	Question     string `json:"question" grom:"not null;"`
	Code         string `json:"code" gorm:"not null"`
	InputSample  string `json:"input_sample"`
	OutputSample string `json:"output_sample"`
	InputSecret  string `json:"input_secret"`
	OutputSecret string `json:"output_secret"`
}
