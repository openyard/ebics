// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h004

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

func (me *ResponseStaticHeaderType) SetNumSegments(value *SegmentNumberType) {
	me.NumSegments = value
}
