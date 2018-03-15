// Generated with goxc vgoxc-0.1.9 - rev 260439f4ef82f3f152002242cdec0bb97e6118c3
package s002

import w3c "github.com/openyard/ebics/3.0/w3c"

// complex type
type UserSignatureDataSigBookType struct {
	OrderSignatureData []*OrderSignatureDataType `xml:"OrderSignatureData"`

	Any []*w3c.Any
}
