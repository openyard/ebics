// Generated with goxc vgoxc-0.1.8 - rev 7e2e945f706bc13e7539c26efd1ec70bc280277e
package h005

// complex type
type StaticHeaderOrderDetailsType struct {
	AdminOrderType StaticHeaderOrderDetailsTypeAdminOrderType `xml:"AdminOrderType"`
	OrderID        OrderIDType                                `xml:"OrderID,omitempty"`
	OrderParams    OrderParams                                `xml:"OrderParams"`
}
