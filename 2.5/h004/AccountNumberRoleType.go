// Generated with goxc v0.1.1 - rev bae2cf01854d664b985cae6986076979716034c7
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

type AccountNumberRoleType w3c.Token

const (
	AccountNumberRoleType_ORIGINATOR AccountNumberRoleType = "Originator"
	AccountNumberRoleType_RECIPIENT  AccountNumberRoleType = "Recipient"
	AccountNumberRoleType_CHARGES    AccountNumberRoleType = "Charges"
	AccountNumberRoleType_OTHER      AccountNumberRoleType = "Other"
)
