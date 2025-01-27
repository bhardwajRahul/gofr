// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go
//
// Generated by this command:
//
//	mockgen -source=interfaces.go -destination=mock_interfaces.go -package=kafka
//

// Package kafka is a generated GoMock package.
package kafka

import (
	context "context"
	reflect "reflect"

	kafka "github.com/segmentio/kafka-go"
	gomock "go.uber.org/mock/gomock"
)

// MockReader is a mock of Reader interface.
type MockReader struct {
	ctrl     *gomock.Controller
	recorder *MockReaderMockRecorder
}

// MockReaderMockRecorder is the mock recorder for MockReader.
type MockReaderMockRecorder struct {
	mock *MockReader
}

// NewMockReader creates a new mock instance.
func NewMockReader(ctrl *gomock.Controller) *MockReader {
	mock := &MockReader{ctrl: ctrl}
	mock.recorder = &MockReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReader) EXPECT() *MockReaderMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockReader) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockReaderMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockReader)(nil).Close))
}

// CommitMessages mocks base method.
func (m *MockReader) CommitMessages(ctx context.Context, msgs ...kafka.Message) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx}
	for _, a := range msgs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CommitMessages", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CommitMessages indicates an expected call of CommitMessages.
func (mr *MockReaderMockRecorder) CommitMessages(ctx any, msgs ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx}, msgs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommitMessages", reflect.TypeOf((*MockReader)(nil).CommitMessages), varargs...)
}

// FetchMessage mocks base method.
func (m *MockReader) FetchMessage(ctx context.Context) (kafka.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchMessage", ctx)
	ret0, _ := ret[0].(kafka.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchMessage indicates an expected call of FetchMessage.
func (mr *MockReaderMockRecorder) FetchMessage(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchMessage", reflect.TypeOf((*MockReader)(nil).FetchMessage), ctx)
}

// ReadMessage mocks base method.
func (m *MockReader) ReadMessage(ctx context.Context) (kafka.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadMessage", ctx)
	ret0, _ := ret[0].(kafka.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadMessage indicates an expected call of ReadMessage.
func (mr *MockReaderMockRecorder) ReadMessage(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadMessage", reflect.TypeOf((*MockReader)(nil).ReadMessage), ctx)
}

// Stats mocks base method.
func (m *MockReader) Stats() kafka.ReaderStats {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stats")
	ret0, _ := ret[0].(kafka.ReaderStats)
	return ret0
}

// Stats indicates an expected call of Stats.
func (mr *MockReaderMockRecorder) Stats() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stats", reflect.TypeOf((*MockReader)(nil).Stats))
}

// MockWriter is a mock of Writer interface.
type MockWriter struct {
	ctrl     *gomock.Controller
	recorder *MockWriterMockRecorder
}

// MockWriterMockRecorder is the mock recorder for MockWriter.
type MockWriterMockRecorder struct {
	mock *MockWriter
}

// NewMockWriter creates a new mock instance.
func NewMockWriter(ctrl *gomock.Controller) *MockWriter {
	mock := &MockWriter{ctrl: ctrl}
	mock.recorder = &MockWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWriter) EXPECT() *MockWriterMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockWriter) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockWriterMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockWriter)(nil).Close))
}

// Stats mocks base method.
func (m *MockWriter) Stats() kafka.WriterStats {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stats")
	ret0, _ := ret[0].(kafka.WriterStats)
	return ret0
}

// Stats indicates an expected call of Stats.
func (mr *MockWriterMockRecorder) Stats() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stats", reflect.TypeOf((*MockWriter)(nil).Stats))
}

// WriteMessages mocks base method.
func (m *MockWriter) WriteMessages(ctx context.Context, msg ...kafka.Message) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx}
	for _, a := range msg {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WriteMessages", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteMessages indicates an expected call of WriteMessages.
func (mr *MockWriterMockRecorder) WriteMessages(ctx any, msg ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx}, msg...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteMessages", reflect.TypeOf((*MockWriter)(nil).WriteMessages), varargs...)
}

// MockConnection is a mock of Connection interface.
type MockConnection struct {
	ctrl     *gomock.Controller
	recorder *MockConnectionMockRecorder
}

// MockConnectionMockRecorder is the mock recorder for MockConnection.
type MockConnectionMockRecorder struct {
	mock *MockConnection
}

// NewMockConnection creates a new mock instance.
func NewMockConnection(ctrl *gomock.Controller) *MockConnection {
	mock := &MockConnection{ctrl: ctrl}
	mock.recorder = &MockConnectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConnection) EXPECT() *MockConnectionMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockConnection) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockConnectionMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockConnection)(nil).Close))
}

// Controller mocks base method.
func (m *MockConnection) Controller() (kafka.Broker, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Controller")
	ret0, _ := ret[0].(kafka.Broker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Controller indicates an expected call of Controller.
func (mr *MockConnectionMockRecorder) Controller() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Controller", reflect.TypeOf((*MockConnection)(nil).Controller))
}

// CreateTopics mocks base method.
func (m *MockConnection) CreateTopics(topics ...kafka.TopicConfig) error {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range topics {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateTopics", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTopics indicates an expected call of CreateTopics.
func (mr *MockConnectionMockRecorder) CreateTopics(topics ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTopics", reflect.TypeOf((*MockConnection)(nil).CreateTopics), topics...)
}

// DeleteTopics mocks base method.
func (m *MockConnection) DeleteTopics(topics ...string) error {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range topics {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteTopics", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTopics indicates an expected call of DeleteTopics.
func (mr *MockConnectionMockRecorder) DeleteTopics(topics ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTopics", reflect.TypeOf((*MockConnection)(nil).DeleteTopics), topics...)
}
