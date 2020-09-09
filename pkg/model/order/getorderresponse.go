package order

type GetOrderResponse struct {
	Status       string `json:"status"`
	ErrorCode    string `json:"err-code"`
	ErrorMessage string `json:"err-msg"`
	Data         *struct {
		AccountId        int     `json:"account-id"`
		Amount_f         float64 `json:"-"`
		Amount           string  `json:"amount"`
		Id               int64   `json:"id"`
		ClientOrderId    string  `json:"client-order-id"`
		Symbol           string  `json:"symbol"`
		Price_f          float64 `json:"-"`
		Price            string  `json:"price"`
		CreatedAt        int64   `json:"created-at"`
		Type             string  `json:"type"`
		FilledAmount     string  `json:"field-amount"`
		FilledCashAmount string  `json:"field-cash-amount"`
		FilledFees       string  `json:"field-fees"`
		Finishedat       uint64  `json:"finished-at"` //订单变为终结态的时间，不是成交时间，包含“已撤单”状态
		Source           string  `json:"source"`
		State            string  `json:"state"`
	}
}
