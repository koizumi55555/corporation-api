package usecase_test

import (
	"context"
	"koizumi55555/go-restapi/internal/controller/http/httperr/apierr"
	"koizumi55555/go-restapi/internal/entity"
	"koizumi55555/go-restapi/internal/usecase"
	"koizumi55555/go-restapi/internal/usecase/mock"
	"testing"

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
			Name:          "Test Taro",
			Domain:        "",
			Number:        0,
			CorpType:      "",
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
					Return(corporationSetting, nil).Times(1)
			},
			args: args{
				corpID: testCorpID,
			},
			want:      []entity.Corporation{},
			wantError: nil,
		},
		{
			name: "[異常系] 指定の企業情報が存在しないときNotFoundが返却されること",
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
			Name:          "Test Taro",
			Domain:        "",
			Number:        0,
			CorpType:      "",
		},
		{
			CorporationID: "efec6797-d0a5-c81f-3ff0-11f2eecf4a02",
			Name:          "Test Taro",
			Domain:        "",
			Number:        0,
			CorpType:      "",
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
					Return(corporationSetting, nil).Times(1)
			},
			want:      []entity.Corporation{},
			wantError: nil,
		},
		{
			name: "[異常系] 企業情報が存在しないときNotFoundが返却されること",
			mock: func() {
				mRepo.EXPECT().GetCorporationList().
					Return(nil, &apierr.ErrorCodeInternalServerError{}).Times(1)
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

	const (
		testCorpName = "TEST"
	)
	var testCorpID = uuid.New().String()
	corporationReq := entity.Corporation{
		CorporationID: testCorpID,
		Name:          "Test Taro",
		Domain:        "",
		Number:        0,
		CorpType:      "",
	}

	corporationSetting := []entity.Corporation{
		{
			CorporationID: testCorpID,
			Name:          "Test Taro",
			Domain:        "",
			Number:        0,
			CorpType:      "",
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
					Return(entity.Corporation{}, &apierr.ErrorCodeInternalServerError{}).Times(1)
			},
			want:      []entity.Corporation{},
			wantError: nil,
		},
		{
			name: "[異常系] 企業情報が存在しないときNotFoundが返却されること",
			mock: func() {
				mRepo.EXPECT().ExistCorporationName(testCorpName).
					Return(nil).Times(1)
				mRepo.EXPECT().CreateCorporation(corporationReq).
					Return(entity.Corporation{}, &apierr.ErrorCodeInternalServerError{}).Times(1)
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
			wantError: &apierr.ErrorCodeInternalServerError{},
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

	const testCorpName = "TEST"
	var testCorpID = uuid.New().String()
	corporationReq := entity.Corporation{
		CorporationID: testCorpID,
		Name:          "Test Taro",
		Domain:        "",
		Number:        0,
		CorpType:      "",
	}

	corporationSetting := []entity.Corporation{
		{
			CorporationID: testCorpID,
			Name:          "Test Taro",
			Domain:        "",
			Number:        0,
			CorpType:      "",
		},
	}

	type args struct {
		corp entity.Corporation
	}
	tests := []struct {
		name      string
		mock      func()
		args      args
		want      []entity.Corporation
		wantError apierr.ApiErrF
	}{
		{
			name: "[正常系] 企業更新成功しその企業情報が返却されること",
			mock: func() {
				mRepo.EXPECT().UpdateCorporation(corporationReq).
					Return(corporationSetting, nil).Times(1)
			},
			want:      corporationSetting,
			wantError: nil,
		},
		{
			name: "[異常系] 企業情報取得時に ErrorCodeInternalServerErrorが返却されること",
			mock: func() {
				mRepo.EXPECT().UpdateCorporation(corporationReq).
					Return(entity.Corporation{}, &apierr.ErrorCodeInternalServerError{}).Times(1)
			},
			want:      []entity.Corporation{},
			wantError: nil,
		},
		{
			name: "[異常系] 企業情報が存在しないときNotFoundが返却されること",
			mock: func() {
				mRepo.EXPECT().GetCorporation(corporationSetting).
					Return(entity.Corporation{}, &apierr.ErrorCodeInternalServerError{}).Times(1)
				mRepo.EXPECT().ExistCorporationName(testCorpName).
					Return(nil).Times(1)
				mRepo.EXPECT().UpdateCorporation(corporationSetting).
					Return(entity.Corporation{}, &apierr.ErrorCodeInternalServerError{}).Times(1)
			},
			want:      []entity.Corporation{},
			wantError: &apierr.ErrorCodeInternalServerError{},
		},
		{
			name: "[異常系] 指定した企業が存在しないとき   が返却されること",
			mock: func() {
				mRepo.EXPECT().GetCorporation(corporationSetting).
					Return(entity.Corporation{}, &apierr.ErrorCodeInternalServerError{}).Times(1)
				mRepo.EXPECT().ExistCorporationName(testCorpName).
					Return(&apierr.ErrorCodeInternalServerError{}).Times(1)
			},
			want:      []entity.Corporation{},
			wantError: &apierr.ErrorCodeInternalServerError{},
		},
		{
			name: "[異常系] 同名の企業が存在するときConflictが返却されること",
			mock: func() {
				mRepo.EXPECT().GetCorporation(corporationSetting).
					Return(entity.Corporation{}, &apierr.ErrorCodeInternalServerError{}).Times(1)
				mRepo.EXPECT().ExistCorporationName(testCorpName).
					Return(&apierr.ErrorCodeConflict{}).Times(1)
			},
			want:      []entity.Corporation{},
			wantError: &apierr.ErrorCodeInternalServerError{},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			uc, mRepo = initCorporationUseCase(t)
			tt.mock()
			got, gotError := uc.UpdateCorporation(
				context.Background(), tt.args.corp)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("CorporationUseCase.GetCorporation() = %s", diff)
			}

			if diff := cmp.Diff(gotError, tt.wantError); diff != "" {
				t.Errorf("CorporationUseCase.GetCorporation() = %s", diff)
			}
		})
	}
}

func Test_DeleteCorporation(t *testing.T) {
	var (
		uc    usecase.CorporationUseCase
		mRepo *mock.MockMasterRepository
	)

	const (
		testCorpID   = "3d407c8d-73e2-4d98-84eb-612e1adb9f29"
		testCorpName = "TEST"
	)

	corporationSetting := entity.Corporation{
		CorporationID: testCorpID,
		Name:          testCorpName,
		Domain:        "",
		Number:        0,
		CorpType:      "",
	}

	tests := []struct {
		name       string
		mock       func()
		argsCropID string
		want       entity.Corporation
		wantError  apierr.ApiErrF
	}{
		{
			name: "[正常系] 企業更新成功しその企業情報が返却されること",
			mock: func() {
				corporationSetting.CorporationID = uuid.NewString()
				mRepo.EXPECT().DeleteCorporation(testCorpID).
					Return(nil).Times(1)
			},
			argsCropID: testCorpID,
			want:       corporationSetting,
			wantError:  nil,
		},
		{
			name: "[異常系] 企業情報取得時に ErrorCodeInternalServerErrorが返却されること",
			mock: func() {
				mRepo.EXPECT().DeleteCorporation(testCorpID).
					Return(entity.Corporation{}, &apierr.ErrorCodeInternalServerError{}).Times(1)
			},
			want:      entity.Corporation{},
			wantError: nil,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			uc, mRepo = initCorporationUseCase(t)
			tt.mock()
			gotError := uc.DeleteCorporation(
				context.Background(), tt.argsCropID)
			if diff := cmp.Diff(gotError, tt.wantError); diff != "" {
				t.Errorf("CorporationUseCase.GetCorporation() = %s", diff)
			}
		})
	}
}
