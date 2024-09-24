package repository

import (
	"api/model"
	"fmt"

	"gorm.io/gorm"
)

func applyFilters(db *gorm.DB, filters []model.Filter) (*gorm.DB, error) {

	for _, filter := range filters {
		switch filter.Operator {
		case "equals":
			db = db.Where(filter.Field+" = ?", filter.Value)
		case "not_equals":
			db = db.Where(filter.Field+" != ?", filter.Value)
		case "contains":
			db = db.Where(filter.Field+" LIKE ?", "%"+filter.Value+"%")
		case "less_than":
			db = db.Where(filter.Field+" < ?", filter.Value)
		case "greater_than":
			db = db.Where(filter.Field+" > ?", filter.Value)
		default:
			return db, fmt.Errorf("invalid filter operator")
		}
	}

	return db, nil
}
