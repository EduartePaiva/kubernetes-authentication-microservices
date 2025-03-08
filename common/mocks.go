package common

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

/*
  Test objects
*/

// MyMockedObject is a mocked object that implements an interface
// that describes an object that the code I am testing relies on.
type HttpResponseWriterMock struct {
	mock.Mock
}

// DoSomething is a method on MyMockedObject that implements some interface
// and just records the activity, and returns what the Mock object tells it to.
//
// In the real object, this method would do something useful, but since this
// is a mocked object - we're just going to stub it out.
//
// NOTE: This method is not being tested here, code that uses this object is.
func (m *HttpResponseWriterMock) Header() http.Header {
	return http.Header{}
}
func (m *HttpResponseWriterMock) Write(data []byte) (int, error) {
	args := m.Called(data)

	// I'm not using the returned value of write
	return args.Int(0), args.Error(1)
}
func (m *HttpResponseWriterMock) WriteHeader(status int) {
	m.Called(status)
}
