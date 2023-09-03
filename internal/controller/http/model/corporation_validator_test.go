package model

import (
	"koizumi55555/go-restapi/internal/controller/http/httperr/apierr"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
)

func Test_ValidatePatchCorporationRequest(t *testing.T) {
	type mockParam struct {
		Params []gin.Param
	}

	tests := []struct {
		name                 string
		mockParam            mockParam
		wantCorporationID    string
		wantCorporationPatch CorporationPatch
		wantErr              apierr.ApiErrF
	}{
		{
			name: "[正常系] 適切なリクエスト情報の場合を返却する [corporationID=2907a563-978c-4383-a65d-64819821f1f1, updateSequenceNumber=1]",
			mockParam: mockParam{
				Params: []gin.Param{
					{
						Key:   "corporation_id",
						Value: "2907a563-978c-4383-a65d-64819821f1f1",
					},
				},
			},
			wantCorporationID:    "2907a563-978c-4383-a65d-64819821f1f1",
			wantCorporationPatch: CorporationPatch{},
			wantErr:              nil,
		},
		{
			name: "[異常系] 不適切なリクエスト情報の場合、エラーを返却する [corporationID=2907a563-978c-4383-a65d-64819821f1f1, updateSequenceNumber=1xxx]",
			mockParam: mockParam{
				Params: []gin.Param{
					{
						Key:   "corporation_id",
						Value: "2907a563-978c-4383-a65d-64819821f1f1",
					},
				},
			},
			wantCorporationID:    "",
			wantCorporationPatch: CorporationPatch{},
			wantErr:              &apierr.ErrorCodeValidationFailed{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRequest, _ := http.NewRequest("GET", "/corporation", nil)
			mockContext := &gin.Context{Request: mockRequest, Params: tt.mockParam.Params}
			gotCorporationID, gotCorporationPatch, gotApiErr :=
				ValidatePatchCorporationRequest(mockContext)

			// CorporationIDの比較
			if gotCorporationID != tt.wantCorporationID {
				t.Errorf("ValidatePatchCorporationRequest() gotCorporationID = %s, want %s",
					gotCorporationID, tt.wantCorporationID)
			}

			// CorporationPatchの比較
			if diff := cmp.Diff(gotCorporationPatch, CorporationPatch{}); diff != "" {
				t.Errorf("ValidatePatchCorporationRequest() gotCorporationPatch (-got +want)\n%s", diff)
			}

			// apierrの比較
			if tt.wantErr == nil && gotApiErr != nil {
				t.Errorf("ValidatePatchCorporationRequest() gotapierr = %#v, want %#v", gotApiErr.Error(), nil)
			}

			if tt.wantErr != nil {
				gotErrorResponse := gotApiErr.Error()
				if diff := cmp.Diff(gotErrorResponse, tt.wantErr); diff != "" {
					t.Errorf("ValidatePatchCorporationRequest() differs: (-got +want)\n%s", diff)
				}
			}
		})
	}
}
