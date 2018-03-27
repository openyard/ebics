// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// ComplexType
type PreValidationRequestType struct {
	DataDigest           []*DataDigestType               `xml:"DataDigest,omitempty"`
	AccountAuthorisation []*PreValidationAccountAuthType `xml:"AccountAuthorisation,omitempty"`

	Any []*w3c.Any
}

func NewPreValidationRequestType() *PreValidationRequestType {
	return new(PreValidationRequestType)
}

func (me *PreValidationRequestType) AddDataDigest(value *DataDigestType) {
	me.DataDigest = append(me.DataDigest, value)
}

func (me *PreValidationRequestType) AddAccountAuthorisation(value *PreValidationAccountAuthType) {
	me.AccountAuthorisation = append(me.AccountAuthorisation, value)
}
