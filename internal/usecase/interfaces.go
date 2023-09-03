package usecase

import (
	"context"
	"koizumi55555/go-restapi/internal/controller/http/httperr/apierr"
	"koizumi55555/go-restapi/internal/entity"
)

type (
	CorporationUseCase interface {
		GetCorporation(ctx context.Context, corp string) ([]entity.Corporation, apierr.ApiErrF)
		GetCorporationList(ctx context.Context) ([]entity.Corporation, apierr.ApiErrF)
		CreateCorporation(ctx context.Context,
			input entity.Corporation) ([]entity.Corporation, apierr.ApiErrF)
		UpdateCorporation(ctx context.Context,
			input entity.Corporation) ([]entity.Corporation, apierr.ApiErrF)
		DeleteCorporation(ctx context.Context, corp string) apierr.ApiErrF
	}

	MasterRepository interface {
		GetCorporation(corp string) ([]entity.Corporation, apierr.ApiErrF)
		GetCorporationList() ([]entity.Corporation, apierr.ApiErrF)
		CreateCorporation(input entity.Corporation) ([]entity.Corporation, apierr.ApiErrF)
		UpdateCorporation(input entity.Corporation) ([]entity.Corporation, apierr.ApiErrF)
		DeleteCorporation(corp string) apierr.ApiErrF
		ExistCorporationID(corp string) apierr.ApiErrF
		ExistCorporationName(name string) apierr.ApiErrF
	}
)
