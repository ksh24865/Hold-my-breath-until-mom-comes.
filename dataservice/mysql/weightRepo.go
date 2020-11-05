package mysql

import (
	"github.com/baby-vibe/domain/model"
	"github.com/jinzhu/gorm"
)

type weightRepo struct {
	db *gorm.DB
}

func NewWeightRepo() *weightRepo {
	return &weightRepo{
		db: dbConn,
	}
}
func (wr *weightRepo) GetByID(id int) (a *model.Weight, err error) {
	a = new(model.Weight)
	return a, wr.db.Where("id=?", id).First(a).Error
}
