// Code generated by counterfeiter. DO NOT EDIT.
package apifakes

import (
	"sync"

	"code.cloudfoundry.org/cli/cf/api"
	"code.cloudfoundry.org/cli/cf/models"
)

type FakeServiceBrokerRepository struct {
	CreateStub        func(string, string, string, string, string) error
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
		arg5 string
	}
	createReturns struct {
		result1 error
	}
	createReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteStub        func(string) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 string
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	FindByGUIDStub        func(string) (models.ServiceBroker, error)
	findByGUIDMutex       sync.RWMutex
	findByGUIDArgsForCall []struct {
		arg1 string
	}
	findByGUIDReturns struct {
		result1 models.ServiceBroker
		result2 error
	}
	findByGUIDReturnsOnCall map[int]struct {
		result1 models.ServiceBroker
		result2 error
	}
	FindByNameStub        func(string) (models.ServiceBroker, error)
	findByNameMutex       sync.RWMutex
	findByNameArgsForCall []struct {
		arg1 string
	}
	findByNameReturns struct {
		result1 models.ServiceBroker
		result2 error
	}
	findByNameReturnsOnCall map[int]struct {
		result1 models.ServiceBroker
		result2 error
	}
	ListServiceBrokersStub        func(func(models.ServiceBroker) bool) error
	listServiceBrokersMutex       sync.RWMutex
	listServiceBrokersArgsForCall []struct {
		arg1 func(models.ServiceBroker) bool
	}
	listServiceBrokersReturns struct {
		result1 error
	}
	listServiceBrokersReturnsOnCall map[int]struct {
		result1 error
	}
	RenameStub        func(string, string) error
	renameMutex       sync.RWMutex
	renameArgsForCall []struct {
		arg1 string
		arg2 string
	}
	renameReturns struct {
		result1 error
	}
	renameReturnsOnCall map[int]struct {
		result1 error
	}
	UpdateStub        func(models.ServiceBroker) error
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		arg1 models.ServiceBroker
	}
	updateReturns struct {
		result1 error
	}
	updateReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeServiceBrokerRepository) Create(arg1 string, arg2 string, arg3 string, arg4 string, arg5 string) error {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
		arg5 string
	}{arg1, arg2, arg3, arg4, arg5})
	fake.recordInvocation("Create", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.createReturns
	return fakeReturns.result1
}

func (fake *FakeServiceBrokerRepository) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeServiceBrokerRepository) CreateCalls(stub func(string, string, string, string, string) error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *FakeServiceBrokerRepository) CreateArgsForCall(i int) (string, string, string, string, string) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeServiceBrokerRepository) CreateReturns(result1 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeServiceBrokerRepository) CreateReturnsOnCall(i int, result1 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeServiceBrokerRepository) Delete(arg1 string) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Delete", []interface{}{arg1})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteReturns
	return fakeReturns.result1
}

func (fake *FakeServiceBrokerRepository) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeServiceBrokerRepository) DeleteCalls(stub func(string) error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *FakeServiceBrokerRepository) DeleteArgsForCall(i int) string {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeServiceBrokerRepository) DeleteReturns(result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeServiceBrokerRepository) DeleteReturnsOnCall(i int, result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeServiceBrokerRepository) FindByGUID(arg1 string) (models.ServiceBroker, error) {
	fake.findByGUIDMutex.Lock()
	ret, specificReturn := fake.findByGUIDReturnsOnCall[len(fake.findByGUIDArgsForCall)]
	fake.findByGUIDArgsForCall = append(fake.findByGUIDArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("FindByGUID", []interface{}{arg1})
	fake.findByGUIDMutex.Unlock()
	if fake.FindByGUIDStub != nil {
		return fake.FindByGUIDStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.findByGUIDReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeServiceBrokerRepository) FindByGUIDCallCount() int {
	fake.findByGUIDMutex.RLock()
	defer fake.findByGUIDMutex.RUnlock()
	return len(fake.findByGUIDArgsForCall)
}

func (fake *FakeServiceBrokerRepository) FindByGUIDCalls(stub func(string) (models.ServiceBroker, error)) {
	fake.findByGUIDMutex.Lock()
	defer fake.findByGUIDMutex.Unlock()
	fake.FindByGUIDStub = stub
}

func (fake *FakeServiceBrokerRepository) FindByGUIDArgsForCall(i int) string {
	fake.findByGUIDMutex.RLock()
	defer fake.findByGUIDMutex.RUnlock()
	argsForCall := fake.findByGUIDArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeServiceBrokerRepository) FindByGUIDReturns(result1 models.ServiceBroker, result2 error) {
	fake.findByGUIDMutex.Lock()
	defer fake.findByGUIDMutex.Unlock()
	fake.FindByGUIDStub = nil
	fake.findByGUIDReturns = struct {
		result1 models.ServiceBroker
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBrokerRepository) FindByGUIDReturnsOnCall(i int, result1 models.ServiceBroker, result2 error) {
	fake.findByGUIDMutex.Lock()
	defer fake.findByGUIDMutex.Unlock()
	fake.FindByGUIDStub = nil
	if fake.findByGUIDReturnsOnCall == nil {
		fake.findByGUIDReturnsOnCall = make(map[int]struct {
			result1 models.ServiceBroker
			result2 error
		})
	}
	fake.findByGUIDReturnsOnCall[i] = struct {
		result1 models.ServiceBroker
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBrokerRepository) FindByName(arg1 string) (models.ServiceBroker, error) {
	fake.findByNameMutex.Lock()
	ret, specificReturn := fake.findByNameReturnsOnCall[len(fake.findByNameArgsForCall)]
	fake.findByNameArgsForCall = append(fake.findByNameArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("FindByName", []interface{}{arg1})
	fake.findByNameMutex.Unlock()
	if fake.FindByNameStub != nil {
		return fake.FindByNameStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.findByNameReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeServiceBrokerRepository) FindByNameCallCount() int {
	fake.findByNameMutex.RLock()
	defer fake.findByNameMutex.RUnlock()
	return len(fake.findByNameArgsForCall)
}

func (fake *FakeServiceBrokerRepository) FindByNameCalls(stub func(string) (models.ServiceBroker, error)) {
	fake.findByNameMutex.Lock()
	defer fake.findByNameMutex.Unlock()
	fake.FindByNameStub = stub
}

func (fake *FakeServiceBrokerRepository) FindByNameArgsForCall(i int) string {
	fake.findByNameMutex.RLock()
	defer fake.findByNameMutex.RUnlock()
	argsForCall := fake.findByNameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeServiceBrokerRepository) FindByNameReturns(result1 models.ServiceBroker, result2 error) {
	fake.findByNameMutex.Lock()
	defer fake.findByNameMutex.Unlock()
	fake.FindByNameStub = nil
	fake.findByNameReturns = struct {
		result1 models.ServiceBroker
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBrokerRepository) FindByNameReturnsOnCall(i int, result1 models.ServiceBroker, result2 error) {
	fake.findByNameMutex.Lock()
	defer fake.findByNameMutex.Unlock()
	fake.FindByNameStub = nil
	if fake.findByNameReturnsOnCall == nil {
		fake.findByNameReturnsOnCall = make(map[int]struct {
			result1 models.ServiceBroker
			result2 error
		})
	}
	fake.findByNameReturnsOnCall[i] = struct {
		result1 models.ServiceBroker
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBrokerRepository) ListServiceBrokers(arg1 func(models.ServiceBroker) bool) error {
	fake.listServiceBrokersMutex.Lock()
	ret, specificReturn := fake.listServiceBrokersReturnsOnCall[len(fake.listServiceBrokersArgsForCall)]
	fake.listServiceBrokersArgsForCall = append(fake.listServiceBrokersArgsForCall, struct {
		arg1 func(models.ServiceBroker) bool
	}{arg1})
	fake.recordInvocation("ListServiceBrokers", []interface{}{arg1})
	fake.listServiceBrokersMutex.Unlock()
	if fake.ListServiceBrokersStub != nil {
		return fake.ListServiceBrokersStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.listServiceBrokersReturns
	return fakeReturns.result1
}

func (fake *FakeServiceBrokerRepository) ListServiceBrokersCallCount() int {
	fake.listServiceBrokersMutex.RLock()
	defer fake.listServiceBrokersMutex.RUnlock()
	return len(fake.listServiceBrokersArgsForCall)
}

func (fake *FakeServiceBrokerRepository) ListServiceBrokersCalls(stub func(func(models.ServiceBroker) bool) error) {
	fake.listServiceBrokersMutex.Lock()
	defer fake.listServiceBrokersMutex.Unlock()
	fake.ListServiceBrokersStub = stub
}

func (fake *FakeServiceBrokerRepository) ListServiceBrokersArgsForCall(i int) func(models.ServiceBroker) bool {
	fake.listServiceBrokersMutex.RLock()
	defer fake.listServiceBrokersMutex.RUnlock()
	argsForCall := fake.listServiceBrokersArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeServiceBrokerRepository) ListServiceBrokersReturns(result1 error) {
	fake.listServiceBrokersMutex.Lock()
	defer fake.listServiceBrokersMutex.Unlock()
	fake.ListServiceBrokersStub = nil
	fake.listServiceBrokersReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeServiceBrokerRepository) ListServiceBrokersReturnsOnCall(i int, result1 error) {
	fake.listServiceBrokersMutex.Lock()
	defer fake.listServiceBrokersMutex.Unlock()
	fake.ListServiceBrokersStub = nil
	if fake.listServiceBrokersReturnsOnCall == nil {
		fake.listServiceBrokersReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.listServiceBrokersReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeServiceBrokerRepository) Rename(arg1 string, arg2 string) error {
	fake.renameMutex.Lock()
	ret, specificReturn := fake.renameReturnsOnCall[len(fake.renameArgsForCall)]
	fake.renameArgsForCall = append(fake.renameArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("Rename", []interface{}{arg1, arg2})
	fake.renameMutex.Unlock()
	if fake.RenameStub != nil {
		return fake.RenameStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.renameReturns
	return fakeReturns.result1
}

func (fake *FakeServiceBrokerRepository) RenameCallCount() int {
	fake.renameMutex.RLock()
	defer fake.renameMutex.RUnlock()
	return len(fake.renameArgsForCall)
}

func (fake *FakeServiceBrokerRepository) RenameCalls(stub func(string, string) error) {
	fake.renameMutex.Lock()
	defer fake.renameMutex.Unlock()
	fake.RenameStub = stub
}

func (fake *FakeServiceBrokerRepository) RenameArgsForCall(i int) (string, string) {
	fake.renameMutex.RLock()
	defer fake.renameMutex.RUnlock()
	argsForCall := fake.renameArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeServiceBrokerRepository) RenameReturns(result1 error) {
	fake.renameMutex.Lock()
	defer fake.renameMutex.Unlock()
	fake.RenameStub = nil
	fake.renameReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeServiceBrokerRepository) RenameReturnsOnCall(i int, result1 error) {
	fake.renameMutex.Lock()
	defer fake.renameMutex.Unlock()
	fake.RenameStub = nil
	if fake.renameReturnsOnCall == nil {
		fake.renameReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.renameReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeServiceBrokerRepository) Update(arg1 models.ServiceBroker) error {
	fake.updateMutex.Lock()
	ret, specificReturn := fake.updateReturnsOnCall[len(fake.updateArgsForCall)]
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		arg1 models.ServiceBroker
	}{arg1})
	fake.recordInvocation("Update", []interface{}{arg1})
	fake.updateMutex.Unlock()
	if fake.UpdateStub != nil {
		return fake.UpdateStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.updateReturns
	return fakeReturns.result1
}

func (fake *FakeServiceBrokerRepository) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeServiceBrokerRepository) UpdateCalls(stub func(models.ServiceBroker) error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = stub
}

func (fake *FakeServiceBrokerRepository) UpdateArgsForCall(i int) models.ServiceBroker {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	argsForCall := fake.updateArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeServiceBrokerRepository) UpdateReturns(result1 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeServiceBrokerRepository) UpdateReturnsOnCall(i int, result1 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	if fake.updateReturnsOnCall == nil {
		fake.updateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeServiceBrokerRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.findByGUIDMutex.RLock()
	defer fake.findByGUIDMutex.RUnlock()
	fake.findByNameMutex.RLock()
	defer fake.findByNameMutex.RUnlock()
	fake.listServiceBrokersMutex.RLock()
	defer fake.listServiceBrokersMutex.RUnlock()
	fake.renameMutex.RLock()
	defer fake.renameMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeServiceBrokerRepository) recordInvocation(key string, args []interface{}) {
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

var _ api.ServiceBrokerRepository = new(FakeServiceBrokerRepository)