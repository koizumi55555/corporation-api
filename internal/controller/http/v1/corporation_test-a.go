package v1

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"koizumi55555/go-restapi/internal/entity"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	gomock "github.com/golang/mock/gomock"
// 	"github.com/google/go-cmp/cmp"
// )

// func initMemberRoutes(t *testing.T) (*memberRoutes, *mock.MockMemberUseCase) {
// 	t.Helper()
// 	mockCtl := gomock.NewController(t) // go1.14 以降はmockCtl.Finish()呼び出しの必要なし
// 	uc := mock.NewMockMemberUseCase(mockCtl)
// 	return &memberRoutes{
// 		memberUC: uc,
// 		l:        mock.GetMockLogger(t),
// 	}, uc
// }

// func Test_GetCorporation(t *testing.T) {
// 	var uc *mock.MockMemberUseCase
// 	var response *httptest.ResponseRecorder

// 	baseUpdateSequenceNumber := int32(1)
// 	internalServerErrorByte, _ := json.Marshal(apperr.ErrorInternalServerError{}.Error())
// 	internalServerErrorString := string(internalServerErrorByte)

// 	conv := func(data time.Time) string {
// 		return data.Format(time.RFC3339)
// 	}

// 	dateParse := func(date string) time.Time {
// 		t, _ := time.Parse(time.DateOnly, date)
// 		return t
// 	}

// 	corporationSetting := []entity.Corporation{
// 		{
// 			CorporationID: "efec6797-d0a5-c81f-3ff0-11f2eecf4a01",
// 			Name:          "Test Taro",
// 			Domain:        "",
// 			Number:        0,
// 			CorpType:      "",
// 		},
// 	}

// 	resMember := []model.Member{
// 		{
// 			MemberId:          "efec6797-d0a5-c81f-3ff0-11f2eecf4a01",
// 			Name:              "Test Taro",
// 			Email:             "user1@test.com",
// 			MemberNumber:      "00006",
// 			Status:            "employed",
// 			WorkingType:       "正社員",
// 			Department:        "開発部",
// 			Position:          "代表",
// 			BizEstablishments: "株式会社〇〇",
// 			JoinDate:          conv(dateParse("2001-01-01")),
// 			LeaveDate:         conv(dateParse("2023-07-14")),
// 			IsDeleted:         false,
// 			DeletedAt:         "",
// 			CreatedAt:         time.Now().Format(time.RFC3339),
// 			UpdatedAt:         time.Now().Format(time.RFC3339),
// 		},
// 		{
// 			MemberId:          "efec6797-d0a5-c81f-3ff0-11f2eecf4a02",
// 			Name:              "Test Jun",
// 			Email:             "user2@test.com",
// 			MemberNumber:      "00007",
// 			Status:            "employed",
// 			WorkingType:       "正社員",
// 			Department:        "開発部",
// 			Position:          "代表",
// 			BizEstablishments: "株式会社〇〇",
// 			JoinDate:          conv(dateParse("2001-01-01")),
// 			LeaveDate:         conv(dateParse("2023-07-14")),
// 			IsDeleted:         true,
// 			DeletedAt:         conv(dateParse("2023-07-18")),
// 			CreatedAt:         time.Now().Format(time.RFC3339),
// 			UpdatedAt:         time.Now().Format(time.RFC3339),
// 		},
// 	}

// 	resMembers := model.Members{
// 		Total:                    int32(2),
// 		LastUpdateSequenceNumber: int32(1),
// 		Members:                  resMember,
// 	}

// 	m, _ := json.Marshal(resMembers)
// 	wantMember := string(m)

// 	tests := []struct {
// 		name         string
// 		argGenFn     func() *gin.Context
// 		mock         func()
// 		wantCode     int
// 		wantMembers  model.Members
// 		wantResponse string
// 	}{
// 		{
// 			name: "[正常系] 全ての値が適切に設定されている場合、200ステータスを返却する",
// 			argGenFn: func() *gin.Context {
// 				response = httptest.NewRecorder()
// 				ctx, _ := gin.CreateTestContext(response)
// 				mockRequest, _ := http.NewRequest("GET",
// 					"/corporations/"+baseCorpID+"/members_diff", nil)
// 				ctx.Request = mockRequest
// 				ctx.Params = []gin.Param{
// 					{
// 						Key:   "corporation_id",
// 						Value: baseCorpID,
// 					},
// 				}
// 				q := ctx.Request.URL.Query()
// 				q.Add("update_sequence_number", fmt.Sprintf("%d", baseUpdateSequenceNumber))
// 				ctx.Request.URL.RawQuery = q.Encode()
// 				return ctx
// 			},
// 			mock: func() {
// 				uc.EXPECT().GetMembersAfterUpdateSequenceNumber(gomock.Any(), baseCorpID,
// 					baseUpdateSequenceNumber).Return(settingMembers, nil).Times(1)
// 			},
// 			wantCode:     200,
// 			wantResponse: wantMember,
// 		},
// 		{
// 			name: "[異常系] クエリパラメーターが不正な場合、validation_failed のレスポンスを返却する",
// 			argGenFn: func() *gin.Context {
// 				response = httptest.NewRecorder()
// 				ctx, _ := gin.CreateTestContext(response)
// 				mockRequest, _ := http.NewRequest("GET",
// 					"/corporations/"+baseCorpID+"/members_diff", nil)
// 				ctx.Request = mockRequest
// 				ctx.Params = []gin.Param{
// 					{
// 						Key:   "corporation_id",
// 						Value: baseCorpID,
// 					},
// 				}
// 				q := ctx.Request.URL.Query()
// 				q.Add("update_sequence_number", "")
// 				ctx.Request.URL.RawQuery = q.Encode()
// 				return ctx
// 			},
// 			mock: func() {
// 				// 呼び出しはされない
// 			},
// 			wantCode:     http.StatusBadRequest,
// 			wantResponse: `{"error_code":"validation_failed","errors":[{"error_code":"validation_failed","message":"update_sequence_number cannot be empty.","field":"update_sequence_number","reason":"presence"}]}`,
// 		},
// 		{
// 			name: "[異常系] メンバー一覧取得で失敗した場合、メンバー一覧取得で発生したエラータイプのレスポンスを返却する",
// 			argGenFn: func() *gin.Context {
// 				response = httptest.NewRecorder()
// 				ctx, _ := gin.CreateTestContext(response)
// 				mockRequest, _ := http.NewRequest("GET",
// 					"/corporations/"+baseCorpID+"/members_diff", nil)
// 				ctx.Request = mockRequest
// 				ctx.Params = []gin.Param{
// 					{
// 						Key:   "corporation_id",
// 						Value: baseCorpID,
// 					},
// 				}
// 				q := ctx.Request.URL.Query()
// 				q.Add("update_sequence_number", fmt.Sprintf("%d", baseUpdateSequenceNumber))
// 				ctx.Request.URL.RawQuery = q.Encode()
// 				return ctx
// 			},
// 			mock: func() {
// 				uc.EXPECT().GetMembersAfterUpdateSequenceNumber(gomock.Any(), baseCorpID,
// 					baseUpdateSequenceNumber).Return(nil, apperr.ErrorInternalServerError{})
// 			},
// 			wantCode:     http.StatusInternalServerError,
// 			wantResponse: internalServerErrorString,
// 		},
// 	}
// 	for _, tt := range tests {
// 		tt := tt
// 		t.Run(tt.name, func(t *testing.T) {
// 			r, newUC := initMemberRoutes(t)
// 			uc = newUC
// 			tt.mock()

// 			// ginのコンテキストの作成とresponseコードの作成
// 			argCtx := tt.argGenFn()

// 			// テスト対象関数呼び出し
// 			r.GetMembersAfterUpdateSequenceNumber(argCtx)

// 			// StatusCode Check
// 			if d := cmp.Diff(response.Code, tt.wantCode); len(d) != 0 {
// 				t.Errorf("status Code diffs: (-got +want)\n%s", d)
// 				return
// 			}

// 			// レスポンスボディーチェック
// 			resBody, _ := io.ReadAll(response.Body)
// 			if d := cmp.Diff(string(resBody), tt.wantResponse); len(d) != 0 {
// 				t.Errorf("body diffs: (-got +want)\n%s", d)
// 				return
// 			}
// 		})
// 	}
// }
