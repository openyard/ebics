// Generated with goxc v - rev
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex element
type AccountTypeAccountNumber struct {
	Value         AccountNumberType `xml:",chardata"`
	International w3c.Boolean       `xml:"International,attr,omitempty"`
}
