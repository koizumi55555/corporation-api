package master_repo

// import (
// 	"github.com/koizumi55555/corporation-api/internal/controller/http/httperr/apierr"
// 	"github.com/koizumi55555/corporation-api/internal/entity"
// 	"github.com/koizumi55555/corporation-api/internal/usecase/master_repo/schema"
// 	"reflect"
// 	"testing"

// 	"github.com/google/go-cmp/cmp"
// 	"gorm.io/gorm"
// )

// func Test_GetCorporationList(t *testing.T) {
// 	mRepo := makeMasterRepo(t)
// 	tests := []struct {
// 		name          string
// 		otherDBAccess func(conn *gorm.DB)
// 		want          []entity.Corporation
// 		wantError     apierr.ApiErrF
// 	}{
// 		{
// 			name: "[正常系]複数人取得",
// 			otherDBAccess: func(tx *gorm.DB) {
// 				// create seed
// 				if err := tx.Create(makeCorporation()).Error; err != nil {
// 					t.Errorf("insert seed error.[%s]", err.Error())
// 				}
// 				tx.Commit()
// 			},
// 			want:      nil,
// 			wantError: nil,
// 		},
// 	}

// 	for _, tt := range tests {
// 		// truncate
// 		TruncateDB(t, mRepo.DBHandler)
// 		SeedData(t, mRepo.DBHandler)

// 		t.Run(tt.name, func(t *testing.T) {
// 			tx := mRepo.DBHandler.Conn.Begin()
// 			defer tx.Rollback()
// 			if tt.otherDBAccess != nil {
// 				tt.otherDBAccess(tx)
// 			}
// 			got, gotError := mRepo.GetCorporationList()
// 			if diff := cmp.Diff(got, tt.want); diff != "" {
// 				t.Errorf("GetCorporationList() = %s", diff)
// 			}

// 			if gotError != tt.wantError &&
// 				reflect.TypeOf(gotError) != reflect.TypeOf(tt.wantError) {
// 				t.Errorf("GetCorporationList() gotError = %v, wantError %v", gotError, tt.wantError)
// 			}
// 		})
// 	}
// }

// func makeCorporation() []schema.Corporation {
// 	return []schema.Corporation{
// 		{
// 			CorporationID: "efec6797-d0a5-c81f-3ff0-11f2eecf4a01",
// 			Name:          "小泉誓約",
// 			Domain:        "koizumi1234",
// 			Number:        123456,
// 			CorpType:      "株式会社",
// 		},
// 		{
// 			CorporationID: "efec6797-d0a5-c81f-3ff0-11f2eecf4a02",
// 			Name:          "小泉製薬",
// 			Domain:        "koizumi1234",
// 			Number:        123456,
// 			CorpType:      "株式会社",
// 		},
// 	}
// }
