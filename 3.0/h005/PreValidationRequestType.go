// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
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
