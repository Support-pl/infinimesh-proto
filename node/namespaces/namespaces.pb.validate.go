// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: node/namespaces/namespaces.proto

package namespaces

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

// Validate checks the field values on Plugin with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Plugin) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Plugin with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in PluginMultiError, or nil if none found.
func (m *Plugin) ValidateAll() error {
	return m.validate(true)
}

func (m *Plugin) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Uuid

	// no validation rules for Vars

	if len(errors) > 0 {
		return PluginMultiError(errors)
	}

	return nil
}

// PluginMultiError is an error wrapping multiple validation errors returned by
// Plugin.ValidateAll() if the designated constraints aren't met.
type PluginMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PluginMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PluginMultiError) AllErrors() []error { return m }

// PluginValidationError is the validation error returned by Plugin.Validate if
// the designated constraints aren't met.
type PluginValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PluginValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PluginValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PluginValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PluginValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PluginValidationError) ErrorName() string { return "PluginValidationError" }

// Error satisfies the builtin error interface
func (e PluginValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPlugin.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PluginValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PluginValidationError{}

// Validate checks the field values on Namespace with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Namespace) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Namespace with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in NamespaceMultiError, or nil
// if none found.
func (m *Namespace) ValidateAll() error {
	return m.validate(true)
}

func (m *Namespace) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Uuid

	// no validation rules for Title

	if all {
		switch v := interface{}(m.GetConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, NamespaceValidationError{
					field:  "Config",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, NamespaceValidationError{
					field:  "Config",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NamespaceValidationError{
				field:  "Config",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.Access != nil {

		if all {
			switch v := interface{}(m.GetAccess()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, NamespaceValidationError{
						field:  "Access",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, NamespaceValidationError{
						field:  "Access",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetAccess()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return NamespaceValidationError{
					field:  "Access",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.Plugin != nil {

		if all {
			switch v := interface{}(m.GetPlugin()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, NamespaceValidationError{
						field:  "Plugin",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, NamespaceValidationError{
						field:  "Plugin",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetPlugin()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return NamespaceValidationError{
					field:  "Plugin",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return NamespaceMultiError(errors)
	}

	return nil
}

// NamespaceMultiError is an error wrapping multiple validation errors returned
// by Namespace.ValidateAll() if the designated constraints aren't met.
type NamespaceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NamespaceMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NamespaceMultiError) AllErrors() []error { return m }

// NamespaceValidationError is the validation error returned by
// Namespace.Validate if the designated constraints aren't met.
type NamespaceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NamespaceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NamespaceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NamespaceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NamespaceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NamespaceValidationError) ErrorName() string { return "NamespaceValidationError" }

// Error satisfies the builtin error interface
func (e NamespaceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNamespace.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NamespaceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NamespaceValidationError{}

// Validate checks the field values on Namespaces with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Namespaces) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Namespaces with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in NamespacesMultiError, or
// nil if none found.
func (m *Namespaces) ValidateAll() error {
	return m.validate(true)
}

func (m *Namespaces) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetNamespaces() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, NamespacesValidationError{
						field:  fmt.Sprintf("Namespaces[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, NamespacesValidationError{
						field:  fmt.Sprintf("Namespaces[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return NamespacesValidationError{
					field:  fmt.Sprintf("Namespaces[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return NamespacesMultiError(errors)
	}

	return nil
}

// NamespacesMultiError is an error wrapping multiple validation errors
// returned by Namespaces.ValidateAll() if the designated constraints aren't met.
type NamespacesMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NamespacesMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NamespacesMultiError) AllErrors() []error { return m }

// NamespacesValidationError is the validation error returned by
// Namespaces.Validate if the designated constraints aren't met.
type NamespacesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NamespacesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NamespacesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NamespacesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NamespacesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NamespacesValidationError) ErrorName() string { return "NamespacesValidationError" }

// Error satisfies the builtin error interface
func (e NamespacesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNamespaces.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NamespacesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NamespacesValidationError{}
