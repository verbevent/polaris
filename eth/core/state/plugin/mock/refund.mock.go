// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"github.com/berachain/stargazer/eth/core/state"
	"sync"
)

// Ensure, that RefundPluginMock does implement state.RefundPlugin.
// If this is not the case, regenerate this file with moq.
var _ state.RefundPlugin = &RefundPluginMock{}

// RefundPluginMock is a mock implementation of state.RefundPlugin.
//
//	func TestSomethingThatUsesRefundPlugin(t *testing.T) {
//
//		// make and configure a mocked state.RefundPlugin
//		mockedRefundPlugin := &RefundPluginMock{
//			AddRefundFunc: func(gas uint64)  {
//				panic("mock out the AddRefund method")
//			},
//			FinalizeFunc: func()  {
//				panic("mock out the Finalize method")
//			},
//			GetRefundFunc: func() uint64 {
//				panic("mock out the GetRefund method")
//			},
//			RegistryKeyFunc: func() string {
//				panic("mock out the RegistryKey method")
//			},
//			RevertToSnapshotFunc: func(n int)  {
//				panic("mock out the RevertToSnapshot method")
//			},
//			SnapshotFunc: func() int {
//				panic("mock out the Snapshot method")
//			},
//			SubRefundFunc: func(gas uint64)  {
//				panic("mock out the SubRefund method")
//			},
//		}
//
//		// use mockedRefundPlugin in code that requires state.RefundPlugin
//		// and then make assertions.
//
//	}
type RefundPluginMock struct {
	// AddRefundFunc mocks the AddRefund method.
	AddRefundFunc func(gas uint64)

	// FinalizeFunc mocks the Finalize method.
	FinalizeFunc func()

	// GetRefundFunc mocks the GetRefund method.
	GetRefundFunc func() uint64

	// RegistryKeyFunc mocks the RegistryKey method.
	RegistryKeyFunc func() string

	// RevertToSnapshotFunc mocks the RevertToSnapshot method.
	RevertToSnapshotFunc func(n int)

	// SnapshotFunc mocks the Snapshot method.
	SnapshotFunc func() int

	// SubRefundFunc mocks the SubRefund method.
	SubRefundFunc func(gas uint64)

	// calls tracks calls to the methods.
	calls struct {
		// AddRefund holds details about calls to the AddRefund method.
		AddRefund []struct {
			// Gas is the gas argument value.
			Gas uint64
		}
		// Finalize holds details about calls to the Finalize method.
		Finalize []struct {
		}
		// GetRefund holds details about calls to the GetRefund method.
		GetRefund []struct {
		}
		// RegistryKey holds details about calls to the RegistryKey method.
		RegistryKey []struct {
		}
		// RevertToSnapshot holds details about calls to the RevertToSnapshot method.
		RevertToSnapshot []struct {
			// N is the n argument value.
			N int
		}
		// Snapshot holds details about calls to the Snapshot method.
		Snapshot []struct {
		}
		// SubRefund holds details about calls to the SubRefund method.
		SubRefund []struct {
			// Gas is the gas argument value.
			Gas uint64
		}
	}
	lockAddRefund        sync.RWMutex
	lockFinalize         sync.RWMutex
	lockGetRefund        sync.RWMutex
	lockRegistryKey      sync.RWMutex
	lockRevertToSnapshot sync.RWMutex
	lockSnapshot         sync.RWMutex
	lockSubRefund        sync.RWMutex
}

// AddRefund calls AddRefundFunc.
func (mock *RefundPluginMock) AddRefund(gas uint64) {
	if mock.AddRefundFunc == nil {
		panic("RefundPluginMock.AddRefundFunc: method is nil but RefundPlugin.AddRefund was just called")
	}
	callInfo := struct {
		Gas uint64
	}{
		Gas: gas,
	}
	mock.lockAddRefund.Lock()
	mock.calls.AddRefund = append(mock.calls.AddRefund, callInfo)
	mock.lockAddRefund.Unlock()
	mock.AddRefundFunc(gas)
}

// AddRefundCalls gets all the calls that were made to AddRefund.
// Check the length with:
//
//	len(mockedRefundPlugin.AddRefundCalls())
func (mock *RefundPluginMock) AddRefundCalls() []struct {
	Gas uint64
} {
	var calls []struct {
		Gas uint64
	}
	mock.lockAddRefund.RLock()
	calls = mock.calls.AddRefund
	mock.lockAddRefund.RUnlock()
	return calls
}

// Finalize calls FinalizeFunc.
func (mock *RefundPluginMock) Finalize() {
	if mock.FinalizeFunc == nil {
		panic("RefundPluginMock.FinalizeFunc: method is nil but RefundPlugin.Finalize was just called")
	}
	callInfo := struct {
	}{}
	mock.lockFinalize.Lock()
	mock.calls.Finalize = append(mock.calls.Finalize, callInfo)
	mock.lockFinalize.Unlock()
	mock.FinalizeFunc()
}

// FinalizeCalls gets all the calls that were made to Finalize.
// Check the length with:
//
//	len(mockedRefundPlugin.FinalizeCalls())
func (mock *RefundPluginMock) FinalizeCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockFinalize.RLock()
	calls = mock.calls.Finalize
	mock.lockFinalize.RUnlock()
	return calls
}

// GetRefund calls GetRefundFunc.
func (mock *RefundPluginMock) GetRefund() uint64 {
	if mock.GetRefundFunc == nil {
		panic("RefundPluginMock.GetRefundFunc: method is nil but RefundPlugin.GetRefund was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetRefund.Lock()
	mock.calls.GetRefund = append(mock.calls.GetRefund, callInfo)
	mock.lockGetRefund.Unlock()
	return mock.GetRefundFunc()
}

// GetRefundCalls gets all the calls that were made to GetRefund.
// Check the length with:
//
//	len(mockedRefundPlugin.GetRefundCalls())
func (mock *RefundPluginMock) GetRefundCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetRefund.RLock()
	calls = mock.calls.GetRefund
	mock.lockGetRefund.RUnlock()
	return calls
}

// RegistryKey calls RegistryKeyFunc.
func (mock *RefundPluginMock) RegistryKey() string {
	if mock.RegistryKeyFunc == nil {
		panic("RefundPluginMock.RegistryKeyFunc: method is nil but RefundPlugin.RegistryKey was just called")
	}
	callInfo := struct {
	}{}
	mock.lockRegistryKey.Lock()
	mock.calls.RegistryKey = append(mock.calls.RegistryKey, callInfo)
	mock.lockRegistryKey.Unlock()
	return mock.RegistryKeyFunc()
}

// RegistryKeyCalls gets all the calls that were made to RegistryKey.
// Check the length with:
//
//	len(mockedRefundPlugin.RegistryKeyCalls())
func (mock *RefundPluginMock) RegistryKeyCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockRegistryKey.RLock()
	calls = mock.calls.RegistryKey
	mock.lockRegistryKey.RUnlock()
	return calls
}

// RevertToSnapshot calls RevertToSnapshotFunc.
func (mock *RefundPluginMock) RevertToSnapshot(n int) {
	if mock.RevertToSnapshotFunc == nil {
		panic("RefundPluginMock.RevertToSnapshotFunc: method is nil but RefundPlugin.RevertToSnapshot was just called")
	}
	callInfo := struct {
		N int
	}{
		N: n,
	}
	mock.lockRevertToSnapshot.Lock()
	mock.calls.RevertToSnapshot = append(mock.calls.RevertToSnapshot, callInfo)
	mock.lockRevertToSnapshot.Unlock()
	mock.RevertToSnapshotFunc(n)
}

// RevertToSnapshotCalls gets all the calls that were made to RevertToSnapshot.
// Check the length with:
//
//	len(mockedRefundPlugin.RevertToSnapshotCalls())
func (mock *RefundPluginMock) RevertToSnapshotCalls() []struct {
	N int
} {
	var calls []struct {
		N int
	}
	mock.lockRevertToSnapshot.RLock()
	calls = mock.calls.RevertToSnapshot
	mock.lockRevertToSnapshot.RUnlock()
	return calls
}

// Snapshot calls SnapshotFunc.
func (mock *RefundPluginMock) Snapshot() int {
	if mock.SnapshotFunc == nil {
		panic("RefundPluginMock.SnapshotFunc: method is nil but RefundPlugin.Snapshot was just called")
	}
	callInfo := struct {
	}{}
	mock.lockSnapshot.Lock()
	mock.calls.Snapshot = append(mock.calls.Snapshot, callInfo)
	mock.lockSnapshot.Unlock()
	return mock.SnapshotFunc()
}

// SnapshotCalls gets all the calls that were made to Snapshot.
// Check the length with:
//
//	len(mockedRefundPlugin.SnapshotCalls())
func (mock *RefundPluginMock) SnapshotCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockSnapshot.RLock()
	calls = mock.calls.Snapshot
	mock.lockSnapshot.RUnlock()
	return calls
}

// SubRefund calls SubRefundFunc.
func (mock *RefundPluginMock) SubRefund(gas uint64) {
	if mock.SubRefundFunc == nil {
		panic("RefundPluginMock.SubRefundFunc: method is nil but RefundPlugin.SubRefund was just called")
	}
	callInfo := struct {
		Gas uint64
	}{
		Gas: gas,
	}
	mock.lockSubRefund.Lock()
	mock.calls.SubRefund = append(mock.calls.SubRefund, callInfo)
	mock.lockSubRefund.Unlock()
	mock.SubRefundFunc(gas)
}

// SubRefundCalls gets all the calls that were made to SubRefund.
// Check the length with:
//
//	len(mockedRefundPlugin.SubRefundCalls())
func (mock *RefundPluginMock) SubRefundCalls() []struct {
	Gas uint64
} {
	var calls []struct {
		Gas uint64
	}
	mock.lockSubRefund.RLock()
	calls = mock.calls.SubRefund
	mock.lockSubRefund.RUnlock()
	return calls
}
