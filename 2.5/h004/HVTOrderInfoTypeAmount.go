// Generated with goxc vgoxc-0.1.8 - rev 7e2e945f706bc13e7539c26efd1ec70bc280277e
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex element
type HVTOrderInfoTypeAmount struct {
    Value AmountValueType `xml:",chardata"`
    IsCredit w3c.Boolean `xml:"IsCredit,attr,omitempty"`
    Currency CurrencyBaseType `xml:"Currency,attr,omitempty"`
    
    }
