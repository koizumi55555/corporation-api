package v1

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/koizumi55555/corporation-api/internal/controller/http/httperr/apierr"
	"github.com/koizumi55555/corporation-api/internal/controller/http/model"
	"github.com/koizumi55555/corporation-api/internal/entity"
	"github.com/koizumi55555/corporation-api/internal/usecase/mock"

	"github.com/AlekSi/pointer"
	"github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

func initCorporationRoutes(t *testing.T) (*corporationRoutes, *mock.MockCorporationUseCase) {
	t.Helper()
	mockCtl := gomock.NewController(t)
	uc := mock.NewMockCorporationUseCase(mockCtl)
	return &corporationRoutes{
		corporationUC: uc,
		l:             mock.GetMockLogger(t),
	}, uc
}

func Test_GetCorporation(t *testing.T) {
	var (
		uc            *mock.MockCorporationUseCase
		response      *httptest.ResponseRecorder
		testCorpID    = "efec6797-d0a5-c81f-3ff0-11f2eecf4a01"
		errTestCorpID = "efec6797-d0a5-c81f-3ff0-11f2eecf4a011111"
	)

	corporationResponse := []entity.Corporation{
		{
			CorporationID: "efec6797-d0a5-c81f-3ff0-11f2eecf4a01",
			Name:          "小泉誓約",
			Domain:        "koizumi1234",
			Number:        123456,
			CorpType:      "株式会社",
		},
	}

	corporationJsonResponse := []model.Corporation{
		{
			CorporationId: "efec6797-d0a5-c81f-3ff0-11f2eecf4a01",
			Name:          "小泉誓約",
			Domain:        "koizumi1234",
			Number:        123456,
			CorpType:      "株式会社",
		},
	}

	m, _ := json.Marshal(corporationJsonResponse)
	wantCorporation := string(m)

	tests := []struct {
		name         string
		argGenFn     func() *gin.Context
		mock         func()
		wantCode     int
		wantResponse string
	}{
		{
			name: "[異常系] 企業情報が取得できない場合",
			argGenFn: func() *gin.Context {
				response = httptest.NewRecorder()
				ctx, _ := gin.CreateTestContext(response)
				mockRequest, _ := http.NewRequest("GET",
					"/corporation/"+testCorpID, nil)
				ctx.Request = mockRequest
				ctx.Params = []gin.Param{
					{
						Key:   "corporation_id",
						Value: testCorpID,
					},
				}
				return ctx
			},
			mock: func() {
				uc.EXPECT().GetCorporation(gomock.Any(), testCorpID).
					Return([]entity.Corporation{}, &apierr.ErrorCodeInternalServerError{}).Times(1)
			},
			wantCode:     http.StatusInternalServerError,
			wantResponse: `{"error_message":"internal_server_error"}`,
		},
		{
			name: "[正常系] GETリクエストが成功し、200ステータスを返却する",
			argGenFn: func() *gin.Context {
				response = httptest.NewRecorder()
				ctx, _ := gin.CreateTestContext(response)
				mockRequest, _ := http.NewRequest("GET",
					"/corporation/"+testCorpID, nil)
				ctx.Request = mockRequest
				ctx.Params = []gin.Param{
					{
						Key:   "corporation_id",
						Value: testCorpID,
					},
				}
				return ctx
			},
			mock: func() {
				uc.EXPECT().GetCorporation(gomock.Any(), testCorpID).
					Return(corporationResponse, nil).Times(1)
			},
			wantCode:     http.StatusOK,
			wantResponse: wantCorporation,
		},
		{
			name: "[異常系] 不正な企業IDの場合",
			argGenFn: func() *gin.Context {
				response = httptest.NewRecorder()
				ctx, _ := gin.CreateTestContext(response)
				mockRequest, _ := http.NewRequest("GET",
					"/corporation/"+errTestCorpID, nil)
				ctx.Request = mockRequest
				ctx.Params = []gin.Param{
					{
						Key:   "corporation_id",
						Value: errTestCorpID,
					},
				}
				return ctx
			},
			mock: func() {
				// 呼び出しなし
			},
			wantCode:     http.StatusBadRequest,
			wantResponse: `{"error_message":"validation_failed"}`,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			r, newUC := initCorporationRoutes(t)
			uc = newUC
			tt.mock()

			argCtx := tt.argGenFn()
			r.GetCorporation(argCtx)

			if d := cmp.Diff(response.Code, tt.wantCode); len(d) != 0 {
				t.Errorf("status Code diffs: (-got +want)\n%s", d)
				return
			}

			resBody, _ := io.ReadAll(response.Body)
			if d := cmp.Diff(string(resBody), tt.wantResponse); len(d) != 0 {
				t.Errorf("body diffs: (-got +want)\n%s", d)
				return
			}
		})
	}
}

func Test_GetCorporationList(t *testing.T) {
	var (
		uc       *mock.MockCorporationUseCase
		response *httptest.ResponseRecorder
	)

	corporationResponse := []entity.Corporation{
		{
			CorporationID: "efec6797-d0a5-c81f-3ff0-11f2eecf4a01",
			Name:          "小泉誓約",
			Domain:        "koizumi1234",
			Number:        123456,
			CorpType:      "株式会社",
		},
		{
			CorporationID: "efec6797-d0a5-c81f-3ff0-11f2eecf4a02",
			Name:          "小泉誓約2",
			Domain:        "koizumi1231",
			Number:        123451,
			CorpType:      "株式会社",
		},
	}

	corporationJsonResponse := []model.Corporation{
		{
			CorporationId: "efec6797-d0a5-c81f-3ff0-11f2eecf4a01",
			Name:          "小泉誓約",
			Domain:        "koizumi1234",
			Number:        123456,
			CorpType:      "株式会社",
		},
		{
			CorporationId: "efec6797-d0a5-c81f-3ff0-11f2eecf4a02",
			Name:          "小泉誓約2",
			Domain:        "koizumi1231",
			Number:        123451,
			CorpType:      "株式会社",
		},
	}

	m, _ := json.Marshal(corporationJsonResponse)
	wantCorporation := string(m)

	tests := []struct {
		name         string
		argGenFn     func() *gin.Context
		mock         func()
		wantCode     int
		wantResponse string
	}{
		{
			name: "[正常系] GETリクエストが成功し、200ステータスを返却する",
			argGenFn: func() *gin.Context {
				response = httptest.NewRecorder()
				ctx, _ := gin.CreateTestContext(response)
				return ctx
			},
			mock: func() {
				uc.EXPECT().GetCorporationList(gomock.Any()).
					Return(corporationResponse, nil).Times(1)
			},
			wantCode:     http.StatusOK,
			wantResponse: wantCorporation,
		},
		{
			name: "[異常系] 企業情報が取得できない場合",
			argGenFn: func() *gin.Context {
				response = httptest.NewRecorder()
				ctx, _ := gin.CreateTestContext(response)
				return ctx
			},
			mock: func() {
				uc.EXPECT().GetCorporationList(gomock.Any()).
					Return([]entity.Corporation{}, &apierr.ErrorCodeInternalServerError{}).Times(1)
			},
			wantCode:     http.StatusInternalServerError,
			wantResponse: `{"error_message":"internal_server_error"}`,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			r, newUC := initCorporationRoutes(t)
			uc = newUC
			tt.mock()

			argCtx := tt.argGenFn()
			r.GetCorporationList(argCtx)

			if d := cmp.Diff(response.Code, tt.wantCode); len(d) != 0 {
				t.Errorf("status Code diffs: (-got +want)\n%s", d)
				return
			}

			resBody, _ := io.ReadAll(response.Body)
			if d := cmp.Diff(string(resBody), tt.wantResponse); len(d) != 0 {
				t.Errorf("body diffs: (-got +want)\n%s", d)
				return
			}
		})
	}
}

func Test_CreateCorporation(t *testing.T) {
	var (
		uc         *mock.MockCorporationUseCase
		response   *httptest.ResponseRecorder
		testCorpID = "efec6797-d0a5-c81f-3ff0-11f2eecf4a01"
	)

	requestBodyStruct := model.CorporationCreate{
		Name:     "小泉誓約",
		Domain:   "koizumi1234",
		Number:   123456,
		CorpType: "株式会社",
	}
	data, _ := json.Marshal(requestBodyStruct)

	errRequestBodyStruct := model.CorporationCreate{
		Name:     "小泉誓約",
		Domain:   "koizumi1234",
		Number:   223456,
		CorpType: "株式会社",
	}
	errData, _ := json.Marshal(errRequestBodyStruct)

	corporationRequest := entity.Corporation{
		CorporationID: "",
		Name:          "小泉誓約",
		Domain:        "koizumi1234",
		Number:        123456,
		CorpType:      "株式会社",
	}

	corporationResponse := []entity.Corporation{
		{
			CorporationID: testCorpID,
			Name:          "小泉誓約",
			Domain:        "koizumi1234",
			Number:        123456,
			CorpType:      "株式会社",
		},
	}

	corporationJsonResponse := []model.Corporation{
		{
			CorporationId: testCorpID,
			Name:          "小泉誓約",
			Domain:        "koizumi1234",
			Number:        123456,
			CorpType:      "株式会社",
		},
	}

	m, _ := json.Marshal(corporationJsonResponse)
	wantCorporation := string(m)

	tests := []struct {
		name         string
		argGenFn     func() *gin.Context
		mock         func()
		wantCode     int
		wantResponse string
	}{
		{
			name: "[正常系] POSTリクエストが成功し、201ステータスを返却する",
			argGenFn: func() *gin.Context {
				response = httptest.NewRecorder()
				ctx, _ := gin.CreateTestContext(response)
				mockRequest, _ := http.NewRequest("POST",
					"/corporation", nil)
				ctx.Request = mockRequest
				reader := strings.NewReader(string(data))
				ctx.Request.Body = io.NopCloser(reader)
				return ctx
			},
			mock: func() {
				uc.EXPECT().CreateCorporation(gomock.Any(), corporationRequest).
					Return(corporationResponse, nil).Times(1)
			},
			wantCode:     http.StatusCreated,
			wantResponse: wantCorporation,
		},
		{
			name: "[異常系] validationErr",
			argGenFn: func() *gin.Context {
				response = httptest.NewRecorder()
				ctx, _ := gin.CreateTestContext(response)
				mockRequest, _ := http.NewRequest("POST",
					"/corporation", nil)
				ctx.Request = mockRequest
				reader := strings.NewReader(string(errData))
				ctx.Request.Body = io.NopCloser(reader)
				return ctx
			},
			mock: func() {
				// 呼び出されない
			},
			wantCode:     http.StatusBadRequest,
			wantResponse: `{"error_message":"validation_failed"}`,
		},
		{
			name: "[異常系] 企業情報が取得できない場合",
			argGenFn: func() *gin.Context {
				response = httptest.NewRecorder()
				ctx, _ := gin.CreateTestContext(response)
				mockRequest, _ := http.NewRequest("POST",
					"/corporation", nil)
				ctx.Request = mockRequest
				reader := strings.NewReader(string(data))
				ctx.Request.Body = io.NopCloser(reader)
				return ctx
			},
			mock: func() {
				uc.EXPECT().CreateCorporation(gomock.Any(), corporationRequest).
					Return([]entity.Corporation{}, &apierr.ErrorCodeInternalServerError{}).Times(1)
			},
			wantCode:     http.StatusInternalServerError,
			wantResponse: `{"error_message":"internal_server_error"}`,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			r, newUC := initCorporationRoutes(t)
			uc = newUC
			tt.mock()

			argCtx := tt.argGenFn()
			r.CreateCorporation(argCtx)

			if d := cmp.Diff(response.Code, tt.wantCode); len(d) != 0 {
				t.Errorf("status Code diffs: (-got +want)\n%s", d)
				return
			}

			resBody, _ := io.ReadAll(response.Body)
			if d := cmp.Diff(string(resBody), tt.wantResponse); len(d) != 0 {
				t.Errorf("body diffs: (-got +want)\n%s", d)
				return
			}
		})
	}
}

func Test_UpdateCorporation(t *testing.T) {
	var (
		uc            *mock.MockCorporationUseCase
		response      *httptest.ResponseRecorder
		testCorpID    = "efec6797-d0a5-c81f-3ff0-11f2eecf4a01"
		errTestCorpID = "efec6797-d0a5-c81f-3ff0-11f2eecf4a011111"
	)

	requestBodyStruct := model.CorporationPatch{
		Name:     pointer.ToString("小泉誓約"),
		Domain:   pointer.ToString("koizumi1234"),
		Number:   pointer.ToInt32(123456),
		CorpType: pointer.ToString("株式会社"),
	}
	data, _ := json.Marshal(requestBodyStruct)

	corporationRequest := entity.Corporation{
		CorporationID: testCorpID,
		Name:          "小泉誓約",
		Domain:        "koizumi1234",
		Number:        123456,
		CorpType:      "株式会社",
	}

	corporationResponse := []entity.Corporation{
		{
			CorporationID: testCorpID,
			Name:          "小泉誓約",
			Domain:        "koizumi1234",
			Number:        123456,
			CorpType:      "株式会社",
		},
	}

	corporationJsonResponse := []model.Corporation{
		{
			CorporationId: testCorpID,
			Name:          "小泉誓約",
			Domain:        "koizumi1234",
			Number:        123456,
			CorpType:      "株式会社",
		},
	}

	m, _ := json.Marshal(corporationJsonResponse)
	wantCorporation := string(m)

	tests := []struct {
		name         string
		argGenFn     func() *gin.Context
		mock         func()
		wantCode     int
		wantResponse string
	}{
		{
			name: "[異常系] 企業情報が取得できない場合",
			argGenFn: func() *gin.Context {
				response = httptest.NewRecorder()
				ctx, _ := gin.CreateTestContext(response)
				mockRequest, _ := http.NewRequest("DELETE",
					"/corporation/"+testCorpID, nil)
				ctx.Request = mockRequest
				ctx.Params = []gin.Param{
					{
						Key:   "corporation_id",
						Value: testCorpID,
					},
				}
				reader := strings.NewReader(string(data))
				ctx.Request.Body = io.NopCloser(reader)
				return ctx
			},
			mock: func() {
				uc.EXPECT().UpdateCorporation(gomock.Any(), corporationRequest).
					Return([]entity.Corporation{}, &apierr.ErrorCodeInternalServerError{}).Times(1)
			},
			wantCode:     http.StatusInternalServerError,
			wantResponse: `{"error_message":"internal_server_error"}`,
		},
		{
			name: "[正常系] PATCHリクエストが成功し、200ステータスを返却する",
			argGenFn: func() *gin.Context {
				response = httptest.NewRecorder()
				ctx, _ := gin.CreateTestContext(response)
				mockRequest, _ := http.NewRequest("PATCH",
					"/corporation/"+testCorpID, nil)
				ctx.Request = mockRequest
				ctx.Params = []gin.Param{
					{
						Key:   "corporation_id",
						Value: testCorpID,
					},
				}
				reader := strings.NewReader(string(data))
				ctx.Request.Body = io.NopCloser(reader)
				return ctx
			},
			mock: func() {
				uc.EXPECT().UpdateCorporation(gomock.Any(), corporationRequest).
					Return(corporationResponse, nil).Times(1)
			},
			wantCode:     http.StatusOK,
			wantResponse: wantCorporation,
		},
		{
			name: "[異常系] 不正な企業IDの場合",
			argGenFn: func() *gin.Context {
				response = httptest.NewRecorder()
				ctx, _ := gin.CreateTestContext(response)
				mockRequest, _ := http.NewRequest("PATCH",
					"/corporation/"+errTestCorpID, nil)
				ctx.Request = mockRequest
				ctx.Params = []gin.Param{
					{
						Key:   "corporation_id",
						Value: errTestCorpID,
					},
				}
				reader := strings.NewReader(string(data))
				ctx.Request.Body = io.NopCloser(reader)
				return ctx
			},
			mock: func() {
				// 呼び出しなし
			},
			wantCode:     http.StatusBadRequest,
			wantResponse: `{"error_message":"validation_failed"}`,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			r, newUC := initCorporationRoutes(t)
			uc = newUC
			tt.mock()

			argCtx := tt.argGenFn()
			r.UpdateCorporation(argCtx)

			if d := cmp.Diff(response.Code, tt.wantCode); len(d) != 0 {
				t.Errorf("status Code diffs: (-got +want)\n%s", d)
				return
			}

			resBody, _ := io.ReadAll(response.Body)
			if d := cmp.Diff(string(resBody), tt.wantResponse); len(d) != 0 {
				t.Errorf("body diffs: (-got +want)\n%s", d)
				return
			}
		})
	}
}

func Test_DeleteCorporation(t *testing.T) {
	var (
		uc            *mock.MockCorporationUseCase
		response      *httptest.ResponseRecorder
		testCorpID    = "efec6797-d0a5-c81f-3ff0-11f2eecf4a01"
		errTestCorpID = "efec6797-d0a5-c81f-3ff0-11f2eecf4a011111"
	)
	tests := []struct {
		name         string
		argGenFn     func() *gin.Context
		mock         func()
		wantCode     int
		wantResponse string
	}{
		{
			name: "[正常系] DELETEリクエストが成功し、204ステータスを返却する",
			argGenFn: func() *gin.Context {
				response = httptest.NewRecorder()
				ctx, _ := gin.CreateTestContext(response)
				mockRequest, _ := http.NewRequest("DELETE",
					"/corporation/"+testCorpID, nil)
				ctx.Request = mockRequest
				ctx.Params = []gin.Param{
					{
						Key:   "corporation_id",
						Value: testCorpID,
					},
				}
				return ctx
			},
			mock: func() {
				uc.EXPECT().DeleteCorporation(gomock.Any(), testCorpID).
					Return(nil).Times(1)
			},
			wantCode:     http.StatusNoContent,
			wantResponse: "",
		},
		{
			name: "[異常系] 不正な企業IDの場合",
			argGenFn: func() *gin.Context {
				response = httptest.NewRecorder()
				ctx, _ := gin.CreateTestContext(response)
				mockRequest, _ := http.NewRequest("DELETE",
					"/corporation/"+errTestCorpID, nil)
				ctx.Request = mockRequest
				ctx.Params = []gin.Param{
					{
						Key:   "corporation_id",
						Value: errTestCorpID,
					},
				}
				return ctx
			},
			mock: func() {
				// 呼び出しなし
			},
			wantCode:     http.StatusBadRequest,
			wantResponse: `{"error_message":"validation_failed"}`,
		},
		{
			name: "[異常系] 企業情報が取得できない場合",
			argGenFn: func() *gin.Context {
				response = httptest.NewRecorder()
				ctx, _ := gin.CreateTestContext(response)
				mockRequest, _ := http.NewRequest("DELETE",
					"/corporation/"+testCorpID, nil)
				ctx.Request = mockRequest
				ctx.Params = []gin.Param{
					{
						Key:   "corporation_id",
						Value: testCorpID,
					},
				}
				return ctx
			},
			mock: func() {
				uc.EXPECT().DeleteCorporation(gomock.Any(), testCorpID).
					Return(&apierr.ErrorCodeInternalServerError{}).Times(1)
			},
			wantCode:     http.StatusInternalServerError,
			wantResponse: `{"error_message":"internal_server_error"}`,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			r, newUC := initCorporationRoutes(t)
			uc = newUC
			tt.mock()

			argCtx := tt.argGenFn()
			r.DeleteCorporation(argCtx)

			if d := cmp.Diff(response.Code, tt.wantCode); len(d) != 0 {
				t.Errorf("status Code diffs: (-got +want)\n%s", d)
				return
			}

			resBody, _ := io.ReadAll(response.Body)
			if d := cmp.Diff(string(resBody), tt.wantResponse); len(d) != 0 {
				t.Errorf("body diffs: (-got +want)\n%s", d)
				return
			}
		})
	}
}
