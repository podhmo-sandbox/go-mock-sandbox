// Code generated by mockery v1.0.0
package tmock

import mock "github.com/stretchr/testify/mock"
import user "github.com/podhmo-sandbox/go-mock-sandbox/user"

// UserService is an autogenerated mock type for the UserService type
type UserService struct {
	mock.Mock
}

// Get provides a mock function with given fields: id
func (_m *UserService) Get(id string) (*user.User, error) {
	ret := _m.Called(id)

	var r0 *user.User
	if rf, ok := ret.Get(0).(func(string) *user.User); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
