package mock

import (
	"fmt"
	"testing"

	"github.com/koizumi55555/corporation-api/pkg/logger"

	gomock "github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

func GetMockLogger(t *testing.T) *logger.Logger {
	t.Helper()
	l := logger.New("debug")
	return l
}

// GetMockMasterRepo
func GetMockMasterRepo(t *testing.T) (*MockMasterRepository, *logger.Logger) {
	t.Helper()
	l := logger.New("debug")

	mockCtl := gomock.NewController(t)
	mRepo := NewMockMasterRepository(mockCtl)

	return mRepo, l
}

// custom matcher
type gocmpMatcher struct {
	want    interface{}
	options []cmp.Option
	diff    string
}

func NewGocmpMatcher(want interface{}, options []cmp.Option) gomock.Matcher {
	return &gocmpMatcher{
		want:    want,
		options: options,
	}
}

func (m *gocmpMatcher) String() string {
	return fmt.Sprintf("diff(-got +want) %s", m.diff)
}

func (m *gocmpMatcher) Matches(x interface{}) bool {
	m.diff = cmp.Diff(x, m.want, m.options...)
	return m.diff == ""
}
