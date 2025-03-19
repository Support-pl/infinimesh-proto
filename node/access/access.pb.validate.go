// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: node/access/access.proto

package access

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

// Validate checks the field values on Access with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Access) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Access with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in AccessMultiError, or nil if none found.
func (m *Access) ValidateAll() error {
	return m.validate(true)
}

func (m *Access) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Level

	// no validation rules for Role

	if m.Namespace != nil {
		// no validation rules for Namespace
	}

	if len(errors) > 0 {
		return AccessMultiError(errors)
	}

	return nil
}

// AccessMultiError is an error wrapping multiple validation errors returned by
// Access.ValidateAll() if the designated constraints aren't met.
type AccessMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AccessMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AccessMultiError) AllErrors() []error { return m }

// AccessValidationError is the validation error returned by Access.Validate if
// the designated constraints aren't met.
type AccessValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AccessValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AccessValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AccessValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AccessValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AccessValidationError) ErrorName() string { return "AccessValidationError" }

// Error satisfies the builtin error interface
func (e AccessValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAccess.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AccessValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AccessValidationError{}

// Validate checks the field values on Node with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Node) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Node with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in NodeMultiError, or nil if none found.
func (m *Node) ValidateAll() error {
	return m.validate(true)
}

func (m *Node) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Node

	// no validation rules for Parent

	// no validation rules for Edge

	if m.Access != nil {

		if all {
			switch v := interface{}(m.GetAccess()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, NodeValidationError{
						field:  "Access",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, NodeValidationError{
						field:  "Access",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetAccess()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return NodeValidationError{
					field:  "Access",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return NodeMultiError(errors)
	}

	return nil
}

// NodeMultiError is an error wrapping multiple validation errors returned by
// Node.ValidateAll() if the designated constraints aren't met.
type NodeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NodeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NodeMultiError) AllErrors() []error { return m }

// NodeValidationError is the validation error returned by Node.Validate if the
// designated constraints aren't met.
type NodeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NodeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NodeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NodeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NodeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NodeValidationError) ErrorName() string { return "NodeValidationError" }

// Error satisfies the builtin error interface
func (e NodeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNode.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NodeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NodeValidationError{}

// Validate checks the field values on Nodes with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Nodes) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Nodes with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in NodesMultiError, or nil if none found.
func (m *Nodes) ValidateAll() error {
	return m.validate(true)
}

func (m *Nodes) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetNodes() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, NodesValidationError{
						field:  fmt.Sprintf("Nodes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, NodesValidationError{
						field:  fmt.Sprintf("Nodes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return NodesValidationError{
					field:  fmt.Sprintf("Nodes[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return NodesMultiError(errors)
	}

	return nil
}

// NodesMultiError is an error wrapping multiple validation errors returned by
// Nodes.ValidateAll() if the designated constraints aren't met.
type NodesMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NodesMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NodesMultiError) AllErrors() []error { return m }

// NodesValidationError is the validation error returned by Nodes.Validate if
// the designated constraints aren't met.
type NodesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NodesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NodesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NodesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NodesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NodesValidationError) ErrorName() string { return "NodesValidationError" }

// Error satisfies the builtin error interface
func (e NodesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNodes.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NodesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NodesValidationError{}
