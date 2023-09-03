package model

import (
	"encoding/json"
	"io"
	"koizumi55555/go-restapi/internal/controller/http/httperr/apierr"
	"regexp"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// ValidatePatchCorporationRequest
func ValidatePatchCorporationRequest(c *gin.Context,
) (string, CorporationPatch, apierr.ApiErrF) {
	// input
	urlCorpID := c.Param("corporation_id")
	var corporationPatch CorporationPatch
	bodyByte, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return "", corporationPatch, apierr.ErrorCodeInvalidRequest{}
	}
	err = json.Unmarshal(bodyByte, &corporationPatch)
	if err != nil {
		return "", corporationPatch, apierr.ErrorCodeInvalidRequest{}
	}

	validateList := map[string]error{}
	// path validation
	addUrlCorporationIDValidate(validateList, urlCorpID)

	// body validation
	fieldName := "name"
	validateList[fieldName] = validation.Validate(
		corporationPatch.Name,
		validation.Length(1, 100).Error(apierr.ErrCodeInvalidRequest),
	)

	fieldDomain := "domain"
	validateList[fieldDomain] = validation.Validate(
		corporationPatch.Domain,
		validation.Match(regexp.MustCompile("^[a-zA-Z0-9]*$")).
			Error("Domain should contain only alphanumeric characters"),
	)

	fieldNumber := "number"
	validateList[fieldNumber] = validation.Validate(
		corporationPatch.Number,
		validation.Match(regexp.MustCompile("^[0-9]{1,6}$")).
			Error("Number should be a numeric value with up to 6 digits"),
	)

	fieldCorpType := "corp_type"
	validateList[fieldCorpType] = validation.Validate(
		corporationPatch.CorpType,
		validation.In("株式会社", "合同会社", "合資会社", "合名会社").
			Error("Invalid CorpType"),
	)

	return urlCorpID, corporationPatch, nil
}

// ValidatePostCorporationRequest
func ValidatePostCorporationRequest(c *gin.Context,
) (CorporationCreate, apierr.ApiErrF) {

	var msg string
	// input
	var corporationCreate CorporationCreate
	bodyByte, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return corporationCreate, nil
	}
	err = json.Unmarshal(bodyByte, &corporationCreate)
	if err != nil {
		return corporationCreate, nil
	}

	validateList := map[string]error{}

	// body validation
	fieldName := "name"
	validateList[fieldName] = validation.Validate(
		corporationCreate.Name,
		validation.Length(1, 100).Error(msg),
	)

	fieldDomain := "domain"
	validateList[fieldDomain] = validation.Validate(
		corporationCreate.Domain,
		validation.Match(regexp.MustCompile("^[a-zA-Z0-9]*$")).
			Error("Domain should contain only alphanumeric characters"),
	)

	fieldNumber := "number"
	validateList[fieldNumber] = validation.Validate(
		corporationCreate.Number,
		validation.Match(regexp.MustCompile("^[0-9]{1,6}$")).
			Error("Number should be a numeric value with up to 6 digits"),
	)

	fieldCorpType := "corp_type"
	validateList[fieldCorpType] = validation.Validate(
		corporationCreate.CorpType,
		validation.In("株式会社", "合同会社", "合資会社", "合名会社").
			Error("Invalid CorpType"),
	)

	// validation
	validationErr := (validation.Errors)(validateList).Filter()
	if validationErr != nil {
		return corporationCreate, nil
	}

	return corporationCreate, nil
}

// ValidateCorporationIdRequest
func ValidateCorporationIdRequest(c *gin.Context) (string, apierr.ApiErrF) {
	// input
	urlCorpID := c.Param("corporation_id")
	validateList := map[string]error{}
	// path validation
	addUrlCorporationIDValidate(validateList, urlCorpID)
	return urlCorpID, nil
}
