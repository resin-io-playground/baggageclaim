// Code generated by counterfeiter. DO NOT EDIT.
package volumefakes

import (
	"io"
	"sync"

	"github.com/concourse/baggageclaim/volume"
)

type FakeRepository struct {
	ListVolumesStub        func(queryProperties volume.Properties) (volume.Volumes, []string, error)
	listVolumesMutex       sync.RWMutex
	listVolumesArgsForCall []struct {
		queryProperties volume.Properties
	}
	listVolumesReturns struct {
		result1 volume.Volumes
		result2 []string
		result3 error
	}
	listVolumesReturnsOnCall map[int]struct {
		result1 volume.Volumes
		result2 []string
		result3 error
	}
	GetVolumeStub        func(handle string) (volume.Volume, bool, error)
	getVolumeMutex       sync.RWMutex
	getVolumeArgsForCall []struct {
		handle string
	}
	getVolumeReturns struct {
		result1 volume.Volume
		result2 bool
		result3 error
	}
	getVolumeReturnsOnCall map[int]struct {
		result1 volume.Volume
		result2 bool
		result3 error
	}
	CreateVolumeStub        func(handle string, strategy volume.Strategy, properties volume.Properties, ttlInSeconds uint, isPrivileged bool) (volume.Volume, error)
	createVolumeMutex       sync.RWMutex
	createVolumeArgsForCall []struct {
		handle       string
		strategy     volume.Strategy
		properties   volume.Properties
		ttlInSeconds uint
		isPrivileged bool
	}
	createVolumeReturns struct {
		result1 volume.Volume
		result2 error
	}
	createVolumeReturnsOnCall map[int]struct {
		result1 volume.Volume
		result2 error
	}
	DestroyVolumeStub        func(handle string) error
	destroyVolumeMutex       sync.RWMutex
	destroyVolumeArgsForCall []struct {
		handle string
	}
	destroyVolumeReturns struct {
		result1 error
	}
	destroyVolumeReturnsOnCall map[int]struct {
		result1 error
	}
	DestroyVolumeAndDescendantsStub        func(handle string) error
	destroyVolumeAndDescendantsMutex       sync.RWMutex
	destroyVolumeAndDescendantsArgsForCall []struct {
		handle string
	}
	destroyVolumeAndDescendantsReturns struct {
		result1 error
	}
	destroyVolumeAndDescendantsReturnsOnCall map[int]struct {
		result1 error
	}
	SetPropertyStub        func(handle string, propertyName string, propertyValue string) error
	setPropertyMutex       sync.RWMutex
	setPropertyArgsForCall []struct {
		handle        string
		propertyName  string
		propertyValue string
	}
	setPropertyReturns struct {
		result1 error
	}
	setPropertyReturnsOnCall map[int]struct {
		result1 error
	}
	SetTTLStub        func(handle string, ttl uint) error
	setTTLMutex       sync.RWMutex
	setTTLArgsForCall []struct {
		handle string
		ttl    uint
	}
	setTTLReturns struct {
		result1 error
	}
	setTTLReturnsOnCall map[int]struct {
		result1 error
	}
	SetPrivilegedStub        func(handle string, privileged bool) error
	setPrivilegedMutex       sync.RWMutex
	setPrivilegedArgsForCall []struct {
		handle     string
		privileged bool
	}
	setPrivilegedReturns struct {
		result1 error
	}
	setPrivilegedReturnsOnCall map[int]struct {
		result1 error
	}
	StreamInStub        func(handle string, path string, stream io.Reader) (bool, error)
	streamInMutex       sync.RWMutex
	streamInArgsForCall []struct {
		handle string
		path   string
		stream io.Reader
	}
	streamInReturns struct {
		result1 bool
		result2 error
	}
	streamInReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	StreamOutStub        func(handle string, path string, dest io.Writer) error
	streamOutMutex       sync.RWMutex
	streamOutArgsForCall []struct {
		handle string
		path   string
		dest   io.Writer
	}
	streamOutReturns struct {
		result1 error
	}
	streamOutReturnsOnCall map[int]struct {
		result1 error
	}
	VolumeParentStub        func(handle string) (volume.Volume, bool, error)
	volumeParentMutex       sync.RWMutex
	volumeParentArgsForCall []struct {
		handle string
	}
	volumeParentReturns struct {
		result1 volume.Volume
		result2 bool
		result3 error
	}
	volumeParentReturnsOnCall map[int]struct {
		result1 volume.Volume
		result2 bool
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRepository) ListVolumes(queryProperties volume.Properties) (volume.Volumes, []string, error) {
	fake.listVolumesMutex.Lock()
	ret, specificReturn := fake.listVolumesReturnsOnCall[len(fake.listVolumesArgsForCall)]
	fake.listVolumesArgsForCall = append(fake.listVolumesArgsForCall, struct {
		queryProperties volume.Properties
	}{queryProperties})
	fake.recordInvocation("ListVolumes", []interface{}{queryProperties})
	fake.listVolumesMutex.Unlock()
	if fake.ListVolumesStub != nil {
		return fake.ListVolumesStub(queryProperties)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.listVolumesReturns.result1, fake.listVolumesReturns.result2, fake.listVolumesReturns.result3
}

func (fake *FakeRepository) ListVolumesCallCount() int {
	fake.listVolumesMutex.RLock()
	defer fake.listVolumesMutex.RUnlock()
	return len(fake.listVolumesArgsForCall)
}

func (fake *FakeRepository) ListVolumesArgsForCall(i int) volume.Properties {
	fake.listVolumesMutex.RLock()
	defer fake.listVolumesMutex.RUnlock()
	return fake.listVolumesArgsForCall[i].queryProperties
}

func (fake *FakeRepository) ListVolumesReturns(result1 volume.Volumes, result2 []string, result3 error) {
	fake.ListVolumesStub = nil
	fake.listVolumesReturns = struct {
		result1 volume.Volumes
		result2 []string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeRepository) ListVolumesReturnsOnCall(i int, result1 volume.Volumes, result2 []string, result3 error) {
	fake.ListVolumesStub = nil
	if fake.listVolumesReturnsOnCall == nil {
		fake.listVolumesReturnsOnCall = make(map[int]struct {
			result1 volume.Volumes
			result2 []string
			result3 error
		})
	}
	fake.listVolumesReturnsOnCall[i] = struct {
		result1 volume.Volumes
		result2 []string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeRepository) GetVolume(handle string) (volume.Volume, bool, error) {
	fake.getVolumeMutex.Lock()
	ret, specificReturn := fake.getVolumeReturnsOnCall[len(fake.getVolumeArgsForCall)]
	fake.getVolumeArgsForCall = append(fake.getVolumeArgsForCall, struct {
		handle string
	}{handle})
	fake.recordInvocation("GetVolume", []interface{}{handle})
	fake.getVolumeMutex.Unlock()
	if fake.GetVolumeStub != nil {
		return fake.GetVolumeStub(handle)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getVolumeReturns.result1, fake.getVolumeReturns.result2, fake.getVolumeReturns.result3
}

func (fake *FakeRepository) GetVolumeCallCount() int {
	fake.getVolumeMutex.RLock()
	defer fake.getVolumeMutex.RUnlock()
	return len(fake.getVolumeArgsForCall)
}

func (fake *FakeRepository) GetVolumeArgsForCall(i int) string {
	fake.getVolumeMutex.RLock()
	defer fake.getVolumeMutex.RUnlock()
	return fake.getVolumeArgsForCall[i].handle
}

func (fake *FakeRepository) GetVolumeReturns(result1 volume.Volume, result2 bool, result3 error) {
	fake.GetVolumeStub = nil
	fake.getVolumeReturns = struct {
		result1 volume.Volume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeRepository) GetVolumeReturnsOnCall(i int, result1 volume.Volume, result2 bool, result3 error) {
	fake.GetVolumeStub = nil
	if fake.getVolumeReturnsOnCall == nil {
		fake.getVolumeReturnsOnCall = make(map[int]struct {
			result1 volume.Volume
			result2 bool
			result3 error
		})
	}
	fake.getVolumeReturnsOnCall[i] = struct {
		result1 volume.Volume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeRepository) CreateVolume(handle string, strategy volume.Strategy, properties volume.Properties, ttlInSeconds uint, isPrivileged bool) (volume.Volume, error) {
	fake.createVolumeMutex.Lock()
	ret, specificReturn := fake.createVolumeReturnsOnCall[len(fake.createVolumeArgsForCall)]
	fake.createVolumeArgsForCall = append(fake.createVolumeArgsForCall, struct {
		handle       string
		strategy     volume.Strategy
		properties   volume.Properties
		ttlInSeconds uint
		isPrivileged bool
	}{handle, strategy, properties, ttlInSeconds, isPrivileged})
	fake.recordInvocation("CreateVolume", []interface{}{handle, strategy, properties, ttlInSeconds, isPrivileged})
	fake.createVolumeMutex.Unlock()
	if fake.CreateVolumeStub != nil {
		return fake.CreateVolumeStub(handle, strategy, properties, ttlInSeconds, isPrivileged)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createVolumeReturns.result1, fake.createVolumeReturns.result2
}

func (fake *FakeRepository) CreateVolumeCallCount() int {
	fake.createVolumeMutex.RLock()
	defer fake.createVolumeMutex.RUnlock()
	return len(fake.createVolumeArgsForCall)
}

func (fake *FakeRepository) CreateVolumeArgsForCall(i int) (string, volume.Strategy, volume.Properties, uint, bool) {
	fake.createVolumeMutex.RLock()
	defer fake.createVolumeMutex.RUnlock()
	return fake.createVolumeArgsForCall[i].handle, fake.createVolumeArgsForCall[i].strategy, fake.createVolumeArgsForCall[i].properties, fake.createVolumeArgsForCall[i].ttlInSeconds, fake.createVolumeArgsForCall[i].isPrivileged
}

func (fake *FakeRepository) CreateVolumeReturns(result1 volume.Volume, result2 error) {
	fake.CreateVolumeStub = nil
	fake.createVolumeReturns = struct {
		result1 volume.Volume
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) CreateVolumeReturnsOnCall(i int, result1 volume.Volume, result2 error) {
	fake.CreateVolumeStub = nil
	if fake.createVolumeReturnsOnCall == nil {
		fake.createVolumeReturnsOnCall = make(map[int]struct {
			result1 volume.Volume
			result2 error
		})
	}
	fake.createVolumeReturnsOnCall[i] = struct {
		result1 volume.Volume
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) DestroyVolume(handle string) error {
	fake.destroyVolumeMutex.Lock()
	ret, specificReturn := fake.destroyVolumeReturnsOnCall[len(fake.destroyVolumeArgsForCall)]
	fake.destroyVolumeArgsForCall = append(fake.destroyVolumeArgsForCall, struct {
		handle string
	}{handle})
	fake.recordInvocation("DestroyVolume", []interface{}{handle})
	fake.destroyVolumeMutex.Unlock()
	if fake.DestroyVolumeStub != nil {
		return fake.DestroyVolumeStub(handle)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.destroyVolumeReturns.result1
}

func (fake *FakeRepository) DestroyVolumeCallCount() int {
	fake.destroyVolumeMutex.RLock()
	defer fake.destroyVolumeMutex.RUnlock()
	return len(fake.destroyVolumeArgsForCall)
}

func (fake *FakeRepository) DestroyVolumeArgsForCall(i int) string {
	fake.destroyVolumeMutex.RLock()
	defer fake.destroyVolumeMutex.RUnlock()
	return fake.destroyVolumeArgsForCall[i].handle
}

func (fake *FakeRepository) DestroyVolumeReturns(result1 error) {
	fake.DestroyVolumeStub = nil
	fake.destroyVolumeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepository) DestroyVolumeReturnsOnCall(i int, result1 error) {
	fake.DestroyVolumeStub = nil
	if fake.destroyVolumeReturnsOnCall == nil {
		fake.destroyVolumeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.destroyVolumeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepository) DestroyVolumeAndDescendants(handle string) error {
	fake.destroyVolumeAndDescendantsMutex.Lock()
	ret, specificReturn := fake.destroyVolumeAndDescendantsReturnsOnCall[len(fake.destroyVolumeAndDescendantsArgsForCall)]
	fake.destroyVolumeAndDescendantsArgsForCall = append(fake.destroyVolumeAndDescendantsArgsForCall, struct {
		handle string
	}{handle})
	fake.recordInvocation("DestroyVolumeAndDescendants", []interface{}{handle})
	fake.destroyVolumeAndDescendantsMutex.Unlock()
	if fake.DestroyVolumeAndDescendantsStub != nil {
		return fake.DestroyVolumeAndDescendantsStub(handle)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.destroyVolumeAndDescendantsReturns.result1
}

func (fake *FakeRepository) DestroyVolumeAndDescendantsCallCount() int {
	fake.destroyVolumeAndDescendantsMutex.RLock()
	defer fake.destroyVolumeAndDescendantsMutex.RUnlock()
	return len(fake.destroyVolumeAndDescendantsArgsForCall)
}

func (fake *FakeRepository) DestroyVolumeAndDescendantsArgsForCall(i int) string {
	fake.destroyVolumeAndDescendantsMutex.RLock()
	defer fake.destroyVolumeAndDescendantsMutex.RUnlock()
	return fake.destroyVolumeAndDescendantsArgsForCall[i].handle
}

func (fake *FakeRepository) DestroyVolumeAndDescendantsReturns(result1 error) {
	fake.DestroyVolumeAndDescendantsStub = nil
	fake.destroyVolumeAndDescendantsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepository) DestroyVolumeAndDescendantsReturnsOnCall(i int, result1 error) {
	fake.DestroyVolumeAndDescendantsStub = nil
	if fake.destroyVolumeAndDescendantsReturnsOnCall == nil {
		fake.destroyVolumeAndDescendantsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.destroyVolumeAndDescendantsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepository) SetProperty(handle string, propertyName string, propertyValue string) error {
	fake.setPropertyMutex.Lock()
	ret, specificReturn := fake.setPropertyReturnsOnCall[len(fake.setPropertyArgsForCall)]
	fake.setPropertyArgsForCall = append(fake.setPropertyArgsForCall, struct {
		handle        string
		propertyName  string
		propertyValue string
	}{handle, propertyName, propertyValue})
	fake.recordInvocation("SetProperty", []interface{}{handle, propertyName, propertyValue})
	fake.setPropertyMutex.Unlock()
	if fake.SetPropertyStub != nil {
		return fake.SetPropertyStub(handle, propertyName, propertyValue)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.setPropertyReturns.result1
}

func (fake *FakeRepository) SetPropertyCallCount() int {
	fake.setPropertyMutex.RLock()
	defer fake.setPropertyMutex.RUnlock()
	return len(fake.setPropertyArgsForCall)
}

func (fake *FakeRepository) SetPropertyArgsForCall(i int) (string, string, string) {
	fake.setPropertyMutex.RLock()
	defer fake.setPropertyMutex.RUnlock()
	return fake.setPropertyArgsForCall[i].handle, fake.setPropertyArgsForCall[i].propertyName, fake.setPropertyArgsForCall[i].propertyValue
}

func (fake *FakeRepository) SetPropertyReturns(result1 error) {
	fake.SetPropertyStub = nil
	fake.setPropertyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepository) SetPropertyReturnsOnCall(i int, result1 error) {
	fake.SetPropertyStub = nil
	if fake.setPropertyReturnsOnCall == nil {
		fake.setPropertyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setPropertyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepository) SetTTL(handle string, ttl uint) error {
	fake.setTTLMutex.Lock()
	ret, specificReturn := fake.setTTLReturnsOnCall[len(fake.setTTLArgsForCall)]
	fake.setTTLArgsForCall = append(fake.setTTLArgsForCall, struct {
		handle string
		ttl    uint
	}{handle, ttl})
	fake.recordInvocation("SetTTL", []interface{}{handle, ttl})
	fake.setTTLMutex.Unlock()
	if fake.SetTTLStub != nil {
		return fake.SetTTLStub(handle, ttl)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.setTTLReturns.result1
}

func (fake *FakeRepository) SetTTLCallCount() int {
	fake.setTTLMutex.RLock()
	defer fake.setTTLMutex.RUnlock()
	return len(fake.setTTLArgsForCall)
}

func (fake *FakeRepository) SetTTLArgsForCall(i int) (string, uint) {
	fake.setTTLMutex.RLock()
	defer fake.setTTLMutex.RUnlock()
	return fake.setTTLArgsForCall[i].handle, fake.setTTLArgsForCall[i].ttl
}

func (fake *FakeRepository) SetTTLReturns(result1 error) {
	fake.SetTTLStub = nil
	fake.setTTLReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepository) SetTTLReturnsOnCall(i int, result1 error) {
	fake.SetTTLStub = nil
	if fake.setTTLReturnsOnCall == nil {
		fake.setTTLReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setTTLReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepository) SetPrivileged(handle string, privileged bool) error {
	fake.setPrivilegedMutex.Lock()
	ret, specificReturn := fake.setPrivilegedReturnsOnCall[len(fake.setPrivilegedArgsForCall)]
	fake.setPrivilegedArgsForCall = append(fake.setPrivilegedArgsForCall, struct {
		handle     string
		privileged bool
	}{handle, privileged})
	fake.recordInvocation("SetPrivileged", []interface{}{handle, privileged})
	fake.setPrivilegedMutex.Unlock()
	if fake.SetPrivilegedStub != nil {
		return fake.SetPrivilegedStub(handle, privileged)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.setPrivilegedReturns.result1
}

func (fake *FakeRepository) SetPrivilegedCallCount() int {
	fake.setPrivilegedMutex.RLock()
	defer fake.setPrivilegedMutex.RUnlock()
	return len(fake.setPrivilegedArgsForCall)
}

func (fake *FakeRepository) SetPrivilegedArgsForCall(i int) (string, bool) {
	fake.setPrivilegedMutex.RLock()
	defer fake.setPrivilegedMutex.RUnlock()
	return fake.setPrivilegedArgsForCall[i].handle, fake.setPrivilegedArgsForCall[i].privileged
}

func (fake *FakeRepository) SetPrivilegedReturns(result1 error) {
	fake.SetPrivilegedStub = nil
	fake.setPrivilegedReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepository) SetPrivilegedReturnsOnCall(i int, result1 error) {
	fake.SetPrivilegedStub = nil
	if fake.setPrivilegedReturnsOnCall == nil {
		fake.setPrivilegedReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setPrivilegedReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepository) StreamIn(handle string, path string, stream io.Reader) (bool, error) {
	fake.streamInMutex.Lock()
	ret, specificReturn := fake.streamInReturnsOnCall[len(fake.streamInArgsForCall)]
	fake.streamInArgsForCall = append(fake.streamInArgsForCall, struct {
		handle string
		path   string
		stream io.Reader
	}{handle, path, stream})
	fake.recordInvocation("StreamIn", []interface{}{handle, path, stream})
	fake.streamInMutex.Unlock()
	if fake.StreamInStub != nil {
		return fake.StreamInStub(handle, path, stream)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.streamInReturns.result1, fake.streamInReturns.result2
}

func (fake *FakeRepository) StreamInCallCount() int {
	fake.streamInMutex.RLock()
	defer fake.streamInMutex.RUnlock()
	return len(fake.streamInArgsForCall)
}

func (fake *FakeRepository) StreamInArgsForCall(i int) (string, string, io.Reader) {
	fake.streamInMutex.RLock()
	defer fake.streamInMutex.RUnlock()
	return fake.streamInArgsForCall[i].handle, fake.streamInArgsForCall[i].path, fake.streamInArgsForCall[i].stream
}

func (fake *FakeRepository) StreamInReturns(result1 bool, result2 error) {
	fake.StreamInStub = nil
	fake.streamInReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) StreamInReturnsOnCall(i int, result1 bool, result2 error) {
	fake.StreamInStub = nil
	if fake.streamInReturnsOnCall == nil {
		fake.streamInReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.streamInReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) StreamOut(handle string, path string, dest io.Writer) error {
	fake.streamOutMutex.Lock()
	ret, specificReturn := fake.streamOutReturnsOnCall[len(fake.streamOutArgsForCall)]
	fake.streamOutArgsForCall = append(fake.streamOutArgsForCall, struct {
		handle string
		path   string
		dest   io.Writer
	}{handle, path, dest})
	fake.recordInvocation("StreamOut", []interface{}{handle, path, dest})
	fake.streamOutMutex.Unlock()
	if fake.StreamOutStub != nil {
		return fake.StreamOutStub(handle, path, dest)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.streamOutReturns.result1
}

func (fake *FakeRepository) StreamOutCallCount() int {
	fake.streamOutMutex.RLock()
	defer fake.streamOutMutex.RUnlock()
	return len(fake.streamOutArgsForCall)
}

func (fake *FakeRepository) StreamOutArgsForCall(i int) (string, string, io.Writer) {
	fake.streamOutMutex.RLock()
	defer fake.streamOutMutex.RUnlock()
	return fake.streamOutArgsForCall[i].handle, fake.streamOutArgsForCall[i].path, fake.streamOutArgsForCall[i].dest
}

func (fake *FakeRepository) StreamOutReturns(result1 error) {
	fake.StreamOutStub = nil
	fake.streamOutReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepository) StreamOutReturnsOnCall(i int, result1 error) {
	fake.StreamOutStub = nil
	if fake.streamOutReturnsOnCall == nil {
		fake.streamOutReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.streamOutReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepository) VolumeParent(handle string) (volume.Volume, bool, error) {
	fake.volumeParentMutex.Lock()
	ret, specificReturn := fake.volumeParentReturnsOnCall[len(fake.volumeParentArgsForCall)]
	fake.volumeParentArgsForCall = append(fake.volumeParentArgsForCall, struct {
		handle string
	}{handle})
	fake.recordInvocation("VolumeParent", []interface{}{handle})
	fake.volumeParentMutex.Unlock()
	if fake.VolumeParentStub != nil {
		return fake.VolumeParentStub(handle)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.volumeParentReturns.result1, fake.volumeParentReturns.result2, fake.volumeParentReturns.result3
}

func (fake *FakeRepository) VolumeParentCallCount() int {
	fake.volumeParentMutex.RLock()
	defer fake.volumeParentMutex.RUnlock()
	return len(fake.volumeParentArgsForCall)
}

func (fake *FakeRepository) VolumeParentArgsForCall(i int) string {
	fake.volumeParentMutex.RLock()
	defer fake.volumeParentMutex.RUnlock()
	return fake.volumeParentArgsForCall[i].handle
}

func (fake *FakeRepository) VolumeParentReturns(result1 volume.Volume, result2 bool, result3 error) {
	fake.VolumeParentStub = nil
	fake.volumeParentReturns = struct {
		result1 volume.Volume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeRepository) VolumeParentReturnsOnCall(i int, result1 volume.Volume, result2 bool, result3 error) {
	fake.VolumeParentStub = nil
	if fake.volumeParentReturnsOnCall == nil {
		fake.volumeParentReturnsOnCall = make(map[int]struct {
			result1 volume.Volume
			result2 bool
			result3 error
		})
	}
	fake.volumeParentReturnsOnCall[i] = struct {
		result1 volume.Volume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.listVolumesMutex.RLock()
	defer fake.listVolumesMutex.RUnlock()
	fake.getVolumeMutex.RLock()
	defer fake.getVolumeMutex.RUnlock()
	fake.createVolumeMutex.RLock()
	defer fake.createVolumeMutex.RUnlock()
	fake.destroyVolumeMutex.RLock()
	defer fake.destroyVolumeMutex.RUnlock()
	fake.destroyVolumeAndDescendantsMutex.RLock()
	defer fake.destroyVolumeAndDescendantsMutex.RUnlock()
	fake.setPropertyMutex.RLock()
	defer fake.setPropertyMutex.RUnlock()
	fake.setTTLMutex.RLock()
	defer fake.setTTLMutex.RUnlock()
	fake.setPrivilegedMutex.RLock()
	defer fake.setPrivilegedMutex.RUnlock()
	fake.streamInMutex.RLock()
	defer fake.streamInMutex.RUnlock()
	fake.streamOutMutex.RLock()
	defer fake.streamOutMutex.RUnlock()
	fake.volumeParentMutex.RLock()
	defer fake.volumeParentMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRepository) recordInvocation(key string, args []interface{}) {
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

var _ volume.Repository = new(FakeRepository)
