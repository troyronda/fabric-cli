// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/hyperledger/fabric-cli/pkg/fabric"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel/invoke"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
)

type Channel struct {
	ExecuteStub        func(request channel.Request, options ...channel.RequestOption) (channel.Response, error)
	executeMutex       sync.RWMutex
	executeArgsForCall []struct {
		request channel.Request
		options []channel.RequestOption
	}
	executeReturns struct {
		result1 channel.Response
		result2 error
	}
	executeReturnsOnCall map[int]struct {
		result1 channel.Response
		result2 error
	}
	InvokeHandlerStub        func(handler invoke.Handler, request channel.Request, options ...channel.RequestOption) (channel.Response, error)
	invokeHandlerMutex       sync.RWMutex
	invokeHandlerArgsForCall []struct {
		handler invoke.Handler
		request channel.Request
		options []channel.RequestOption
	}
	invokeHandlerReturns struct {
		result1 channel.Response
		result2 error
	}
	invokeHandlerReturnsOnCall map[int]struct {
		result1 channel.Response
		result2 error
	}
	QueryStub        func(request channel.Request, options ...channel.RequestOption) (channel.Response, error)
	queryMutex       sync.RWMutex
	queryArgsForCall []struct {
		request channel.Request
		options []channel.RequestOption
	}
	queryReturns struct {
		result1 channel.Response
		result2 error
	}
	queryReturnsOnCall map[int]struct {
		result1 channel.Response
		result2 error
	}
	RegisterChaincodeEventStub        func(chainCodeID string, eventFilter string) (fab.Registration, <-chan *fab.CCEvent, error)
	registerChaincodeEventMutex       sync.RWMutex
	registerChaincodeEventArgsForCall []struct {
		chainCodeID string
		eventFilter string
	}
	registerChaincodeEventReturns struct {
		result1 fab.Registration
		result2 <-chan *fab.CCEvent
		result3 error
	}
	registerChaincodeEventReturnsOnCall map[int]struct {
		result1 fab.Registration
		result2 <-chan *fab.CCEvent
		result3 error
	}
	UnregisterChaincodeEventStub        func(registration fab.Registration)
	unregisterChaincodeEventMutex       sync.RWMutex
	unregisterChaincodeEventArgsForCall []struct {
		registration fab.Registration
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Channel) Execute(request channel.Request, options ...channel.RequestOption) (channel.Response, error) {
	fake.executeMutex.Lock()
	ret, specificReturn := fake.executeReturnsOnCall[len(fake.executeArgsForCall)]
	fake.executeArgsForCall = append(fake.executeArgsForCall, struct {
		request channel.Request
		options []channel.RequestOption
	}{request, options})
	fake.recordInvocation("Execute", []interface{}{request, options})
	fake.executeMutex.Unlock()
	if fake.ExecuteStub != nil {
		return fake.ExecuteStub(request, options...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.executeReturns.result1, fake.executeReturns.result2
}

func (fake *Channel) ExecuteCallCount() int {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return len(fake.executeArgsForCall)
}

func (fake *Channel) ExecuteArgsForCall(i int) (channel.Request, []channel.RequestOption) {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return fake.executeArgsForCall[i].request, fake.executeArgsForCall[i].options
}

func (fake *Channel) ExecuteReturns(result1 channel.Response, result2 error) {
	fake.ExecuteStub = nil
	fake.executeReturns = struct {
		result1 channel.Response
		result2 error
	}{result1, result2}
}

func (fake *Channel) ExecuteReturnsOnCall(i int, result1 channel.Response, result2 error) {
	fake.ExecuteStub = nil
	if fake.executeReturnsOnCall == nil {
		fake.executeReturnsOnCall = make(map[int]struct {
			result1 channel.Response
			result2 error
		})
	}
	fake.executeReturnsOnCall[i] = struct {
		result1 channel.Response
		result2 error
	}{result1, result2}
}

func (fake *Channel) InvokeHandler(handler invoke.Handler, request channel.Request, options ...channel.RequestOption) (channel.Response, error) {
	fake.invokeHandlerMutex.Lock()
	ret, specificReturn := fake.invokeHandlerReturnsOnCall[len(fake.invokeHandlerArgsForCall)]
	fake.invokeHandlerArgsForCall = append(fake.invokeHandlerArgsForCall, struct {
		handler invoke.Handler
		request channel.Request
		options []channel.RequestOption
	}{handler, request, options})
	fake.recordInvocation("InvokeHandler", []interface{}{handler, request, options})
	fake.invokeHandlerMutex.Unlock()
	if fake.InvokeHandlerStub != nil {
		return fake.InvokeHandlerStub(handler, request, options...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.invokeHandlerReturns.result1, fake.invokeHandlerReturns.result2
}

func (fake *Channel) InvokeHandlerCallCount() int {
	fake.invokeHandlerMutex.RLock()
	defer fake.invokeHandlerMutex.RUnlock()
	return len(fake.invokeHandlerArgsForCall)
}

func (fake *Channel) InvokeHandlerArgsForCall(i int) (invoke.Handler, channel.Request, []channel.RequestOption) {
	fake.invokeHandlerMutex.RLock()
	defer fake.invokeHandlerMutex.RUnlock()
	return fake.invokeHandlerArgsForCall[i].handler, fake.invokeHandlerArgsForCall[i].request, fake.invokeHandlerArgsForCall[i].options
}

func (fake *Channel) InvokeHandlerReturns(result1 channel.Response, result2 error) {
	fake.InvokeHandlerStub = nil
	fake.invokeHandlerReturns = struct {
		result1 channel.Response
		result2 error
	}{result1, result2}
}

func (fake *Channel) InvokeHandlerReturnsOnCall(i int, result1 channel.Response, result2 error) {
	fake.InvokeHandlerStub = nil
	if fake.invokeHandlerReturnsOnCall == nil {
		fake.invokeHandlerReturnsOnCall = make(map[int]struct {
			result1 channel.Response
			result2 error
		})
	}
	fake.invokeHandlerReturnsOnCall[i] = struct {
		result1 channel.Response
		result2 error
	}{result1, result2}
}

func (fake *Channel) Query(request channel.Request, options ...channel.RequestOption) (channel.Response, error) {
	fake.queryMutex.Lock()
	ret, specificReturn := fake.queryReturnsOnCall[len(fake.queryArgsForCall)]
	fake.queryArgsForCall = append(fake.queryArgsForCall, struct {
		request channel.Request
		options []channel.RequestOption
	}{request, options})
	fake.recordInvocation("Query", []interface{}{request, options})
	fake.queryMutex.Unlock()
	if fake.QueryStub != nil {
		return fake.QueryStub(request, options...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.queryReturns.result1, fake.queryReturns.result2
}

func (fake *Channel) QueryCallCount() int {
	fake.queryMutex.RLock()
	defer fake.queryMutex.RUnlock()
	return len(fake.queryArgsForCall)
}

func (fake *Channel) QueryArgsForCall(i int) (channel.Request, []channel.RequestOption) {
	fake.queryMutex.RLock()
	defer fake.queryMutex.RUnlock()
	return fake.queryArgsForCall[i].request, fake.queryArgsForCall[i].options
}

func (fake *Channel) QueryReturns(result1 channel.Response, result2 error) {
	fake.QueryStub = nil
	fake.queryReturns = struct {
		result1 channel.Response
		result2 error
	}{result1, result2}
}

func (fake *Channel) QueryReturnsOnCall(i int, result1 channel.Response, result2 error) {
	fake.QueryStub = nil
	if fake.queryReturnsOnCall == nil {
		fake.queryReturnsOnCall = make(map[int]struct {
			result1 channel.Response
			result2 error
		})
	}
	fake.queryReturnsOnCall[i] = struct {
		result1 channel.Response
		result2 error
	}{result1, result2}
}

func (fake *Channel) RegisterChaincodeEvent(chainCodeID string, eventFilter string) (fab.Registration, <-chan *fab.CCEvent, error) {
	fake.registerChaincodeEventMutex.Lock()
	ret, specificReturn := fake.registerChaincodeEventReturnsOnCall[len(fake.registerChaincodeEventArgsForCall)]
	fake.registerChaincodeEventArgsForCall = append(fake.registerChaincodeEventArgsForCall, struct {
		chainCodeID string
		eventFilter string
	}{chainCodeID, eventFilter})
	fake.recordInvocation("RegisterChaincodeEvent", []interface{}{chainCodeID, eventFilter})
	fake.registerChaincodeEventMutex.Unlock()
	if fake.RegisterChaincodeEventStub != nil {
		return fake.RegisterChaincodeEventStub(chainCodeID, eventFilter)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.registerChaincodeEventReturns.result1, fake.registerChaincodeEventReturns.result2, fake.registerChaincodeEventReturns.result3
}

func (fake *Channel) RegisterChaincodeEventCallCount() int {
	fake.registerChaincodeEventMutex.RLock()
	defer fake.registerChaincodeEventMutex.RUnlock()
	return len(fake.registerChaincodeEventArgsForCall)
}

func (fake *Channel) RegisterChaincodeEventArgsForCall(i int) (string, string) {
	fake.registerChaincodeEventMutex.RLock()
	defer fake.registerChaincodeEventMutex.RUnlock()
	return fake.registerChaincodeEventArgsForCall[i].chainCodeID, fake.registerChaincodeEventArgsForCall[i].eventFilter
}

func (fake *Channel) RegisterChaincodeEventReturns(result1 fab.Registration, result2 <-chan *fab.CCEvent, result3 error) {
	fake.RegisterChaincodeEventStub = nil
	fake.registerChaincodeEventReturns = struct {
		result1 fab.Registration
		result2 <-chan *fab.CCEvent
		result3 error
	}{result1, result2, result3}
}

func (fake *Channel) RegisterChaincodeEventReturnsOnCall(i int, result1 fab.Registration, result2 <-chan *fab.CCEvent, result3 error) {
	fake.RegisterChaincodeEventStub = nil
	if fake.registerChaincodeEventReturnsOnCall == nil {
		fake.registerChaincodeEventReturnsOnCall = make(map[int]struct {
			result1 fab.Registration
			result2 <-chan *fab.CCEvent
			result3 error
		})
	}
	fake.registerChaincodeEventReturnsOnCall[i] = struct {
		result1 fab.Registration
		result2 <-chan *fab.CCEvent
		result3 error
	}{result1, result2, result3}
}

func (fake *Channel) UnregisterChaincodeEvent(registration fab.Registration) {
	fake.unregisterChaincodeEventMutex.Lock()
	fake.unregisterChaincodeEventArgsForCall = append(fake.unregisterChaincodeEventArgsForCall, struct {
		registration fab.Registration
	}{registration})
	fake.recordInvocation("UnregisterChaincodeEvent", []interface{}{registration})
	fake.unregisterChaincodeEventMutex.Unlock()
	if fake.UnregisterChaincodeEventStub != nil {
		fake.UnregisterChaincodeEventStub(registration)
	}
}

func (fake *Channel) UnregisterChaincodeEventCallCount() int {
	fake.unregisterChaincodeEventMutex.RLock()
	defer fake.unregisterChaincodeEventMutex.RUnlock()
	return len(fake.unregisterChaincodeEventArgsForCall)
}

func (fake *Channel) UnregisterChaincodeEventArgsForCall(i int) fab.Registration {
	fake.unregisterChaincodeEventMutex.RLock()
	defer fake.unregisterChaincodeEventMutex.RUnlock()
	return fake.unregisterChaincodeEventArgsForCall[i].registration
}

func (fake *Channel) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	fake.invokeHandlerMutex.RLock()
	defer fake.invokeHandlerMutex.RUnlock()
	fake.queryMutex.RLock()
	defer fake.queryMutex.RUnlock()
	fake.registerChaincodeEventMutex.RLock()
	defer fake.registerChaincodeEventMutex.RUnlock()
	fake.unregisterChaincodeEventMutex.RLock()
	defer fake.unregisterChaincodeEventMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Channel) recordInvocation(key string, args []interface{}) {
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

var _ fabric.Channel = new(Channel)
