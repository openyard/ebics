// Generated with goxc v0.1.1 - rev bae2cf01854d664b985cae6986076979716034c7
package h004

import ds "github.com/openyard/ebics/2.5/xmldsig"
import w3c "github.com/openyard/ebics/2.5/w3c"

// complex type
type PubKeyValueType struct {
	TimeStamp   TimestampType  `xml:"TimeStamp,omitempty"`
	RSAKeyValue ds.RSAKeyValue `xml:"http://www.w3.org/2000/09/xmldsig# RSAKeyValue"`

	Any []*w3c.Any
}
