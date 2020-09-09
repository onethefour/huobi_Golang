package account

type GetAccountInfoResponse struct {
	Status string        `json:"status"`
	Data   []AccountInfo `json:"data"`
}
type AccountInfo struct {
	Id      int64  `json:"id"`
	Type    string `json:"type"` //spot：现货账户， margin：逐仓杠杆账户，otc：OTC 账户，point：点卡账户，super-margin：全仓杠杆账户, investment: C2C杠杆借出账户, borrow: C2C杠杆借入账户
	Subtype string `json:"subtype"`
	State   string `json:"state"`
}
