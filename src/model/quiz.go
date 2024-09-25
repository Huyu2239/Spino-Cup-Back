package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type StringMatrix [][]string

// データベースからの読み込み時
func (sm *StringMatrix) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), sm)
}

// データベースに保存する前に呼び出される
func (sm StringMatrix) Value() (driver.Value, error) {
	return json.Marshal(sm)
}

type Quiz struct {
	ID          uint         `json:"id" gorm:"primaryKey"`
	Difficulty  string       `json:"difficulty"`                  //難易度
	Language    string       `json:"language"`                    //クイズに使用されているプログラミング言語
	Question    StringMatrix `json:"question" grom:"not null"`    //内容
	AnswerX     uint         `json:"answer_x" gorm:"not null;"`   //正解のx座標
	AnswerY     uint         `json:"answer_y" gorm:"not null;"`   //正解のy座標
	EditedText  string       `json:"edited_text"`                 //編集した内容
	Explanation string       `json:"explanation" gorm:"not null"` //解説

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type QuizResponse struct {
	ID         uint         `json:"id" gorm:"primaryKey"`
	Question   StringMatrix `json:"question" grom:"not null;"`
	Difficulty string       `json:"difficulty"`
	Language   string       `json:"language"`
}

type CheckResponse struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	IsCorrect bool `json:"is_correct"`
}

type AnswerResponse struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	AnswerX     uint   `json:"answer_x" gorm:"not null;"` //正解のx座標
	AnswerY     uint   `json:"answer_y" gorm:"not null;"` //正解のy座標
	EditedText  string `json:"edited_text"`
	Explanation string `json:"explanation" gorm:"not null"`
}

type LanguageResponse struct {
	Languages []string `json:"languages"`
}
