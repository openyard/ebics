// Generated with goxc vgoxc-0.1.9 - rev 260439f4ef82f3f152002242cdec0bb97e6118c3
package w3c

import "encoding/xml"

// any type
type Any struct {
	XMLName xml.Name
	Content string `xml:",innerxml,omitempty"`
}
