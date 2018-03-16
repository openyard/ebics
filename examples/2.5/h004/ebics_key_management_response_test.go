package h004_test

import (
	"encoding/xml"
	"fmt"
	"github.com/openyard/ebics/2.5/h004"
	"github.com/openyard/ebics/2.5/w3c"
)

func ExampleMarshalEbicsKeyManagementResponse() {
	req := h004.NewEbicsKeyManagementResponse()
	req.AddHeader().AddMutable().SetOrderID(h004.NewOrderIDType(w3c.NewToken("A001")))
	req.Header.Mutable.SetReturnCode(h004.NewReturnCodeType(w3c.NewToken("000000")))
	req.Header.Mutable.SetReportText(h004.NewReportTextType(w3c.NewNormalizedString("[EBICS_OK] OK")))
	req.AddBody().AddReturnCode().Value = h004.NewReturnCodeType(w3c.NewToken("000000"))
	req.Body.ReturnCode.SetAuthenticate(w3c.NewBoolean(true))
	raw, _ := xml.MarshalIndent(&req, "", "  ")
	fmt.Println(fmt.Sprintf("\n%s", string(raw)))
	// Output:
	//
	// <EbicsKeyManagementResponse>
	//   <header>
	//     <mutable>
	//       <OrderID>A001</OrderID>
	//       <ReturnCode>000000</ReturnCode>
	//       <ReportText>[EBICS_OK] OK</ReportText>
	//     </mutable>
	//   </header>
	//   <body>
	//     <ReturnCode authenticate="true">000000</ReturnCode>
	//   </body>
	// </EbicsKeyManagementResponse>
}
