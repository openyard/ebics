// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package s001

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
