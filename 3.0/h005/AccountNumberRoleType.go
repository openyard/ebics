// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

type AccountNumberRoleType w3c.Token

const (
	AccountNumberRoleType_ORIGINATOR AccountNumberRoleType = "Originator"
	AccountNumberRoleType_RECIPIENT  AccountNumberRoleType = "Recipient"
	AccountNumberRoleType_CHARGES    AccountNumberRoleType = "Charges"
	AccountNumberRoleType_OTHER      AccountNumberRoleType = "Other"
)
