package master_repo

import (
	"github.com/koizumi55555/corporation-api/internal/controller/http/httperr/apierr"
	"github.com/koizumi55555/corporation-api/internal/entity"
	"github.com/koizumi55555/corporation-api/internal/usecase/master_repo/schema"

	"github.com/google/uuid"
)

var updateCorporationColumns = []string{
	"name",
	"domain",
	"number",
	"corp_type",
}

type WriteCorporation struct {
	CorporationID string `gorm:"column:corporation_id;comment:企業ID"`
	Name          string `gorm:"column:name;comment:企業名"`
	Domain        string `gorm:"column:domain;comment:企業ドメイン"`
	Number        int32  `gorm:"column:number;comment:企業番号"`
	CorpType      string `gorm:"column:corp_type;comment:企業種別"`
}

func (*WriteCorporation) TableName() string {
	return "corporation"
}

// Get Corporation
func (r *MasterRepository) GetCorporation(corpID string,
) ([]entity.Corporation, apierr.ApiErrF) {
	var rtnCorporation []schema.Corporation
	result := r.DBHandler.Conn.
		Model(&schema.Corporation{}).
		Where(&schema.Corporation{CorporationID: corpID}).
		Find(&rtnCorporation)

	if result.Error != nil {
		r.l.Errorf("Get Corporation Error", result.Error)
		return []entity.Corporation{}, apierr.ErrorCodeInternalServerError{}
	}

	return makeCorporationRes(&rtnCorporation), nil
}

// Get Corporation List
func (r *MasterRepository) GetCorporationList() (
	[]entity.Corporation, apierr.ApiErrF) {
	var rtnCorporation []schema.Corporation
	result := r.DBHandler.Conn.
		Model(&schema.Corporation{}).
		Find(&rtnCorporation)
	if result.Error != nil {
		r.l.Errorf("Get Corporation List Error", result.Error)
		return []entity.Corporation{}, apierr.ErrorCodeInternalServerError{}
	}

	return makeCorporationRes(&rtnCorporation), nil
}

// Create Corporation
func (r *MasterRepository) CreateCorporation(
	input entity.Corporation,
) ([]entity.Corporation, apierr.ApiErrF) {
	var rtnCorporation []schema.Corporation
	createSetting := makeCorporationReq(input)
	result := r.DBHandler.Conn.
		Create(createSetting).
		Where(&schema.Corporation{CorporationID: createSetting.CorporationID}).
		Find(&rtnCorporation)
	if result.Error != nil {
		r.l.Errorf("Create Corporation Error", result.Error)
		return []entity.Corporation{}, apierr.ErrorCodeInternalServerError{}
	}
	return makeCorporationRes(&rtnCorporation), nil
}

// Update Corporation
func (r *MasterRepository) UpdateCorporation(
	input entity.Corporation,
) ([]entity.Corporation, apierr.ApiErrF) {
	var rtnCorporation []schema.Corporation
	updateSetting := makeCorporationReq(input)
	result := r.DBHandler.Conn.
		Model(updateSetting).
		Select(updateCorporationColumns).
		Where(&WriteCorporation{
			CorporationID: updateSetting.CorporationID,
		}).
		Updates(updateSetting)
	if result.Error != nil {
		return []entity.Corporation{}, apierr.ErrorCodeInternalServerError{}
	}

	response := r.DBHandler.Conn.
		Model(&schema.Corporation{}).
		Where(&WriteCorporation{
			CorporationID: updateSetting.CorporationID,
		}).
		Find(&rtnCorporation)
	if response.Error != nil {
		return []entity.Corporation{}, apierr.ErrorCodeInternalServerError{}
	}

	return makeCorporationRes(&rtnCorporation), nil
}

// Delete Corporation
func (r *MasterRepository) DeleteCorporation(corpID string) apierr.ApiErrF {
	var deleteSetting []schema.Corporation
	result := r.DBHandler.Conn.Debug().
		Where(&WriteCorporation{
			CorporationID: corpID,
		}).Delete(deleteSetting)
	if result.Error != nil {
		return apierr.ErrorCodeInternalServerError{}
	}
	return nil
}

// exist CorporationIDw
func (r *MasterRepository) ExistCorporationID(corpID string,
) apierr.ApiErrF {
	var rtnCorporation []schema.Corporation
	result := r.DBHandler.Conn.
		Model(&schema.Corporation{}).
		Where(&schema.Corporation{CorporationID: corpID}).
		Find(&rtnCorporation)
	if result.Error != nil {
		r.l.Errorf("Exist CorporationID Error", result.Error)
		return apierr.ErrorCodeInternalServerError{}
	}
	//存在しない場合404を返す
	if len(rtnCorporation) == 0 {
		return apierr.ErrorCodeResourceNotFound{}
	}
	return nil
}

// exist CorporationName
func (r *MasterRepository) ExistCorporationName(name string,
) apierr.ApiErrF {
	var rtnCorporation []schema.Corporation
	result := r.DBHandler.Conn.
		Model(&schema.Corporation{}).
		Where(&schema.Corporation{Name: name}).
		Find(&rtnCorporation)
	if result.Error != nil {
		return apierr.ErrorCodeInternalServerError{}
	}

	if len(rtnCorporation) > 0 && rtnCorporation[0].Name == name {
		return nil
	}
	// 存在する場合は409を返す
	if len(rtnCorporation) != 0 {
		return apierr.ErrorCodeConflict{}
	}

	return nil
}

func makeCorporationReq(input entity.Corporation) *WriteCorporation {
	if input.CorporationID == "" {
		input.CorporationID = uuid.New().String()
	}
	return &WriteCorporation{
		CorporationID: input.CorporationID,
		Name:          input.Name,
		Domain:        input.Domain,
		Number:        input.Number,
		CorpType:      input.CorpType,
	}
}

func makeCorporationRes(corporation *[]schema.Corporation) []entity.Corporation {
	if corporation != nil {
		res := make([]entity.Corporation, len(*corporation))
		for i, m := range *corporation {
			res[i] = entity.Corporation{
				CorporationID: m.CorporationID,
				Name:          m.Name,
				Domain:        m.Domain,
				Number:        m.Number,
				CorpType:      m.CorpType,
			}
		}
		return res
	} else {
		return []entity.Corporation{}
	}

}
