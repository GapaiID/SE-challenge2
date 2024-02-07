package repositories

import (
	"github.com/GapaiID/SE-challenge2/api/dto"
	"gorm.io/gorm"
	"math"
)

func QueryPagination(db *gorm.DB, pp dto.PaginationParams, out interface{}) (*dto.Pagination, error) {
	pagination := new(dto.Pagination)

	total, err := QueryPage(db, pp, out)
	if err != nil {
		return pagination, err
	}

	pagination.Current = pp.Current
	pagination.PageSize = pp.PageSize
	pagination.Total = total
	pagination.LastPage = int(math.Ceil(float64(total) / float64(pp.PageSize)))

	return pagination, nil
}

func QueryPage(db *gorm.DB, pp dto.PaginationParams, out interface{}) (int64, error) {
	n, err := QueryCount(db)
	if err != nil || n == 0 {
		return 0, nil
	}

	current, pageSize := pp.Current, pp.PageSize
	if current > 0 && pageSize > 0 {
		db = db.Offset((current - 1) * pageSize).Limit(pageSize)
	} else if pageSize > 0 {
		db = db.Limit(pageSize)
	}

	err = db.Find(out).Error
	return n, err
}

func QueryCount(db *gorm.DB) (int64, error) {
	var n int64 = 0
	result := db.Count(&n)
	if err := result.Error; err != nil {
		return n, err
	}
	return n, nil
}
