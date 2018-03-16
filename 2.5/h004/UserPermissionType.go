// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
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

func (me *UserPermissionType) SetFileFormat(value *FileFormatType) {
	me.FileFormat = value
}

func (me *UserPermissionType) SetAccountID(value *AccountIDType) {
	me.AccountID = value
}

func (me *UserPermissionType) SetMaxAmount(value *AmountType) {
	me.MaxAmount = value
}
