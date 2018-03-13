// Generated with goxc vgoxc-0.1.8 - rev 7e2e945f706bc13e7539c26efd1ec70bc280277e
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// complex type
type BankInfoType struct {
	HostID    HostIDType `xml:"HostID"`
	Parameter Parameter  `xml:"Parameter,omitempty"`

	Any []*w3c.Any
}
