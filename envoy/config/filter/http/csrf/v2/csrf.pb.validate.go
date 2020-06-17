// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/filter/http/csrf/v2/csrf.proto

package envoy_config_filter_http_csrf_v2

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _csrf_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on CsrfPolicy with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *CsrfPolicy) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetFilterEnabled() == nil {
		return CsrfPolicyValidationError{
			field:  "FilterEnabled",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetFilterEnabled()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CsrfPolicyValidationError{
				field:  "FilterEnabled",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetShadowEnabled()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CsrfPolicyValidationError{
				field:  "ShadowEnabled",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetAdditionalOrigins() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CsrfPolicyValidationError{
					field:  fmt.Sprintf("AdditionalOrigins[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// CsrfPolicyValidationError is the validation error returned by
// CsrfPolicy.Validate if the designated constraints aren't met.
type CsrfPolicyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CsrfPolicyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CsrfPolicyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CsrfPolicyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CsrfPolicyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CsrfPolicyValidationError) ErrorName() string { return "CsrfPolicyValidationError" }

// Error satisfies the builtin error interface
func (e CsrfPolicyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCsrfPolicy.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CsrfPolicyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CsrfPolicyValidationError{}
