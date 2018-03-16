// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

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
