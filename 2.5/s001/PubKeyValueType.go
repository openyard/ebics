// Generated with goxc v0.1.2 - rev bbe25b23ba83e8f8464e681ca3e514eee51fd2ba
package s001

import ds "github.com/openyard/ebics/2.5/xmldsig"
import w3c "github.com/openyard/ebics/2.5/w3c"

// complex type
type PubKeyValueType struct {
	TimeStamp   TimestampType  `xml:"TimeStamp,omitempty"`
	RSAKeyValue ds.RSAKeyValue `xml:"http://www.w3.org/2000/09/xmldsig# RSAKeyValue"`

	Any []*w3c.Any
}
