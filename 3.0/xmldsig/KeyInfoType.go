// Generated with goxc vgoxc-0.1.8 - rev 7e2e945f706bc13e7539c26efd1ec70bc280277e
package xmldsig

import w3c "github.com/openyard/ebics/3.0/w3c"

// complex type
type KeyInfoType struct {
	Id              w3c.ID          `xml:"Id,attr,omitempty"`
	KeyName         KeyName         `xml:"KeyName"`
	KeyValue        KeyValue        `xml:"KeyValue"`
	RetrievalMethod RetrievalMethod `xml:"RetrievalMethod"`
	X509Data        X509Data        `xml:"X509Data"`
	PGPData         PGPData         `xml:"PGPData"`
	SPKIData        SPKIData        `xml:"SPKIData"`
	MgmtData        MgmtData        `xml:"MgmtData"`

	Any []*w3c.Any
}
