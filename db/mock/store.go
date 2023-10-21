// Code generated by mockery v2.33.3. DO NOT EDIT.

package mockdb

import (
	context "context"

	db "github.com/fauzanfebrian/simplebank/db/sqlc"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// MockStore is an autogenerated mock type for the Store type
type MockStore struct {
	mock.Mock
}

type MockStore_Expecter struct {
	mock *mock.Mock
}

func (_m *MockStore) EXPECT() *MockStore_Expecter {
	return &MockStore_Expecter{mock: &_m.Mock}
}

// AddAccountBalance provides a mock function with given fields: ctx, arg
func (_m *MockStore) AddAccountBalance(ctx context.Context, arg db.AddAccountBalanceParams) (db.Account, error) {
	ret := _m.Called(ctx, arg)

	var r0 db.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.AddAccountBalanceParams) (db.Account, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.AddAccountBalanceParams) db.Account); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(db.Account)
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.AddAccountBalanceParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStore_AddAccountBalance_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddAccountBalance'
type MockStore_AddAccountBalance_Call struct {
	*mock.Call
}

// AddAccountBalance is a helper method to define mock.On call
//   - ctx context.Context
//   - arg db.AddAccountBalanceParams
func (_e *MockStore_Expecter) AddAccountBalance(ctx interface{}, arg interface{}) *MockStore_AddAccountBalance_Call {
	return &MockStore_AddAccountBalance_Call{Call: _e.mock.On("AddAccountBalance", ctx, arg)}
}

func (_c *MockStore_AddAccountBalance_Call) Run(run func(ctx context.Context, arg db.AddAccountBalanceParams)) *MockStore_AddAccountBalance_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(db.AddAccountBalanceParams))
	})
	return _c
}

func (_c *MockStore_AddAccountBalance_Call) Return(_a0 db.Account, _a1 error) *MockStore_AddAccountBalance_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStore_AddAccountBalance_Call) RunAndReturn(run func(context.Context, db.AddAccountBalanceParams) (db.Account, error)) *MockStore_AddAccountBalance_Call {
	_c.Call.Return(run)
	return _c
}

// CreateAccount provides a mock function with given fields: ctx, arg
func (_m *MockStore) CreateAccount(ctx context.Context, arg db.CreateAccountParams) (db.Account, error) {
	ret := _m.Called(ctx, arg)

	var r0 db.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.CreateAccountParams) (db.Account, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.CreateAccountParams) db.Account); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(db.Account)
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.CreateAccountParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStore_CreateAccount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateAccount'
type MockStore_CreateAccount_Call struct {
	*mock.Call
}

// CreateAccount is a helper method to define mock.On call
//   - ctx context.Context
//   - arg db.CreateAccountParams
func (_e *MockStore_Expecter) CreateAccount(ctx interface{}, arg interface{}) *MockStore_CreateAccount_Call {
	return &MockStore_CreateAccount_Call{Call: _e.mock.On("CreateAccount", ctx, arg)}
}

func (_c *MockStore_CreateAccount_Call) Run(run func(ctx context.Context, arg db.CreateAccountParams)) *MockStore_CreateAccount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(db.CreateAccountParams))
	})
	return _c
}

func (_c *MockStore_CreateAccount_Call) Return(_a0 db.Account, _a1 error) *MockStore_CreateAccount_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStore_CreateAccount_Call) RunAndReturn(run func(context.Context, db.CreateAccountParams) (db.Account, error)) *MockStore_CreateAccount_Call {
	_c.Call.Return(run)
	return _c
}

// CreateEntry provides a mock function with given fields: ctx, arg
func (_m *MockStore) CreateEntry(ctx context.Context, arg db.CreateEntryParams) (db.Entry, error) {
	ret := _m.Called(ctx, arg)

	var r0 db.Entry
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.CreateEntryParams) (db.Entry, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.CreateEntryParams) db.Entry); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(db.Entry)
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.CreateEntryParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStore_CreateEntry_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateEntry'
type MockStore_CreateEntry_Call struct {
	*mock.Call
}

// CreateEntry is a helper method to define mock.On call
//   - ctx context.Context
//   - arg db.CreateEntryParams
func (_e *MockStore_Expecter) CreateEntry(ctx interface{}, arg interface{}) *MockStore_CreateEntry_Call {
	return &MockStore_CreateEntry_Call{Call: _e.mock.On("CreateEntry", ctx, arg)}
}

func (_c *MockStore_CreateEntry_Call) Run(run func(ctx context.Context, arg db.CreateEntryParams)) *MockStore_CreateEntry_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(db.CreateEntryParams))
	})
	return _c
}

func (_c *MockStore_CreateEntry_Call) Return(_a0 db.Entry, _a1 error) *MockStore_CreateEntry_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStore_CreateEntry_Call) RunAndReturn(run func(context.Context, db.CreateEntryParams) (db.Entry, error)) *MockStore_CreateEntry_Call {
	_c.Call.Return(run)
	return _c
}

// CreateSession provides a mock function with given fields: ctx, arg
func (_m *MockStore) CreateSession(ctx context.Context, arg db.CreateSessionParams) (db.Session, error) {
	ret := _m.Called(ctx, arg)

	var r0 db.Session
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.CreateSessionParams) (db.Session, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.CreateSessionParams) db.Session); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(db.Session)
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.CreateSessionParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStore_CreateSession_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateSession'
type MockStore_CreateSession_Call struct {
	*mock.Call
}

// CreateSession is a helper method to define mock.On call
//   - ctx context.Context
//   - arg db.CreateSessionParams
func (_e *MockStore_Expecter) CreateSession(ctx interface{}, arg interface{}) *MockStore_CreateSession_Call {
	return &MockStore_CreateSession_Call{Call: _e.mock.On("CreateSession", ctx, arg)}
}

func (_c *MockStore_CreateSession_Call) Run(run func(ctx context.Context, arg db.CreateSessionParams)) *MockStore_CreateSession_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(db.CreateSessionParams))
	})
	return _c
}

func (_c *MockStore_CreateSession_Call) Return(_a0 db.Session, _a1 error) *MockStore_CreateSession_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStore_CreateSession_Call) RunAndReturn(run func(context.Context, db.CreateSessionParams) (db.Session, error)) *MockStore_CreateSession_Call {
	_c.Call.Return(run)
	return _c
}

// CreateTransfer provides a mock function with given fields: ctx, arg
func (_m *MockStore) CreateTransfer(ctx context.Context, arg db.CreateTransferParams) (db.Transfer, error) {
	ret := _m.Called(ctx, arg)

	var r0 db.Transfer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.CreateTransferParams) (db.Transfer, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.CreateTransferParams) db.Transfer); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(db.Transfer)
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.CreateTransferParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStore_CreateTransfer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateTransfer'
type MockStore_CreateTransfer_Call struct {
	*mock.Call
}

// CreateTransfer is a helper method to define mock.On call
//   - ctx context.Context
//   - arg db.CreateTransferParams
func (_e *MockStore_Expecter) CreateTransfer(ctx interface{}, arg interface{}) *MockStore_CreateTransfer_Call {
	return &MockStore_CreateTransfer_Call{Call: _e.mock.On("CreateTransfer", ctx, arg)}
}

func (_c *MockStore_CreateTransfer_Call) Run(run func(ctx context.Context, arg db.CreateTransferParams)) *MockStore_CreateTransfer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(db.CreateTransferParams))
	})
	return _c
}

func (_c *MockStore_CreateTransfer_Call) Return(_a0 db.Transfer, _a1 error) *MockStore_CreateTransfer_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStore_CreateTransfer_Call) RunAndReturn(run func(context.Context, db.CreateTransferParams) (db.Transfer, error)) *MockStore_CreateTransfer_Call {
	_c.Call.Return(run)
	return _c
}

// CreateUser provides a mock function with given fields: ctx, arg
func (_m *MockStore) CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error) {
	ret := _m.Called(ctx, arg)

	var r0 db.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.CreateUserParams) (db.User, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.CreateUserParams) db.User); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(db.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.CreateUserParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStore_CreateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateUser'
type MockStore_CreateUser_Call struct {
	*mock.Call
}

// CreateUser is a helper method to define mock.On call
//   - ctx context.Context
//   - arg db.CreateUserParams
func (_e *MockStore_Expecter) CreateUser(ctx interface{}, arg interface{}) *MockStore_CreateUser_Call {
	return &MockStore_CreateUser_Call{Call: _e.mock.On("CreateUser", ctx, arg)}
}

func (_c *MockStore_CreateUser_Call) Run(run func(ctx context.Context, arg db.CreateUserParams)) *MockStore_CreateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(db.CreateUserParams))
	})
	return _c
}

func (_c *MockStore_CreateUser_Call) Return(_a0 db.User, _a1 error) *MockStore_CreateUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStore_CreateUser_Call) RunAndReturn(run func(context.Context, db.CreateUserParams) (db.User, error)) *MockStore_CreateUser_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteAccount provides a mock function with given fields: ctx, id
func (_m *MockStore) DeleteAccount(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStore_DeleteAccount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteAccount'
type MockStore_DeleteAccount_Call struct {
	*mock.Call
}

// DeleteAccount is a helper method to define mock.On call
//   - ctx context.Context
//   - id int64
func (_e *MockStore_Expecter) DeleteAccount(ctx interface{}, id interface{}) *MockStore_DeleteAccount_Call {
	return &MockStore_DeleteAccount_Call{Call: _e.mock.On("DeleteAccount", ctx, id)}
}

func (_c *MockStore_DeleteAccount_Call) Run(run func(ctx context.Context, id int64)) *MockStore_DeleteAccount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockStore_DeleteAccount_Call) Return(_a0 error) *MockStore_DeleteAccount_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStore_DeleteAccount_Call) RunAndReturn(run func(context.Context, int64) error) *MockStore_DeleteAccount_Call {
	_c.Call.Return(run)
	return _c
}

// GetAccount provides a mock function with given fields: ctx, id
func (_m *MockStore) GetAccount(ctx context.Context, id int64) (db.Account, error) {
	ret := _m.Called(ctx, id)

	var r0 db.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (db.Account, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) db.Account); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(db.Account)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStore_GetAccount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAccount'
type MockStore_GetAccount_Call struct {
	*mock.Call
}

// GetAccount is a helper method to define mock.On call
//   - ctx context.Context
//   - id int64
func (_e *MockStore_Expecter) GetAccount(ctx interface{}, id interface{}) *MockStore_GetAccount_Call {
	return &MockStore_GetAccount_Call{Call: _e.mock.On("GetAccount", ctx, id)}
}

func (_c *MockStore_GetAccount_Call) Run(run func(ctx context.Context, id int64)) *MockStore_GetAccount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockStore_GetAccount_Call) Return(_a0 db.Account, _a1 error) *MockStore_GetAccount_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStore_GetAccount_Call) RunAndReturn(run func(context.Context, int64) (db.Account, error)) *MockStore_GetAccount_Call {
	_c.Call.Return(run)
	return _c
}

// GetAccountForUpdate provides a mock function with given fields: ctx, id
func (_m *MockStore) GetAccountForUpdate(ctx context.Context, id int64) (db.Account, error) {
	ret := _m.Called(ctx, id)

	var r0 db.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (db.Account, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) db.Account); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(db.Account)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStore_GetAccountForUpdate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAccountForUpdate'
type MockStore_GetAccountForUpdate_Call struct {
	*mock.Call
}

// GetAccountForUpdate is a helper method to define mock.On call
//   - ctx context.Context
//   - id int64
func (_e *MockStore_Expecter) GetAccountForUpdate(ctx interface{}, id interface{}) *MockStore_GetAccountForUpdate_Call {
	return &MockStore_GetAccountForUpdate_Call{Call: _e.mock.On("GetAccountForUpdate", ctx, id)}
}

func (_c *MockStore_GetAccountForUpdate_Call) Run(run func(ctx context.Context, id int64)) *MockStore_GetAccountForUpdate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockStore_GetAccountForUpdate_Call) Return(_a0 db.Account, _a1 error) *MockStore_GetAccountForUpdate_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStore_GetAccountForUpdate_Call) RunAndReturn(run func(context.Context, int64) (db.Account, error)) *MockStore_GetAccountForUpdate_Call {
	_c.Call.Return(run)
	return _c
}

// GetEntry provides a mock function with given fields: ctx, id
func (_m *MockStore) GetEntry(ctx context.Context, id int64) (db.Entry, error) {
	ret := _m.Called(ctx, id)

	var r0 db.Entry
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (db.Entry, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) db.Entry); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(db.Entry)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStore_GetEntry_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetEntry'
type MockStore_GetEntry_Call struct {
	*mock.Call
}

// GetEntry is a helper method to define mock.On call
//   - ctx context.Context
//   - id int64
func (_e *MockStore_Expecter) GetEntry(ctx interface{}, id interface{}) *MockStore_GetEntry_Call {
	return &MockStore_GetEntry_Call{Call: _e.mock.On("GetEntry", ctx, id)}
}

func (_c *MockStore_GetEntry_Call) Run(run func(ctx context.Context, id int64)) *MockStore_GetEntry_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockStore_GetEntry_Call) Return(_a0 db.Entry, _a1 error) *MockStore_GetEntry_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStore_GetEntry_Call) RunAndReturn(run func(context.Context, int64) (db.Entry, error)) *MockStore_GetEntry_Call {
	_c.Call.Return(run)
	return _c
}

// GetSession provides a mock function with given fields: ctx, id
func (_m *MockStore) GetSession(ctx context.Context, id uuid.UUID) (db.Session, error) {
	ret := _m.Called(ctx, id)

	var r0 db.Session
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (db.Session, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) db.Session); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(db.Session)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStore_GetSession_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSession'
type MockStore_GetSession_Call struct {
	*mock.Call
}

// GetSession is a helper method to define mock.On call
//   - ctx context.Context
//   - id uuid.UUID
func (_e *MockStore_Expecter) GetSession(ctx interface{}, id interface{}) *MockStore_GetSession_Call {
	return &MockStore_GetSession_Call{Call: _e.mock.On("GetSession", ctx, id)}
}

func (_c *MockStore_GetSession_Call) Run(run func(ctx context.Context, id uuid.UUID)) *MockStore_GetSession_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *MockStore_GetSession_Call) Return(_a0 db.Session, _a1 error) *MockStore_GetSession_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStore_GetSession_Call) RunAndReturn(run func(context.Context, uuid.UUID) (db.Session, error)) *MockStore_GetSession_Call {
	_c.Call.Return(run)
	return _c
}

// GetTransfer provides a mock function with given fields: ctx, id
func (_m *MockStore) GetTransfer(ctx context.Context, id int64) (db.Transfer, error) {
	ret := _m.Called(ctx, id)

	var r0 db.Transfer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (db.Transfer, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) db.Transfer); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(db.Transfer)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStore_GetTransfer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTransfer'
type MockStore_GetTransfer_Call struct {
	*mock.Call
}

// GetTransfer is a helper method to define mock.On call
//   - ctx context.Context
//   - id int64
func (_e *MockStore_Expecter) GetTransfer(ctx interface{}, id interface{}) *MockStore_GetTransfer_Call {
	return &MockStore_GetTransfer_Call{Call: _e.mock.On("GetTransfer", ctx, id)}
}

func (_c *MockStore_GetTransfer_Call) Run(run func(ctx context.Context, id int64)) *MockStore_GetTransfer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockStore_GetTransfer_Call) Return(_a0 db.Transfer, _a1 error) *MockStore_GetTransfer_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStore_GetTransfer_Call) RunAndReturn(run func(context.Context, int64) (db.Transfer, error)) *MockStore_GetTransfer_Call {
	_c.Call.Return(run)
	return _c
}

// GetUser provides a mock function with given fields: ctx, username
func (_m *MockStore) GetUser(ctx context.Context, username string) (db.User, error) {
	ret := _m.Called(ctx, username)

	var r0 db.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (db.User, error)); ok {
		return rf(ctx, username)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) db.User); ok {
		r0 = rf(ctx, username)
	} else {
		r0 = ret.Get(0).(db.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStore_GetUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUser'
type MockStore_GetUser_Call struct {
	*mock.Call
}

// GetUser is a helper method to define mock.On call
//   - ctx context.Context
//   - username string
func (_e *MockStore_Expecter) GetUser(ctx interface{}, username interface{}) *MockStore_GetUser_Call {
	return &MockStore_GetUser_Call{Call: _e.mock.On("GetUser", ctx, username)}
}

func (_c *MockStore_GetUser_Call) Run(run func(ctx context.Context, username string)) *MockStore_GetUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockStore_GetUser_Call) Return(_a0 db.User, _a1 error) *MockStore_GetUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStore_GetUser_Call) RunAndReturn(run func(context.Context, string) (db.User, error)) *MockStore_GetUser_Call {
	_c.Call.Return(run)
	return _c
}

// ListAccounts provides a mock function with given fields: ctx, arg
func (_m *MockStore) ListAccounts(ctx context.Context, arg db.ListAccountsParams) ([]db.Account, error) {
	ret := _m.Called(ctx, arg)

	var r0 []db.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.ListAccountsParams) ([]db.Account, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.ListAccountsParams) []db.Account); ok {
		r0 = rf(ctx, arg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]db.Account)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.ListAccountsParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStore_ListAccounts_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAccounts'
type MockStore_ListAccounts_Call struct {
	*mock.Call
}

// ListAccounts is a helper method to define mock.On call
//   - ctx context.Context
//   - arg db.ListAccountsParams
func (_e *MockStore_Expecter) ListAccounts(ctx interface{}, arg interface{}) *MockStore_ListAccounts_Call {
	return &MockStore_ListAccounts_Call{Call: _e.mock.On("ListAccounts", ctx, arg)}
}

func (_c *MockStore_ListAccounts_Call) Run(run func(ctx context.Context, arg db.ListAccountsParams)) *MockStore_ListAccounts_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(db.ListAccountsParams))
	})
	return _c
}

func (_c *MockStore_ListAccounts_Call) Return(_a0 []db.Account, _a1 error) *MockStore_ListAccounts_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStore_ListAccounts_Call) RunAndReturn(run func(context.Context, db.ListAccountsParams) ([]db.Account, error)) *MockStore_ListAccounts_Call {
	_c.Call.Return(run)
	return _c
}

// ListEntries provides a mock function with given fields: ctx, arg
func (_m *MockStore) ListEntries(ctx context.Context, arg db.ListEntriesParams) ([]db.Entry, error) {
	ret := _m.Called(ctx, arg)

	var r0 []db.Entry
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.ListEntriesParams) ([]db.Entry, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.ListEntriesParams) []db.Entry); ok {
		r0 = rf(ctx, arg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]db.Entry)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.ListEntriesParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStore_ListEntries_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListEntries'
type MockStore_ListEntries_Call struct {
	*mock.Call
}

// ListEntries is a helper method to define mock.On call
//   - ctx context.Context
//   - arg db.ListEntriesParams
func (_e *MockStore_Expecter) ListEntries(ctx interface{}, arg interface{}) *MockStore_ListEntries_Call {
	return &MockStore_ListEntries_Call{Call: _e.mock.On("ListEntries", ctx, arg)}
}

func (_c *MockStore_ListEntries_Call) Run(run func(ctx context.Context, arg db.ListEntriesParams)) *MockStore_ListEntries_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(db.ListEntriesParams))
	})
	return _c
}

func (_c *MockStore_ListEntries_Call) Return(_a0 []db.Entry, _a1 error) *MockStore_ListEntries_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStore_ListEntries_Call) RunAndReturn(run func(context.Context, db.ListEntriesParams) ([]db.Entry, error)) *MockStore_ListEntries_Call {
	_c.Call.Return(run)
	return _c
}

// ListTransfers provides a mock function with given fields: ctx, arg
func (_m *MockStore) ListTransfers(ctx context.Context, arg db.ListTransfersParams) ([]db.Transfer, error) {
	ret := _m.Called(ctx, arg)

	var r0 []db.Transfer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.ListTransfersParams) ([]db.Transfer, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.ListTransfersParams) []db.Transfer); ok {
		r0 = rf(ctx, arg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]db.Transfer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.ListTransfersParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStore_ListTransfers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListTransfers'
type MockStore_ListTransfers_Call struct {
	*mock.Call
}

// ListTransfers is a helper method to define mock.On call
//   - ctx context.Context
//   - arg db.ListTransfersParams
func (_e *MockStore_Expecter) ListTransfers(ctx interface{}, arg interface{}) *MockStore_ListTransfers_Call {
	return &MockStore_ListTransfers_Call{Call: _e.mock.On("ListTransfers", ctx, arg)}
}

func (_c *MockStore_ListTransfers_Call) Run(run func(ctx context.Context, arg db.ListTransfersParams)) *MockStore_ListTransfers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(db.ListTransfersParams))
	})
	return _c
}

func (_c *MockStore_ListTransfers_Call) Return(_a0 []db.Transfer, _a1 error) *MockStore_ListTransfers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStore_ListTransfers_Call) RunAndReturn(run func(context.Context, db.ListTransfersParams) ([]db.Transfer, error)) *MockStore_ListTransfers_Call {
	_c.Call.Return(run)
	return _c
}

// TransferTx provides a mock function with given fields: ctx, arg
func (_m *MockStore) TransferTx(ctx context.Context, arg db.TransferTXParams) (db.TransferTxResult, error) {
	ret := _m.Called(ctx, arg)

	var r0 db.TransferTxResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.TransferTXParams) (db.TransferTxResult, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.TransferTXParams) db.TransferTxResult); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(db.TransferTxResult)
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.TransferTXParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStore_TransferTx_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TransferTx'
type MockStore_TransferTx_Call struct {
	*mock.Call
}

// TransferTx is a helper method to define mock.On call
//   - ctx context.Context
//   - arg db.TransferTXParams
func (_e *MockStore_Expecter) TransferTx(ctx interface{}, arg interface{}) *MockStore_TransferTx_Call {
	return &MockStore_TransferTx_Call{Call: _e.mock.On("TransferTx", ctx, arg)}
}

func (_c *MockStore_TransferTx_Call) Run(run func(ctx context.Context, arg db.TransferTXParams)) *MockStore_TransferTx_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(db.TransferTXParams))
	})
	return _c
}

func (_c *MockStore_TransferTx_Call) Return(_a0 db.TransferTxResult, _a1 error) *MockStore_TransferTx_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStore_TransferTx_Call) RunAndReturn(run func(context.Context, db.TransferTXParams) (db.TransferTxResult, error)) *MockStore_TransferTx_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateAccount provides a mock function with given fields: ctx, arg
func (_m *MockStore) UpdateAccount(ctx context.Context, arg db.UpdateAccountParams) (db.Account, error) {
	ret := _m.Called(ctx, arg)

	var r0 db.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.UpdateAccountParams) (db.Account, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.UpdateAccountParams) db.Account); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(db.Account)
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.UpdateAccountParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStore_UpdateAccount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateAccount'
type MockStore_UpdateAccount_Call struct {
	*mock.Call
}

// UpdateAccount is a helper method to define mock.On call
//   - ctx context.Context
//   - arg db.UpdateAccountParams
func (_e *MockStore_Expecter) UpdateAccount(ctx interface{}, arg interface{}) *MockStore_UpdateAccount_Call {
	return &MockStore_UpdateAccount_Call{Call: _e.mock.On("UpdateAccount", ctx, arg)}
}

func (_c *MockStore_UpdateAccount_Call) Run(run func(ctx context.Context, arg db.UpdateAccountParams)) *MockStore_UpdateAccount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(db.UpdateAccountParams))
	})
	return _c
}

func (_c *MockStore_UpdateAccount_Call) Return(_a0 db.Account, _a1 error) *MockStore_UpdateAccount_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStore_UpdateAccount_Call) RunAndReturn(run func(context.Context, db.UpdateAccountParams) (db.Account, error)) *MockStore_UpdateAccount_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUser provides a mock function with given fields: ctx, arg
func (_m *MockStore) UpdateUser(ctx context.Context, arg db.UpdateUserParams) (db.User, error) {
	ret := _m.Called(ctx, arg)

	var r0 db.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.UpdateUserParams) (db.User, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.UpdateUserParams) db.User); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(db.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.UpdateUserParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStore_UpdateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUser'
type MockStore_UpdateUser_Call struct {
	*mock.Call
}

// UpdateUser is a helper method to define mock.On call
//   - ctx context.Context
//   - arg db.UpdateUserParams
func (_e *MockStore_Expecter) UpdateUser(ctx interface{}, arg interface{}) *MockStore_UpdateUser_Call {
	return &MockStore_UpdateUser_Call{Call: _e.mock.On("UpdateUser", ctx, arg)}
}

func (_c *MockStore_UpdateUser_Call) Run(run func(ctx context.Context, arg db.UpdateUserParams)) *MockStore_UpdateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(db.UpdateUserParams))
	})
	return _c
}

func (_c *MockStore_UpdateUser_Call) Return(_a0 db.User, _a1 error) *MockStore_UpdateUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStore_UpdateUser_Call) RunAndReturn(run func(context.Context, db.UpdateUserParams) (db.User, error)) *MockStore_UpdateUser_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockStore creates a new instance of MockStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockStore {
	mock := &MockStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
