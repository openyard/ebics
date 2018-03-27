// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
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

func (me *AuthOrderInfoType) AddAdminOrderType() *OrderTBaseType {
	me.AdminOrderType = new(OrderTBaseType)
	return me.AdminOrderType
}

func (me *AuthOrderInfoType) SetService(value *RestrictedServiceType) {
	me.Service = value
}

func (me *AuthOrderInfoType) AddService() *RestrictedServiceType {
	me.Service = new(RestrictedServiceType)
	return me.Service
}

func (me *AuthOrderInfoType) SetDescription(value *OrderDescriptionType) {
	me.Description = value
}

func (me *AuthOrderInfoType) AddDescription() *OrderDescriptionType {
	me.Description = new(OrderDescriptionType)
	return me.Description
}

func (me *AuthOrderInfoType) SetNumSigRequired(value *w3c.NonNegativeInteger) {
	me.NumSigRequired = value
}

func (me *AuthOrderInfoType) AddNumSigRequired() *w3c.NonNegativeInteger {
	me.NumSigRequired = new(w3c.NonNegativeInteger)
	return me.NumSigRequired
}
