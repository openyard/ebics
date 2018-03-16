// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
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

func (me *UserPermissionType) SetService(value *RestrictedServiceType) {
	me.Service = value
}

func (me *UserPermissionType) SetAccountID(value *AccountIDType) {
	me.AccountID = value
}

func (me *UserPermissionType) SetMaxAmount(value *AmountType) {
	me.MaxAmount = value
}
