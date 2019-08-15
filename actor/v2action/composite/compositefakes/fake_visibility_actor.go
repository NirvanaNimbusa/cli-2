// Code generated by counterfeiter. DO NOT EDIT.
package compositefakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v2action"
	"code.cloudfoundry.org/cli/actor/v2action/composite"
)

type FakeVisibilityActor struct {
	GetServicePlanVisibilitiesStub        func(string) ([]v2action.ServicePlanVisibility, v2action.Warnings, error)
	getServicePlanVisibilitiesMutex       sync.RWMutex
	getServicePlanVisibilitiesArgsForCall []struct {
		arg1 string
	}
	getServicePlanVisibilitiesReturns struct {
		result1 []v2action.ServicePlanVisibility
		result2 v2action.Warnings
		result3 error
	}
	getServicePlanVisibilitiesReturnsOnCall map[int]struct {
		result1 []v2action.ServicePlanVisibility
		result2 v2action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeVisibilityActor) GetServicePlanVisibilities(arg1 string) ([]v2action.ServicePlanVisibility, v2action.Warnings, error) {
	fake.getServicePlanVisibilitiesMutex.Lock()
	ret, specificReturn := fake.getServicePlanVisibilitiesReturnsOnCall[len(fake.getServicePlanVisibilitiesArgsForCall)]
	fake.getServicePlanVisibilitiesArgsForCall = append(fake.getServicePlanVisibilitiesArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetServicePlanVisibilities", []interface{}{arg1})
	fake.getServicePlanVisibilitiesMutex.Unlock()
	if fake.GetServicePlanVisibilitiesStub != nil {
		return fake.GetServicePlanVisibilitiesStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getServicePlanVisibilitiesReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeVisibilityActor) GetServicePlanVisibilitiesCallCount() int {
	fake.getServicePlanVisibilitiesMutex.RLock()
	defer fake.getServicePlanVisibilitiesMutex.RUnlock()
	return len(fake.getServicePlanVisibilitiesArgsForCall)
}

func (fake *FakeVisibilityActor) GetServicePlanVisibilitiesCalls(stub func(string) ([]v2action.ServicePlanVisibility, v2action.Warnings, error)) {
	fake.getServicePlanVisibilitiesMutex.Lock()
	defer fake.getServicePlanVisibilitiesMutex.Unlock()
	fake.GetServicePlanVisibilitiesStub = stub
}

func (fake *FakeVisibilityActor) GetServicePlanVisibilitiesArgsForCall(i int) string {
	fake.getServicePlanVisibilitiesMutex.RLock()
	defer fake.getServicePlanVisibilitiesMutex.RUnlock()
	argsForCall := fake.getServicePlanVisibilitiesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeVisibilityActor) GetServicePlanVisibilitiesReturns(result1 []v2action.ServicePlanVisibility, result2 v2action.Warnings, result3 error) {
	fake.getServicePlanVisibilitiesMutex.Lock()
	defer fake.getServicePlanVisibilitiesMutex.Unlock()
	fake.GetServicePlanVisibilitiesStub = nil
	fake.getServicePlanVisibilitiesReturns = struct {
		result1 []v2action.ServicePlanVisibility
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeVisibilityActor) GetServicePlanVisibilitiesReturnsOnCall(i int, result1 []v2action.ServicePlanVisibility, result2 v2action.Warnings, result3 error) {
	fake.getServicePlanVisibilitiesMutex.Lock()
	defer fake.getServicePlanVisibilitiesMutex.Unlock()
	fake.GetServicePlanVisibilitiesStub = nil
	if fake.getServicePlanVisibilitiesReturnsOnCall == nil {
		fake.getServicePlanVisibilitiesReturnsOnCall = make(map[int]struct {
			result1 []v2action.ServicePlanVisibility
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.getServicePlanVisibilitiesReturnsOnCall[i] = struct {
		result1 []v2action.ServicePlanVisibility
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeVisibilityActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getServicePlanVisibilitiesMutex.RLock()
	defer fake.getServicePlanVisibilitiesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeVisibilityActor) recordInvocation(key string, args []interface{}) {
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

var _ composite.VisibilityActor = new(FakeVisibilityActor)