package repo

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"wfm_api/internal/model"
)

func (r * Repo) CreateGood(ctx context.Context, name, description string, count uint ) (model.Good, error) {
	good := model.Good{
		Name:        name,
		Description: description,
		Count:       count,
	}
	return good, r.db.WithContext(ctx).Create(&good).Error
}

func (r * Repo) GetByName(ctx context.Context, name string) (model.Good, error) {
	var good model.Good
	res := r.db.WithContext(ctx).Where("name = ?", name).First(&good)
	return good, res.Error
}


func (r * Repo) GetAllGoods(ctx context.Context, limit, offset int) ([] model.Good, error) {
	if limit < -1 {
		return nil, fmt.Errorf("limit must be not lower than -1, but %d", limit)
	}
	if offset < -1 {
		return nil, fmt.Errorf("offset must be not lower than -1, but %d", offset)
	}
	var goods [] model.Good
	result := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&goods)
	return goods, result.Error
}

func (r * Repo) UpdateGoodCount(ctx context.Context, id uint, count uint) error {
	// check if exist
	var good model.Good
	err := r.db.WithContext(ctx).First(&good, id).Error
	if err != nil {
		return err
	}
	good.Count = count
	return r.db.Save(&good).Error
 }

func (r * Repo) ChangeGoodCount(ctx context.Context, id uint, diffCount int) error {
	var good model.Good
	tx := r.db.Session(&gorm.Session{Context:ctx})
	res := tx.First(&good, id)
	if int(good.Count) -diffCount < 0 {
		good.Count = 0
	} else {
		good.Count = uint(int(good.Count) + diffCount)
	}
	res = tx.Save(good)
	return res.Error
}


