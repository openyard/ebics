// Generated with goxc v0.1.3 - rev 0e63342ac0a4d5f52582ea0065a462e700069839
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

type AccountNumberRoleType w3c.Token

const (
	AccountNumberRoleType_ORIGINATOR AccountNumberRoleType = "Originator"
	AccountNumberRoleType_RECIPIENT  AccountNumberRoleType = "Recipient"
	AccountNumberRoleType_CHARGES    AccountNumberRoleType = "Charges"
	AccountNumberRoleType_OTHER      AccountNumberRoleType = "Other"
)
