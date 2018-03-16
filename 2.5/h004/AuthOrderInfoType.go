// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
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

func (me *AuthOrderInfoType) AddOrderType() *OrderTBaseType {
	me.OrderType = new(OrderTBaseType)
	return me.OrderType
}

func (me *AuthOrderInfoType) SetFileFormat(value *FileFormatType) {
	me.FileFormat = value
}

func (me *AuthOrderInfoType) AddFileFormat() *FileFormatType {
	me.FileFormat = new(FileFormatType)
	return me.FileFormat
}

func (me *AuthOrderInfoType) SetTransferType(value *TransferType) {
	me.TransferType = value
}

func (me *AuthOrderInfoType) AddTransferType() *TransferType {
	me.TransferType = new(TransferType)
	return me.TransferType
}

func (me *AuthOrderInfoType) SetOrderFormat(value *OrderFormatType) {
	me.OrderFormat = value
}

func (me *AuthOrderInfoType) AddOrderFormat() *OrderFormatType {
	me.OrderFormat = new(OrderFormatType)
	return me.OrderFormat
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
