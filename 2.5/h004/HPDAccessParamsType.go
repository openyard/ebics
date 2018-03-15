// Generated with goxc vgoxc-0.1.9 - rev 260439f4ef82f3f152002242cdec0bb97e6118c3
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex type
type HPDAccessParamsType struct {
	URL       []*HPDAccessParamsTypeURL `xml:"URL"`
	Institute *w3c.NormalizedString     `xml:"Institute"`
	HostID    *HostIDType               `xml:"HostID,omitempty"`

	Any []*w3c.Any
}
