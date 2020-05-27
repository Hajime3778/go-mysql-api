package models

import "time"

// CommonColumn すべてのDBに共通する項目
type CommonColumn struct {
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at" sql:"not null;type:datetime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at" sql:"not null;type:datetime"`
}
