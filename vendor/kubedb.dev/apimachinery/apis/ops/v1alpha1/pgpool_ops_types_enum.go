// Code generated by go-enum DO NOT EDIT.
// Version:
// Revision:
// Build Date:
// Built By:

package v1alpha1

import (
	"fmt"
	"strings"
)

const (
	// PgpoolOpsRequestTypeUpdateVersion is a PgpoolOpsRequestType of type UpdateVersion.
	PgpoolOpsRequestTypeUpdateVersion PgpoolOpsRequestType = "UpdateVersion"
	// PgpoolOpsRequestTypeRestart is a PgpoolOpsRequestType of type Restart.
	PgpoolOpsRequestTypeRestart PgpoolOpsRequestType = "Restart"
	// PgpoolOpsRequestTypeReconfigure is a PgpoolOpsRequestType of type Reconfigure.
	PgpoolOpsRequestTypeReconfigure PgpoolOpsRequestType = "Reconfigure"
	// PgpoolOpsRequestTypeVerticalScaling is a PgpoolOpsRequestType of type VerticalScaling.
	PgpoolOpsRequestTypeVerticalScaling PgpoolOpsRequestType = "VerticalScaling"
	// PgpoolOpsRequestTypeHorizontalScaling is a PgpoolOpsRequestType of type HorizontalScaling.
	PgpoolOpsRequestTypeHorizontalScaling PgpoolOpsRequestType = "HorizontalScaling"
	// PgpoolOpsRequestTypeReconfigureTLS is a PgpoolOpsRequestType of type ReconfigureTLS.
	PgpoolOpsRequestTypeReconfigureTLS PgpoolOpsRequestType = "ReconfigureTLS"
)

var ErrInvalidPgpoolOpsRequestType = fmt.Errorf("not a valid PgpoolOpsRequestType, try [%s]", strings.Join(_PgpoolOpsRequestTypeNames, ", "))

var _PgpoolOpsRequestTypeNames = []string{
	string(PgpoolOpsRequestTypeUpdateVersion),
	string(PgpoolOpsRequestTypeRestart),
	string(PgpoolOpsRequestTypeReconfigure),
	string(PgpoolOpsRequestTypeVerticalScaling),
	string(PgpoolOpsRequestTypeHorizontalScaling),
	string(PgpoolOpsRequestTypeReconfigureTLS),
}

// PgpoolOpsRequestTypeNames returns a list of possible string values of PgpoolOpsRequestType.
func PgpoolOpsRequestTypeNames() []string {
	tmp := make([]string, len(_PgpoolOpsRequestTypeNames))
	copy(tmp, _PgpoolOpsRequestTypeNames)
	return tmp
}

// PgpoolOpsRequestTypeValues returns a list of the values for PgpoolOpsRequestType
func PgpoolOpsRequestTypeValues() []PgpoolOpsRequestType {
	return []PgpoolOpsRequestType{
		PgpoolOpsRequestTypeUpdateVersion,
		PgpoolOpsRequestTypeRestart,
		PgpoolOpsRequestTypeReconfigure,
		PgpoolOpsRequestTypeVerticalScaling,
		PgpoolOpsRequestTypeHorizontalScaling,
		PgpoolOpsRequestTypeReconfigureTLS,
	}
}

// String implements the Stringer interface.
func (x PgpoolOpsRequestType) String() string {
	return string(x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x PgpoolOpsRequestType) IsValid() bool {
	_, err := ParsePgpoolOpsRequestType(string(x))
	return err == nil
}

var _PgpoolOpsRequestTypeValue = map[string]PgpoolOpsRequestType{
	"UpdateVersion":     PgpoolOpsRequestTypeUpdateVersion,
	"Restart":           PgpoolOpsRequestTypeRestart,
	"Reconfigure":       PgpoolOpsRequestTypeReconfigure,
	"VerticalScaling":   PgpoolOpsRequestTypeVerticalScaling,
	"HorizontalScaling": PgpoolOpsRequestTypeHorizontalScaling,
	"ReconfigureTLS":    PgpoolOpsRequestTypeReconfigureTLS,
}

// ParsePgpoolOpsRequestType attempts to convert a string to a PgpoolOpsRequestType.
func ParsePgpoolOpsRequestType(name string) (PgpoolOpsRequestType, error) {
	if x, ok := _PgpoolOpsRequestTypeValue[name]; ok {
		return x, nil
	}
	return PgpoolOpsRequestType(""), fmt.Errorf("%s is %w", name, ErrInvalidPgpoolOpsRequestType)
}

// MustParsePgpoolOpsRequestType converts a string to a PgpoolOpsRequestType, and panics if is not valid.
func MustParsePgpoolOpsRequestType(name string) PgpoolOpsRequestType {
	val, err := ParsePgpoolOpsRequestType(name)
	if err != nil {
		panic(err)
	}
	return val
}
