package mysql

import (
	"github.com/baby-vibe/domain/model"
	"github.com/jinzhu/gorm"
)

type signalRepo struct {
	db *gorm.DB
}

func NewSignalRepo() *signalRepo {
	return &signalRepo{
		db: dbConn,
	}
}
func (ar *signalRepo) GetByID(id int) (a *model.Signal, err error) {
	a = new(model.Signal)
	return a, ar.db.Where("id=?", id).First(a).Error
}
