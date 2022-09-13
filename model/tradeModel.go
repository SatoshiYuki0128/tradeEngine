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
	//Id         int    `json:"id"`
	//Type       string `json:"type"`
	Quantity int `json:"quantity"`
	Price    int `json:"price"`
	//Close      int    `json:"close"`
	//UpdateTime string `json:"updateTime"`
}
