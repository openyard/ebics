// Generated with goxc vgoxc-0.1.8 - rev 7e2e945f706bc13e7539c26efd1ec70bc280277e
package h000

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex type
type HEVRequestDataType struct {
	HostID HostIDType `xml:"HostID"`

	Any []*w3c.Any
}
