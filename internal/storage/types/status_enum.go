// SPDX-FileCopyrightText: 2023 PK Lab AG <contact@pklab.io>
// SPDX-License-Identifier: MIT

// Code generated by go-enum DO NOT EDIT.
// Version: 0.5.7
// Revision: bf63e108589bbd2327b13ec2c5da532aad234029
// Build Date: 2023-07-25T23:27:55Z
// Built By: goreleaser

package types

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

const (
	// StatusSuccess is a Status of type success.
	StatusSuccess Status = "success"
	// StatusFailed is a Status of type failed.
	StatusFailed Status = "failed"
)

var ErrInvalidStatus = fmt.Errorf("not a valid Status, try [%s]", strings.Join(_StatusNames, ", "))

var _StatusNames = []string{
	string(StatusSuccess),
	string(StatusFailed),
}

// StatusNames returns a list of possible string values of Status.
func StatusNames() []string {
	tmp := make([]string, len(_StatusNames))
	copy(tmp, _StatusNames)
	return tmp
}

// StatusValues returns a list of the values for Status
func StatusValues() []Status {
	return []Status{
		StatusSuccess,
		StatusFailed,
	}
}

// String implements the Stringer interface.
func (x Status) String() string {
	return string(x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x Status) IsValid() bool {
	_, err := ParseStatus(string(x))
	return err == nil
}

var _StatusValue = map[string]Status{
	"success": StatusSuccess,
	"failed":  StatusFailed,
}

// ParseStatus attempts to convert a string to a Status.
func ParseStatus(name string) (Status, error) {
	if x, ok := _StatusValue[name]; ok {
		return x, nil
	}
	return Status(""), fmt.Errorf("%s is %w", name, ErrInvalidStatus)
}

// MarshalText implements the text marshaller method.
func (x Status) MarshalText() ([]byte, error) {
	return []byte(string(x)), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *Status) UnmarshalText(text []byte) error {
	tmp, err := ParseStatus(string(text))
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

var errStatusNilPtr = errors.New("value pointer is nil") // one per type for package clashes

// Scan implements the Scanner interface.
func (x *Status) Scan(value interface{}) (err error) {
	if value == nil {
		*x = Status("")
		return
	}

	// A wider range of scannable types.
	// driver.Value values at the top of the list for expediency
	switch v := value.(type) {
	case string:
		*x, err = ParseStatus(v)
	case []byte:
		*x, err = ParseStatus(string(v))
	case Status:
		*x = v
	case *Status:
		if v == nil {
			return errStatusNilPtr
		}
		*x = *v
	case *string:
		if v == nil {
			return errStatusNilPtr
		}
		*x, err = ParseStatus(*v)
	default:
		return errors.New("invalid type for Status")
	}

	return
}

// Value implements the driver Valuer interface.
func (x Status) Value() (driver.Value, error) {
	return x.String(), nil
}
