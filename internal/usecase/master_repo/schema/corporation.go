package schema

const TableNameCorporation = "corporation"

type Corporation struct {
	CorporationID string `gorm:"column:corporation_id;comment:企業ID"`
	Name          string `gorm:"column:name;comment:企業名"`
	Domain        string `gorm:"column:domain;comment:企業ドメイン"`
	Number        int32  `gorm:"column:number;comment:企業番号"`
	CorpType      string `gorm:"column:corp_type;comment:企業種別"`
}

func (*Corporation) TableName() string {
	return TableNameCorporation
}
