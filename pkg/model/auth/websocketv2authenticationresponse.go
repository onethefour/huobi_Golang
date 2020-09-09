package auth

import (
	"encoding/json"
	"huobi_Golang/pkg/model/base"
)

type WebSocketV2AuthenticationResponse struct {
	base.WebSocketV2ResponseBase
}

func ParseWSV2AuthResp(message string) *WebSocketV2AuthenticationResponse {
	result := &WebSocketV2AuthenticationResponse{}
	err := json.Unmarshal([]byte(message), result)
	if err != nil {
		return nil
	}

	return result
}
