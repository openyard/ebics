// Generated with goxc vgoxc-0.1.8 - rev 7e2e945f706bc13e7539c26efd1ec70bc280277e
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// complex type
type HVTOrderFlagsType struct {
	CompleteOrderData w3c.Boolean            `xml:"completeOrderData,attr"`
	FetchLimit        w3c.NonNegativeInteger `xml:"fetchLimit,attr"`
	FetchOffset       w3c.NonNegativeInteger `xml:"fetchOffset,attr"`
}
