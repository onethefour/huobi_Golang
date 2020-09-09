package subuserclientexample

import (
	"huobi_Golang/common/config"
	"huobi_Golang/common/log"
	"huobi_Golang/pkg/client"
	"huobi_Golang/pkg/model/subuser"
)

func RunAllExamples() {
	createSubUser()
	lockSubUser()
	unlockSubUser()
	getSubUserDepositAddress()
	querySubUserDepositHistory()
}

func createSubUser() {
	client := new(client.SubUserClient).Init(config.AccessKey, config.SecretKey, config.Host)
	request := subuser.CreateSubUserRequest{
		UserList: []subuser.Users{
			subuser.Users{"subuser1412", "sub-user-1-note"},
			subuser.Users{"subuser1413", "sub-user-2-note"},
		},
	}

	resp, err := client.CreateSubUser(request)
	if err != nil {
		log.Error("Create sub user error: %s", err)
	} else {
		log.Info("Create sub user, count=%d", len(resp))
		for _, result := range resp {
			log.Info("sub user: %+v", result)
		}
	}
}

func lockSubUser() {
	client := new(client.SubUserClient).Init(config.AccessKey, config.SecretKey, config.Host)
	subUserManagementRequest := subuser.SubUserManagementRequest{SubUid: config.SubUid, Action: "lock"}
	resp, err := client.SubUserManagement(subUserManagementRequest)
	if err != nil {
		log.Error("Lock sub user error: %s", err)
	} else {
		log.Info("Lock sub user: %+v", resp)
	}
}

func unlockSubUser() {
	client := new(client.SubUserClient).Init(config.AccessKey, config.SecretKey, config.Host)
	subUserManagementRequest := subuser.SubUserManagementRequest{SubUid: config.SubUid, Action: "unlock"}
	resp, err := client.SubUserManagement(subUserManagementRequest)
	if err != nil {
		log.Error("Unlock sub user error: %s", err)
	} else {
		log.Info("Unlock sub user: %+v", resp)
	}
}

func getSubUserDepositAddress() {
	client := new(client.SubUserClient).Init(config.AccessKey, config.SecretKey, config.Host)
	currency := "usdt"
	resp, err := client.GetSubUserDepositAddress(config.SubUid, currency)
	if err != nil {
		log.Error("Get sub user deposit address error: %s", err)
	} else {
		log.Info("Get sub user deposit address, count=%d", len(resp))
		for _, result := range resp {
			log.Info("DepositAddress: %+v", result)
		}
	}
}

func querySubUserDepositHistory() {
	client := new(client.SubUserClient).Init(config.AccessKey, config.SecretKey, config.Host)
	optionalRequest := subuser.QuerySubUserDepositHistoryOptionalRequest{Currency: "usdt"}
	resp, err := client.QuerySubUserDepositHistory(config.SubUid, optionalRequest)
	if err != nil {
		log.Error("Query sub user deposit history error: %s", err)
	} else {
		log.Info("Query sub user deposit history, count=%d", len(resp))
		for _, result := range resp {
			log.Info("resp: %+v", result)
		}
	}
}
