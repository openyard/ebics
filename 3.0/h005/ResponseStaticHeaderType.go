// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

// ComplexType
type ResponseStaticHeaderType struct {
	TransactionID *TransactionIDType `xml:"TransactionID,omitempty"`
	NumSegments   *SegmentNumberType `xml:"NumSegments,omitempty"`
}

func NewResponseStaticHeaderType() *ResponseStaticHeaderType {
	return new(ResponseStaticHeaderType)
}

func (me *ResponseStaticHeaderType) SetTransactionID(value *TransactionIDType) {
	me.TransactionID = value
}

func (me *ResponseStaticHeaderType) AddTransactionID() *TransactionIDType {
	me.TransactionID = new(TransactionIDType)
	return me.TransactionID
}

func (me *ResponseStaticHeaderType) SetNumSegments(value *SegmentNumberType) {
	me.NumSegments = value
}

func (me *ResponseStaticHeaderType) AddNumSegments() *SegmentNumberType {
	me.NumSegments = new(SegmentNumberType)
	return me.NumSegments
}
