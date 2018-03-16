// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexType
type AuthOrderInfoType struct {
	AdminOrderType *OrderTBaseType         `xml:"AdminOrderType"`
	Service        *RestrictedServiceType  `xml:"Service,omitempty"`
	Description    *OrderDescriptionType   `xml:"Description"`
	NumSigRequired *w3c.NonNegativeInteger `xml:"NumSigRequired,omitempty"`

	Any []*w3c.Any
}

func NewAuthOrderInfoType() *AuthOrderInfoType {
	return new(AuthOrderInfoType)
}

func (me *AuthOrderInfoType) SetAdminOrderType(value *OrderTBaseType) {
	me.AdminOrderType = value
}

func (me *AuthOrderInfoType) SetService(value *RestrictedServiceType) {
	me.Service = value
}

func (me *AuthOrderInfoType) SetDescription(value *OrderDescriptionType) {
	me.Description = value
}

func (me *AuthOrderInfoType) SetNumSigRequired(value *w3c.NonNegativeInteger) {
	me.NumSigRequired = value
}
