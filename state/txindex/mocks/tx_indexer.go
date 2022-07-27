// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	query "github.com/tendermint/tendermint/libs/pubsub/query"

	txindex "github.com/tendermint/tendermint/state/txindex"

	types "github.com/tendermint/tendermint/abci/types"
)

// TxIndexer is an autogenerated mock type for the TxIndexer type
type TxIndexer struct {
	mock.Mock
}

// AddBatch provides a mock function with given fields: b
func (_m *TxIndexer) AddBatch(b *txindex.Batch) error {
	ret := _m.Called(b)

	var r0 error
	if rf, ok := ret.Get(0).(func(*txindex.Batch) error); ok {
		r0 = rf(b)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: hash
func (_m *TxIndexer) Get(hash []byte) (*types.TxResult, error) {
	ret := _m.Called(hash)

	var r0 *types.TxResult
	if rf, ok := ret.Get(0).(func([]byte) *types.TxResult); ok {
		r0 = rf(hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.TxResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Index provides a mock function with given fields: result
func (_m *TxIndexer) Index(result *types.TxResult) error {
	ret := _m.Called(result)

	var r0 error
	if rf, ok := ret.Get(0).(func(*types.TxResult) error); ok {
		r0 = rf(result)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Search provides a mock function with given fields: ctx, q
func (_m *TxIndexer) Search(ctx context.Context, q *query.Query) ([]*types.TxResult, error) {
	ret := _m.Called(ctx, q)

	var r0 []*types.TxResult
	if rf, ok := ret.Get(0).(func(context.Context, *query.Query) []*types.TxResult); ok {
		r0 = rf(ctx, q)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.TxResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *query.Query) error); ok {
		r1 = rf(ctx, q)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTxIndexer interface {
	mock.TestingT
	Cleanup(func())
}

// NewTxIndexer creates a new instance of TxIndexer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTxIndexer(t mockConstructorTestingTNewTxIndexer) *TxIndexer {
	mock := &TxIndexer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
