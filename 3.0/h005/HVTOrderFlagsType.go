// Generated with goxc vgoxc-0.1.9 - rev 260439f4ef82f3f152002242cdec0bb97e6118c3
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// complex type
type HVTOrderFlagsType struct {
	CompleteOrderData *w3c.Boolean            `xml:"completeOrderData,attr"`
	FetchLimit        *w3c.NonNegativeInteger `xml:"fetchLimit,attr"`
	FetchOffset       *w3c.NonNegativeInteger `xml:"fetchOffset,attr"`
}
