// Code generated by mockery v2.50.4. DO NOT EDIT.

package spotify

import (
	spotify "github.com/nogurenn/practice/golang/music-api/internal/spotify"
	mock "github.com/stretchr/testify/mock"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

type Client_Expecter struct {
	mock *mock.Mock
}

func (_m *Client) EXPECT() *Client_Expecter {
	return &Client_Expecter{mock: &_m.Mock}
}

// GetAccessToken provides a mock function with no fields
func (_m *Client) GetAccessToken() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAccessToken")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Client_GetAccessToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAccessToken'
type Client_GetAccessToken_Call struct {
	*mock.Call
}

// GetAccessToken is a helper method to define mock.On call
func (_e *Client_Expecter) GetAccessToken() *Client_GetAccessToken_Call {
	return &Client_GetAccessToken_Call{Call: _e.mock.On("GetAccessToken")}
}

func (_c *Client_GetAccessToken_Call) Run(run func()) *Client_GetAccessToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Client_GetAccessToken_Call) Return(_a0 error) *Client_GetAccessToken_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Client_GetAccessToken_Call) RunAndReturn(run func() error) *Client_GetAccessToken_Call {
	_c.Call.Return(run)
	return _c
}

// GetArtistBySpotifyArtistID provides a mock function with given fields: artistID
func (_m *Client) GetArtistBySpotifyArtistID(artistID string) (*spotify.Artist, error) {
	ret := _m.Called(artistID)

	if len(ret) == 0 {
		panic("no return value specified for GetArtistBySpotifyArtistID")
	}

	var r0 *spotify.Artist
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*spotify.Artist, error)); ok {
		return rf(artistID)
	}
	if rf, ok := ret.Get(0).(func(string) *spotify.Artist); ok {
		r0 = rf(artistID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*spotify.Artist)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(artistID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_GetArtistBySpotifyArtistID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetArtistBySpotifyArtistID'
type Client_GetArtistBySpotifyArtistID_Call struct {
	*mock.Call
}

// GetArtistBySpotifyArtistID is a helper method to define mock.On call
//   - artistID string
func (_e *Client_Expecter) GetArtistBySpotifyArtistID(artistID interface{}) *Client_GetArtistBySpotifyArtistID_Call {
	return &Client_GetArtistBySpotifyArtistID_Call{Call: _e.mock.On("GetArtistBySpotifyArtistID", artistID)}
}

func (_c *Client_GetArtistBySpotifyArtistID_Call) Run(run func(artistID string)) *Client_GetArtistBySpotifyArtistID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Client_GetArtistBySpotifyArtistID_Call) Return(_a0 *spotify.Artist, _a1 error) *Client_GetArtistBySpotifyArtistID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_GetArtistBySpotifyArtistID_Call) RunAndReturn(run func(string) (*spotify.Artist, error)) *Client_GetArtistBySpotifyArtistID_Call {
	_c.Call.Return(run)
	return _c
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
