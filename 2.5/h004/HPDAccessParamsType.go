// Generated with goxc vgoxc-0.1.8 - rev 7e2e945f706bc13e7539c26efd1ec70bc280277e
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"
// complex type
type HPDAccessParamsType struct {
    
    URL []HPDAccessParamsTypeURL `xml:"URL"`
    Institute w3c.NormalizedString `xml:"Institute"`
    HostID HostIDType `xml:"HostID,omitempty"`
    
        Any []*w3c.Any
    }
