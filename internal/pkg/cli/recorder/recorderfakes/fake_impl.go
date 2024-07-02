//go:build linux && !no_bpf
// +build linux,!no_bpf

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by counterfeiter. DO NOT EDIT.
package recorderfakes

import (
	"context"
	"io"
	"os"
	"sync"

	"github.com/aquasecurity/libbpfgo"
	seccompa "github.com/containers/common/pkg/seccomp"
	seccomp "github.com/seccomp/libseccomp-golang"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/cli-runtime/pkg/printers"
	"sigs.k8s.io/security-profiles-operator/internal/pkg/cli/command"
	"sigs.k8s.io/security-profiles-operator/internal/pkg/daemon/bpfrecorder"
)

type FakeImpl struct {
	CommandRunStub        func(*command.Command) (uint32, error)
	commandRunMutex       sync.RWMutex
	commandRunArgsForCall []struct {
		arg1 *command.Command
	}
	commandRunReturns struct {
		result1 uint32
		result2 error
	}
	commandRunReturnsOnCall map[int]struct {
		result1 uint32
		result2 error
	}
	CommandWaitStub        func(*command.Command) error
	commandWaitMutex       sync.RWMutex
	commandWaitArgsForCall []struct {
		arg1 *command.Command
	}
	commandWaitReturns struct {
		result1 error
	}
	commandWaitReturnsOnCall map[int]struct {
		result1 error
	}
	CreateStub        func(string) (io.WriteCloser, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 string
	}
	createReturns struct {
		result1 io.WriteCloser
		result2 error
	}
	createReturnsOnCall map[int]struct {
		result1 io.WriteCloser
		result2 error
	}
	FindProcMountNamespaceStub        func(*bpfrecorder.BpfRecorder, uint32) (uint32, error)
	findProcMountNamespaceMutex       sync.RWMutex
	findProcMountNamespaceArgsForCall []struct {
		arg1 *bpfrecorder.BpfRecorder
		arg2 uint32
	}
	findProcMountNamespaceReturns struct {
		result1 uint32
		result2 error
	}
	findProcMountNamespaceReturnsOnCall map[int]struct {
		result1 uint32
		result2 error
	}
	GetNameStub        func(seccomp.ScmpSyscall) (string, error)
	getNameMutex       sync.RWMutex
	getNameArgsForCall []struct {
		arg1 seccomp.ScmpSyscall
	}
	getNameReturns struct {
		result1 string
		result2 error
	}
	getNameReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	GoArchToSeccompArchStub        func(string) (seccompa.Arch, error)
	goArchToSeccompArchMutex       sync.RWMutex
	goArchToSeccompArchArgsForCall []struct {
		arg1 string
	}
	goArchToSeccompArchReturns struct {
		result1 seccompa.Arch
		result2 error
	}
	goArchToSeccompArchReturnsOnCall map[int]struct {
		result1 seccompa.Arch
		result2 error
	}
	IteratorKeyStub        func(*libbpfgo.BPFMapIterator) []byte
	iteratorKeyMutex       sync.RWMutex
	iteratorKeyArgsForCall []struct {
		arg1 *libbpfgo.BPFMapIterator
	}
	iteratorKeyReturns struct {
		result1 []byte
	}
	iteratorKeyReturnsOnCall map[int]struct {
		result1 []byte
	}
	IteratorNextStub        func(*libbpfgo.BPFMapIterator) bool
	iteratorNextMutex       sync.RWMutex
	iteratorNextArgsForCall []struct {
		arg1 *libbpfgo.BPFMapIterator
	}
	iteratorNextReturns struct {
		result1 bool
	}
	iteratorNextReturnsOnCall map[int]struct {
		result1 bool
	}
	LoadBpfRecorderStub        func(*bpfrecorder.BpfRecorder) error
	loadBpfRecorderMutex       sync.RWMutex
	loadBpfRecorderArgsForCall []struct {
		arg1 *bpfrecorder.BpfRecorder
	}
	loadBpfRecorderReturns struct {
		result1 error
	}
	loadBpfRecorderReturnsOnCall map[int]struct {
		result1 error
	}
	MarshalIndentStub        func(any, string, string) ([]byte, error)
	marshalIndentMutex       sync.RWMutex
	marshalIndentArgsForCall []struct {
		arg1 any
		arg2 string
		arg3 string
	}
	marshalIndentReturns struct {
		result1 []byte
		result2 error
	}
	marshalIndentReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	NotifyStub        func(chan<- os.Signal, ...os.Signal)
	notifyMutex       sync.RWMutex
	notifyArgsForCall []struct {
		arg1 chan<- os.Signal
		arg2 []os.Signal
	}
	PrintObjStub        func(printers.YAMLPrinter, runtime.Object, io.Writer) error
	printObjMutex       sync.RWMutex
	printObjArgsForCall []struct {
		arg1 printers.YAMLPrinter
		arg2 runtime.Object
		arg3 io.Writer
	}
	printObjReturns struct {
		result1 error
	}
	printObjReturnsOnCall map[int]struct {
		result1 error
	}
	SyscallsGetValueStub        func(*bpfrecorder.BpfRecorder, uint32) ([]byte, error)
	syscallsGetValueMutex       sync.RWMutex
	syscallsGetValueArgsForCall []struct {
		arg1 *bpfrecorder.BpfRecorder
		arg2 uint32
	}
	syscallsGetValueReturns struct {
		result1 []byte
		result2 error
	}
	syscallsGetValueReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	SyscallsIteratorStub        func(*bpfrecorder.BpfRecorder) *libbpfgo.BPFMapIterator
	syscallsIteratorMutex       sync.RWMutex
	syscallsIteratorArgsForCall []struct {
		arg1 *bpfrecorder.BpfRecorder
	}
	syscallsIteratorReturns struct {
		result1 *libbpfgo.BPFMapIterator
	}
	syscallsIteratorReturnsOnCall map[int]struct {
		result1 *libbpfgo.BPFMapIterator
	}
	UnloadBpfRecorderStub        func(*bpfrecorder.BpfRecorder)
	unloadBpfRecorderMutex       sync.RWMutex
	unloadBpfRecorderArgsForCall []struct {
		arg1 *bpfrecorder.BpfRecorder
	}
	WaitForPidExitStub        func(*bpfrecorder.BpfRecorder, context.Context, uint32) error
	waitForPidExitMutex       sync.RWMutex
	waitForPidExitArgsForCall []struct {
		arg1 *bpfrecorder.BpfRecorder
		arg2 context.Context
		arg3 uint32
	}
	waitForPidExitReturns struct {
		result1 error
	}
	waitForPidExitReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeImpl) CommandRun(arg1 *command.Command) (uint32, error) {
	fake.commandRunMutex.Lock()
	ret, specificReturn := fake.commandRunReturnsOnCall[len(fake.commandRunArgsForCall)]
	fake.commandRunArgsForCall = append(fake.commandRunArgsForCall, struct {
		arg1 *command.Command
	}{arg1})
	stub := fake.CommandRunStub
	fakeReturns := fake.commandRunReturns
	fake.recordInvocation("CommandRun", []interface{}{arg1})
	fake.commandRunMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) CommandRunCallCount() int {
	fake.commandRunMutex.RLock()
	defer fake.commandRunMutex.RUnlock()
	return len(fake.commandRunArgsForCall)
}

func (fake *FakeImpl) CommandRunCalls(stub func(*command.Command) (uint32, error)) {
	fake.commandRunMutex.Lock()
	defer fake.commandRunMutex.Unlock()
	fake.CommandRunStub = stub
}

func (fake *FakeImpl) CommandRunArgsForCall(i int) *command.Command {
	fake.commandRunMutex.RLock()
	defer fake.commandRunMutex.RUnlock()
	argsForCall := fake.commandRunArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) CommandRunReturns(result1 uint32, result2 error) {
	fake.commandRunMutex.Lock()
	defer fake.commandRunMutex.Unlock()
	fake.CommandRunStub = nil
	fake.commandRunReturns = struct {
		result1 uint32
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) CommandRunReturnsOnCall(i int, result1 uint32, result2 error) {
	fake.commandRunMutex.Lock()
	defer fake.commandRunMutex.Unlock()
	fake.CommandRunStub = nil
	if fake.commandRunReturnsOnCall == nil {
		fake.commandRunReturnsOnCall = make(map[int]struct {
			result1 uint32
			result2 error
		})
	}
	fake.commandRunReturnsOnCall[i] = struct {
		result1 uint32
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) CommandWait(arg1 *command.Command) error {
	fake.commandWaitMutex.Lock()
	ret, specificReturn := fake.commandWaitReturnsOnCall[len(fake.commandWaitArgsForCall)]
	fake.commandWaitArgsForCall = append(fake.commandWaitArgsForCall, struct {
		arg1 *command.Command
	}{arg1})
	stub := fake.CommandWaitStub
	fakeReturns := fake.commandWaitReturns
	fake.recordInvocation("CommandWait", []interface{}{arg1})
	fake.commandWaitMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) CommandWaitCallCount() int {
	fake.commandWaitMutex.RLock()
	defer fake.commandWaitMutex.RUnlock()
	return len(fake.commandWaitArgsForCall)
}

func (fake *FakeImpl) CommandWaitCalls(stub func(*command.Command) error) {
	fake.commandWaitMutex.Lock()
	defer fake.commandWaitMutex.Unlock()
	fake.CommandWaitStub = stub
}

func (fake *FakeImpl) CommandWaitArgsForCall(i int) *command.Command {
	fake.commandWaitMutex.RLock()
	defer fake.commandWaitMutex.RUnlock()
	argsForCall := fake.commandWaitArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) CommandWaitReturns(result1 error) {
	fake.commandWaitMutex.Lock()
	defer fake.commandWaitMutex.Unlock()
	fake.CommandWaitStub = nil
	fake.commandWaitReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) CommandWaitReturnsOnCall(i int, result1 error) {
	fake.commandWaitMutex.Lock()
	defer fake.commandWaitMutex.Unlock()
	fake.CommandWaitStub = nil
	if fake.commandWaitReturnsOnCall == nil {
		fake.commandWaitReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.commandWaitReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) Create(arg1 string) (io.WriteCloser, error) {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.CreateStub
	fakeReturns := fake.createReturns
	fake.recordInvocation("Create", []interface{}{arg1})
	fake.createMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeImpl) CreateCalls(stub func(string) (io.WriteCloser, error)) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *FakeImpl) CreateArgsForCall(i int) string {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) CreateReturns(result1 io.WriteCloser, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 io.WriteCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) CreateReturnsOnCall(i int, result1 io.WriteCloser, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 io.WriteCloser
			result2 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 io.WriteCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) FindProcMountNamespace(arg1 *bpfrecorder.BpfRecorder, arg2 uint32) (uint32, error) {
	fake.findProcMountNamespaceMutex.Lock()
	ret, specificReturn := fake.findProcMountNamespaceReturnsOnCall[len(fake.findProcMountNamespaceArgsForCall)]
	fake.findProcMountNamespaceArgsForCall = append(fake.findProcMountNamespaceArgsForCall, struct {
		arg1 *bpfrecorder.BpfRecorder
		arg2 uint32
	}{arg1, arg2})
	stub := fake.FindProcMountNamespaceStub
	fakeReturns := fake.findProcMountNamespaceReturns
	fake.recordInvocation("FindProcMountNamespace", []interface{}{arg1, arg2})
	fake.findProcMountNamespaceMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) FindProcMountNamespaceCallCount() int {
	fake.findProcMountNamespaceMutex.RLock()
	defer fake.findProcMountNamespaceMutex.RUnlock()
	return len(fake.findProcMountNamespaceArgsForCall)
}

func (fake *FakeImpl) FindProcMountNamespaceCalls(stub func(*bpfrecorder.BpfRecorder, uint32) (uint32, error)) {
	fake.findProcMountNamespaceMutex.Lock()
	defer fake.findProcMountNamespaceMutex.Unlock()
	fake.FindProcMountNamespaceStub = stub
}

func (fake *FakeImpl) FindProcMountNamespaceArgsForCall(i int) (*bpfrecorder.BpfRecorder, uint32) {
	fake.findProcMountNamespaceMutex.RLock()
	defer fake.findProcMountNamespaceMutex.RUnlock()
	argsForCall := fake.findProcMountNamespaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImpl) FindProcMountNamespaceReturns(result1 uint32, result2 error) {
	fake.findProcMountNamespaceMutex.Lock()
	defer fake.findProcMountNamespaceMutex.Unlock()
	fake.FindProcMountNamespaceStub = nil
	fake.findProcMountNamespaceReturns = struct {
		result1 uint32
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) FindProcMountNamespaceReturnsOnCall(i int, result1 uint32, result2 error) {
	fake.findProcMountNamespaceMutex.Lock()
	defer fake.findProcMountNamespaceMutex.Unlock()
	fake.FindProcMountNamespaceStub = nil
	if fake.findProcMountNamespaceReturnsOnCall == nil {
		fake.findProcMountNamespaceReturnsOnCall = make(map[int]struct {
			result1 uint32
			result2 error
		})
	}
	fake.findProcMountNamespaceReturnsOnCall[i] = struct {
		result1 uint32
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) GetName(arg1 seccomp.ScmpSyscall) (string, error) {
	fake.getNameMutex.Lock()
	ret, specificReturn := fake.getNameReturnsOnCall[len(fake.getNameArgsForCall)]
	fake.getNameArgsForCall = append(fake.getNameArgsForCall, struct {
		arg1 seccomp.ScmpSyscall
	}{arg1})
	stub := fake.GetNameStub
	fakeReturns := fake.getNameReturns
	fake.recordInvocation("GetName", []interface{}{arg1})
	fake.getNameMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) GetNameCallCount() int {
	fake.getNameMutex.RLock()
	defer fake.getNameMutex.RUnlock()
	return len(fake.getNameArgsForCall)
}

func (fake *FakeImpl) GetNameCalls(stub func(seccomp.ScmpSyscall) (string, error)) {
	fake.getNameMutex.Lock()
	defer fake.getNameMutex.Unlock()
	fake.GetNameStub = stub
}

func (fake *FakeImpl) GetNameArgsForCall(i int) seccomp.ScmpSyscall {
	fake.getNameMutex.RLock()
	defer fake.getNameMutex.RUnlock()
	argsForCall := fake.getNameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) GetNameReturns(result1 string, result2 error) {
	fake.getNameMutex.Lock()
	defer fake.getNameMutex.Unlock()
	fake.GetNameStub = nil
	fake.getNameReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) GetNameReturnsOnCall(i int, result1 string, result2 error) {
	fake.getNameMutex.Lock()
	defer fake.getNameMutex.Unlock()
	fake.GetNameStub = nil
	if fake.getNameReturnsOnCall == nil {
		fake.getNameReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getNameReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) GoArchToSeccompArch(arg1 string) (seccompa.Arch, error) {
	fake.goArchToSeccompArchMutex.Lock()
	ret, specificReturn := fake.goArchToSeccompArchReturnsOnCall[len(fake.goArchToSeccompArchArgsForCall)]
	fake.goArchToSeccompArchArgsForCall = append(fake.goArchToSeccompArchArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GoArchToSeccompArchStub
	fakeReturns := fake.goArchToSeccompArchReturns
	fake.recordInvocation("GoArchToSeccompArch", []interface{}{arg1})
	fake.goArchToSeccompArchMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) GoArchToSeccompArchCallCount() int {
	fake.goArchToSeccompArchMutex.RLock()
	defer fake.goArchToSeccompArchMutex.RUnlock()
	return len(fake.goArchToSeccompArchArgsForCall)
}

func (fake *FakeImpl) GoArchToSeccompArchCalls(stub func(string) (seccompa.Arch, error)) {
	fake.goArchToSeccompArchMutex.Lock()
	defer fake.goArchToSeccompArchMutex.Unlock()
	fake.GoArchToSeccompArchStub = stub
}

func (fake *FakeImpl) GoArchToSeccompArchArgsForCall(i int) string {
	fake.goArchToSeccompArchMutex.RLock()
	defer fake.goArchToSeccompArchMutex.RUnlock()
	argsForCall := fake.goArchToSeccompArchArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) GoArchToSeccompArchReturns(result1 seccompa.Arch, result2 error) {
	fake.goArchToSeccompArchMutex.Lock()
	defer fake.goArchToSeccompArchMutex.Unlock()
	fake.GoArchToSeccompArchStub = nil
	fake.goArchToSeccompArchReturns = struct {
		result1 seccompa.Arch
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) GoArchToSeccompArchReturnsOnCall(i int, result1 seccompa.Arch, result2 error) {
	fake.goArchToSeccompArchMutex.Lock()
	defer fake.goArchToSeccompArchMutex.Unlock()
	fake.GoArchToSeccompArchStub = nil
	if fake.goArchToSeccompArchReturnsOnCall == nil {
		fake.goArchToSeccompArchReturnsOnCall = make(map[int]struct {
			result1 seccompa.Arch
			result2 error
		})
	}
	fake.goArchToSeccompArchReturnsOnCall[i] = struct {
		result1 seccompa.Arch
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) IteratorKey(arg1 *libbpfgo.BPFMapIterator) []byte {
	fake.iteratorKeyMutex.Lock()
	ret, specificReturn := fake.iteratorKeyReturnsOnCall[len(fake.iteratorKeyArgsForCall)]
	fake.iteratorKeyArgsForCall = append(fake.iteratorKeyArgsForCall, struct {
		arg1 *libbpfgo.BPFMapIterator
	}{arg1})
	stub := fake.IteratorKeyStub
	fakeReturns := fake.iteratorKeyReturns
	fake.recordInvocation("IteratorKey", []interface{}{arg1})
	fake.iteratorKeyMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) IteratorKeyCallCount() int {
	fake.iteratorKeyMutex.RLock()
	defer fake.iteratorKeyMutex.RUnlock()
	return len(fake.iteratorKeyArgsForCall)
}

func (fake *FakeImpl) IteratorKeyCalls(stub func(*libbpfgo.BPFMapIterator) []byte) {
	fake.iteratorKeyMutex.Lock()
	defer fake.iteratorKeyMutex.Unlock()
	fake.IteratorKeyStub = stub
}

func (fake *FakeImpl) IteratorKeyArgsForCall(i int) *libbpfgo.BPFMapIterator {
	fake.iteratorKeyMutex.RLock()
	defer fake.iteratorKeyMutex.RUnlock()
	argsForCall := fake.iteratorKeyArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) IteratorKeyReturns(result1 []byte) {
	fake.iteratorKeyMutex.Lock()
	defer fake.iteratorKeyMutex.Unlock()
	fake.IteratorKeyStub = nil
	fake.iteratorKeyReturns = struct {
		result1 []byte
	}{result1}
}

func (fake *FakeImpl) IteratorKeyReturnsOnCall(i int, result1 []byte) {
	fake.iteratorKeyMutex.Lock()
	defer fake.iteratorKeyMutex.Unlock()
	fake.IteratorKeyStub = nil
	if fake.iteratorKeyReturnsOnCall == nil {
		fake.iteratorKeyReturnsOnCall = make(map[int]struct {
			result1 []byte
		})
	}
	fake.iteratorKeyReturnsOnCall[i] = struct {
		result1 []byte
	}{result1}
}

func (fake *FakeImpl) IteratorNext(arg1 *libbpfgo.BPFMapIterator) bool {
	fake.iteratorNextMutex.Lock()
	ret, specificReturn := fake.iteratorNextReturnsOnCall[len(fake.iteratorNextArgsForCall)]
	fake.iteratorNextArgsForCall = append(fake.iteratorNextArgsForCall, struct {
		arg1 *libbpfgo.BPFMapIterator
	}{arg1})
	stub := fake.IteratorNextStub
	fakeReturns := fake.iteratorNextReturns
	fake.recordInvocation("IteratorNext", []interface{}{arg1})
	fake.iteratorNextMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) IteratorNextCallCount() int {
	fake.iteratorNextMutex.RLock()
	defer fake.iteratorNextMutex.RUnlock()
	return len(fake.iteratorNextArgsForCall)
}

func (fake *FakeImpl) IteratorNextCalls(stub func(*libbpfgo.BPFMapIterator) bool) {
	fake.iteratorNextMutex.Lock()
	defer fake.iteratorNextMutex.Unlock()
	fake.IteratorNextStub = stub
}

func (fake *FakeImpl) IteratorNextArgsForCall(i int) *libbpfgo.BPFMapIterator {
	fake.iteratorNextMutex.RLock()
	defer fake.iteratorNextMutex.RUnlock()
	argsForCall := fake.iteratorNextArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) IteratorNextReturns(result1 bool) {
	fake.iteratorNextMutex.Lock()
	defer fake.iteratorNextMutex.Unlock()
	fake.IteratorNextStub = nil
	fake.iteratorNextReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeImpl) IteratorNextReturnsOnCall(i int, result1 bool) {
	fake.iteratorNextMutex.Lock()
	defer fake.iteratorNextMutex.Unlock()
	fake.IteratorNextStub = nil
	if fake.iteratorNextReturnsOnCall == nil {
		fake.iteratorNextReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.iteratorNextReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeImpl) LoadBpfRecorder(arg1 *bpfrecorder.BpfRecorder) error {
	fake.loadBpfRecorderMutex.Lock()
	ret, specificReturn := fake.loadBpfRecorderReturnsOnCall[len(fake.loadBpfRecorderArgsForCall)]
	fake.loadBpfRecorderArgsForCall = append(fake.loadBpfRecorderArgsForCall, struct {
		arg1 *bpfrecorder.BpfRecorder
	}{arg1})
	stub := fake.LoadBpfRecorderStub
	fakeReturns := fake.loadBpfRecorderReturns
	fake.recordInvocation("LoadBpfRecorder", []interface{}{arg1})
	fake.loadBpfRecorderMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) LoadBpfRecorderCallCount() int {
	fake.loadBpfRecorderMutex.RLock()
	defer fake.loadBpfRecorderMutex.RUnlock()
	return len(fake.loadBpfRecorderArgsForCall)
}

func (fake *FakeImpl) LoadBpfRecorderCalls(stub func(*bpfrecorder.BpfRecorder) error) {
	fake.loadBpfRecorderMutex.Lock()
	defer fake.loadBpfRecorderMutex.Unlock()
	fake.LoadBpfRecorderStub = stub
}

func (fake *FakeImpl) LoadBpfRecorderArgsForCall(i int) *bpfrecorder.BpfRecorder {
	fake.loadBpfRecorderMutex.RLock()
	defer fake.loadBpfRecorderMutex.RUnlock()
	argsForCall := fake.loadBpfRecorderArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) LoadBpfRecorderReturns(result1 error) {
	fake.loadBpfRecorderMutex.Lock()
	defer fake.loadBpfRecorderMutex.Unlock()
	fake.LoadBpfRecorderStub = nil
	fake.loadBpfRecorderReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) LoadBpfRecorderReturnsOnCall(i int, result1 error) {
	fake.loadBpfRecorderMutex.Lock()
	defer fake.loadBpfRecorderMutex.Unlock()
	fake.LoadBpfRecorderStub = nil
	if fake.loadBpfRecorderReturnsOnCall == nil {
		fake.loadBpfRecorderReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.loadBpfRecorderReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) MarshalIndent(arg1 any, arg2 string, arg3 string) ([]byte, error) {
	fake.marshalIndentMutex.Lock()
	ret, specificReturn := fake.marshalIndentReturnsOnCall[len(fake.marshalIndentArgsForCall)]
	fake.marshalIndentArgsForCall = append(fake.marshalIndentArgsForCall, struct {
		arg1 any
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.MarshalIndentStub
	fakeReturns := fake.marshalIndentReturns
	fake.recordInvocation("MarshalIndent", []interface{}{arg1, arg2, arg3})
	fake.marshalIndentMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) MarshalIndentCallCount() int {
	fake.marshalIndentMutex.RLock()
	defer fake.marshalIndentMutex.RUnlock()
	return len(fake.marshalIndentArgsForCall)
}

func (fake *FakeImpl) MarshalIndentCalls(stub func(any, string, string) ([]byte, error)) {
	fake.marshalIndentMutex.Lock()
	defer fake.marshalIndentMutex.Unlock()
	fake.MarshalIndentStub = stub
}

func (fake *FakeImpl) MarshalIndentArgsForCall(i int) (any, string, string) {
	fake.marshalIndentMutex.RLock()
	defer fake.marshalIndentMutex.RUnlock()
	argsForCall := fake.marshalIndentArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeImpl) MarshalIndentReturns(result1 []byte, result2 error) {
	fake.marshalIndentMutex.Lock()
	defer fake.marshalIndentMutex.Unlock()
	fake.MarshalIndentStub = nil
	fake.marshalIndentReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) MarshalIndentReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.marshalIndentMutex.Lock()
	defer fake.marshalIndentMutex.Unlock()
	fake.MarshalIndentStub = nil
	if fake.marshalIndentReturnsOnCall == nil {
		fake.marshalIndentReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.marshalIndentReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) Notify(arg1 chan<- os.Signal, arg2 ...os.Signal) {
	fake.notifyMutex.Lock()
	fake.notifyArgsForCall = append(fake.notifyArgsForCall, struct {
		arg1 chan<- os.Signal
		arg2 []os.Signal
	}{arg1, arg2})
	stub := fake.NotifyStub
	fake.recordInvocation("Notify", []interface{}{arg1, arg2})
	fake.notifyMutex.Unlock()
	if stub != nil {
		fake.NotifyStub(arg1, arg2...)
	}
}

func (fake *FakeImpl) NotifyCallCount() int {
	fake.notifyMutex.RLock()
	defer fake.notifyMutex.RUnlock()
	return len(fake.notifyArgsForCall)
}

func (fake *FakeImpl) NotifyCalls(stub func(chan<- os.Signal, ...os.Signal)) {
	fake.notifyMutex.Lock()
	defer fake.notifyMutex.Unlock()
	fake.NotifyStub = stub
}

func (fake *FakeImpl) NotifyArgsForCall(i int) (chan<- os.Signal, []os.Signal) {
	fake.notifyMutex.RLock()
	defer fake.notifyMutex.RUnlock()
	argsForCall := fake.notifyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImpl) PrintObj(arg1 printers.YAMLPrinter, arg2 runtime.Object, arg3 io.Writer) error {
	fake.printObjMutex.Lock()
	ret, specificReturn := fake.printObjReturnsOnCall[len(fake.printObjArgsForCall)]
	fake.printObjArgsForCall = append(fake.printObjArgsForCall, struct {
		arg1 printers.YAMLPrinter
		arg2 runtime.Object
		arg3 io.Writer
	}{arg1, arg2, arg3})
	stub := fake.PrintObjStub
	fakeReturns := fake.printObjReturns
	fake.recordInvocation("PrintObj", []interface{}{arg1, arg2, arg3})
	fake.printObjMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) PrintObjCallCount() int {
	fake.printObjMutex.RLock()
	defer fake.printObjMutex.RUnlock()
	return len(fake.printObjArgsForCall)
}

func (fake *FakeImpl) PrintObjCalls(stub func(printers.YAMLPrinter, runtime.Object, io.Writer) error) {
	fake.printObjMutex.Lock()
	defer fake.printObjMutex.Unlock()
	fake.PrintObjStub = stub
}

func (fake *FakeImpl) PrintObjArgsForCall(i int) (printers.YAMLPrinter, runtime.Object, io.Writer) {
	fake.printObjMutex.RLock()
	defer fake.printObjMutex.RUnlock()
	argsForCall := fake.printObjArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeImpl) PrintObjReturns(result1 error) {
	fake.printObjMutex.Lock()
	defer fake.printObjMutex.Unlock()
	fake.PrintObjStub = nil
	fake.printObjReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) PrintObjReturnsOnCall(i int, result1 error) {
	fake.printObjMutex.Lock()
	defer fake.printObjMutex.Unlock()
	fake.PrintObjStub = nil
	if fake.printObjReturnsOnCall == nil {
		fake.printObjReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.printObjReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) SyscallsGetValue(arg1 *bpfrecorder.BpfRecorder, arg2 uint32) ([]byte, error) {
	fake.syscallsGetValueMutex.Lock()
	ret, specificReturn := fake.syscallsGetValueReturnsOnCall[len(fake.syscallsGetValueArgsForCall)]
	fake.syscallsGetValueArgsForCall = append(fake.syscallsGetValueArgsForCall, struct {
		arg1 *bpfrecorder.BpfRecorder
		arg2 uint32
	}{arg1, arg2})
	stub := fake.SyscallsGetValueStub
	fakeReturns := fake.syscallsGetValueReturns
	fake.recordInvocation("SyscallsGetValue", []interface{}{arg1, arg2})
	fake.syscallsGetValueMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) SyscallsGetValueCallCount() int {
	fake.syscallsGetValueMutex.RLock()
	defer fake.syscallsGetValueMutex.RUnlock()
	return len(fake.syscallsGetValueArgsForCall)
}

func (fake *FakeImpl) SyscallsGetValueCalls(stub func(*bpfrecorder.BpfRecorder, uint32) ([]byte, error)) {
	fake.syscallsGetValueMutex.Lock()
	defer fake.syscallsGetValueMutex.Unlock()
	fake.SyscallsGetValueStub = stub
}

func (fake *FakeImpl) SyscallsGetValueArgsForCall(i int) (*bpfrecorder.BpfRecorder, uint32) {
	fake.syscallsGetValueMutex.RLock()
	defer fake.syscallsGetValueMutex.RUnlock()
	argsForCall := fake.syscallsGetValueArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImpl) SyscallsGetValueReturns(result1 []byte, result2 error) {
	fake.syscallsGetValueMutex.Lock()
	defer fake.syscallsGetValueMutex.Unlock()
	fake.SyscallsGetValueStub = nil
	fake.syscallsGetValueReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) SyscallsGetValueReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.syscallsGetValueMutex.Lock()
	defer fake.syscallsGetValueMutex.Unlock()
	fake.SyscallsGetValueStub = nil
	if fake.syscallsGetValueReturnsOnCall == nil {
		fake.syscallsGetValueReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.syscallsGetValueReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) SyscallsIterator(arg1 *bpfrecorder.BpfRecorder) *libbpfgo.BPFMapIterator {
	fake.syscallsIteratorMutex.Lock()
	ret, specificReturn := fake.syscallsIteratorReturnsOnCall[len(fake.syscallsIteratorArgsForCall)]
	fake.syscallsIteratorArgsForCall = append(fake.syscallsIteratorArgsForCall, struct {
		arg1 *bpfrecorder.BpfRecorder
	}{arg1})
	stub := fake.SyscallsIteratorStub
	fakeReturns := fake.syscallsIteratorReturns
	fake.recordInvocation("SyscallsIterator", []interface{}{arg1})
	fake.syscallsIteratorMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) SyscallsIteratorCallCount() int {
	fake.syscallsIteratorMutex.RLock()
	defer fake.syscallsIteratorMutex.RUnlock()
	return len(fake.syscallsIteratorArgsForCall)
}

func (fake *FakeImpl) SyscallsIteratorCalls(stub func(*bpfrecorder.BpfRecorder) *libbpfgo.BPFMapIterator) {
	fake.syscallsIteratorMutex.Lock()
	defer fake.syscallsIteratorMutex.Unlock()
	fake.SyscallsIteratorStub = stub
}

func (fake *FakeImpl) SyscallsIteratorArgsForCall(i int) *bpfrecorder.BpfRecorder {
	fake.syscallsIteratorMutex.RLock()
	defer fake.syscallsIteratorMutex.RUnlock()
	argsForCall := fake.syscallsIteratorArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) SyscallsIteratorReturns(result1 *libbpfgo.BPFMapIterator) {
	fake.syscallsIteratorMutex.Lock()
	defer fake.syscallsIteratorMutex.Unlock()
	fake.SyscallsIteratorStub = nil
	fake.syscallsIteratorReturns = struct {
		result1 *libbpfgo.BPFMapIterator
	}{result1}
}

func (fake *FakeImpl) SyscallsIteratorReturnsOnCall(i int, result1 *libbpfgo.BPFMapIterator) {
	fake.syscallsIteratorMutex.Lock()
	defer fake.syscallsIteratorMutex.Unlock()
	fake.SyscallsIteratorStub = nil
	if fake.syscallsIteratorReturnsOnCall == nil {
		fake.syscallsIteratorReturnsOnCall = make(map[int]struct {
			result1 *libbpfgo.BPFMapIterator
		})
	}
	fake.syscallsIteratorReturnsOnCall[i] = struct {
		result1 *libbpfgo.BPFMapIterator
	}{result1}
}

func (fake *FakeImpl) UnloadBpfRecorder(arg1 *bpfrecorder.BpfRecorder) {
	fake.unloadBpfRecorderMutex.Lock()
	fake.unloadBpfRecorderArgsForCall = append(fake.unloadBpfRecorderArgsForCall, struct {
		arg1 *bpfrecorder.BpfRecorder
	}{arg1})
	stub := fake.UnloadBpfRecorderStub
	fake.recordInvocation("UnloadBpfRecorder", []interface{}{arg1})
	fake.unloadBpfRecorderMutex.Unlock()
	if stub != nil {
		fake.UnloadBpfRecorderStub(arg1)
	}
}

func (fake *FakeImpl) UnloadBpfRecorderCallCount() int {
	fake.unloadBpfRecorderMutex.RLock()
	defer fake.unloadBpfRecorderMutex.RUnlock()
	return len(fake.unloadBpfRecorderArgsForCall)
}

func (fake *FakeImpl) UnloadBpfRecorderCalls(stub func(*bpfrecorder.BpfRecorder)) {
	fake.unloadBpfRecorderMutex.Lock()
	defer fake.unloadBpfRecorderMutex.Unlock()
	fake.UnloadBpfRecorderStub = stub
}

func (fake *FakeImpl) UnloadBpfRecorderArgsForCall(i int) *bpfrecorder.BpfRecorder {
	fake.unloadBpfRecorderMutex.RLock()
	defer fake.unloadBpfRecorderMutex.RUnlock()
	argsForCall := fake.unloadBpfRecorderArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) WaitForPidExit(arg1 *bpfrecorder.BpfRecorder, arg2 context.Context, arg3 uint32) error {
	fake.waitForPidExitMutex.Lock()
	ret, specificReturn := fake.waitForPidExitReturnsOnCall[len(fake.waitForPidExitArgsForCall)]
	fake.waitForPidExitArgsForCall = append(fake.waitForPidExitArgsForCall, struct {
		arg1 *bpfrecorder.BpfRecorder
		arg2 context.Context
		arg3 uint32
	}{arg1, arg2, arg3})
	stub := fake.WaitForPidExitStub
	fakeReturns := fake.waitForPidExitReturns
	fake.recordInvocation("WaitForPidExit", []interface{}{arg1, arg2, arg3})
	fake.waitForPidExitMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) WaitForPidExitCallCount() int {
	fake.waitForPidExitMutex.RLock()
	defer fake.waitForPidExitMutex.RUnlock()
	return len(fake.waitForPidExitArgsForCall)
}

func (fake *FakeImpl) WaitForPidExitCalls(stub func(*bpfrecorder.BpfRecorder, context.Context, uint32) error) {
	fake.waitForPidExitMutex.Lock()
	defer fake.waitForPidExitMutex.Unlock()
	fake.WaitForPidExitStub = stub
}

func (fake *FakeImpl) WaitForPidExitArgsForCall(i int) (*bpfrecorder.BpfRecorder, context.Context, uint32) {
	fake.waitForPidExitMutex.RLock()
	defer fake.waitForPidExitMutex.RUnlock()
	argsForCall := fake.waitForPidExitArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeImpl) WaitForPidExitReturns(result1 error) {
	fake.waitForPidExitMutex.Lock()
	defer fake.waitForPidExitMutex.Unlock()
	fake.WaitForPidExitStub = nil
	fake.waitForPidExitReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) WaitForPidExitReturnsOnCall(i int, result1 error) {
	fake.waitForPidExitMutex.Lock()
	defer fake.waitForPidExitMutex.Unlock()
	fake.WaitForPidExitStub = nil
	if fake.waitForPidExitReturnsOnCall == nil {
		fake.waitForPidExitReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.waitForPidExitReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.commandRunMutex.RLock()
	defer fake.commandRunMutex.RUnlock()
	fake.commandWaitMutex.RLock()
	defer fake.commandWaitMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.findProcMountNamespaceMutex.RLock()
	defer fake.findProcMountNamespaceMutex.RUnlock()
	fake.getNameMutex.RLock()
	defer fake.getNameMutex.RUnlock()
	fake.goArchToSeccompArchMutex.RLock()
	defer fake.goArchToSeccompArchMutex.RUnlock()
	fake.iteratorKeyMutex.RLock()
	defer fake.iteratorKeyMutex.RUnlock()
	fake.iteratorNextMutex.RLock()
	defer fake.iteratorNextMutex.RUnlock()
	fake.loadBpfRecorderMutex.RLock()
	defer fake.loadBpfRecorderMutex.RUnlock()
	fake.marshalIndentMutex.RLock()
	defer fake.marshalIndentMutex.RUnlock()
	fake.notifyMutex.RLock()
	defer fake.notifyMutex.RUnlock()
	fake.printObjMutex.RLock()
	defer fake.printObjMutex.RUnlock()
	fake.syscallsGetValueMutex.RLock()
	defer fake.syscallsGetValueMutex.RUnlock()
	fake.syscallsIteratorMutex.RLock()
	defer fake.syscallsIteratorMutex.RUnlock()
	fake.unloadBpfRecorderMutex.RLock()
	defer fake.unloadBpfRecorderMutex.RUnlock()
	fake.waitForPidExitMutex.RLock()
	defer fake.waitForPidExitMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeImpl) recordInvocation(key string, args []interface{}) {
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
