// Generated with goxc vgoxc-0.1.8 - rev 7e2e945f706bc13e7539c26efd1ec70bc280277e
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"
// complex type
type HVDResponseOrderDataType struct {
    
    DataDigest DataDigestType `xml:"DataDigest"`
    DisplayFile w3c.Base64Binary `xml:"DisplayFile"`
    OrderDataAvailable w3c.Boolean `xml:"OrderDataAvailable"`
    OrderDataSize w3c.PositiveInteger `xml:"OrderDataSize"`
    OrderDetailsAvailable w3c.Boolean `xml:"OrderDetailsAvailable"`
    BankSignature SignatureType `xml:"BankSignature,omitempty"`
    SignerInfo []SignerInfoType `xml:"SignerInfo,omitempty"`
    
        Any []*w3c.Any
    }
