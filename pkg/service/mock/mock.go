package mock_service

import (
        models "gin_news/pkg/models"
        reflect "reflect"

        gomock "github.com/golang/mock/gomock"
)

// MockAuthorization is a mock of Authorization interface.
type MockAuthorization struct {
        ctrl     *gomock.Controller
        recorder *MockAuthorizationMockRecorder
}

// MockAuthorizationMockRecorder is the mock recorder for MockAuthorization.
type MockAuthorizationMockRecorder struct {
        mock *MockAuthorization
}

// NewMockAuthorization creates a new mock instance.
func NewMockAuthorization(ctrl *gomock.Controller) *MockAuthorization {
        mock := &MockAuthorization{ctrl: ctrl}
        mock.recorder = &MockAuthorizationMockRecorder{mock}
        return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorization) EXPECT() *MockAuthorizationMockRecorder {
        return m.recorder
}

// CheckAccess mocks base method.
func (m *MockAuthorization) CheckAccess(id int) (bool, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "CheckAccess", id)
        ret0, _ := ret[0].(bool)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// CheckAccess indicates an expected call of CheckAccess.
func (mr *MockAuthorizationMockRecorder) CheckAccess(id interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckAccess", reflect.TypeOf((*MockAuthorization)(nil).CheckAccess), id)
}

// CreateUser mocks base method.
func (m *MockAuthorization) CreateUser(user models.User) (int, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "CreateUser", user)
        ret0, _ := ret[0].(int)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockAuthorizationMockRecorder) CreateUser(user interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockAuthorization)(nil).CreateUser), user)
}

// GenerateToken mocks base method.
func (m *MockAuthorization) GenerateToken(username, password string) (string, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GenerateToken", username, password)
        ret0, _ := ret[0].(string)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockAuthorizationMockRecorder) GenerateToken(username, password interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockAuthorization)(nil).GenerateToken), username, password)
}

// ParseToken mocks base method.
func (m *MockAuthorization) ParseToken(token string) (int, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "ParseToken", token)
        ret0, _ := ret[0].(int)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// ParseToken indicates an expected call of ParseToken.
func (mr *MockAuthorizationMockRecorder) ParseToken(token interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseToken", reflect.TypeOf((*MockAuthorization)(nil).ParseToken), token)
}

// MockNewslist is a mock of Newslist interface.
type MockNewslist struct {
        ctrl     *gomock.Controller
        recorder *MockNewslistMockRecorder
}

// MockNewslistMockRecorder is the mock recorder for MockNewslist.
type MockNewslistMockRecorder struct {
        mock *MockNewslist
}

// NewMockNewslist creates a new mock instance.
func NewMockNewslist(ctrl *gomock.Controller) *MockNewslist {
        mock := &MockNewslist{ctrl: ctrl}
        mock.recorder = &MockNewslistMockRecorder{mock}
        return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNewslist) EXPECT() *MockNewslistMockRecorder {
        return m.recorder
}

// Create mocks base method.
func (m *MockNewslist) Create(news models.News) (int, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "Create", news)
        ret0, _ := ret[0].(int)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockNewslistMockRecorder) Create(news interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockNewslist)(nil).Create), news)
}

// DeleteNews mocks base method.
func (m *MockNewslist) DeleteNews(id int) error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "DeleteNews", id)
        ret0, _ := ret[0].(error)
        return ret0
}

// DeleteNews indicates an expected call of DeleteNews.
func (mr *MockNewslistMockRecorder) DeleteNews(id interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNews", reflect.TypeOf((*MockNewslist)(nil).DeleteNews), id)
}

// GetAll mocks base method.
func (m *MockNewslist) GetAll() ([]models.News, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetAll")
        ret0, _ := ret[0].([]models.News)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockNewslistMockRecorder) GetAll() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockNewslist)(nil).GetAll))
}

// GetByIdNews mocks base method.
func (m *MockNewslist) GetByIdNews(id int) (models.News, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetByIdNews", id)
        ret0, _ := ret[0].(models.News)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetByIdNews indicates an expected call of GetByIdNews.
func (mr *MockNewslistMockRecorder) GetByIdNews(id interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIdNews", reflect.TypeOf((*MockNewslist)(nil).GetByIdNews), id)
}

// UpdateNews mocks base method.
func (m *MockNewslist) UpdateNews(id int, input models.UpdateNews) error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "UpdateNews", id, input)
        ret0, _ := ret[0].(error)
        return ret0
}

// UpdateNews indicates an expected call of UpdateNews.
func (mr *MockNewslistMockRecorder) UpdateNews(id, input interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNews", reflect.TypeOf((*MockNewslist)(nil).UpdateNews), id, input)
}