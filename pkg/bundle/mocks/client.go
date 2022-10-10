// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/bundle/client.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	v1alpha1 "github.com/aws/eks-anywhere-packages/api/v1alpha1"
	gomock "github.com/golang/mock/gomock"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// CreateBundle mocks base method.
func (m *MockClient) CreateBundle(ctx context.Context, bundle *v1alpha1.PackageBundle) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBundle", ctx, bundle)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateBundle indicates an expected call of CreateBundle.
func (mr *MockClientMockRecorder) CreateBundle(ctx, bundle interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBundle", reflect.TypeOf((*MockClient)(nil).CreateBundle), ctx, bundle)
}

// CreateClusterConfigMap mocks base method.
func (m *MockClient) CreateClusterConfigMap(ctx context.Context, clusterName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateClusterConfigMap", ctx, clusterName)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateClusterConfigMap indicates an expected call of CreateClusterConfigMap.
func (mr *MockClientMockRecorder) CreateClusterConfigMap(ctx, clusterName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateClusterConfigMap", reflect.TypeOf((*MockClient)(nil).CreateClusterConfigMap), ctx, clusterName)
}

// CreateClusterNamespace mocks base method.
func (m *MockClient) CreateClusterNamespace(ctx context.Context, clusterName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateClusterNamespace", ctx, clusterName)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateClusterNamespace indicates an expected call of CreateClusterNamespace.
func (mr *MockClientMockRecorder) CreateClusterNamespace(ctx, clusterName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateClusterNamespace", reflect.TypeOf((*MockClient)(nil).CreateClusterNamespace), ctx, clusterName)
}

// GetActiveBundle mocks base method.
func (m *MockClient) GetActiveBundle(ctx context.Context, clusterName string) (*v1alpha1.PackageBundle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActiveBundle", ctx, clusterName)
	ret0, _ := ret[0].(*v1alpha1.PackageBundle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActiveBundle indicates an expected call of GetActiveBundle.
func (mr *MockClientMockRecorder) GetActiveBundle(ctx, clusterName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActiveBundle", reflect.TypeOf((*MockClient)(nil).GetActiveBundle), ctx, clusterName)
}

// GetBundle mocks base method.
func (m *MockClient) GetBundle(ctx context.Context, name string) (*v1alpha1.PackageBundle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBundle", ctx, name)
	ret0, _ := ret[0].(*v1alpha1.PackageBundle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBundle indicates an expected call of GetBundle.
func (mr *MockClientMockRecorder) GetBundle(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBundle", reflect.TypeOf((*MockClient)(nil).GetBundle), ctx, name)
}

// GetBundleList mocks base method.
func (m *MockClient) GetBundleList(ctx context.Context) ([]v1alpha1.PackageBundle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBundleList", ctx)
	ret0, _ := ret[0].([]v1alpha1.PackageBundle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBundleList indicates an expected call of GetBundleList.
func (mr *MockClientMockRecorder) GetBundleList(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBundleList", reflect.TypeOf((*MockClient)(nil).GetBundleList), ctx)
}

// GetPackageBundleController mocks base method.
func (m *MockClient) GetPackageBundleController(ctx context.Context, clusterName string) (*v1alpha1.PackageBundleController, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPackageBundleController", ctx, clusterName)
	ret0, _ := ret[0].(*v1alpha1.PackageBundleController)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPackageBundleController indicates an expected call of GetPackageBundleController.
func (mr *MockClientMockRecorder) GetPackageBundleController(ctx, clusterName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPackageBundleController", reflect.TypeOf((*MockClient)(nil).GetPackageBundleController), ctx, clusterName)
}

// Save mocks base method.
func (m *MockClient) Save(ctx context.Context, object client.Object) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, object)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockClientMockRecorder) Save(ctx, object interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockClient)(nil).Save), ctx, object)
}

// SaveStatus mocks base method.
func (m *MockClient) SaveStatus(ctx context.Context, object client.Object) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveStatus", ctx, object)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveStatus indicates an expected call of SaveStatus.
func (mr *MockClientMockRecorder) SaveStatus(ctx, object interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveStatus", reflect.TypeOf((*MockClient)(nil).SaveStatus), ctx, object)
}
