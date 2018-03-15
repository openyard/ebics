// Generated with goxc vgoxc-0.1.9 - rev 260439f4ef82f3f152002242cdec0bb97e6118c3
package xmldsig

import w3c "github.com/openyard/ebics/3.0/w3c"

// complex type
type KeyInfoType struct {
	Id              *w3c.ID          `xml:"Id,attr,omitempty"`
	KeyName         *KeyName         `xml:"KeyName"`
	KeyValue        *KeyValue        `xml:"KeyValue"`
	RetrievalMethod *RetrievalMethod `xml:"RetrievalMethod"`
	X509Data        *X509Data        `xml:"X509Data"`
	PGPData         *PGPData         `xml:"PGPData"`
	SPKIData        *SPKIData        `xml:"SPKIData"`
	MgmtData        *MgmtData        `xml:"MgmtData"`

	Any []*w3c.Any
}
