// Generated with goxc v0.1.7 - rev 2ae97d253f5eaa17bab360dad75945920dfceef4
package h005

// complex type
type BTFParamsTyp struct {
	FileName      FileNameStringType    `xml:"fileName,attr"`
	Service       RestrictedServiceType `xml:"Service"`
	SignatureFlag SignatureFlagType     `xml:"SignatureFlag,omitempty"`
	DateRange     DateRangeType         `xml:"DateRange,omitempty"`
	Parameter     Parameter             `xml:"Parameter,omitempty"`
}
