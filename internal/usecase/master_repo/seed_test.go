package master_repo

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// table name
const TableNameCorporation = "corporation"

// seed data
var SeedCorporationData = CorporationSeedModel{
	CorporationID: "2907a563-978c-4383-a65d-64819821f1f1",
	Name:          "小泉製薬_seed",
	Domain:        "domain_seed",
	Number:        123456,
	CorpType:      "株式会社",
}

type CorporationSeedModel struct {
	CorporationID string `gorm:"column:corporation_id;comment:企業ID"`
	Name          string `gorm:"column:name;comment:企業名"`
	Domain        string `gorm:"column:domain;comment:企業ドメイン"`
	Number        int32  `gorm:"column:number;comment:企業番号"`
	CorpType      string `gorm:"column:corp_type;comment:企業種別"`
}

func (*CorporationSeedModel) TableName() string {
	return TableNameCorporation
}

func SeedCorporation(db *gorm.DB) error {
	if err := db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&SeedCorporationData).Error; err != nil {
		return err
	}
	return nil
}
