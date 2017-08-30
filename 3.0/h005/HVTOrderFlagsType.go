// Generated with goxc v0.1.3 - rev 0e63342ac0a4d5f52582ea0065a462e700069839
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// complex type
type HVTOrderFlagsType struct {
	CompleteOrderData w3c.Boolean            `xml:"completeOrderData,attr"`
	FetchLimit        w3c.NonNegativeInteger `xml:"fetchLimit,attr"`
	FetchOffset       w3c.NonNegativeInteger `xml:"fetchOffset,attr"`
}
