package master_repo

import (
	"context"
	"errors"
	"testing"
	"time"

	"koizumi55555/go-restapi/internal/controller/http/httperr/apierr"
	"koizumi55555/go-restapi/internal/entity"
	"koizumi55555/go-restapi/internal/usecase/master_repo"
	"reflect"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Test_GetCorporation(t *testing.T) {
	// テスト用のSQLiteデータベースを作成
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	gdb, err := gorm.Open(sqlite.New(sqlite.Config{
		Conn:                   db,
		SkipDefaultTransaction: true,
	}), &gorm.Config{})
	require.NoError(t, err)

	mRepo := master_repo.NewMasterRepository(gdb)

	ctx := context.TODO()

	corpID := "1111"
	members := makeMembers()

	tests := []struct {
		name         string
		argCorpID    string
		mockBehavior func()
		want         []entity.Corporation
		wantErr      error
	}{
		{
			name:      "[正常系] 指定した企業情報が取得できること (1件)",
			argCorpID: corpID,
			mockBehavior: func() {
				mock.ExpectQuery(`SELECT \* FROM "corporation" WHERE "corporation_id" = \?`).
					WithArgs(corpID).
					WillReturnRows(sqlmock.NewRows([]string{"corporation_id", "name", "domain", "corp_type"}).
						AddRow("1111", "企業名1", "example.com", "種別1"))
			},
			want: []entity.Corporation{
				{
					CorporationID: "1111",
					Name:          "企業名1",
					Domain:        "example.com",
					CorpType:      "種別1",
				},
			},
			wantErr: nil,
		},
		{
			name:      "[正常系] 指定した企業情報が存在しないこと (0件)",
			argCorpID: corpID,
			mockBehavior: func() {
				mock.ExpectQuery(`SELECT \* FROM "corporation" WHERE "corporation_id" = \?`).
					WithArgs(corpID).
					WillReturnRows(sqlmock.NewRows([]string{"corporation_id", "name", "domain", "corp_type"}))
			},
			want:    []entity.Corporation{},
			wantErr: nil,
		},
		{
			name:      "[異常系] データベースエラーが発生した場合",
			argCorpID: corpID,
			mockBehavior: func() {
				mock.ExpectQuery(`SELECT \* FROM "corporation" WHERE "corporation_id" = \?`).
					WithArgs(corpID).
					WillReturnError(errors.New("DBエラー"))
			},
			want:    nil,
			wantErr: apierr.ErrorCodeInternalServerError{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			got, gotErr := mRepo.GetCorporation(ctx, tt.argCorpID)

			if tt.wantErr != nil {
				assert.EqualError(t, gotErr, tt.wantErr.Error())
			} else {
				require.NoError(t, gotErr)
				assert.True(t, reflect.DeepEqual(got, tt.want))
			}
		})
	}

	// モックのアサート
	require.NoError(t, mock.ExpectationsWereMet())
}

func Test_GetCorporationList(t *testing.T) {
	// テスト用のSQLiteデータベースを作成
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	gdb, err := gorm.Open(sqlite.New(sqlite.Config{
		Conn:                   db,
		SkipDefaultTransaction: true,
	}), &gorm.Config{})
	require.NoError(t, err)

	mRepo := master_repo.NewMasterRepository(gdb)

	ctx := context.TODO()

	members := makeMembers()

	tests := []struct {
		name         string
		mockBehavior func()
		want         []entity.Corporation
		wantErr      error
	}{
		{
			name: "正常系: 企業情報が取得できる場合",
			mockBehavior: func() {
				mock.ExpectQuery(`SELECT \* FROM "corporation"`).
					WillReturnRows(sqlmock.NewRows([]string{"corporation_id", "name", "domain", "corp_type"}).
						AddRow("1111", "企業名1", "example.com", "種別1").
						AddRow("2222", "企業名2", "example2.com", "種別2"))
			},
			want: []entity.Corporation{
				{
					CorporationID: "1111",
					Name:          "企業名1",
					Domain:        "example.com",
					CorpType:      "種別1",
				},
				{
					CorporationID: "2222",
					Name:          "企業名2",
					Domain:        "example2.com",
					CorpType:      "種別2",
				},
			},
			wantErr: nil,
		},
		{
			name: "正常系: 企業情報が0件の場合",
			mockBehavior: func() {
				mock.ExpectQuery(`SELECT \* FROM "corporation"`).
					WillReturnRows(sqlmock.NewRows([]string{"corporation_id", "name", "domain", "corp_type"}))
			},
			want:    []entity.Corporation{},
			wantErr: nil,
		},
		{
			name: "異常系: データベースエラーが発生した場合",
			mockBehavior: func() {
				mock.ExpectQuery(`SELECT \* FROM "corporation"`).
					WillReturnError(errors.New("DBエラー"))
			},
			want:    nil,
			wantErr: apierr.ErrorCodeInternalServerError{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			got, gotErr := mRepo.GetCorporationList(ctx)

			if tt.wantErr != nil {
				assert.EqualError(t, gotErr, tt.wantErr.Error())
			} else {
				require.NoError(t, gotErr)
				assert.True(t, reflect.DeepEqual(got, tt.want))
			}
		})
	}

	// モックのアサート
	require.NoError(t, mock.ExpectationsWereMet())
}

func Test_CreateCorporation(t *testing.T) {
	// テスト用のSQLiteデータベースを作成
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	gdb, err := gorm.Open(sqlite.New(sqlite.Config{
		Conn:                   db,
		SkipDefaultTransaction: true,
	}), &gorm.Config{})
	require.NoError(t, err)

	mRepo := master_repo.NewMasterRepository(gdb)

	ctx := context.TODO()

	tests := []struct {
		name         string
		argInput     entity.Corporation
		mockBehavior func()
		want         []entity.Corporation
		wantErr      error
	}{
		{
			name: "正常系: 企業情報が正常に作成される場合",
			argInput: entity.Corporation{
				CorporationID: "1111",
				Name:          "企業名1",
				Domain:        "example.com",
				Number:        1,
				CorpType:      "種別1",
			},
			mockBehavior: func() {
				mock.ExpectBegin()
				mock.ExpectExec(`INSERT INTO "corporation" .*`).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			want: []entity.Corporation{
				{
					CorporationID: "1111",
					Name:          "企業名1",
					Domain:        "example.com",
					Number:        1,
					CorpType:      "種別1",
				},
			},
			wantErr: nil,
		},
		{
			name: "異常系: データベースエラーが発生した場合",
			argInput: entity.Corporation{
				CorporationID: "1111",
				Name:          "企業名1",
				Domain:        "example.com",
				Number:        1,
				CorpType:      "種別1",
			},
			mockBehavior: func() {
				mock.ExpectBegin()
				mock.ExpectExec(`INSERT INTO "corporation" .*`).
					WillReturnError(errors.New("DBエラー"))
				mock.ExpectRollback()
			},
			want:    nil,
			wantErr: apierr.ErrorCodeInternalServerError{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			got, gotErr := mRepo.CreateCorporation(ctx, tt.argInput)

			if tt.wantErr != nil {
				assert.EqualError(t, gotErr, tt.wantErr.Error())
			} else {
				require.NoError(t, gotErr)
				assert.True(t, reflect.DeepEqual(got, tt.want))
			}
		})
	}

	// モックのアサート
	require.NoError(t, mock.ExpectationsWereMet())
}

func Test_UpdateCorporation(t *testing.T) {
	// テスト用のSQLiteデータベースを作成
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	gdb, err := gorm.Open(sqlite.New(sqlite.Config{
		Conn:                   db,
		SkipDefaultTransaction: true,
	}), &gorm.Config{})
	require.NoError(t, err)

	mRepo := master_repo.NewMasterRepository(gdb)

	ctx := context.TODO()

	tests := []struct {
		name         string
		argInput     entity.Corporation
		mockBehavior func()
		want         []entity.Corporation
		wantErr      error
	}{
		{
			name: "正常系: 企業情報が正常に更新される場合",
			argInput: entity.Corporation{
				CorporationID: "1111",
				Name:          "企業名1",
				Domain:        "example.com",
				Number:        1,
				CorpType:      "種別1",
			},
			mockBehavior: func() {
				mock.ExpectBegin()
				mock.ExpectExec(`UPDATE "corporation" SET .* WHERE "corporation_id" = .*`).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			want: []entity.Corporation{
				{
					CorporationID: "1111",
					Name:          "企業名1",
					Domain:        "example.com",
					Number:        1,
					CorpType:      "種別1",
				},
			},
			wantErr: nil,
		},
		{
			name: "異常系: データベースエラーが発生した場合",
			argInput: entity.Corporation{
				CorporationID: "1111",
				Name:          "企業名1",
				Domain:        "example.com",
				Number:        1,
				CorpType:      "種別1",
			},
			mockBehavior: func() {
				mock.ExpectBegin()
				mock.ExpectExec(`UPDATE "corporation" SET .* WHERE "corporation_id" = .*`).
					WillReturnError(errors.New("DBエラー"))
				mock.ExpectRollback()
			},
			want:    nil,
			wantErr: apierr.ErrorCodeInternalServerError{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			got, gotErr := mRepo.UpdateCorporation(ctx, tt.argInput)

			if tt.wantErr != nil {
				assert.EqualError(t, gotErr, tt.wantErr.Error())
			} else {
				require.NoError(t, gotErr)
				assert.True(t, reflect.DeepEqual(got, tt.want))
			}
		})
	}

	// モックのアサート
	require.NoError(t, mock.ExpectationsWereMet())
}
func Test_DeleteCorporation(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	mRepo := makeMasterRepo(t)
	corpID := "1111"

	tests := []struct {
		name         string
		argCorpID    string
		mockBehavior func()
		wantErr      error
	}{
		{
			name:      "正常系: 企業情報が正常に削除される場合",
			argCorpID: corpID,
			mockBehavior: func() {
				mock.ExpectBegin()
				mock.ExpectExec(`DELETE FROM "corporation" WHERE "corporation_id" = .*`).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			wantErr: nil,
		},
		{
			name:      "異常系: データベースエラーが発生した場合",
			argCorpID: corpID,
			mockBehavior: func() {
				mock.ExpectBegin()
				mock.ExpectExec(`DELETE FROM "corporation" WHERE "corporation_id" = .*`).
					WillReturnError(errors.New("DBエラー"))
				mock.ExpectRollback()
			},
			wantErr: apierr.ErrorCodeInternalServerError{},
		},
		{
			name:      "異常系: 削除対象のデータが存在しない場合",
			argCorpID: "2222",
			mockBehavior: func() {
				mock.ExpectBegin()
				mock.ExpectExec(`DELETE FROM "corporation" WHERE "corporation_id" = .*`).
					WillReturnResult(sqlmock.NewResult(0, 0)) // RowsAffected = 0
				mock.ExpectRollback()
			},
			wantErr: apierr.ErrorCodeResourceNotFound{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()
			gotErr := mRepo.DeleteCorporation(ctx, tt.argCorpID)

			if tt.wantErr != nil {
				assert.EqualError(t, gotErr, tt.wantErr.Error())
			} else {
				require.NoError(t, gotErr)
			}
		})
	}

	// モックのアサート
	require.NoError(t, mock.ExpectationsWereMet())
}
