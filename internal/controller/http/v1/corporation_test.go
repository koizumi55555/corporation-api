package v1

import (
	"io"
	"koizumi55555/corporation-api/internal/entity"
	"koizumi55555/corporation-api/internal/usecase/mock"
	"net/http"
	"net/http/httptest"
	"testing"

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

func TestGetCorporation(t *testing.T) {
	var uc *mock.MockCorporationUseCase
	var response *httptest.ResponseRecorder

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
				// テスト用のリクエストコンテキストをここに設定する
				return ctx
			},
			mock: func() {
				uc.EXPECT().GetCorporation(gomock.Any(), gomock.Any()).Return([]entity.Corporation{}, nil).Times(1)
			},
			wantCode:     http.StatusOK,
			wantResponse: "ここに期待されるレスポンスデータを設定する",
		},
		// 他のGETのテストケースを追加する
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

func TestCreateCorporation(t *testing.T) {
	var uc *mock.MockCorporationUseCase
	var response *httptest.ResponseRecorder

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
				// テスト用のリクエストコンテキストをここに設定する
				return ctx
			},
			mock: func() {
				uc.EXPECT().CreateCorporation(gomock.Any(), gomock.Any()).Return([]entity.Corporation{}, nil).Times(1)
			},
			wantCode:     http.StatusCreated,
			wantResponse: "ここに期待されるレスポンスデータを設定する",
		},
		// 他のPOSTのテストケースを追加する
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

func TestUpdateCorporation(t *testing.T) {
	var uc *mock.MockCorporationUseCase
	var response *httptest.ResponseRecorder

	tests := []struct {
		name         string
		argGenFn     func() *gin.Context
		mock         func()
		wantCode     int
		wantResponse string
	}{
		{
			name: "[正常系] PATCHリクエストが成功し、200ステータスを返却する",
			argGenFn: func() *gin.Context {
				response = httptest.NewRecorder()
				ctx, _ := gin.CreateTestContext(response)
				// テスト用のリクエストコンテキストをここに設定する
				return ctx
			},
			mock: func() {
				uc.EXPECT().UpdateCorporation(gomock.Any(), gomock.Any()).Return([]entity.Corporation{}, nil).Times(1)
			},
			wantCode:     http.StatusOK,
			wantResponse: "ここに期待されるレスポンスデータを設定する",
		},
		// 他のPATCHのテストケースを追加する
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

func TestDeleteCorporation(t *testing.T) {
	var uc *mock.MockCorporationUseCase
	var response *httptest.ResponseRecorder

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
				// テスト用のリクエストコンテキストをここに設定する
				return ctx
			},
			mock: func() {
				uc.EXPECT().DeleteCorporation(gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
			wantCode:     http.StatusNoContent,
			wantResponse: "",
		},
		// 他のDELETEのテストケースを追加する
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
		})
	}
}
