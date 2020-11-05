package manageWeight

import (
	"github.com/baby-vibe/domain/model"
	"github.com/baby-vibe/domain/repository"
)

type manageWeight struct {
	wr repository.WeightRepo
}

func NewManageWeightUsecase(wr repository.WeightRepo) *manageWeight {
	return &manageWeight{
		wr: wr,
	}
}

func (mwuc *manageWeight) GetWeightByID(id int) (*model.Weight, error) {
	return mwuc.wr.GetByID(id)
}
