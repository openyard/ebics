// Generated with goxc vgoxc-0.1.8 - rev 7e2e945f706bc13e7539c26efd1ec70bc280277e
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"
// complex type
type TransferReceiptRequestType struct {
    
    ReceiptCode ReceiptCodeType `xml:"ReceiptCode"`
    
        Any []*w3c.Any
    }
