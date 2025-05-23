// Code generated by MockGen. DO NOT EDIT.
// Source: blog/internal/blog/biz/post (interfaces: PostBiz)
//
// Generated by this command:
//
//	mockgen -package=post -destination=./internal/blog/biz/post/mock_post.go blog/internal/blog/biz/post PostBiz
//

// Package post is a generated GoMock package.
package post

import (
	v1 "blog/pkg/api/blog/v1"
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockPostBiz is a mock of PostBiz interface.
type MockPostBiz struct {
	ctrl     *gomock.Controller
	recorder *MockPostBizMockRecorder
	isgomock struct{}
}

// MockPostBizMockRecorder is the mock recorder for MockPostBiz.
type MockPostBizMockRecorder struct {
	mock *MockPostBiz
}

// NewMockPostBiz creates a new mock instance.
func NewMockPostBiz(ctrl *gomock.Controller) *MockPostBiz {
	mock := &MockPostBiz{ctrl: ctrl}
	mock.recorder = &MockPostBizMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPostBiz) EXPECT() *MockPostBizMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockPostBiz) Create(ctx context.Context, username string, r *v1.CreatePostRequest) (*v1.CreatePostResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, username, r)
	ret0, _ := ret[0].(*v1.CreatePostResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockPostBizMockRecorder) Create(ctx, username, r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPostBiz)(nil).Create), ctx, username, r)
}

// Delete mocks base method.
func (m *MockPostBiz) Delete(ctx context.Context, username, postID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, username, postID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockPostBizMockRecorder) Delete(ctx, username, postID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPostBiz)(nil).Delete), ctx, username, postID)
}

// DeleteCollection mocks base method.
func (m *MockPostBiz) DeleteCollection(ctx context.Context, username string, postIDs []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCollection", ctx, username, postIDs)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection.
func (mr *MockPostBizMockRecorder) DeleteCollection(ctx, username, postIDs any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*MockPostBiz)(nil).DeleteCollection), ctx, username, postIDs)
}

// Get mocks base method.
func (m *MockPostBiz) Get(ctx context.Context, username, postID string) (*v1.GetPostResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, username, postID)
	ret0, _ := ret[0].(*v1.GetPostResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockPostBizMockRecorder) Get(ctx, username, postID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPostBiz)(nil).Get), ctx, username, postID)
}

// List mocks base method.
func (m *MockPostBiz) List(ctx context.Context, username string, offset, limit int) (*v1.ListPostResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, username, offset, limit)
	ret0, _ := ret[0].(*v1.ListPostResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockPostBizMockRecorder) List(ctx, username, offset, limit any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockPostBiz)(nil).List), ctx, username, offset, limit)
}

// Update mocks base method.
func (m *MockPostBiz) Update(ctx context.Context, username, postID string, r *v1.UpdatePostRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, username, postID, r)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockPostBizMockRecorder) Update(ctx, username, postID, r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPostBiz)(nil).Update), ctx, username, postID, r)
}
