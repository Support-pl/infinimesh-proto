// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: shadow/shadow.proto

package shadow

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on State with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *State) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on State with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in StateMultiError, or nil if none found.
func (m *State) ValidateAll() error {
	return m.validate(true)
}

func (m *State) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetTimestamp()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, StateValidationError{
					field:  "Timestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, StateValidationError{
					field:  "Timestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimestamp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StateValidationError{
				field:  "Timestamp",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, StateValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, StateValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StateValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return StateMultiError(errors)
	}

	return nil
}

// StateMultiError is an error wrapping multiple validation errors returned by
// State.ValidateAll() if the designated constraints aren't met.
type StateMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StateMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StateMultiError) AllErrors() []error { return m }

// StateValidationError is the validation error returned by State.Validate if
// the designated constraints aren't met.
type StateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StateValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StateValidationError) ErrorName() string { return "StateValidationError" }

// Error satisfies the builtin error interface
func (e StateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sState.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StateValidationError{}

// Validate checks the field values on ConnectionState with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ConnectionState) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ConnectionState with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ConnectionStateMultiError, or nil if none found.
func (m *ConnectionState) ValidateAll() error {
	return m.validate(true)
}

func (m *ConnectionState) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetTimestamp()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ConnectionStateValidationError{
					field:  "Timestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ConnectionStateValidationError{
					field:  "Timestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimestamp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ConnectionStateValidationError{
				field:  "Timestamp",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Connected

	// no validation rules for ClientId

	if len(errors) > 0 {
		return ConnectionStateMultiError(errors)
	}

	return nil
}

// ConnectionStateMultiError is an error wrapping multiple validation errors
// returned by ConnectionState.ValidateAll() if the designated constraints
// aren't met.
type ConnectionStateMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ConnectionStateMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ConnectionStateMultiError) AllErrors() []error { return m }

// ConnectionStateValidationError is the validation error returned by
// ConnectionState.Validate if the designated constraints aren't met.
type ConnectionStateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConnectionStateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConnectionStateValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConnectionStateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConnectionStateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConnectionStateValidationError) ErrorName() string { return "ConnectionStateValidationError" }

// Error satisfies the builtin error interface
func (e ConnectionStateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConnectionState.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConnectionStateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConnectionStateValidationError{}

// Validate checks the field values on Shadow with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Shadow) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Shadow with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ShadowMultiError, or nil if none found.
func (m *Shadow) ValidateAll() error {
	return m.validate(true)
}

func (m *Shadow) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Device

	if all {
		switch v := interface{}(m.GetReported()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ShadowValidationError{
					field:  "Reported",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ShadowValidationError{
					field:  "Reported",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetReported()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ShadowValidationError{
				field:  "Reported",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetDesired()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ShadowValidationError{
					field:  "Desired",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ShadowValidationError{
					field:  "Desired",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDesired()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ShadowValidationError{
				field:  "Desired",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetConnection()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ShadowValidationError{
					field:  "Connection",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ShadowValidationError{
					field:  "Connection",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetConnection()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ShadowValidationError{
				field:  "Connection",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ShadowMultiError(errors)
	}

	return nil
}

// ShadowMultiError is an error wrapping multiple validation errors returned by
// Shadow.ValidateAll() if the designated constraints aren't met.
type ShadowMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ShadowMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ShadowMultiError) AllErrors() []error { return m }

// ShadowValidationError is the validation error returned by Shadow.Validate if
// the designated constraints aren't met.
type ShadowValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ShadowValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ShadowValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ShadowValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ShadowValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ShadowValidationError) ErrorName() string { return "ShadowValidationError" }

// Error satisfies the builtin error interface
func (e ShadowValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sShadow.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ShadowValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ShadowValidationError{}

// Validate checks the field values on GetRequest with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetRequestMultiError, or
// nil if none found.
func (m *GetRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetRequestMultiError(errors)
	}

	return nil
}

// GetRequestMultiError is an error wrapping multiple validation errors
// returned by GetRequest.ValidateAll() if the designated constraints aren't met.
type GetRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetRequestMultiError) AllErrors() []error { return m }

// GetRequestValidationError is the validation error returned by
// GetRequest.Validate if the designated constraints aren't met.
type GetRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetRequestValidationError) ErrorName() string { return "GetRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetRequestValidationError{}

// Validate checks the field values on GetResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetResponseMultiError, or
// nil if none found.
func (m *GetResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetShadows() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetResponseValidationError{
						field:  fmt.Sprintf("Shadows[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetResponseValidationError{
						field:  fmt.Sprintf("Shadows[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetResponseValidationError{
					field:  fmt.Sprintf("Shadows[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetResponseMultiError(errors)
	}

	return nil
}

// GetResponseMultiError is an error wrapping multiple validation errors
// returned by GetResponse.ValidateAll() if the designated constraints aren't met.
type GetResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetResponseMultiError) AllErrors() []error { return m }

// GetResponseValidationError is the validation error returned by
// GetResponse.Validate if the designated constraints aren't met.
type GetResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetResponseValidationError) ErrorName() string { return "GetResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetResponseValidationError{}

// Validate checks the field values on RemoveRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RemoveRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RemoveRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RemoveRequestMultiError, or
// nil if none found.
func (m *RemoveRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *RemoveRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Device

	// no validation rules for StateKey

	// no validation rules for Key

	if len(errors) > 0 {
		return RemoveRequestMultiError(errors)
	}

	return nil
}

// RemoveRequestMultiError is an error wrapping multiple validation errors
// returned by RemoveRequest.ValidateAll() if the designated constraints
// aren't met.
type RemoveRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RemoveRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RemoveRequestMultiError) AllErrors() []error { return m }

// RemoveRequestValidationError is the validation error returned by
// RemoveRequest.Validate if the designated constraints aren't met.
type RemoveRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveRequestValidationError) ErrorName() string { return "RemoveRequestValidationError" }

// Error satisfies the builtin error interface
func (e RemoveRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveRequestValidationError{}

// Validate checks the field values on StreamShadowRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *StreamShadowRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StreamShadowRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StreamShadowRequestMultiError, or nil if none found.
func (m *StreamShadowRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *StreamShadowRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for OnlyDelta

	// no validation rules for Sync

	if len(errors) > 0 {
		return StreamShadowRequestMultiError(errors)
	}

	return nil
}

// StreamShadowRequestMultiError is an error wrapping multiple validation
// errors returned by StreamShadowRequest.ValidateAll() if the designated
// constraints aren't met.
type StreamShadowRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StreamShadowRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StreamShadowRequestMultiError) AllErrors() []error { return m }

// StreamShadowRequestValidationError is the validation error returned by
// StreamShadowRequest.Validate if the designated constraints aren't met.
type StreamShadowRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StreamShadowRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StreamShadowRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StreamShadowRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StreamShadowRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StreamShadowRequestValidationError) ErrorName() string {
	return "StreamShadowRequestValidationError"
}

// Error satisfies the builtin error interface
func (e StreamShadowRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStreamShadowRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StreamShadowRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StreamShadowRequestValidationError{}
