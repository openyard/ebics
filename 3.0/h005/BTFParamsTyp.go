// Generated with goxc vgoxc-0.1.8 - rev 7e2e945f706bc13e7539c26efd1ec70bc280277e
package h005

// complex type
type BTFParamsTyp struct {
	FileName      FileNameStringType    `xml:"fileName,attr"`
	Service       RestrictedServiceType `xml:"Service"`
	SignatureFlag SignatureFlagType     `xml:"SignatureFlag,omitempty"`
	DateRange     DateRangeType         `xml:"DateRange,omitempty"`
	Parameter     Parameter             `xml:"Parameter,omitempty"`
}
