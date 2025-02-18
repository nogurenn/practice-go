// Code generated by mockery v2.50.4. DO NOT EDIT.

package transaction

import (
	sql "database/sql"

	mock "github.com/stretchr/testify/mock"

	sqlx "github.com/jmoiron/sqlx"
)

// Transaction is an autogenerated mock type for the Transaction type
type Transaction struct {
	mock.Mock
}

type Transaction_Expecter struct {
	mock *mock.Mock
}

func (_m *Transaction) EXPECT() *Transaction_Expecter {
	return &Transaction_Expecter{mock: &_m.Mock}
}

// Commit provides a mock function with no fields
func (_m *Transaction) Commit() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Commit")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Transaction_Commit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Commit'
type Transaction_Commit_Call struct {
	*mock.Call
}

// Commit is a helper method to define mock.On call
func (_e *Transaction_Expecter) Commit() *Transaction_Commit_Call {
	return &Transaction_Commit_Call{Call: _e.mock.On("Commit")}
}

func (_c *Transaction_Commit_Call) Run(run func()) *Transaction_Commit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Transaction_Commit_Call) Return(_a0 error) *Transaction_Commit_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Transaction_Commit_Call) RunAndReturn(run func() error) *Transaction_Commit_Call {
	_c.Call.Return(run)
	return _c
}

// Exec provides a mock function with given fields: query, args
func (_m *Transaction) Exec(query string, args ...interface{}) (sql.Result, error) {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Exec")
	}

	var r0 sql.Result
	var r1 error
	if rf, ok := ret.Get(0).(func(string, ...interface{}) (sql.Result, error)); ok {
		return rf(query, args...)
	}
	if rf, ok := ret.Get(0).(func(string, ...interface{}) sql.Result); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sql.Result)
		}
	}

	if rf, ok := ret.Get(1).(func(string, ...interface{}) error); ok {
		r1 = rf(query, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Transaction_Exec_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Exec'
type Transaction_Exec_Call struct {
	*mock.Call
}

// Exec is a helper method to define mock.On call
//   - query string
//   - args ...interface{}
func (_e *Transaction_Expecter) Exec(query interface{}, args ...interface{}) *Transaction_Exec_Call {
	return &Transaction_Exec_Call{Call: _e.mock.On("Exec",
		append([]interface{}{query}, args...)...)}
}

func (_c *Transaction_Exec_Call) Run(run func(query string, args ...interface{})) *Transaction_Exec_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *Transaction_Exec_Call) Return(_a0 sql.Result, _a1 error) *Transaction_Exec_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Transaction_Exec_Call) RunAndReturn(run func(string, ...interface{}) (sql.Result, error)) *Transaction_Exec_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: dest, query, args
func (_m *Transaction) Get(dest interface{}, query string, args ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, dest, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, string, ...interface{}) error); ok {
		r0 = rf(dest, query, args...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Transaction_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type Transaction_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - dest interface{}
//   - query string
//   - args ...interface{}
func (_e *Transaction_Expecter) Get(dest interface{}, query interface{}, args ...interface{}) *Transaction_Get_Call {
	return &Transaction_Get_Call{Call: _e.mock.On("Get",
		append([]interface{}{dest, query}, args...)...)}
}

func (_c *Transaction_Get_Call) Run(run func(dest interface{}, query string, args ...interface{})) *Transaction_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(interface{}), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *Transaction_Get_Call) Return(_a0 error) *Transaction_Get_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Transaction_Get_Call) RunAndReturn(run func(interface{}, string, ...interface{}) error) *Transaction_Get_Call {
	_c.Call.Return(run)
	return _c
}

// NamedExec provides a mock function with given fields: query, arg
func (_m *Transaction) NamedExec(query string, arg interface{}) (sql.Result, error) {
	ret := _m.Called(query, arg)

	if len(ret) == 0 {
		panic("no return value specified for NamedExec")
	}

	var r0 sql.Result
	var r1 error
	if rf, ok := ret.Get(0).(func(string, interface{}) (sql.Result, error)); ok {
		return rf(query, arg)
	}
	if rf, ok := ret.Get(0).(func(string, interface{}) sql.Result); ok {
		r0 = rf(query, arg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sql.Result)
		}
	}

	if rf, ok := ret.Get(1).(func(string, interface{}) error); ok {
		r1 = rf(query, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Transaction_NamedExec_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NamedExec'
type Transaction_NamedExec_Call struct {
	*mock.Call
}

// NamedExec is a helper method to define mock.On call
//   - query string
//   - arg interface{}
func (_e *Transaction_Expecter) NamedExec(query interface{}, arg interface{}) *Transaction_NamedExec_Call {
	return &Transaction_NamedExec_Call{Call: _e.mock.On("NamedExec", query, arg)}
}

func (_c *Transaction_NamedExec_Call) Run(run func(query string, arg interface{})) *Transaction_NamedExec_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(interface{}))
	})
	return _c
}

func (_c *Transaction_NamedExec_Call) Return(_a0 sql.Result, _a1 error) *Transaction_NamedExec_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Transaction_NamedExec_Call) RunAndReturn(run func(string, interface{}) (sql.Result, error)) *Transaction_NamedExec_Call {
	_c.Call.Return(run)
	return _c
}

// Queryx provides a mock function with given fields: query, args
func (_m *Transaction) Queryx(query string, args ...interface{}) (*sqlx.Rows, error) {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Queryx")
	}

	var r0 *sqlx.Rows
	var r1 error
	if rf, ok := ret.Get(0).(func(string, ...interface{}) (*sqlx.Rows, error)); ok {
		return rf(query, args...)
	}
	if rf, ok := ret.Get(0).(func(string, ...interface{}) *sqlx.Rows); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sqlx.Rows)
		}
	}

	if rf, ok := ret.Get(1).(func(string, ...interface{}) error); ok {
		r1 = rf(query, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Transaction_Queryx_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Queryx'
type Transaction_Queryx_Call struct {
	*mock.Call
}

// Queryx is a helper method to define mock.On call
//   - query string
//   - args ...interface{}
func (_e *Transaction_Expecter) Queryx(query interface{}, args ...interface{}) *Transaction_Queryx_Call {
	return &Transaction_Queryx_Call{Call: _e.mock.On("Queryx",
		append([]interface{}{query}, args...)...)}
}

func (_c *Transaction_Queryx_Call) Run(run func(query string, args ...interface{})) *Transaction_Queryx_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *Transaction_Queryx_Call) Return(_a0 *sqlx.Rows, _a1 error) *Transaction_Queryx_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Transaction_Queryx_Call) RunAndReturn(run func(string, ...interface{}) (*sqlx.Rows, error)) *Transaction_Queryx_Call {
	_c.Call.Return(run)
	return _c
}

// Rebind provides a mock function with given fields: query
func (_m *Transaction) Rebind(query string) string {
	ret := _m.Called(query)

	if len(ret) == 0 {
		panic("no return value specified for Rebind")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(query)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Transaction_Rebind_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Rebind'
type Transaction_Rebind_Call struct {
	*mock.Call
}

// Rebind is a helper method to define mock.On call
//   - query string
func (_e *Transaction_Expecter) Rebind(query interface{}) *Transaction_Rebind_Call {
	return &Transaction_Rebind_Call{Call: _e.mock.On("Rebind", query)}
}

func (_c *Transaction_Rebind_Call) Run(run func(query string)) *Transaction_Rebind_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Transaction_Rebind_Call) Return(_a0 string) *Transaction_Rebind_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Transaction_Rebind_Call) RunAndReturn(run func(string) string) *Transaction_Rebind_Call {
	_c.Call.Return(run)
	return _c
}

// Rollback provides a mock function with no fields
func (_m *Transaction) Rollback() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Rollback")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Transaction_Rollback_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Rollback'
type Transaction_Rollback_Call struct {
	*mock.Call
}

// Rollback is a helper method to define mock.On call
func (_e *Transaction_Expecter) Rollback() *Transaction_Rollback_Call {
	return &Transaction_Rollback_Call{Call: _e.mock.On("Rollback")}
}

func (_c *Transaction_Rollback_Call) Run(run func()) *Transaction_Rollback_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Transaction_Rollback_Call) Return(_a0 error) *Transaction_Rollback_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Transaction_Rollback_Call) RunAndReturn(run func() error) *Transaction_Rollback_Call {
	_c.Call.Return(run)
	return _c
}

// Select provides a mock function with given fields: dest, query, args
func (_m *Transaction) Select(dest interface{}, query string, args ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, dest, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Select")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, string, ...interface{}) error); ok {
		r0 = rf(dest, query, args...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Transaction_Select_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Select'
type Transaction_Select_Call struct {
	*mock.Call
}

// Select is a helper method to define mock.On call
//   - dest interface{}
//   - query string
//   - args ...interface{}
func (_e *Transaction_Expecter) Select(dest interface{}, query interface{}, args ...interface{}) *Transaction_Select_Call {
	return &Transaction_Select_Call{Call: _e.mock.On("Select",
		append([]interface{}{dest, query}, args...)...)}
}

func (_c *Transaction_Select_Call) Run(run func(dest interface{}, query string, args ...interface{})) *Transaction_Select_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(interface{}), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *Transaction_Select_Call) Return(_a0 error) *Transaction_Select_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Transaction_Select_Call) RunAndReturn(run func(interface{}, string, ...interface{}) error) *Transaction_Select_Call {
	_c.Call.Return(run)
	return _c
}

// NewTransaction creates a new instance of Transaction. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTransaction(t interface {
	mock.TestingT
	Cleanup(func())
}) *Transaction {
	mock := &Transaction{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
