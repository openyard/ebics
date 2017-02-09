// Generated with goxc v0.1.1 - rev bae2cf01854d664b985cae6986076979716034c7
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex type
type HVUSigningInfoType struct {
	ReadyToBeSigned w3c.Boolean            `xml:"readyToBeSigned,attr"`
	NumSigRequired  w3c.PositiveInteger    `xml:"NumSigRequired,attr"`
	NumSigDone      w3c.NonNegativeInteger `xml:"NumSigDone,attr"`
}
