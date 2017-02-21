// Generated with goxc v0.1.2 - rev bbe25b23ba83e8f8464e681ca3e514eee51fd2ba
package xmldsig

import w3c "github.com/openyard/ebics/2.5/w3c"

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
