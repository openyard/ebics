// Generated with goxc v0.1.3 - rev 0e63342ac0a4d5f52582ea0065a462e700069839
package h005

// complex type
type StaticHeaderOrderDetailsType struct {
	AdminOrderType StaticHeaderOrderDetailsTypeAdminOrderType `xml:"AdminOrderType"`
	OrderID        OrderIDType                                `xml:"OrderID,omitempty"`
	OrderParams    OrderParams                                `xml:"OrderParams"`
}
