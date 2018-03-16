// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package w3c

import "encoding/xml"

// Any Type
type Any struct {
	XMLName xml.Name
	Content string `xml:",innerxml,omitempty"`
}
