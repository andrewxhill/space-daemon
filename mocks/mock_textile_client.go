// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	config "github.com/FleekHQ/space-daemon/config"
	client "github.com/textileio/go-threads/api/client"

	context "context"

	crypto "github.com/libp2p/go-libp2p-core/crypto"

	domain "github.com/FleekHQ/space-daemon/core/space/domain"

	mock "github.com/stretchr/testify/mock"

	textile "github.com/FleekHQ/space-daemon/core/textile"

	thread "github.com/textileio/go-threads/core/thread"

	usersclient "github.com/textileio/textile/api/users/client"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// CreateBucket provides a mock function with given fields: ctx, bucketSlug
func (_m *Client) CreateBucket(ctx context.Context, bucketSlug string) (textile.Bucket, error) {
	ret := _m.Called(ctx, bucketSlug)

	var r0 textile.Bucket
	if rf, ok := ret.Get(0).(func(context.Context, string) textile.Bucket); ok {
		r0 = rf(ctx, bucketSlug)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(textile.Bucket)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, bucketSlug)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBucket provides a mock function with given fields: ctx, slug
func (_m *Client) GetBucket(ctx context.Context, slug string) (textile.Bucket, error) {
	ret := _m.Called(ctx, slug)

	var r0 textile.Bucket
	if rf, ok := ret.Get(0).(func(context.Context, string) textile.Bucket); ok {
		r0 = rf(ctx, slug)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(textile.Bucket)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, slug)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBucketContext provides a mock function with given fields: ctx, bucketSlug
func (_m *Client) GetBucketContext(ctx context.Context, bucketSlug string) (context.Context, *thread.ID, error) {
	ret := _m.Called(ctx, bucketSlug)

	var r0 context.Context
	if rf, ok := ret.Get(0).(func(context.Context, string) context.Context); ok {
		r0 = rf(ctx, bucketSlug)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	var r1 *thread.ID
	if rf, ok := ret.Get(1).(func(context.Context, string) *thread.ID); ok {
		r1 = rf(ctx, bucketSlug)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*thread.ID)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, bucketSlug)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetDefaultBucket provides a mock function with given fields: ctx
func (_m *Client) GetDefaultBucket(ctx context.Context) (textile.Bucket, error) {
	ret := _m.Called(ctx)

	var r0 textile.Bucket
	if rf, ok := ret.Get(0).(func(context.Context) textile.Bucket); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(textile.Bucket)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetThreadsConnection provides a mock function with given fields:
func (_m *Client) GetThreadsConnection() (*client.Client, error) {
	ret := _m.Called()

	var r0 *client.Client
	if rf, ok := ret.Get(0).(func() *client.Client); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.Client)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsRunning provides a mock function with given fields:
func (_m *Client) IsRunning() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// JoinBucket provides a mock function with given fields: ctx, slug, ti
func (_m *Client) JoinBucket(ctx context.Context, slug string, ti *domain.ThreadInfo) (bool, error) {
	ret := _m.Called(ctx, slug, ti)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string, *domain.ThreadInfo) bool); ok {
		r0 = rf(ctx, slug, ti)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *domain.ThreadInfo) error); ok {
		r1 = rf(ctx, slug, ti)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListBuckets provides a mock function with given fields: ctx
func (_m *Client) ListBuckets(ctx context.Context) ([]textile.Bucket, error) {
	ret := _m.Called(ctx)

	var r0 []textile.Bucket
	if rf, ok := ret.Get(0).(func(context.Context) []textile.Bucket); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]textile.Bucket)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendMessage provides a mock function with given fields: ctx, recipient, body
func (_m *Client) SendMessage(ctx context.Context, recipient crypto.PubKey, body []byte) (*usersclient.Message, error) {
	ret := _m.Called(ctx, recipient, body)

	var r0 *usersclient.Message
	if rf, ok := ret.Get(0).(func(context.Context, crypto.PubKey, []byte) *usersclient.Message); ok {
		r0 = rf(ctx, recipient, body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*usersclient.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, crypto.PubKey, []byte) error); ok {
		r1 = rf(ctx, recipient, body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ShareBucket provides a mock function with given fields: ctx, bucketSlug
func (_m *Client) ShareBucket(ctx context.Context, bucketSlug string) (*client.DBInfo, error) {
	ret := _m.Called(ctx, bucketSlug)

	var r0 *client.DBInfo
	if rf, ok := ret.Get(0).(func(context.Context, string) *client.DBInfo); ok {
		r0 = rf(ctx, bucketSlug)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.DBInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, bucketSlug)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Shutdown provides a mock function with given fields:
func (_m *Client) Shutdown() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Start provides a mock function with given fields: ctx, cfg
func (_m *Client) Start(ctx context.Context, cfg config.Config) error {
	ret := _m.Called(ctx, cfg)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, config.Config) error); ok {
		r0 = rf(ctx, cfg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ToggleBucketBackup provides a mock function with given fields: ctx, bucketSlug, bucketBackup
func (_m *Client) ToggleBucketBackup(ctx context.Context, bucketSlug string, bucketBackup bool) (bool, error) {
	ret := _m.Called(ctx, bucketSlug, bucketBackup)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string, bool) bool); ok {
		r0 = rf(ctx, bucketSlug, bucketBackup)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, bool) error); ok {
		r1 = rf(ctx, bucketSlug, bucketBackup)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WaitForReady provides a mock function with given fields:
func (_m *Client) WaitForReady() chan bool {
	ret := _m.Called()

	var r0 chan bool
	if rf, ok := ret.Get(0).(func() chan bool); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan bool)
		}
	}

	return r0
}
