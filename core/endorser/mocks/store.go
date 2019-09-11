// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import fabric_protos_gotransientstore "github.com/hyperledger/fabric-protos-go/transientstore"
import ledger "github.com/hyperledger/fabric/core/ledger"
import mock "github.com/stretchr/testify/mock"
import transientstore "github.com/hyperledger/fabric/core/transientstore"

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

// GetMinTransientBlkHt provides a mock function with given fields:
func (_m *Store) GetMinTransientBlkHt() (uint64, error) {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTxPvtRWSetByTxid provides a mock function with given fields: txid, filter
func (_m *Store) GetTxPvtRWSetByTxid(txid string, filter ledger.PvtNsCollFilter) (transientstore.RWSetScanner, error) {
	ret := _m.Called(txid, filter)

	var r0 transientstore.RWSetScanner
	if rf, ok := ret.Get(0).(func(string, ledger.PvtNsCollFilter) transientstore.RWSetScanner); ok {
		r0 = rf(txid, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(transientstore.RWSetScanner)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, ledger.PvtNsCollFilter) error); ok {
		r1 = rf(txid, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PersistWithConfig provides a mock function with given fields: txid, blockHeight, privateSimulationResultsWithConfig
func (_m *Store) PersistWithConfig(txid string, blockHeight uint64, privateSimulationResultsWithConfig *fabric_protos_gotransientstore.TxPvtReadWriteSetWithConfigInfo) error {
	ret := _m.Called(txid, blockHeight, privateSimulationResultsWithConfig)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, uint64, *fabric_protos_gotransientstore.TxPvtReadWriteSetWithConfigInfo) error); ok {
		r0 = rf(txid, blockHeight, privateSimulationResultsWithConfig)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PurgeByHeight provides a mock function with given fields: maxBlockNumToRetain
func (_m *Store) PurgeByHeight(maxBlockNumToRetain uint64) error {
	ret := _m.Called(maxBlockNumToRetain)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64) error); ok {
		r0 = rf(maxBlockNumToRetain)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PurgeByTxids provides a mock function with given fields: txids
func (_m *Store) PurgeByTxids(txids []string) error {
	ret := _m.Called(txids)

	var r0 error
	if rf, ok := ret.Get(0).(func([]string) error); ok {
		r0 = rf(txids)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Shutdown provides a mock function with given fields:
func (_m *Store) Shutdown() {
	_m.Called()
}
