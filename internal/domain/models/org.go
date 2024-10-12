package models

import (
	"time"

	"gorm.io/gorm"
)

type Partner struct {
	gorm.Model
	Label           string `gorm:"not null"`
	ManualOperation []ManualOperation
}

type IncomeAccount struct {
	gorm.Model
	Label           string `gorm:"not null"`
	Bank            string
	UnionID         string
	ManualOperation []ManualOperation
}

type Sprint struct {
	gorm.Model
	Number    uint      `gorm:"not null"`
	StartDate time.Time `gorm:"not null"`
	EndDate   time.Time `gorm:"not null"`
	TeamID    string
}

type Team struct {
	gorm.Model
	Label           string `gorm:"not null"`
	UnionID         string `gorm:"not null"`
	Sprint          []Sprint
	ManualOperation []ManualOperation
}

type Union struct {
	gorm.Model
	Label         string `gorm:"not null"`
	IncomeAccount []IncomeAccount
	Team          []Team
}
