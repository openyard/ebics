// Generated with goxc v0.1.1 - rev bae2cf01854d664b985cae6986076979716034c7
package h004

// complex type
type StaticHeaderOrderDetailsType struct {
	OrderType      StaticHeaderOrderDetailsTypeOrderType `xml:"OrderType"`
	OrderID        OrderIDType                           `xml:"OrderID,omitempty"`
	OrderAttribute OrderAttributeType                    `xml:"OrderAttribute"`
	OrderParams    OrderParams                           `xml:"OrderParams"`
}
