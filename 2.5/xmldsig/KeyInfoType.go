// Generated with goxc v0.1.7 - rev 2ae97d253f5eaa17bab360dad75945920dfceef4
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
