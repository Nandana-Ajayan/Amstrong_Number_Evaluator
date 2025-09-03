package main

import "time"

type User struct {
	UserID    uint      `gorm:"primaryKey" json:"user_id"`
	Email     string    `gorm:"uniqueIndex;not null" json:"email"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	Numbers   []ArmstrongNumber `json:"numbers,omitempty"`
}

type ArmstrongNumber struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"index;not null" json:"user_id"`
	Number    int64     `gorm:"not null" json:"number"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
