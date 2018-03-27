// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package xmldsig

import w3c "github.com/openyard/ebics/2.5/w3c"

// ComplexType
type SignedInfoType struct {
	Id                     *w3c.ID                 `xml:"Id,attr,omitempty"`
	CanonicalizationMethod *CanonicalizationMethod `xml:"CanonicalizationMethod"`
	SignatureMethod        *SignatureMethod        `xml:"SignatureMethod"`
	Reference              *Reference              `xml:"Reference"`
}

func NewSignedInfoType() *SignedInfoType {
	return new(SignedInfoType)
}

func (me *SignedInfoType) SetId(value *w3c.ID) {
	me.Id = value
}

func (me *SignedInfoType) SetCanonicalizationMethod(value *CanonicalizationMethod) {
	me.CanonicalizationMethod = value
}

func (me *SignedInfoType) SetSignatureMethod(value *SignatureMethod) {
	me.SignatureMethod = value
}

func (me *SignedInfoType) SetReference(value *Reference) {
	me.Reference = value
}
