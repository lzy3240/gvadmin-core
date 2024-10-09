package basedto

import "gorm.io/gorm"

// 分页参数
type PageParams struct {
	PageNum  int `json:"pageNum" form:"pageNum"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

func SetPaginate(pageSize, pageNum int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// 设置默认值
		if pageSize == 0 {
			pageSize = 10
		}

		if pageNum == 0 {
			pageNum = 1
		}

		offset := (pageNum - 1) * pageSize
		if offset < 0 {
			offset = 0
		}
		return db.Offset(offset).Limit(pageSize)
	}
}
