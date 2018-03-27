// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// ComplexType
type AddressInfoType struct {
	Name     *NameType  `xml:"Name,omitempty"`
	Street   *NameType  `xml:"Street,omitempty"`
	PostCode *w3c.Token `xml:"PostCode,omitempty"`
	City     *NameType  `xml:"City,omitempty"`
	Region   *NameType  `xml:"Region,omitempty"`
	Country  *NameType  `xml:"Country,omitempty"`

	Any []*w3c.Any
}

func NewAddressInfoType() *AddressInfoType {
	return new(AddressInfoType)
}

func (me *AddressInfoType) SetName(value *NameType) {
	me.Name = value
}

func (me *AddressInfoType) AddName() *NameType {
	me.Name = new(NameType)
	return me.Name
}

func (me *AddressInfoType) SetStreet(value *NameType) {
	me.Street = value
}

func (me *AddressInfoType) AddStreet() *NameType {
	me.Street = new(NameType)
	return me.Street
}

func (me *AddressInfoType) SetPostCode(value *w3c.Token) {
	me.PostCode = value
}

func (me *AddressInfoType) AddPostCode() *w3c.Token {
	me.PostCode = new(w3c.Token)
	return me.PostCode
}

func (me *AddressInfoType) SetCity(value *NameType) {
	me.City = value
}

func (me *AddressInfoType) AddCity() *NameType {
	me.City = new(NameType)
	return me.City
}

func (me *AddressInfoType) SetRegion(value *NameType) {
	me.Region = value
}

func (me *AddressInfoType) AddRegion() *NameType {
	me.Region = new(NameType)
	return me.Region
}

func (me *AddressInfoType) SetCountry(value *NameType) {
	me.Country = value
}

func (me *AddressInfoType) AddCountry() *NameType {
	me.Country = new(NameType)
	return me.Country
}
