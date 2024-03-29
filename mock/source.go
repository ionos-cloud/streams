// Code generated by go-mockgen 1.3.7; DO NOT EDIT.

package mock

import (
	streams "github.com/ionos-cloud/streams"
	msg "github.com/ionos-cloud/streams/msg"
	"sync"
)

// MockSource is a mock implementation of the Source interface (from the
// package github.com/ionos-cloud/streams) used for unit testing.
type MockSource[K interface{}, V interface{}] struct {
	// CommitFunc is an instance of a mock function object controlling the
	// behavior of the method Commit.
	CommitFunc *SourceCommitFunc[K, V]
	// ErrorFunc is an instance of a mock function object controlling the
	// behavior of the method Error.
	ErrorFunc *SourceErrorFunc[K, V]
	// MessagesFunc is an instance of a mock function object controlling the
	// behavior of the method Messages.
	MessagesFunc *SourceMessagesFunc[K, V]
}

// NewMockSource creates a new mock of the Source interface. All methods
// return zero values for all results, unless overwritten.
func NewMockSource[K interface{}, V interface{}]() *MockSource[K, V] {
	return &MockSource[K, V]{
		CommitFunc: &SourceCommitFunc[K, V]{
			defaultHook: func(...msg.Message[K, V]) (r0 error) {
				return
			},
		},
		ErrorFunc: &SourceErrorFunc[K, V]{
			defaultHook: func() (r0 error) {
				return
			},
		},
		MessagesFunc: &SourceMessagesFunc[K, V]{
			defaultHook: func() (r0 chan msg.Message[K, V]) {
				return
			},
		},
	}
}

// NewStrictMockSource creates a new mock of the Source interface. All
// methods panic on invocation, unless overwritten.
func NewStrictMockSource[K interface{}, V interface{}]() *MockSource[K, V] {
	return &MockSource[K, V]{
		CommitFunc: &SourceCommitFunc[K, V]{
			defaultHook: func(...msg.Message[K, V]) error {
				panic("unexpected invocation of MockSource.Commit")
			},
		},
		ErrorFunc: &SourceErrorFunc[K, V]{
			defaultHook: func() error {
				panic("unexpected invocation of MockSource.Error")
			},
		},
		MessagesFunc: &SourceMessagesFunc[K, V]{
			defaultHook: func() chan msg.Message[K, V] {
				panic("unexpected invocation of MockSource.Messages")
			},
		},
	}
}

// NewMockSourceFrom creates a new mock of the MockSource interface. All
// methods delegate to the given implementation, unless overwritten.
func NewMockSourceFrom[K interface{}, V interface{}](i streams.Source[K, V]) *MockSource[K, V] {
	return &MockSource[K, V]{
		CommitFunc: &SourceCommitFunc[K, V]{
			defaultHook: i.Commit,
		},
		ErrorFunc: &SourceErrorFunc[K, V]{
			defaultHook: i.Error,
		},
		MessagesFunc: &SourceMessagesFunc[K, V]{
			defaultHook: i.Messages,
		},
	}
}

// SourceCommitFunc describes the behavior when the Commit method of the
// parent MockSource instance is invoked.
type SourceCommitFunc[K interface{}, V interface{}] struct {
	defaultHook func(...msg.Message[K, V]) error
	hooks       []func(...msg.Message[K, V]) error
	history     []SourceCommitFuncCall[K, V]
	mutex       sync.Mutex
}

// Commit delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockSource[K, V]) Commit(v0 ...msg.Message[K, V]) error {
	r0 := m.CommitFunc.nextHook()(v0...)
	m.CommitFunc.appendCall(SourceCommitFuncCall[K, V]{v0, r0})
	return r0
}

// SetDefaultHook sets function that is called when the Commit method of the
// parent MockSource instance is invoked and the hook queue is empty.
func (f *SourceCommitFunc[K, V]) SetDefaultHook(hook func(...msg.Message[K, V]) error) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Commit method of the parent MockSource instance invokes the hook at the
// front of the queue and discards it. After the queue is empty, the default
// hook function is invoked for any future action.
func (f *SourceCommitFunc[K, V]) PushHook(hook func(...msg.Message[K, V]) error) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *SourceCommitFunc[K, V]) SetDefaultReturn(r0 error) {
	f.SetDefaultHook(func(...msg.Message[K, V]) error {
		return r0
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *SourceCommitFunc[K, V]) PushReturn(r0 error) {
	f.PushHook(func(...msg.Message[K, V]) error {
		return r0
	})
}

func (f *SourceCommitFunc[K, V]) nextHook() func(...msg.Message[K, V]) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *SourceCommitFunc[K, V]) appendCall(r0 SourceCommitFuncCall[K, V]) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of SourceCommitFuncCall objects describing the
// invocations of this function.
func (f *SourceCommitFunc[K, V]) History() []SourceCommitFuncCall[K, V] {
	f.mutex.Lock()
	history := make([]SourceCommitFuncCall[K, V], len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// SourceCommitFuncCall is an object that describes an invocation of method
// Commit on an instance of MockSource.
type SourceCommitFuncCall[K interface{}, V interface{}] struct {
	// Arg0 is a slice containing the values of the variadic arguments
	// passed to this method invocation.
	Arg0 []msg.Message[K, V]
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 error
}

// Args returns an interface slice containing the arguments of this
// invocation. The variadic slice argument is flattened in this array such
// that one positional argument and three variadic arguments would result in
// a slice of four, not two.
func (c SourceCommitFuncCall[K, V]) Args() []interface{} {
	trailing := []interface{}{}
	for _, val := range c.Arg0 {
		trailing = append(trailing, val)
	}

	return append([]interface{}{}, trailing...)
}

// Results returns an interface slice containing the results of this
// invocation.
func (c SourceCommitFuncCall[K, V]) Results() []interface{} {
	return []interface{}{c.Result0}
}

// SourceErrorFunc describes the behavior when the Error method of the
// parent MockSource instance is invoked.
type SourceErrorFunc[K interface{}, V interface{}] struct {
	defaultHook func() error
	hooks       []func() error
	history     []SourceErrorFuncCall[K, V]
	mutex       sync.Mutex
}

// Error delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockSource[K, V]) Error() error {
	r0 := m.ErrorFunc.nextHook()()
	m.ErrorFunc.appendCall(SourceErrorFuncCall[K, V]{r0})
	return r0
}

// SetDefaultHook sets function that is called when the Error method of the
// parent MockSource instance is invoked and the hook queue is empty.
func (f *SourceErrorFunc[K, V]) SetDefaultHook(hook func() error) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Error method of the parent MockSource instance invokes the hook at the
// front of the queue and discards it. After the queue is empty, the default
// hook function is invoked for any future action.
func (f *SourceErrorFunc[K, V]) PushHook(hook func() error) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *SourceErrorFunc[K, V]) SetDefaultReturn(r0 error) {
	f.SetDefaultHook(func() error {
		return r0
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *SourceErrorFunc[K, V]) PushReturn(r0 error) {
	f.PushHook(func() error {
		return r0
	})
}

func (f *SourceErrorFunc[K, V]) nextHook() func() error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *SourceErrorFunc[K, V]) appendCall(r0 SourceErrorFuncCall[K, V]) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of SourceErrorFuncCall objects describing the
// invocations of this function.
func (f *SourceErrorFunc[K, V]) History() []SourceErrorFuncCall[K, V] {
	f.mutex.Lock()
	history := make([]SourceErrorFuncCall[K, V], len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// SourceErrorFuncCall is an object that describes an invocation of method
// Error on an instance of MockSource.
type SourceErrorFuncCall[K interface{}, V interface{}] struct {
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c SourceErrorFuncCall[K, V]) Args() []interface{} {
	return []interface{}{}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c SourceErrorFuncCall[K, V]) Results() []interface{} {
	return []interface{}{c.Result0}
}

// SourceMessagesFunc describes the behavior when the Messages method of the
// parent MockSource instance is invoked.
type SourceMessagesFunc[K interface{}, V interface{}] struct {
	defaultHook func() chan msg.Message[K, V]
	hooks       []func() chan msg.Message[K, V]
	history     []SourceMessagesFuncCall[K, V]
	mutex       sync.Mutex
}

// Messages delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockSource[K, V]) Messages() chan msg.Message[K, V] {
	r0 := m.MessagesFunc.nextHook()()
	m.MessagesFunc.appendCall(SourceMessagesFuncCall[K, V]{r0})
	return r0
}

// SetDefaultHook sets function that is called when the Messages method of
// the parent MockSource instance is invoked and the hook queue is empty.
func (f *SourceMessagesFunc[K, V]) SetDefaultHook(hook func() chan msg.Message[K, V]) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Messages method of the parent MockSource instance invokes the hook at the
// front of the queue and discards it. After the queue is empty, the default
// hook function is invoked for any future action.
func (f *SourceMessagesFunc[K, V]) PushHook(hook func() chan msg.Message[K, V]) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *SourceMessagesFunc[K, V]) SetDefaultReturn(r0 chan msg.Message[K, V]) {
	f.SetDefaultHook(func() chan msg.Message[K, V] {
		return r0
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *SourceMessagesFunc[K, V]) PushReturn(r0 chan msg.Message[K, V]) {
	f.PushHook(func() chan msg.Message[K, V] {
		return r0
	})
}

func (f *SourceMessagesFunc[K, V]) nextHook() func() chan msg.Message[K, V] {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *SourceMessagesFunc[K, V]) appendCall(r0 SourceMessagesFuncCall[K, V]) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of SourceMessagesFuncCall objects describing
// the invocations of this function.
func (f *SourceMessagesFunc[K, V]) History() []SourceMessagesFuncCall[K, V] {
	f.mutex.Lock()
	history := make([]SourceMessagesFuncCall[K, V], len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// SourceMessagesFuncCall is an object that describes an invocation of
// method Messages on an instance of MockSource.
type SourceMessagesFuncCall[K interface{}, V interface{}] struct {
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 chan msg.Message[K, V]
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c SourceMessagesFuncCall[K, V]) Args() []interface{} {
	return []interface{}{}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c SourceMessagesFuncCall[K, V]) Results() []interface{} {
	return []interface{}{c.Result0}
}
