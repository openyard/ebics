// Generated with goxc v0.1.7 - rev 2ae97d253f5eaa17bab360dad75945920dfceef4
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex element
type HVTOrderInfoTypeDescription struct {
	Value w3c.String                      `xml:",chardata"`
	Type  HVTOrderInfoTypeDescriptionType `xml:"Type,attr"`
}
