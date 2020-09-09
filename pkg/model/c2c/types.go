package c2c

type AccountBalanceData struct {
	AccountId       int              `json:"accountId"`
	AccountStatus   string           `json:"accountStatus"`
	Symbol          string           `json:"symbol"`
	RiskRate        string           `json:"riskRate"`
	SubAccountTypes []SubAccountType `json:"subAccountTypes"`
}

type SubAccountType struct {
	Currency       string `json:"currency"`
	SubAccountType string `json:"subAccountType"`
	AcctBalance    string `json:"acctBalance"`  //账户余额
	AvailBalance   string `json:"availBalance"` //可用余额 （仅对借入账户下trade子类型有效
	Transferable   string `json:"transferable"` //可转出金额 （仅对借入账户下trade子类型有效）
	Borrowable     string `json:"borrowable"`   //可借入金额 （仅对借入账户下trade子类型有效）
}
type GetAccountBalanceResponse struct {
	Code    int64               `json:"code"`
	Message string              `json:"message"`
	Success bool                `json:"success"`
	Data    *AccountBalanceData `json:"data"`
}
