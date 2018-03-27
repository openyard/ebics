// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

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

func (me *HPDVersionType) AddProtocol() *string {
	me.Protocol = new(string)
	return me.Protocol
}

func (me *HPDVersionType) SetAuthentication(value *string) {
	me.Authentication = value
}

func (me *HPDVersionType) AddAuthentication() *string {
	me.Authentication = new(string)
	return me.Authentication
}

func (me *HPDVersionType) SetEncryption(value *string) {
	me.Encryption = value
}

func (me *HPDVersionType) AddEncryption() *string {
	me.Encryption = new(string)
	return me.Encryption
}

func (me *HPDVersionType) SetSignature(value *string) {
	me.Signature = value
}

func (me *HPDVersionType) AddSignature() *string {
	me.Signature = new(string)
	return me.Signature
}
