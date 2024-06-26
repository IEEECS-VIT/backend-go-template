package models

import (
	"math/rand"
	"time"
)

type Team struct {
	TeamName    string    `gorm:"unique;not null"`
	TeamMembers []string  `gorm:"type:varchar(255)[]"`
	TeamLeader  string    `gorm:"not null"`
	TeamCode    string    `gorm:"unique;not null"`
	TeamLength  int       `gorm:"default:0"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

func GenerateTeamCode() string {
	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	code := make([]byte, 4)
	for i := range code {
		code[i] = letters[rand.Intn(len(letters))]
	}
	return string(code)
}