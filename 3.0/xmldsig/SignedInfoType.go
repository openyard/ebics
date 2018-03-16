// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package xmldsig

import w3c "github.com/openyard/ebics/3.0/w3c"

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
