// Generated with goxc v0.1.3 - rev 0e63342ac0a4d5f52582ea0065a462e700069839
package h000

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex type
type HEVRequestDataType struct {
	HostID HostIDType `xml:"HostID"`

	Any []*w3c.Any
}