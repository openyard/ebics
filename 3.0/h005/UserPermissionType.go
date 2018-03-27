// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexType
type UserPermissionType struct {
	AuthorisationLevel *AuthorisationLevelType `xml:"AuthorisationLevel,attr"`
	AdminOrderType     *OrderTBaseType         `xml:"AdminOrderType"`
	Service            *RestrictedServiceType  `xml:"Service,omitempty"`
	AccountID          *AccountIDType          `xml:"AccountID,omitempty"`
	MaxAmount          *AmountType             `xml:"MaxAmount,omitempty"`

	Any []*w3c.Any
}

func NewUserPermissionType() *UserPermissionType {
	return new(UserPermissionType)
}

func (me *UserPermissionType) SetAuthorisationLevel(value *AuthorisationLevelType) {
	me.AuthorisationLevel = value
}

func (me *UserPermissionType) SetAdminOrderType(value *OrderTBaseType) {
	me.AdminOrderType = value
}

func (me *UserPermissionType) AddAdminOrderType() *OrderTBaseType {
	me.AdminOrderType = new(OrderTBaseType)
	return me.AdminOrderType
}

func (me *UserPermissionType) SetService(value *RestrictedServiceType) {
	me.Service = value
}

func (me *UserPermissionType) AddService() *RestrictedServiceType {
	me.Service = new(RestrictedServiceType)
	return me.Service
}

func (me *UserPermissionType) SetAccountID(value *AccountIDType) {
	me.AccountID = value
}

func (me *UserPermissionType) AddAccountID() *AccountIDType {
	me.AccountID = new(AccountIDType)
	return me.AccountID
}

func (me *UserPermissionType) SetMaxAmount(value *AmountType) {
	me.MaxAmount = value
}

func (me *UserPermissionType) AddMaxAmount() *AmountType {
	me.MaxAmount = new(AmountType)
	return me.MaxAmount
}
