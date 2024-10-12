package sqlite

import (
	"context"
	"errors"
	"finance_service/internal/domain/models"
	"finance_service/internal/storage"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

// Конструктор Storage
func New(storagePath string) (*Storage, error) {
	const op = "storage.sqlite.New"

	// Указываем путь до файла БД
	db, err := gorm.Open(sqlite.Open(storagePath), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

// SaveUser saves user to db.
func (s *Storage) SaveFundConfig(
	ctx context.Context,
	unionID int64,
	label string,
	priority uint64,
	check_child bool,
	rule_value int64,
	is_tmp bool,
) (int64, error) {
	const op = "storage.sqlite.SaveFundConfig"

	var union []models.Union
	err := s.db.Model(&models.Union{}).First(&union, unionID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, fmt.Errorf("%s: %w", op, storage.ErrInvalidUnionID)
	}

	fundConfig := models.Fund{
		Label:      label,
		Priority:   uint(priority),
		Union:      union,
		CheckChild: check_child,
		RuleValue:  float32(rule_value),
		IsTmp:      is_tmp,
	}

	result := s.db.Create(&fundConfig)
	if result.Error != nil {
		return 0, fmt.Errorf("%s: %w", op, result.Error)
	}

	return int64(fundConfig.ID), nil
}

// User returns user by email.
func (s *Storage) FundConfig(ctx context.Context, fundID int64) (models.Fund, error) {
	const op = "storage.sqlite.User"

	var fund models.Fund
	err := s.db.Model(&models.Fund{}).Preload("Child").Preload("Goals").First(&fund, fund.ID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Fund{}, fmt.Errorf("%s: %w", op, storage.ErrInvalidFundID)
	}

	return fund, nil
}
