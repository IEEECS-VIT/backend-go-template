package models

import (
	"math/rand"
	"time"
)

// Team represents the team model
type Team struct {
	TeamName    string    `gorm:"unique;not null"`
	TeamMembers []string  `gorm:"type:varchar(255)[]"`
	TeamLeader  string    `gorm:"not null"`
	TeamCode    string    `gorm:"unique;not null"`
	TeamLength  int       `gorm:"default:0"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

// GenerateTeamCode generates a random 4-character alphanumeric code
func GenerateTeamCode() string {
	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	code := make([]byte, 4)
	for i := range code {
		code[i] = letters[rand.Intn(len(letters))]
	}
	return string(code)
}

// ISTTime returns the current time in IST
func ISTTime() time.Time {
	istLocation, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		istLocation = time.FixedZone("IST", 5*60*60+30*60) // Fallback to fixed zone if loading fails
	}
	return time.Now().In(istLocation)
}
