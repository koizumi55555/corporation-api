package model

import (
	"encoding/json"
	"io"
	"regexp"
	"strconv"

	"github.com/koizumi55555/corporation-api/internal/controller/http/httperr/apierr"
	"github.com/koizumi55555/corporation-api/internal/entity"

	"github.com/AlekSi/pointer"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	is "github.com/go-ozzo/ozzo-validation/v4/is"
)

// ValidatePatchCorporationRequest
func ValidatePatchCorporationRequest(c *gin.Context,
) (string, CorporationPatch, apierr.ApiErrF) {
	// input
	urlCorpID := c.Param("corporation_id")
	var corporationPatch CorporationPatch
	bodyByte, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return "", CorporationPatch{}, apierr.ErrorCodeInvalidRequest{}
	}
	err = json.Unmarshal(bodyByte, &corporationPatch)
	if err != nil {
		return "", CorporationPatch{}, apierr.ErrorCodeInvalidRequest{}
	}

	// corporation type
	corporationTypeItem := []interface{}{}
	for key := range entity.CorporationTypeStrMap {
		corporationTypeItem = append(corporationTypeItem, key)
	}

	// validation
	validateList := map[string]error{}
	fieldCorporation := "corporation_id"
	validateList[fieldCorporation] = validation.Validate(
		urlCorpID,
		validation.Required,
		is.UUID,
	)

	fieldName := "name"
	validateList[fieldName] = validation.Validate(
		corporationPatch.Name,
		validation.Required,
		validation.Length(1, 100),
	)

	fieldDomain := "domain"
	if len(pointer.GetString(corporationPatch.Domain)) > 0 {
		validateList[fieldDomain] = validation.Validate(
			corporationPatch.Domain,
			validation.Match(regexp.MustCompile("^[a-zA-Z0-9]*$")),
		)
	}

	fieldNumber := "number"
	validateList[fieldNumber] = validation.Validate(
		strconv.Itoa(int(pointer.GetInt32(corporationPatch.Number))),
		validation.Length(0, 6),
	)

	fieldCorpType := "corp_type"
	validateList[fieldCorpType] = validation.Validate(
		corporationPatch.CorpType,
		validation.Required,
		validation.In(corporationTypeItem...),
	)

	// validation
	validationErr := (validation.Errors)(validateList).Filter()
	if validationErr != nil {
		return "", CorporationPatch{}, apierr.ErrorCodeValidationFailed{}
	}

	if corporationPatch.Number != nil && corporationPatch.CorpType != nil {
		corporationNumberStr := strconv.Itoa(int(*corporationPatch.Number))
		var coLtd = "株式会社"
		switch *corporationPatch.CorpType {
		case coLtd:
			if corporationNumberStr[0:1] != "1" {
				// エラー処理: 株式会社の場合、企業番号は1である必要がある
				return "", CorporationPatch{}, apierr.ErrorCodeValidationFailed{}

			}
		default:
			if corporationNumberStr[0:1] == "1" {
				// エラー処理: その他の企業タイプの場合、企業番号は2以上である必要がある
				return "", CorporationPatch{}, apierr.ErrorCodeValidationFailed{}
			}
		}
	}

	return urlCorpID, corporationPatch, nil
}

// ValidatePostCorporationRequest
func ValidatePostCorporationRequest(c *gin.Context,
) (CorporationCreate, apierr.ApiErrF) {
	// input
	var corporationCreate CorporationCreate
	bodyByte, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return corporationCreate, apierr.ErrorCodeValidationFailed{}
	}
	err = json.Unmarshal(bodyByte, &corporationCreate)
	if err != nil {
		return corporationCreate, apierr.ErrorCodeValidationFailed{}
	}

	// corporation type
	corporationTypeItem := []interface{}{}
	for key := range entity.CorporationTypeStrMap {
		corporationTypeItem = append(corporationTypeItem, key)
	}

	validateList := map[string]error{}
	// body validation
	fieldName := "name"
	validateList[fieldName] = validation.Validate(
		corporationCreate.Name,
		validation.Length(1, 100),
	)

	fieldDomain := "domain"
	validateList[fieldDomain] = validation.Validate(
		corporationCreate.Domain,
		validation.Match(regexp.MustCompile("^[a-zA-Z0-9]*$")),
	)

	fieldNumber := "number"
	validateList[fieldNumber] = validation.Validate(
		strconv.Itoa(int(pointer.GetInt32(&corporationCreate.Number))),
		validation.Length(1, 6),
	)

	fieldCorpType := "corp_type"
	validateList[fieldCorpType] = validation.Validate(
		corporationCreate.CorpType,
		validation.In(corporationTypeItem...),
	)

	corporationNumberStr := strconv.Itoa(int(corporationCreate.Number))
	var coLtd = "株式会社"
	switch corporationCreate.CorpType {
	case coLtd:
		if corporationNumberStr[0:1] != "1" {
			// エラー処理: 株式会社の場合、企業番号は1である必要がある
			return corporationCreate, apierr.ErrorCodeValidationFailed{}

		}
	default:
		if corporationNumberStr[0:1] == "1" {
			// エラー処理: その他の企業タイプの場合、企業番号は2以上である必要がある
			return corporationCreate, apierr.ErrorCodeValidationFailed{}
		}
	}

	// validation
	validationErr := (validation.Errors)(validateList).Filter()
	if validationErr != nil {
		return corporationCreate, apierr.ErrorCodeValidationFailed{}
	}

	return corporationCreate, nil
}

// ValidateCorporationIdRequest
func ValidateCorporationIdRequest(c *gin.Context) (string, apierr.ApiErrF) {
	// input
	urlCorpID := c.Param("corporation_id")
	// validation
	validationErr := validation.Errors{
		"corporation_id": validation.Validate(
			urlCorpID,
			validation.Required,
			is.UUID,
		),
	}.Filter()

	if validationErr != nil {
		return "", apierr.ErrorCodeValidationFailed{}
	}
	return urlCorpID, nil
}
