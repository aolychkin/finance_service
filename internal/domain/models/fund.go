package models

import (
	"time"

	"gorm.io/gorm"
)

// RuleValue = процент, который от родительского фонда / операции (!= null) распределяется по детям
// Если RuleValue = -1 и CheckChild = true, вычисляется от детей. То есть сумма их значений
// Если в детях есть правило %, то смотрится родитель фонда, у которого был CheckChild = true
// ___________________
// Если RuleValue = -1 и CheckChild = false и нет Goals, вычисляется по остаточному принципу
// Если RuleValue = -1 и CheckChild = false и есть Goals, вычисляется по целям
// Если RuleValue > 0 и CheckChild = false, то от заданного RuleValue, рассчитывается значение детей
// ___________________
// Если RuleValue = 0. то в фонд НЕ идет отчислений, проверку детей - пропускаем
// (!) Goal = ограничитель пополнения фонда. То есть больше не требуется
// TODO: добавить показатель since - для даты начала учета фонда
type Fund struct {
	gorm.Model
	Label           string  `gorm:"not null"`
	Priority        uint    `gorm:"not null"`
	Union           []Union `gorm:"many2many:fund_union"`
	CheckChild      bool
	RuleValue       float32
	Goals           []Goals
	Child           []*Fund `gorm:"many2many:fund_child"`
	ManualOperation []ManualOperation
	AutoOperation   []AutoOperation
	IsTmp           bool
}

// ExpireDate = срок возврата инвестиций
type Goals struct {
	gorm.Model
	Label         string
	Total         float32
	ExpireDate    time.Time
	FundID        string
	AutoOperation []AutoOperation
	IsTmp         bool
}
