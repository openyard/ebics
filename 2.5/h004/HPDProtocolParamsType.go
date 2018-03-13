// Generated with goxc vgoxc-0.1.8 - rev 7e2e945f706bc13e7539c26efd1ec70bc280277e
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"
// complex type
type HPDProtocolParamsType struct {
    
    Version HPDVersionType `xml:"Version"`
    Recovery HPDProtocolParamsTypeRecovery `xml:"Recovery,omitempty"`
    PreValidation HPDProtocolParamsTypePreValidation `xml:"PreValidation,omitempty"`
    X509Data HPDProtocolParamsTypeX509Data `xml:"X509Data,omitempty"`
    ClientDataDownload HPDProtocolParamsTypeClientDataDownload `xml:"ClientDataDownload,omitempty"`
    DownloadableOrderData HPDProtocolParamsTypeDownloadableOrderData `xml:"DownloadableOrderData,omitempty"`
    
        Any []*w3c.Any
    }
