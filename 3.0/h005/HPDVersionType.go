// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexType
type HPDVersionType struct {
	Protocol       *string `xml:"Protocol"`
	Authentication *string `xml:"Authentication"`
	Encryption     *string `xml:"Encryption"`
	Signature      *string `xml:"Signature"`

	Any []*w3c.Any
}

func NewHPDVersionType() *HPDVersionType {
	return new(HPDVersionType)
}

func (me *HPDVersionType) SetProtocol(value *string) {
	me.Protocol = value
}

func (me *HPDVersionType) SetAuthentication(value *string) {
	me.Authentication = value
}

func (me *HPDVersionType) SetEncryption(value *string) {
	me.Encryption = value
}

func (me *HPDVersionType) SetSignature(value *string) {
	me.Signature = value
}
