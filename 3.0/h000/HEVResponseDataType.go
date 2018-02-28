// Generated with goxc v0.1.7 - rev 2ae97d253f5eaa17bab360dad75945920dfceef4
package h000

import w3c "github.com/openyard/ebics/3.0/w3c"

// complex type
type HEVResponseDataType struct {
	SystemReturnCode SystemReturnCodeType             `xml:"SystemReturnCode"`
	VersionNumber    HEVResponseDataTypeVersionNumber `xml:"VersionNumber,omitempty"`

	Any []*w3c.Any
}
