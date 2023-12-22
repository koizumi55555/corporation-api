package usecase_test

import (
	"context"
	"testing"

	"github.com/koizumi55555/corporation-api/internal/controller/http/httperr/apierr"
	"github.com/koizumi55555/corporation-api/internal/entity"
	"github.com/koizumi55555/corporation-api/internal/usecase"
	"github.com/koizumi55555/corporation-api/internal/usecase/mock"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
)

func initCorporationUseCase(t *testing.T) (usecase.CorporationUseCase, *mock.MockMasterRepository) {
	t.Helper()
	mRepo, _ := mock.GetMockMasterRepo(t)
	uc := usecase.NewCorporationUsecase(mRepo)
	return uc, mRepo
}

func Test_GetCorporation(t *testing.T) {
	var (
		uc    usecase.CorporationUseCase
		mRepo *mock.MockMasterRepository
	)

	const testCorpID = "3d407c8d-73e2-4d98-84eb-612e1adb9f29"

	corporationSetting := []entity.Corporation{
		{
			CorporationID: "efec6797-d0a5-c81f-3ff0-11f2eecf4a01",
			Name:          "小泉誓約",
			Domain:        "koizumi1234",
			Number:        123456,
			CorpType:      "株式会社",
		},
	}
	type args struct {
		corpID string
	}
	tests := []struct {
		name      string
		mock      func()
		args      args
		want      []entity.Corporation
		wantError apierr.ApiErrF
	}{
		{
			name: "[正常系] 指定の企業情報が取得できること",
			mock: func() {
				mRepo.EXPECT().GetCorporation(testCorpID).
					Return(corporationSetting, nil).Times(1)
			},
			args: args{
				corpID: testCorpID,
			},
			want:      corporationSetting,
			wantError: nil,
		},
		{
			name: "[異常系] 企業情報取得時に ErrorCodeInternalServerErrorが返却されること",
			mock: func() {
				mRepo.EXPECT().GetCorporation(testCorpID).
					Return(nil, &apierr.ErrorCodeInternalServerError{}).Times(1)
			},
			args: args{
				corpID: testCorpID,
			},
			want:      []entity.Corporation{},
			wantError: &apierr.ErrorCodeInternalServerError{},
		},
		{
			name: "[異常系] 指定の企業情報が存在しないときNotFoundが返却されること",
			mock: func() {
				mRepo.EXPECT().GetCorporation(testCorpID).
					Return(nil, &apierr.ErrorCodeResourceNotFound{}).Times(1)
			},
			args: args{
				corpID: testCorpID,
			},
			want:      []entity.Corporation{},
			wantError: &apierr.ErrorCodeResourceNotFound{},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			uc, mRepo = initCorporationUseCase(t)
			tt.mock()
			got, gotError := uc.GetCorporation(
				context.Background(), tt.args.corpID)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("CorporationUseCase.GetCorporation() = %s", diff)
			}

			if diff := cmp.Diff(gotError, tt.wantError); diff != "" {
				t.Errorf("CorporationUseCase.GetCorporation() = %s", diff)
			}
		})
	}
}

func Test_GetCorporationList(t *testing.T) {
	var (
		uc    usecase.CorporationUseCase
		mRepo *mock.MockMasterRepository
	)

	corporationSetting := []entity.Corporation{
		{
			CorporationID: "efec6797-d0a5-c81f-3ff0-11f2eecf4a01",
			Name:          "小泉誓約",
			Domain:        "koizumi1234",
			Number:        123456,
			CorpType:      "株式会社",
		},
		{
			CorporationID: "efec6797-d0a5-c81f-3ff0-11f2eecf4a02",
			Name:          "小泉製薬",
			Domain:        "koizumiS",
			Number:        123457,
			CorpType:      "株式会社",
		},
	}

	tests := []struct {
		name      string
		mock      func()
		want      []entity.Corporation
		wantError apierr.ApiErrF
	}{
		{
			name: "[正常系] 企業情報のリストが取得できること",
			mock: func() {
				mRepo.EXPECT().GetCorporationList().
					Return(corporationSetting, nil).Times(1)
			},
			want:      corporationSetting,
			wantError: nil,
		},
		{
			name: "[異常系] 企業情報取得時に ErrorCodeInternalServerErrorが返却されること",
			mock: func() {
				mRepo.EXPECT().GetCorporationList().
					Return(nil, &apierr.ErrorCodeInternalServerError{}).Times(1)
			},
			want:      []entity.Corporation{},
			wantError: &apierr.ErrorCodeInternalServerError{},
		},
		{
			name: "[異常系] 企業情報が存在しないときNotFoundが返却されること",
			mock: func() {
				mRepo.EXPECT().GetCorporationList().
					Return(nil, &apierr.ErrorCodeResourceNotFound{}).Times(1)
			},
			want:      []entity.Corporation{},
			wantError: &apierr.ErrorCodeResourceNotFound{},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			uc, mRepo = initCorporationUseCase(t)
			tt.mock()
			got, gotError := uc.GetCorporationList(context.Background())
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("CorporationUseCase.GetCorporationList() = %s", diff)
			}

			if diff := cmp.Diff(gotError, tt.wantError); diff != "" {
				t.Errorf("CorporationUseCase.GetCorporationList() = %s", diff)
			}
		})
	}
}

func Test_CreateCorporation(t *testing.T) {
	var (
		uc    usecase.CorporationUseCase
		mRepo *mock.MockMasterRepository
	)

	var testCorpName = "小泉誓約"
	var testCorpID = uuid.New().String()

	corporationReq := entity.Corporation{
		CorporationID: testCorpID,
		Name:          testCorpName,
		Domain:        "koizumi1234",
		Number:        123456,
		CorpType:      "株式会社",
	}

	corporationSetting := []entity.Corporation{
		{
			CorporationID: testCorpID,
			Name:          testCorpName,
			Domain:        "koizumi1234",
			Number:        123456,
			CorpType:      "株式会社",
		},
	}

	tests := []struct {
		name      string
		mock      func()
		want      []entity.Corporation
		wantError apierr.ApiErrF
	}{
		{
			name: "[正常系] 企業登録成功しその企業情報が返却されること",
			mock: func() {
				mRepo.EXPECT().ExistCorporationName(testCorpName).
					Return(nil).Times(1)
				mRepo.EXPECT().CreateCorporation(corporationReq).
					Return(corporationSetting, nil).Times(1)
			},
			want:      corporationSetting,
			wantError: nil,
		},
		{
			name: "[異常系] 企業情報取得時に ErrorCodeInternalServerErrorが返却されること",
			mock: func() {
				mRepo.EXPECT().ExistCorporationName(testCorpName).
					Return(nil).Times(1)
				mRepo.EXPECT().CreateCorporation(corporationReq).
					Return([]entity.Corporation{}, &apierr.ErrorCodeInternalServerError{}).Times(1)
			},
			want:      []entity.Corporation{},
			wantError: &apierr.ErrorCodeInternalServerError{},
		},
		{
			name: "[異常系] 同名の企業が存在するときConflictが返却されること",
			mock: func() {
				mRepo.EXPECT().ExistCorporationName(testCorpName).
					Return(&apierr.ErrorCodeConflict{}).Times(1)
			},
			want:      []entity.Corporation{},
			wantError: &apierr.ErrorCodeConflict{},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			uc, mRepo = initCorporationUseCase(t)
			tt.mock()
			got, gotError := uc.CreateCorporation(
				context.Background(), corporationReq)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("CorporationUseCase.CreateCorporation() = %s", diff)
			}

			if diff := cmp.Diff(gotError, tt.wantError); diff != "" {
				t.Errorf("CorporationUseCase.CreateCorporation() = %s", diff)
			}
		})
	}
}

func Test_UpdateCorporation(t *testing.T) {
	var (
		uc    usecase.CorporationUseCase
		mRepo *mock.MockMasterRepository
	)

	var testCorpName = "小泉誓約"
	var testCorpID = "efec6797-d0a5-c81f-3ff0-11f2eecf4a01"
	corporationReq := entity.Corporation{
		CorporationID: testCorpID,
		Name:          testCorpName,
		Domain:        "koizumi1234",
		Number:        123456,
		CorpType:      "株式会社",
	}

	corporationSetting := []entity.Corporation{
		{
			CorporationID: testCorpID,
			Name:          testCorpName,
			Domain:        "koizumi1234",
			Number:        123456,
			CorpType:      "株式会社",
		},
	}

	tests := []struct {
		name      string
		mock      func()
		want      []entity.Corporation
		wantError apierr.ApiErrF
	}{
		{
			name: "[正常系] 企業更新成功しその企業情報が返却されること",
			mock: func() {
				mRepo.EXPECT().ExistCorporationID(testCorpID).
					Return(nil).Times(1)
				mRepo.EXPECT().ExistCorporationName(testCorpName).
					Return(nil).Times(1)
				mRepo.EXPECT().UpdateCorporation(corporationReq).
					Return(corporationSetting, nil).Times(1)
			},
			want:      corporationSetting,
			wantError: nil,
		},
		{
			name: "[異常系] 企業情報取得時にErrorCodeInternalServerErrorが返却されること",
			mock: func() {
				mRepo.EXPECT().ExistCorporationID(testCorpID).
					Return(nil).Times(1)
				mRepo.EXPECT().ExistCorporationName(testCorpName).
					Return(nil).Times(1)
				mRepo.EXPECT().UpdateCorporation(corporationReq).
					Return([]entity.Corporation{}, &apierr.ErrorCodeInternalServerError{}).Times(1)
			},
			want:      []entity.Corporation{},
			wantError: &apierr.ErrorCodeInternalServerError{},
		},
		{
			name: "[異常系] 指定した企業が存在しないときNotFoundが返却されること",
			mock: func() {
				mRepo.EXPECT().ExistCorporationID(testCorpID).
					Return(&apierr.ErrorCodeResourceNotFound{}).Times(1)
			},
			want:      []entity.Corporation{},
			wantError: &apierr.ErrorCodeResourceNotFound{},
		},
		{
			name: "[異常系] 同名の企業が存在するときConflictが返却されること",
			mock: func() {
				mRepo.EXPECT().ExistCorporationID(testCorpID).
					Return(nil).Times(1)
				mRepo.EXPECT().ExistCorporationName(testCorpName).
					Return(&apierr.ErrorCodeConflict{}).Times(1)
			},
			want:      []entity.Corporation{},
			wantError: &apierr.ErrorCodeConflict{},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			uc, mRepo = initCorporationUseCase(t)
			tt.mock()
			got, gotError := uc.UpdateCorporation(
				context.Background(), corporationReq)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("CorporationUseCase.UpdateCorporation() = %s", diff)
			}

			if diff := cmp.Diff(gotError, tt.wantError); diff != "" {
				t.Errorf("CorporationUseCase.UpdateCorporation() = %s", diff)
			}
		})
	}
}

func Test_DeleteCorporation(t *testing.T) {
	var (
		uc    usecase.CorporationUseCase
		mRepo *mock.MockMasterRepository
	)

	var testCorpID = "efec6797-d0a5-c81f-3ff0-11f2eecf4a01"

	tests := []struct {
		name      string
		mock      func()
		wantError apierr.ApiErrF
	}{
		{
			name: "[正常系] 企業削除されること",
			mock: func() {
				mRepo.EXPECT().ExistCorporationID(testCorpID).
					Return(nil).Times(1)
				mRepo.EXPECT().DeleteCorporation(testCorpID).
					Return(nil).Times(1)
			},
			wantError: nil,
		},
		{
			name: "[異常系] 企業削除時にErrorCodeInternalServerErrorが返却されること",
			mock: func() {
				mRepo.EXPECT().ExistCorporationID(testCorpID).
					Return(nil).Times(1)
				mRepo.EXPECT().DeleteCorporation(testCorpID).
					Return(&apierr.ErrorCodeInternalServerError{}).Times(1)
			},
			wantError: &apierr.ErrorCodeInternalServerError{},
		},
		{
			name: "[異常系] 指定した企業が存在しないときNotFoundが返却されること",
			mock: func() {
				mRepo.EXPECT().ExistCorporationID(testCorpID).
					Return(&apierr.ErrorCodeResourceNotFound{}).Times(1)
			},
			wantError: &apierr.ErrorCodeResourceNotFound{},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			uc, mRepo = initCorporationUseCase(t)
			tt.mock()
			gotError := uc.DeleteCorporation(
				context.Background(), testCorpID)
			if diff := cmp.Diff(gotError, tt.wantError); diff != "" {
				t.Errorf("CorporationUseCase.DeleteCorporation() = %s", diff)
			}
		})
	}
}
