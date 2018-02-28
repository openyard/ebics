// Generated with goxc v0.1.7 - rev 2ae97d253f5eaa17bab360dad75945920dfceef4
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// complex type
type HVUSigningInfoType struct {
	ReadyToBeSigned w3c.Boolean            `xml:"readyToBeSigned,attr"`
	NumSigRequired  w3c.PositiveInteger    `xml:"NumSigRequired,attr"`
	NumSigDone      w3c.NonNegativeInteger `xml:"NumSigDone,attr"`
}
