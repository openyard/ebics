package s001_test

import (
	"encoding/xml"
	"fmt"

	"github.com/openyard/ebics/2.5/s001"
)

func ExampleUnmarshalSignaturePubKeyOrderData() {
	var doc s001.SignaturePubKeyOrderData
	xml.Unmarshal([]byte(validSignaturePubKeyOrderDataDoc), &doc)
	fmt.Println(fmt.Sprintf("Version      :%10s", *doc.SignaturePubKeyInfo.SignatureVersion))
	fmt.Println(fmt.Sprintf("UserID       :%10s", *doc.UserID))
	fmt.Println(fmt.Sprintf("PartnerID    :%10s", *doc.PartnerID))
	fmt.Println(fmt.Sprintf("Modulus      :%10s", *doc.SignaturePubKeyInfo.PubKeyValue.RSAKeyValue.Modulus))
	fmt.Println(fmt.Sprintf("Exponent     :%10s", *doc.SignaturePubKeyInfo.PubKeyValue.RSAKeyValue.Exponent))
	// Output:
	//Version      :      A005
	//UserID       :  USER0001
	//PartnerID    :  CUSTOMER
	//Modulus      :AJIiZnAKeubfcRj7bA87iWthtS3gzaYPViKLqcQ9YzTZlTKyJoHbFCH151XcXceaUmLg37vd+B9wyZQZXRyv8+o4BC4GIWzcq9c3E1H9lilR8OIffYteZKYV5bvb6wBSzoQE+JykiL+YShjprYiFMKYz/KChQLJ7LLTKZXG8FpmbokCu/6bgjPuW9i0sTx4WwkAnUHxwpYdfrjIaVWDxWsqOFGseEhRNj1wYre/cCVOZaQvPvWyxNVysxF18SnbGfQ==
	//Exponent     :      AQAB

}

const validSignaturePubKeyOrderDataDoc = `<?xml version="1.0" encoding="UTF-8"?>
<s001:SignaturePubKeyOrderData xmlns:s001="http://www.ebics.org/S001">
  <s001:SignaturePubKeyInfo>
    <s001:PubKeyValue>
      <xd:RSAKeyValue xmlns:xd="http://www.w3.org/2000/09/xmldsig#">
        <xd:Modulus>AJIiZnAKeubfcRj7bA87iWthtS3gzaYPViKLqcQ9YzTZlTKyJoHbFCH151XcXceaUmLg37vd+B9wyZQZXRyv8+o4BC4GIWzcq9c3E1H9lilR8OIffYteZKYV5bvb6wBSzoQE+JykiL+YShjprYiFMKYz/KChQLJ7LLTKZXG8FpmbokCu/6bgjPuW9i0sTx4WwkAnUHxwpYdfrjIaVWDxWsqOFGseEhRNj1wYre/cCVOZaQvPvWyxNVysxF18SnbGfQ==</xd:Modulus>
        <xd:Exponent>AQAB</xd:Exponent>
      </xd:RSAKeyValue>
    </s001:PubKeyValue>
    <s001:SignatureVersion>A005</s001:SignatureVersion>
  </s001:SignaturePubKeyInfo>
  <s001:PartnerID>CUSTOMER</s001:PartnerID>
  <s001:UserID>USER0001</s001:UserID>
</s001:SignaturePubKeyOrderData>`
