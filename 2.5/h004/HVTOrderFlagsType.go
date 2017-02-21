// Generated with goxc v0.1.2 - rev bbe25b23ba83e8f8464e681ca3e514eee51fd2ba
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex type
type HVTOrderFlagsType struct {
	CompleteOrderData w3c.Boolean            `xml:"completeOrderData,attr"`
	FetchLimit        w3c.NonNegativeInteger `xml:"fetchLimit,attr"`
	FetchOffset       w3c.NonNegativeInteger `xml:"fetchOffset,attr"`
}
