package repository

import "github.com/baby-vibe/domain/model"

type SignalRepo interface {
	GetByID(id int) (*model.Signal, error)
}
type WeightRepo interface {
	GetByID(id int) (*model.Weight, error)
}

