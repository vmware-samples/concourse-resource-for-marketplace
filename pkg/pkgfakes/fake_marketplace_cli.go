// Code generated by counterfeiter. DO NOT EDIT.
package pkgfakes

import (
	"sync"
	"vmware-samples/concourse-resource-for-marketplace/m/v2/pkg"
)

type FakeMarketplaceCLI struct {
	DownloadAssetStub        func(string) error
	downloadAssetMutex       sync.RWMutex
	downloadAssetArgsForCall []struct {
		arg1 string
	}
	downloadAssetReturns struct {
		result1 error
	}
	downloadAssetReturnsOnCall map[int]struct {
		result1 error
	}
	GetInputSlugStub        func() string
	getInputSlugMutex       sync.RWMutex
	getInputSlugArgsForCall []struct {
	}
	getInputSlugReturns struct {
		result1 string
	}
	getInputSlugReturnsOnCall map[int]struct {
		result1 string
	}
	GetInputVersionStub        func() *pkg.Version
	getInputVersionMutex       sync.RWMutex
	getInputVersionArgsForCall []struct {
	}
	getInputVersionReturns struct {
		result1 *pkg.Version
	}
	getInputVersionReturnsOnCall map[int]struct {
		result1 *pkg.Version
	}
	GetProductJSONStub        func() ([]byte, error)
	getProductJSONMutex       sync.RWMutex
	getProductJSONArgsForCall []struct {
	}
	getProductJSONReturns struct {
		result1 []byte
		result2 error
	}
	getProductJSONReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	GetVersionsStub        func() ([]*pkg.Version, error)
	getVersionsMutex       sync.RWMutex
	getVersionsArgsForCall []struct {
	}
	getVersionsReturns struct {
		result1 []*pkg.Version
		result2 error
	}
	getVersionsReturnsOnCall map[int]struct {
		result1 []*pkg.Version
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeMarketplaceCLI) DownloadAsset(arg1 string) error {
	fake.downloadAssetMutex.Lock()
	ret, specificReturn := fake.downloadAssetReturnsOnCall[len(fake.downloadAssetArgsForCall)]
	fake.downloadAssetArgsForCall = append(fake.downloadAssetArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.DownloadAssetStub
	fakeReturns := fake.downloadAssetReturns
	fake.recordInvocation("DownloadAsset", []interface{}{arg1})
	fake.downloadAssetMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeMarketplaceCLI) DownloadAssetCallCount() int {
	fake.downloadAssetMutex.RLock()
	defer fake.downloadAssetMutex.RUnlock()
	return len(fake.downloadAssetArgsForCall)
}

func (fake *FakeMarketplaceCLI) DownloadAssetCalls(stub func(string) error) {
	fake.downloadAssetMutex.Lock()
	defer fake.downloadAssetMutex.Unlock()
	fake.DownloadAssetStub = stub
}

func (fake *FakeMarketplaceCLI) DownloadAssetArgsForCall(i int) string {
	fake.downloadAssetMutex.RLock()
	defer fake.downloadAssetMutex.RUnlock()
	argsForCall := fake.downloadAssetArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeMarketplaceCLI) DownloadAssetReturns(result1 error) {
	fake.downloadAssetMutex.Lock()
	defer fake.downloadAssetMutex.Unlock()
	fake.DownloadAssetStub = nil
	fake.downloadAssetReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeMarketplaceCLI) DownloadAssetReturnsOnCall(i int, result1 error) {
	fake.downloadAssetMutex.Lock()
	defer fake.downloadAssetMutex.Unlock()
	fake.DownloadAssetStub = nil
	if fake.downloadAssetReturnsOnCall == nil {
		fake.downloadAssetReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.downloadAssetReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeMarketplaceCLI) GetInputSlug() string {
	fake.getInputSlugMutex.Lock()
	ret, specificReturn := fake.getInputSlugReturnsOnCall[len(fake.getInputSlugArgsForCall)]
	fake.getInputSlugArgsForCall = append(fake.getInputSlugArgsForCall, struct {
	}{})
	stub := fake.GetInputSlugStub
	fakeReturns := fake.getInputSlugReturns
	fake.recordInvocation("GetInputSlug", []interface{}{})
	fake.getInputSlugMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeMarketplaceCLI) GetInputSlugCallCount() int {
	fake.getInputSlugMutex.RLock()
	defer fake.getInputSlugMutex.RUnlock()
	return len(fake.getInputSlugArgsForCall)
}

func (fake *FakeMarketplaceCLI) GetInputSlugCalls(stub func() string) {
	fake.getInputSlugMutex.Lock()
	defer fake.getInputSlugMutex.Unlock()
	fake.GetInputSlugStub = stub
}

func (fake *FakeMarketplaceCLI) GetInputSlugReturns(result1 string) {
	fake.getInputSlugMutex.Lock()
	defer fake.getInputSlugMutex.Unlock()
	fake.GetInputSlugStub = nil
	fake.getInputSlugReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeMarketplaceCLI) GetInputSlugReturnsOnCall(i int, result1 string) {
	fake.getInputSlugMutex.Lock()
	defer fake.getInputSlugMutex.Unlock()
	fake.GetInputSlugStub = nil
	if fake.getInputSlugReturnsOnCall == nil {
		fake.getInputSlugReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.getInputSlugReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeMarketplaceCLI) GetInputVersion() *pkg.Version {
	fake.getInputVersionMutex.Lock()
	ret, specificReturn := fake.getInputVersionReturnsOnCall[len(fake.getInputVersionArgsForCall)]
	fake.getInputVersionArgsForCall = append(fake.getInputVersionArgsForCall, struct {
	}{})
	stub := fake.GetInputVersionStub
	fakeReturns := fake.getInputVersionReturns
	fake.recordInvocation("GetInputVersion", []interface{}{})
	fake.getInputVersionMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeMarketplaceCLI) GetInputVersionCallCount() int {
	fake.getInputVersionMutex.RLock()
	defer fake.getInputVersionMutex.RUnlock()
	return len(fake.getInputVersionArgsForCall)
}

func (fake *FakeMarketplaceCLI) GetInputVersionCalls(stub func() *pkg.Version) {
	fake.getInputVersionMutex.Lock()
	defer fake.getInputVersionMutex.Unlock()
	fake.GetInputVersionStub = stub
}

func (fake *FakeMarketplaceCLI) GetInputVersionReturns(result1 *pkg.Version) {
	fake.getInputVersionMutex.Lock()
	defer fake.getInputVersionMutex.Unlock()
	fake.GetInputVersionStub = nil
	fake.getInputVersionReturns = struct {
		result1 *pkg.Version
	}{result1}
}

func (fake *FakeMarketplaceCLI) GetInputVersionReturnsOnCall(i int, result1 *pkg.Version) {
	fake.getInputVersionMutex.Lock()
	defer fake.getInputVersionMutex.Unlock()
	fake.GetInputVersionStub = nil
	if fake.getInputVersionReturnsOnCall == nil {
		fake.getInputVersionReturnsOnCall = make(map[int]struct {
			result1 *pkg.Version
		})
	}
	fake.getInputVersionReturnsOnCall[i] = struct {
		result1 *pkg.Version
	}{result1}
}

func (fake *FakeMarketplaceCLI) GetProductJSON() ([]byte, error) {
	fake.getProductJSONMutex.Lock()
	ret, specificReturn := fake.getProductJSONReturnsOnCall[len(fake.getProductJSONArgsForCall)]
	fake.getProductJSONArgsForCall = append(fake.getProductJSONArgsForCall, struct {
	}{})
	stub := fake.GetProductJSONStub
	fakeReturns := fake.getProductJSONReturns
	fake.recordInvocation("GetProductJSON", []interface{}{})
	fake.getProductJSONMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeMarketplaceCLI) GetProductJSONCallCount() int {
	fake.getProductJSONMutex.RLock()
	defer fake.getProductJSONMutex.RUnlock()
	return len(fake.getProductJSONArgsForCall)
}

func (fake *FakeMarketplaceCLI) GetProductJSONCalls(stub func() ([]byte, error)) {
	fake.getProductJSONMutex.Lock()
	defer fake.getProductJSONMutex.Unlock()
	fake.GetProductJSONStub = stub
}

func (fake *FakeMarketplaceCLI) GetProductJSONReturns(result1 []byte, result2 error) {
	fake.getProductJSONMutex.Lock()
	defer fake.getProductJSONMutex.Unlock()
	fake.GetProductJSONStub = nil
	fake.getProductJSONReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeMarketplaceCLI) GetProductJSONReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.getProductJSONMutex.Lock()
	defer fake.getProductJSONMutex.Unlock()
	fake.GetProductJSONStub = nil
	if fake.getProductJSONReturnsOnCall == nil {
		fake.getProductJSONReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.getProductJSONReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeMarketplaceCLI) GetVersions() ([]*pkg.Version, error) {
	fake.getVersionsMutex.Lock()
	ret, specificReturn := fake.getVersionsReturnsOnCall[len(fake.getVersionsArgsForCall)]
	fake.getVersionsArgsForCall = append(fake.getVersionsArgsForCall, struct {
	}{})
	stub := fake.GetVersionsStub
	fakeReturns := fake.getVersionsReturns
	fake.recordInvocation("GetVersions", []interface{}{})
	fake.getVersionsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeMarketplaceCLI) GetVersionsCallCount() int {
	fake.getVersionsMutex.RLock()
	defer fake.getVersionsMutex.RUnlock()
	return len(fake.getVersionsArgsForCall)
}

func (fake *FakeMarketplaceCLI) GetVersionsCalls(stub func() ([]*pkg.Version, error)) {
	fake.getVersionsMutex.Lock()
	defer fake.getVersionsMutex.Unlock()
	fake.GetVersionsStub = stub
}

func (fake *FakeMarketplaceCLI) GetVersionsReturns(result1 []*pkg.Version, result2 error) {
	fake.getVersionsMutex.Lock()
	defer fake.getVersionsMutex.Unlock()
	fake.GetVersionsStub = nil
	fake.getVersionsReturns = struct {
		result1 []*pkg.Version
		result2 error
	}{result1, result2}
}

func (fake *FakeMarketplaceCLI) GetVersionsReturnsOnCall(i int, result1 []*pkg.Version, result2 error) {
	fake.getVersionsMutex.Lock()
	defer fake.getVersionsMutex.Unlock()
	fake.GetVersionsStub = nil
	if fake.getVersionsReturnsOnCall == nil {
		fake.getVersionsReturnsOnCall = make(map[int]struct {
			result1 []*pkg.Version
			result2 error
		})
	}
	fake.getVersionsReturnsOnCall[i] = struct {
		result1 []*pkg.Version
		result2 error
	}{result1, result2}
}

func (fake *FakeMarketplaceCLI) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.downloadAssetMutex.RLock()
	defer fake.downloadAssetMutex.RUnlock()
	fake.getInputSlugMutex.RLock()
	defer fake.getInputSlugMutex.RUnlock()
	fake.getInputVersionMutex.RLock()
	defer fake.getInputVersionMutex.RUnlock()
	fake.getProductJSONMutex.RLock()
	defer fake.getProductJSONMutex.RUnlock()
	fake.getVersionsMutex.RLock()
	defer fake.getVersionsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeMarketplaceCLI) recordInvocation(key string, args []interface{}) {
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

var _ pkg.MarketplaceCLI = new(FakeMarketplaceCLI)