package model

import (
    "time"
    "github.com/google/uuid"
)

// Base Model with Increament Primary Key 
type Base struct {
    ID        uint64    `gorm:"column:id;autoIncrement;primaryKey;" json:"id"`
    CreatedAt time.Time `gorm:"column:created_at;type:datetime;not null;" json:"created_at"`
    UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;not null;" json:"updated_at"`
}

// Base Model with UUID Primary Key 
type Base struct {
    ID        uuid.UUID `gorm:"column:id;autoIncrement;primaryKey;" json:"id"`
    CreatedAt time.Time `gorm:"column:created_at;type:datetime;not null;" json:"created_at"`
    UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;not null;" json:"updated_at"`
}
