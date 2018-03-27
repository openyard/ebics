// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// ComplexType
type UserPermissionType struct {
	AuthorisationLevel *AuthorisationLevelType `xml:"AuthorisationLevel,attr"`
	OrderTypes         *OrderTListType         `xml:"OrderTypes"`
	FileFormat         *FileFormatType         `xml:"FileFormat,omitempty"`
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

func (me *UserPermissionType) SetOrderTypes(value *OrderTListType) {
	me.OrderTypes = value
}

func (me *UserPermissionType) AddOrderTypes() *OrderTListType {
	me.OrderTypes = new(OrderTListType)
	return me.OrderTypes
}

func (me *UserPermissionType) SetFileFormat(value *FileFormatType) {
	me.FileFormat = value
}

func (me *UserPermissionType) AddFileFormat() *FileFormatType {
	me.FileFormat = new(FileFormatType)
	return me.FileFormat
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
