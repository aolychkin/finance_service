package models

import (
	"time"

	"gorm.io/gorm"
)

type ManualOperation struct {
	gorm.Model
	Date              time.Time `gorm:"not null"`
	SprintID          string
	Value             float32 `gorm:"not null"`
	Details           string
	TeamID            string
	FundID            string
	IncomeAccountID   string
	PartnerID         string
	OperationStatusID string
	AutoOperation     []AutoOperation
	IsTmp             bool
}

type AutoOperation struct {
	gorm.Model
	Date              time.Time `gorm:"not null"`
	SprintID          string
	Value             float32 `gorm:"not null"`
	ManualOperationID string
	FundID            string
	GoalsID           string
	OperationStatusID string
}

type OperationStatus struct {
	gorm.Model
	Label           string `gorm:"not null"`
	ManualOperation []ManualOperation
}
