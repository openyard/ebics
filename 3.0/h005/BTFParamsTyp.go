// Generated with goxc vgoxc-0.1.9 - rev 260439f4ef82f3f152002242cdec0bb97e6118c3
package h005

// complex type
type BTFParamsTyp struct {
	FileName      *FileNameStringType    `xml:"fileName,attr"`
	Service       *RestrictedServiceType `xml:"Service"`
	SignatureFlag *SignatureFlagType     `xml:"SignatureFlag,omitempty"`
	DateRange     *DateRangeType         `xml:"DateRange,omitempty"`
	Parameter     *Parameter             `xml:"Parameter,omitempty"`
}
