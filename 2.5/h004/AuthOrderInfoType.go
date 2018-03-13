// Generated with goxc vgoxc-0.1.8 - rev 7e2e945f706bc13e7539c26efd1ec70bc280277e
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"
// complex type
type AuthOrderInfoType struct {
    
    OrderType OrderTBaseType `xml:"OrderType"`
    FileFormat FileFormatType `xml:"FileFormat,omitempty"`
    TransferType TransferType `xml:"TransferType"`
    OrderFormat OrderFormatType `xml:"OrderFormat,omitempty"`
    Description OrderDescriptionType `xml:"Description"`
    NumSigRequired w3c.NonNegativeInteger `xml:"NumSigRequired,omitempty"`
    
        Any []*w3c.Any
    }
