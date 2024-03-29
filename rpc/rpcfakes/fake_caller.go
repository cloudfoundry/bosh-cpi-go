// This file was generated by counterfeiter
package rpcfakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-cpi-go/rpc"
)

type FakeCaller struct {
	CallStub        func(interface{}, []interface{}) (interface{}, error)
	callMutex       sync.RWMutex
	callArgsForCall []struct {
		arg1 interface{}
		arg2 []interface{}
	}
	callReturns struct {
		result1 interface{}
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCaller) Call(arg1 interface{}, arg2 []interface{}) (interface{}, error) {
	var arg2Copy []interface{}
	if arg2 != nil {
		arg2Copy = make([]interface{}, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.callMutex.Lock()
	fake.callArgsForCall = append(fake.callArgsForCall, struct {
		arg1 interface{}
		arg2 []interface{}
	}{arg1, arg2Copy})
	fake.recordInvocation("Call", []interface{}{arg1, arg2Copy})
	fake.callMutex.Unlock()
	if fake.CallStub != nil {
		return fake.CallStub(arg1, arg2)
	}
	return fake.callReturns.result1, fake.callReturns.result2
}

func (fake *FakeCaller) CallCallCount() int {
	fake.callMutex.RLock()
	defer fake.callMutex.RUnlock()
	return len(fake.callArgsForCall)
}

func (fake *FakeCaller) CallArgsForCall(i int) (interface{}, []interface{}) {
	fake.callMutex.RLock()
	defer fake.callMutex.RUnlock()
	return fake.callArgsForCall[i].arg1, fake.callArgsForCall[i].arg2
}

func (fake *FakeCaller) CallReturns(result1 interface{}, result2 error) {
	fake.CallStub = nil
	fake.callReturns = struct {
		result1 interface{}
		result2 error
	}{result1, result2}
}

func (fake *FakeCaller) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.callMutex.RLock()
	defer fake.callMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeCaller) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ rpc.Caller = new(FakeCaller)
