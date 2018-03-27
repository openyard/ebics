// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
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
