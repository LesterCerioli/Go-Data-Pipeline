package models

import "time"

type Pipeline struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"not null;uniqueIndex"`
	Description string         `json:"description"`
	Source      map[string]any `json:"source" gorm:"serializer:json;type:text;not null"`
	Destination map[string]any `json:"destination" gorm:"serializer:json;type:text;not null"`
	Status      string         `json:"status" gorm:"not null;default:draft;index"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
}
