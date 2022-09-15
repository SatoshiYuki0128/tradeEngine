package model

type Trade struct {
	Id         int    `json:"id"`
	Type       string `json:"type"`
	Quantity   int    `json:"quantity"`
	Price      int    `json:"price"`
	Close      int    `json:"close"`
	UpdateTime string `json:"updateTime"`
}

type Sell struct {
	Quantity int `json:"quantity"`
	Price    int `json:"price"`
}

type SellResponse struct {
	Message string `json:"message"`
	BuyerId []int  `json:"tradeId"`
}

type BuyResponse struct {
	Message  string `json:"message"`
	SellerId []int  `json:"tradeId"`
}

type Buy struct {
	Quantity int `json:"quantity"`
	Price    int `json:"price"`
}

type Search struct {
	Id              int    `json:"id"`
	Type            string `json:"type"`
	Quantity        int    `json:"quantity"`
	Price           int    `json:"price"`
	Close           int    `json:"close"`
	UpdateTimeStart string `json:"updateTimeStart"`
	UpdateTimeEnd   string `json:"updateTimeEnd"`
}
