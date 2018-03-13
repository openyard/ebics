// Generated with goxc vgoxc-0.1.8 - rev 7e2e945f706bc13e7539c26efd1ec70bc280277e
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"
// complex type
type HVUOrderDetailsType struct {
    
    OrderType OrderTBaseType `xml:"OrderType"`
    FileFormat FileFormatType `xml:"FileFormat,omitempty"`
    OrderID OrderIDType `xml:"OrderID"`
    OrderDataSize w3c.PositiveInteger `xml:"OrderDataSize"`
    SigningInfo HVUSigningInfoType `xml:"SigningInfo"`
    SignerInfo []SignerInfoType `xml:"SignerInfo,omitempty"`
    OriginatorInfo HVUOriginatorInfoType `xml:"OriginatorInfo"`
    
        Any []*w3c.Any
    }
