package contract

type ContractInfoResponse struct {
	Status string                      `json:"status"`
	Ts     int64                       `json:"ts"`
	Ch     string                      `json:"ch"`
	Data   []ContractInfoResponse_Data `json:"data"`
}
type ContractInfoResponse_Data struct {
	Symbol         string  `json:"symbol"`
	ContractCode   string  `json:"contract_code"`
	ContractType   string  `json:"contract_type"`
	ContractSize   float64 `json:"contract_size"`
	PriceTick      float64 `json:"price_tick"`
	DeliveryDate   string  `json:"delivery_date"`
	CreateDate     string  `json:"create_date"`
	ContractStatus int64   `json:"contract_status"`
}

//type ContractInfoResponse_Data2 struct {
//	Ts   int64                            `json:"ts"`
//	Id   int64                            `json:"id"`
//	Data []ContractInfoResponse_data_data `json:"data"`
//}
//type ContractInfoResponse_data_data struct {
//	Amount    int64   `json:"amount"`
//	Direction string  `json:"direction"`
//	Id        int64   `json:"id"`
//	Price     float64 `json:"price"`
//	Ts        int64   `json:"ts"`
//}
type ContractBasis struct {
	Status string               `json:"status"`
	Ts     int64                `json:"ts"`
	Ch     string               `json:"ch"`
	Data   []ContractBasis_Data `json:"data"`
}
type ContractBasis_Data struct {
	Basis         string `json:"basis"`
	BasisRate     string `json:"basis_rate"`
	ContractPrice string `json:"contract_price"`
	Id            int64  `json:"id"`
	IndexPrice    string `json:"index_price"`
}

type DetailMerged struct {
	Status string            `json:"status"`
	Ts     int64             `json:"ts"`
	Ch     string            `json:"ch"`
	Tick   DetailMerged_tick `json:"tick"`
}
type DetailMerged_tick struct {
	Amount string    `json:"amount"`
	Ask    []float64 `json:"ask"`
	Bid    []float64 `json:"bid"`
	Close  string    `json:"close"`
	Count  int64     `json:"count"`
	High   string    `json:"high"`
	Id     int64     `json:"id"`
	Low    string    `json:"low"`
	Open   string    `json:"open"`
	Ts     int64     `json:"ts"`
	Vol    string    `json:"vol"`
}
