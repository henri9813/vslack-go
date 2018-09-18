// Code generated by mockery v1.0.0. DO NOT EDIT.
package vslack

import mock "github.com/stretchr/testify/mock"

// MockAttachmentInterface is an autogenerated mock type for the AttachmentInterface type
type MockAttachmentInterface struct {
	mock.Mock
}

// SetColor provides a mock function with given fields: c
func (_m *MockAttachmentInterface) SetColor(c string) Attachment {
	ret := _m.Called(c)

	var r0 Attachment
	if rf, ok := ret.Get(0).(func(string) Attachment); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Get(0).(Attachment)
	}

	return r0
}

// SetFields provides a mock function with given fields: f
func (_m *MockAttachmentInterface) SetFields(f ...Field) Attachment {
	_va := make([]interface{}, len(f))
	for _i := range f {
		_va[_i] = f[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 Attachment
	if rf, ok := ret.Get(0).(func(...Field) Attachment); ok {
		r0 = rf(f...)
	} else {
		r0 = ret.Get(0).(Attachment)
	}

	return r0
}

// SetMarkdown provides a mock function with given fields: m
func (_m *MockAttachmentInterface) SetMarkdown(m bool) Attachment {
	ret := _m.Called(m)

	var r0 Attachment
	if rf, ok := ret.Get(0).(func(bool) Attachment); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Get(0).(Attachment)
	}

	return r0
}

// SetMarkdownIn provides a mock function with given fields: opts
func (_m *MockAttachmentInterface) SetMarkdownIn(opts ...MarkdownOption) Attachment {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 Attachment
	if rf, ok := ret.Get(0).(func(...MarkdownOption) Attachment); ok {
		r0 = rf(opts...)
	} else {
		r0 = ret.Get(0).(Attachment)
	}

	return r0
}

// SetText provides a mock function with given fields: t
func (_m *MockAttachmentInterface) SetText(t string) Attachment {
	ret := _m.Called(t)

	var r0 Attachment
	if rf, ok := ret.Get(0).(func(string) Attachment); ok {
		r0 = rf(t)
	} else {
		r0 = ret.Get(0).(Attachment)
	}

	return r0
}

// SetTitle provides a mock function with given fields: t
func (_m *MockAttachmentInterface) SetTitle(t string) Attachment {
	ret := _m.Called(t)

	var r0 Attachment
	if rf, ok := ret.Get(0).(func(string) Attachment); ok {
		r0 = rf(t)
	} else {
		r0 = ret.Get(0).(Attachment)
	}

	return r0
}
