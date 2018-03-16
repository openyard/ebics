// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

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

func (me *AddressInfoType) SetStreet(value *NameType) {
	me.Street = value
}

func (me *AddressInfoType) SetPostCode(value *w3c.Token) {
	me.PostCode = value
}

func (me *AddressInfoType) SetCity(value *NameType) {
	me.City = value
}

func (me *AddressInfoType) SetRegion(value *NameType) {
	me.Region = value
}

func (me *AddressInfoType) SetCountry(value *NameType) {
	me.Country = value
}
