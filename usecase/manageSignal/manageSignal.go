package manageSignal

import (
	"github.com/baby-vibe/domain/model"
	"github.com/baby-vibe/domain/repository"
)

type manageSignal struct {
	sr repository.SignalRepo
}

func NewManageSignalUsecase(sr repository.SignalRepo) *manageSignal {
	return &manageSignal{
		sr: sr,
	}
}

func (msuc *manageSignal) GetSignalByID(id int) (*model.Signal, error) {
	return msuc.sr.GetByID(id)
}
