// Generated with goxc vgoxc-0.1.8 - rev 7e2e945f706bc13e7539c26efd1ec70bc280277e
package h004

// complex type
type StaticHeaderOrderDetailsType struct {
    
    OrderType StaticHeaderOrderDetailsTypeOrderType `xml:"OrderType"`
    OrderID OrderIDType `xml:"OrderID,omitempty"`
    OrderAttribute OrderAttributeType `xml:"OrderAttribute"`
    OrderParams OrderParams `xml:"OrderParams"`
    }
