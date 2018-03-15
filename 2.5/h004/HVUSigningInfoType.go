// Generated with goxc vgoxc-0.1.9 - rev 260439f4ef82f3f152002242cdec0bb97e6118c3
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex type
type HVUSigningInfoType struct {
	ReadyToBeSigned *w3c.Boolean            `xml:"readyToBeSigned,attr"`
	NumSigRequired  *w3c.PositiveInteger    `xml:"NumSigRequired,attr"`
	NumSigDone      *w3c.NonNegativeInteger `xml:"NumSigDone,attr"`
}
