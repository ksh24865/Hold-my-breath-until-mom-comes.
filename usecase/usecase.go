package usecase

import "github.com/baby-vibe/domain/model"

type ManageSignalUsecase interface {
	GetSignalByID(id int) (*model.Signal, error)
}
