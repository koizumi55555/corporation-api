package entity

type CorporationType int32

const (
	CorporationTypeJointStock CorporationType = iota + 1
	CorporationTypeLimitedLiability
	CorporationTypeLimitedPartnership
	CorporationTypeGeneralPartnership
)

var (
	CorporationTypeMap = map[CorporationType]string{
		CorporationTypeJointStock:         "株式会社",
		CorporationTypeLimitedLiability:   "合同会社",
		CorporationTypeLimitedPartnership: "合資会社",
		CorporationTypeGeneralPartnership: "合名会社",
	}

	CorporationTypeStrMap = map[string]CorporationType{
		"株式会社": CorporationTypeJointStock,
		"合同会社": CorporationTypeLimitedLiability,
		"合資会社": CorporationTypeLimitedPartnership,
		"合名会社": CorporationTypeGeneralPartnership,
	}
)
