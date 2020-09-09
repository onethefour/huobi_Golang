package models

var Accounts map[string]string

func init() {
	Accounts = make(map[string]string)
	Accounts["xutonghua"] = "123456"
}
