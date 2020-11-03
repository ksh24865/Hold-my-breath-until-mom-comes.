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
/*
func (ar *signalRepo) Create(signal *model.Signal) (*model.Signal, error) {
	return Signal, ar.db.Create(signal).Error
}

func (ar *signalRepo) Delete(signal *model.Signal) error {
	println("call signalRepo-Delete")
	return ar.db.Delete(signal).Error
}
*/