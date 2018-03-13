// Generated with goxc vgoxc-0.1.8 - rev 7e2e945f706bc13e7539c26efd1ec70bc280277e
package h004

import ds "github.com/openyard/ebics/2.5/xmldsig"
import w3c "github.com/openyard/ebics/2.5/w3c"
// complex type
type CertificateInfoType struct {
    
    X509Data ds.X509Data `xml:"http://www.w3.org/2000/09/xmldsig# X509Data"`
    
        Any []*w3c.Any
    }
