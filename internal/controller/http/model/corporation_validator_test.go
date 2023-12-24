package model

import (
	"bytes"
	"io"
	"net/http"
	"reflect"
	"testing"

	"github.com/koizumi55555/corporation-api/internal/controller/http/httperr/apierr"

	"github.com/AlekSi/pointer"
	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
)

func Test_ValidatePatchCorporationRequest(t *testing.T) {
	type mockParam struct {
		Params []gin.Param
	}

	var corporationPatch CorporationPatch
	testCorpID := "2907a563-978c-4383-a65d-64819821f1f1"
	baseCorporationPatch := CorporationPatch{
		Name:     pointer.ToString("小泉誓約"),
		Domain:   pointer.ToString("koizumi1234"),
		Number:   pointer.ToInt32(123456),
		CorpType: pointer.ToString("株式会社"),
	}

	baseCorporationPatchNotDomain := CorporationPatch{
		Name:     pointer.ToString("小泉誓約"),
		Number:   pointer.ToInt32(123456),
		CorpType: pointer.ToString("株式会社"),
	}

	baseCorporationPatchNotNumber := CorporationPatch{
		Name:     pointer.ToString("小泉誓約"),
		Domain:   pointer.ToString("koizumi1234"),
		CorpType: pointer.ToString("株式会社"),
	}

	tests := []struct {
		name                  string
		mockParam             mockParam
		requestBody           io.Reader
		want_corporationID    string
		want_CorporationPatch CorporationPatch
		wantApiErrResponse    apierr.ApiErrF
	}{
		{
			name: "[正常系] 適切なリクエスト情報の場合、企業IDを返却する path[corporationID=2907a563-978c-4383-a65d-64819821f1f1]",
			mockParam: mockParam{
				Params: []gin.Param{
					{
						Key:   "corporation_id",
						Value: testCorpID,
					},
				},
			},
			requestBody: bytes.NewBuffer([]byte(`
				{
					"name": "小泉誓約",
					"domain": "koizumi1234",
					"number": 123456,
					"corp_type":"株式会社"
				}`)),
			want_corporationID:    testCorpID,
			want_CorporationPatch: baseCorporationPatch,
			wantApiErrResponse:    nil,
		},
		{
			name: "[正常系] 適切なリクエスト情報の場合、企業IDを返却する(body domain無) path[corporationID=2907a563-978c-4383-a65d-64819821f1f1]",
			mockParam: mockParam{
				Params: []gin.Param{
					{
						Key:   "corporation_id",
						Value: testCorpID,
					},
				},
			},
			requestBody: bytes.NewBuffer([]byte(`
				{
					"name": "小泉誓約",
					"number": 123456,
					"corp_type":"株式会社"
				}`)),
			want_corporationID:    testCorpID,
			want_CorporationPatch: baseCorporationPatchNotDomain,
			wantApiErrResponse:    nil,
		},
		{
			name: "[正常系] 適切なリクエスト情報の場合、企業IDを返却する(body number無) path[corporationID=2907a563-978c-4383-a65d-64819821f1f1]",
			mockParam: mockParam{
				Params: []gin.Param{
					{
						Key:   "corporation_id",
						Value: testCorpID,
					},
				},
			},
			requestBody: bytes.NewBuffer([]byte(`
				{
					"name": "小泉誓約",
					"domain": "koizumi1234",
					"corp_type":"株式会社"
				}`)),
			want_corporationID:    testCorpID,
			want_CorporationPatch: baseCorporationPatchNotNumber,
			wantApiErrResponse:    nil,
		},
		{
			name: "[異常系] 不正なリクエスト情報の場合ErrorCodeInvalidRequest",
			mockParam: mockParam{
				Params: []gin.Param{
					{
						Key:   "corporation_id",
						Value: testCorpID,
					},
				},
			},
			requestBody: bytes.NewBuffer([]byte(`
				{
					"name": "小泉誓約",
					"domain": "koizumi1234",
					"number": 123456,
				}`)),
			want_corporationID:    "",
			want_CorporationPatch: corporationPatch,
			wantApiErrResponse:    apierr.ErrorCodeInvalidRequest{},
		},
		{
			name: "[異常系] validation error not corporation id",
			mockParam: mockParam{
				Params: []gin.Param{
					{
						Key:   "corporation_id",
						Value: "",
					},
				},
			},
			requestBody: bytes.NewBuffer([]byte(`
				{
					"name": "小泉誓約",
					"domain": "koizumi1234",
					"number": 123456,
					"corp_type":"株式会社"
				}`)),
			want_corporationID:    "",
			want_CorporationPatch: corporationPatch,
			wantApiErrResponse:    apierr.ErrorCodeValidationFailed{},
		},
		{
			name: "[異常系] validation error not name",
			mockParam: mockParam{
				Params: []gin.Param{
					{
						Key:   "corporation_id",
						Value: testCorpID,
					},
				},
			},
			requestBody: bytes.NewBuffer([]byte(`
				{
					"domain": "koizumi1234",
					"number": 123456,
					"corp_type":"株式会社"
				}`)),
			want_corporationID:    "",
			want_CorporationPatch: corporationPatch,
			wantApiErrResponse:    apierr.ErrorCodeValidationFailed{},
		},
		{
			name: "[異常系] validation error not corp_type",
			mockParam: mockParam{
				Params: []gin.Param{
					{
						Key:   "corporation_id",
						Value: testCorpID,
					},
				},
			},
			requestBody: bytes.NewBuffer([]byte(`
				{
					"name": "小泉誓約",
					"domain": "koizumi1234",
					"number": 123456
				}`)),
			want_corporationID:    "",
			want_CorporationPatch: corporationPatch,
			wantApiErrResponse:    apierr.ErrorCodeValidationFailed{},
		},
		{
			name: "[異常系] validation error name over",
			mockParam: mockParam{
				Params: []gin.Param{
					{
						Key:   "corporation_id",
						Value: testCorpID,
					},
				},
			},
			requestBody: bytes.NewBuffer([]byte(`
				{
					"name": "12345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890",
					"domain": "koizumi1234",
					"number": 123456,
					"corp_type":"株式会社"
				}`)),
			want_corporationID:    "",
			want_CorporationPatch: corporationPatch,
			wantApiErrResponse:    apierr.ErrorCodeValidationFailed{},
		},
		{
			name: "[異常系] validation error name 0",
			mockParam: mockParam{
				Params: []gin.Param{
					{
						Key:   "corporation_id",
						Value: testCorpID,
					},
				},
			},
			requestBody: bytes.NewBuffer([]byte(`
				{
					"name": "",
					"domain": "domain",
					"number": 123456,
					"corp_type":"株式会社"
				}`)),
			want_corporationID:    "",
			want_CorporationPatch: corporationPatch,
			wantApiErrResponse:    apierr.ErrorCodeValidationFailed{},
		},
		{
			name: "[異常系] validation error domain not alphanumeric",
			mockParam: mockParam{
				Params: []gin.Param{
					{
						Key:   "corporation_id",
						Value: testCorpID,
					},
				},
			},
			requestBody: bytes.NewBuffer([]byte(`
				{
					"name": "",
					"domain": "てすと",
					"number": 123456,
					"corp_type":"株式会社"
				}`)),
			want_corporationID:    "",
			want_CorporationPatch: corporationPatch,
			wantApiErrResponse:    apierr.ErrorCodeValidationFailed{},
		},

		{
			name: "[異常系] validation error number over",
			mockParam: mockParam{
				Params: []gin.Param{
					{
						Key:   "corporation_id",
						Value: testCorpID,
					},
				},
			},
			requestBody: bytes.NewBuffer([]byte(`
				{
					"name": "小泉誓約",
					"domain": "koizumi1234",
					"number": 1234567,
					"corp_type":"株式会社"
				}`)),
			want_corporationID:    "",
			want_CorporationPatch: corporationPatch,
			wantApiErrResponse:    apierr.ErrorCodeValidationFailed{},
		},
		{
			name: "[異常系] validation error not corp type  ",
			mockParam: mockParam{
				Params: []gin.Param{
					{
						Key:   "corporation_id",
						Value: testCorpID,
					},
				},
			},
			requestBody: bytes.NewBuffer([]byte(`
				{
					"name": "小泉誓約",
					"domain": "koizumi1234",
					"number": 123456
				}`)),
			want_corporationID:    "",
			want_CorporationPatch: corporationPatch,
			wantApiErrResponse:    apierr.ErrorCodeValidationFailed{},
		},
		{
			name: "[異常系] validation error corp type not included ",
			mockParam: mockParam{
				Params: []gin.Param{
					{
						Key:   "corporation_id",
						Value: testCorpID,
					},
				},
			},
			requestBody: bytes.NewBuffer([]byte(`
				{
					"name": "小泉誓約",
					"domain": "koizumi1234",
					"number": 123456,
					"corp_type":"TEST"
				}`)),
			want_corporationID:    "",
			want_CorporationPatch: corporationPatch,
			wantApiErrResponse:    apierr.ErrorCodeValidationFailed{},
		},
		{
			name: "[異常系] 株式会社で企業番号が1以外から始まること ",
			mockParam: mockParam{
				Params: []gin.Param{
					{
						Key:   "corporation_id",
						Value: testCorpID,
					},
				},
			},
			requestBody: bytes.NewBuffer([]byte(`
				{
					"name": "小泉誓約",
					"domain": "koizumi1234",
					"number": 223456,
					"corp_type":"株式会社"
				}`)),
			want_corporationID:    "",
			want_CorporationPatch: corporationPatch,
			wantApiErrResponse:    apierr.ErrorCodeValidationFailed{},
		},
		{
			name: "[異常系] 株式会社以外で企業番号が1から始まること ",
			mockParam: mockParam{
				Params: []gin.Param{
					{
						Key:   "corporation_id",
						Value: testCorpID,
					},
				},
			},
			requestBody: bytes.NewBuffer([]byte(`
				{
					"name": "小泉誓約",
					"domain": "koizumi1234",
					"number": 123456,
					"corp_type":"合同会社"
				}`)),
			want_corporationID:    "",
			want_CorporationPatch: corporationPatch,
			wantApiErrResponse:    apierr.ErrorCodeValidationFailed{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRequest, _ := http.NewRequest("PATCH",
				"/corporation/"+testCorpID, tt.requestBody)
			mockContext := &gin.Context{Request: mockRequest, Params: tt.mockParam.Params}
			got_corporationID, gotCorporationPatch, gotApiErr :=
				ValidatePatchCorporationRequest(mockContext)
			if got_corporationID != tt.want_corporationID {
				t.Errorf("ValidatePostWorkspaceRequest() got_corporationID = %v, want %v",
					got_corporationID, tt.want_corporationID)
			}

			// nilの場合の判定
			if tt.wantApiErrResponse == nil {
				if d := cmp.Diff(gotCorporationPatch, tt.want_CorporationPatch); len(d) != 0 {
					t.Errorf("differs: (-got +want)\n%s", d)
				}
			}

			// nil以外の判定
			if tt.wantApiErrResponse != nil {
				if gotApiErr != tt.wantApiErrResponse &&
					reflect.TypeOf(gotApiErr) != reflect.TypeOf(tt.wantApiErrResponse) {
					t.Errorf("ValidatePatchCorporationRequest  gotError = %v, wantError %v", gotApiErr, tt.wantApiErrResponse)
				}
			}

		})
	}
}
func Test_ValidateCorporationIdRequest(t *testing.T) {
	type mockParam struct {
		Params []gin.Param
	}

	testCorpID := "2907a563-978c-4383-a65d-64819821f1f1"

	tests := []struct {
		name                  string
		mockParam             mockParam
		requestBody           io.Reader
		want_corporationID    string
		want_CorporationPatch CorporationPatch
		wantApiErrResponse    apierr.ApiErrF
	}{
		{
			name: "[正常系] 適切なリクエスト情報の場合、企業IDを返却する path[corporationID=2907a563-978c-4383-a65d-64819821f1f1]",
			mockParam: mockParam{
				Params: []gin.Param{
					{
						Key:   "corporation_id",
						Value: testCorpID,
					},
				},
			},
			want_corporationID: testCorpID,
			wantApiErrResponse: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRequest, _ := http.NewRequest("GET",
				"/corporation/"+testCorpID, tt.requestBody)
			mockContext := &gin.Context{Request: mockRequest, Params: tt.mockParam.Params}
			got_corporationID, gotApiErr :=
				ValidateCorporationIdRequest(mockContext)
			if got_corporationID != tt.want_corporationID {
				t.Errorf("ValidatePostWorkspaceRequest() got_corporationID = %v, want %v",
					got_corporationID, tt.want_corporationID)
			}

			// nil以外の判定
			if tt.wantApiErrResponse != nil {
				if gotApiErr != tt.wantApiErrResponse &&
					reflect.TypeOf(gotApiErr) != reflect.TypeOf(tt.wantApiErrResponse) {
					t.Errorf("ValidatePatchCorporationRequest  gotError = %v, wantError %v", gotApiErr, tt.wantApiErrResponse)
				}
			}

		})
	}
}
