// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

import ds "github.com/openyard/ebics/2.5/xmldsig"
import w3c "github.com/openyard/ebics/2.5/w3c"

// ComplexType
type PubKeyValueType struct {
	TimeStamp   *TimestampType  `xml:"TimeStamp,omitempty"`
	RSAKeyValue *ds.RSAKeyValue `xml:"http://www.w3.org/2000/09/xmldsig# RSAKeyValue"`

	Any []*w3c.Any
}

func NewPubKeyValueType() *PubKeyValueType {
	return new(PubKeyValueType)
}

func (me *PubKeyValueType) SetTimeStamp(value *TimestampType) {
	me.TimeStamp = value
}

func (me *PubKeyValueType) AddTimeStamp() *TimestampType {
	me.TimeStamp = new(TimestampType)
	return me.TimeStamp
}

func (me *PubKeyValueType) SetRSAKeyValue(value *ds.RSAKeyValue) {
	me.RSAKeyValue = value
}
