// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package xmldsig

import w3c "github.com/openyard/ebics/2.5/w3c"

// Attribute
type MimeType struct {
	Value *w3c.String `xml:"MimeType,attr"`
}

func NewMimeType() *MimeType {
	return new(MimeType)
}

func (me *MimeType) SetValue(value *w3c.String) {
	me.Value = value
}
