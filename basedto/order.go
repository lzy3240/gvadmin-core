package basedto

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gvadmin_v3/core/global/E"
)

// 排序参数
type OrderParams struct {
	OrderByColumn string `json:"orderByColumn" form:"orderByColumn"`
	IsAsc         string `json:"isAsc" form:"isAsc"`
}

func SetOrder(orderByColumn string, isAsc string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if orderByColumn == "" {
			orderByColumn = "id"
		}

		var desc bool
		switch isAsc {
		case "ascending":
			desc = false
		case "descending":
			desc = true
		default:
			desc = true
		}

		return db.Order(clause.OrderByColumn{Column: clause.Column{Name: E.OrderKey[orderByColumn]}, Desc: desc})
	}
}
