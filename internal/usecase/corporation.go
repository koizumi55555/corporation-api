package usecase

import (
	"context"
	"koizumi55555/corporation-api/internal/controller/http/httperr/apierr"
	"koizumi55555/corporation-api/internal/entity"
	"koizumi55555/corporation-api/pkg/logger"

	"github.com/google/uuid"
)

type corporationUseCase struct {
	masterRepo MasterRepository
	l          *logger.Logger
}

func NewCorporationUsecase(mRepo MasterRepository) CorporationUseCase {
	return &corporationUseCase{
		masterRepo: mRepo,
		//l:          l,
	}
}

// Get Corporation
func (cuc *corporationUseCase) GetCorporation(
	tx context.Context, corp string,
) ([]entity.Corporation, apierr.ApiErrF) {
	corporation, err := cuc.masterRepo.GetCorporation(corp)
	if err != nil {
		return []entity.Corporation{}, err
	}
	return corporation, nil
}

// Get Corporation List
func (cuc *corporationUseCase) GetCorporationList(
	ctx context.Context,
) ([]entity.Corporation, apierr.ApiErrF) {
	corporations, err := cuc.masterRepo.GetCorporationList()
	if err != nil {
		return nil, err
	}
	return corporations, nil
}

// Create Corporation
func (cuc *corporationUseCase) CreateCorporation(
	ctx context.Context, input entity.Corporation,
) ([]entity.Corporation, apierr.ApiErrF) {
	input.CorporationID = uuid.NewString()
	// 同名の企業がないか確認
	err := cuc.masterRepo.ExistCorporationName(input.Name)
	if err != nil {
		return []entity.Corporation{}, err
	}

	// create
	corp, err := cuc.masterRepo.CreateCorporation(input)
	if err != nil {
		return []entity.Corporation{}, err
	}
	return corp, nil
}

// Update Corporation
func (cuc *corporationUseCase) UpdateCorporation(
	ctx context.Context, input entity.Corporation,
) ([]entity.Corporation, apierr.ApiErrF) {
	// 指定されたCorporationIDが存在するか確認
	err := cuc.masterRepo.ExistCorporationID(input.CorporationID)
	if err == nil {
		return []entity.Corporation{}, err
	}

	// 同名の企業がないか確認
	err = cuc.masterRepo.ExistCorporationName(input.Name)
	if err == nil {
		return []entity.Corporation{}, err
	}
	// update
	corp, err := cuc.masterRepo.UpdateCorporation(input)
	if err != nil {
		return []entity.Corporation{}, err
	}
	return corp, nil
}

// Delete Corporation
func (cuc *corporationUseCase) DeleteCorporation(
	ctx context.Context, corp string,
) apierr.ApiErrF {
	// 指定されたCorporationIDが存在するか確認
	err := cuc.masterRepo.ExistCorporationID(corp)
	if err == nil {
		return err
	}

	// delete
	err = cuc.masterRepo.DeleteCorporation(corp)
	if err != nil {
		return err
	}
	return nil
}
