package repository

import "github.com/baby-vibe/domain/model"

type SignalRepo interface {
	GetByID(id int) (*model.Signal, error)
	//Create(signal *model.Signal) (*model.Signal, error)
	//Delete(signal *model.Signal) error
}

