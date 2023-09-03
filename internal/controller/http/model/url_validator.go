package model

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	is "github.com/go-ozzo/ozzo-validation/v4/is"
)

var fieldCorporationID = "corporation_id"

func addUrlCorporationIDValidate(validateList map[string]error, urlCorpID string) {
	var str string
	validateList[fieldCorporationID] = validation.Validate(
		urlCorpID,
		validation.Required.Error(str),
		is.UUID.Error(str),
	)
}
