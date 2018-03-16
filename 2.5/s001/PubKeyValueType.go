// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
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

func (me *PubKeyValueType) SetRSAKeyValue(value *ds.RSAKeyValue) {
	me.RSAKeyValue = value
}
