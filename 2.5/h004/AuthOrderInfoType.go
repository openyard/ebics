// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// ComplexType
type AuthOrderInfoType struct {
	OrderType      *OrderTBaseType         `xml:"OrderType"`
	FileFormat     *FileFormatType         `xml:"FileFormat,omitempty"`
	TransferType   *TransferType           `xml:"TransferType"`
	OrderFormat    *OrderFormatType        `xml:"OrderFormat,omitempty"`
	Description    *OrderDescriptionType   `xml:"Description"`
	NumSigRequired *w3c.NonNegativeInteger `xml:"NumSigRequired,omitempty"`

	Any []*w3c.Any
}

func NewAuthOrderInfoType() *AuthOrderInfoType {
	return new(AuthOrderInfoType)
}

func (me *AuthOrderInfoType) SetOrderType(value *OrderTBaseType) {
	me.OrderType = value
}

func (me *AuthOrderInfoType) SetFileFormat(value *FileFormatType) {
	me.FileFormat = value
}

func (me *AuthOrderInfoType) SetTransferType(value *TransferType) {
	me.TransferType = value
}

func (me *AuthOrderInfoType) SetOrderFormat(value *OrderFormatType) {
	me.OrderFormat = value
}

func (me *AuthOrderInfoType) SetDescription(value *OrderDescriptionType) {
	me.Description = value
}

func (me *AuthOrderInfoType) SetNumSigRequired(value *w3c.NonNegativeInteger) {
	me.NumSigRequired = value
}
