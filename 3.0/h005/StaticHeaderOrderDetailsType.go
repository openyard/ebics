// Generated with goxc v0.1.7 - rev 2ae97d253f5eaa17bab360dad75945920dfceef4
package h005

// complex type
type StaticHeaderOrderDetailsType struct {
	AdminOrderType StaticHeaderOrderDetailsTypeAdminOrderType `xml:"AdminOrderType"`
	OrderID        OrderIDType                                `xml:"OrderID,omitempty"`
	OrderParams    OrderParams                                `xml:"OrderParams"`
}
